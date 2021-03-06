# ssm-cache

Simple ssm-cache extracted from [lambda-cache-example](github.com/wolfeidau/lambda-cache-example).

Caches a SSM parameter (defaults for 30 seconds).
This ensures that you only hit the [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html) API to refresh and not every time the lambda is triggered, therefore avoiding rate limiting your self.

## example

```go
package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/wolviecb/go-ssm/ssmcache"
)

var (
  cache = ssmcache.New(
    session.Must(
      session.NewSessionWithOptions(
        session.Options{
          Config:            aws.Config{Region: aws.String("eu-west-1")},
          SharedConfigState: session.SharedConfigEnable,
        },
      ),
    ),
  )
)

func ssmCheck() string {
  keyname := "/some/param"
  ssmvc, err := cache.GetKey(keyname)
  if err != nil {
    return "Error getting SSM parameter"
  }
  return ssmvc
}
```

to set the default expiry time run, the default is 30s

```go
ssmcache.SetDefaultExpiry(300 * time.Second)
```

to fetch the parameter with decryption run, the default is false

```go
ssmcache.SetDefaultDecryption(true)
```

## license

This code is released under MIT License.
