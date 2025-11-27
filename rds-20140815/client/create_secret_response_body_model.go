// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSecretResponseBody
	GetRequestId() *string
	SetSecretArn(v string) *CreateSecretResponseBody
	GetSecretArn() *string
	SetSecretName(v string) *CreateSecretResponseBody
	GetSecretName() *string
	SetSuccess(v bool) *CreateSecretResponseBody
	GetSuccess() *bool
}

type CreateSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DF4961DD-16F5-5B24-BD4C-0C7788F7ADAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the credential for the created Data API account.
	//
	// example:
	//
	// acs:rds:cn-hangzhou:1335786***:dbInstance/rm-bp1m7l3j63****
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// Foo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretResponseBody) GetSecretArn() *string {
	return s.SecretArn
}

func (s *CreateSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSecretResponseBody) SetRequestId(v string) *CreateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretArn(v string) *CreateSecretResponseBody {
	s.SecretArn = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretName(v string) *CreateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *CreateSecretResponseBody) SetSuccess(v bool) *CreateSecretResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
