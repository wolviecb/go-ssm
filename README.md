# ssm-cache

Simple ssm-cache extracted from [lambda-cache-example](github.com/wolfeidau/lambda-cache-example).

Caches a SSM parameter (defaults for 30 seconds).
This ensures that you only hit the [AWS Systems Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html) API to refresh and not every time the lambda is triggered, therefore avoiding rate limiting your self.

## example

The below example is based on the api gateway code in [Announcing Go Support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/).

```go
region := "eu-west-1"

var (
  cache = ssmcache.New(
    session.Must(
      session.NewSessionWithOptions(
        session.Options{
          Config:            aws.Config{Region: aws.String(region)},
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

## license

This code is released under MIT License.
