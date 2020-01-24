package ssmcache

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/stretchr/testify/mock"
	"github.com/wolviecb/go-ssm/mocks"
)

func TestGetKey(t *testing.T) {

	ssmMock := &mocks.SSMAPI{}

	gpo := &ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Name:  aws.String("testtest"),
			Value: aws.String("sup"),
		},
	}

	ssmMock.On("GetParameter", mock.AnythingOfType("*ssm.GetParameterInput")).Return(gpo, nil)

	cache := &cache{
		ssmSvc:    ssmMock,
		ssmValues: make(map[string]*Entry),
	}

	val, err := cache.GetKey("testtest")
	require.Nil(t, err)
	require.Equal(t, "sup", val)
}

func TestGetKeyWithEncryption(t *testing.T) {

	ssmMock := &mocks.SSMAPI{}

	gpo := &ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Name:  aws.String("testtest"),
			Value: aws.String("sup"),
		},
	}

	ssmMock.On("GetParameter", mock.AnythingOfType("*ssm.GetParameterInput")).Return(gpo, nil)

	cache := &cache{
		ssmSvc:    ssmMock,
		ssmValues: make(map[string]*Entry),
	}

	val, err := cache.GetKeyWithEncryption("testtest", true)
	require.Nil(t, err)
	require.Equal(t, "sup", val)
}
