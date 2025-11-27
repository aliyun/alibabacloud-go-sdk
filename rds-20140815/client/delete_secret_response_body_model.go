// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecretResponseBody
	GetRequestId() *string
	SetSecretArn(v string) *DeleteSecretResponseBody
	GetSecretArn() *string
	SetSecretName(v string) *DeleteSecretResponseBody
	GetSecretName() *string
	SetSuccess(v bool) *DeleteSecretResponseBody
	GetSuccess() *bool
}

type DeleteSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DF4961DD-16F5-5B24-BD4C-0C7788F7ADAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the credential for the Data API account.
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretResponseBody) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DeleteSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSecretArn(v string) *DeleteSecretResponseBody {
	s.SecretArn = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSecretName(v string) *DeleteSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSuccess(v bool) *DeleteSecretResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
