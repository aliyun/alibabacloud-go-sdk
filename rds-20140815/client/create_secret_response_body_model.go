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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretArn  *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
