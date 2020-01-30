package main

import (
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sts.New(sess)

	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
	}

	result, err := svc.GetSessionToken(input)
	check(err)

	string := "[default]\naws_access_key_id=" + *result.Credentials.AccessKeyId +
		"\naws_secret_access_key=" + *result.Credentials.SecretAccessKey +
		"\naws_session_token=" + *result.Credentials.SessionToken + "\n"
	creds := []byte(string)
	err = ioutil.WriteFile(os.Getenv("HOME")+"/.aws/credentials", creds, 0644)
	check(err)
}
