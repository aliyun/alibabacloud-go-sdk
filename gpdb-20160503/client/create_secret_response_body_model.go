// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecretResponseBody
	GetRequestId() *string
	SetSecretArn(v string) *CreateSecretResponseBody
	GetSecretArn() *string
	SetSecretName(v string) *CreateSecretResponseBody
	GetSecretName() *string
	SetStatus(v string) *CreateSecretResponseBody
	GetStatus() *string
}

type CreateSecretResponseBody struct {
	// The detailed information returned by the operation.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN that uniquely identifies the logon credential and permissions of the Data API account created by the user.
	//
	// Format: ```acs:gpdb:{{region}}:{{accountId}}:secret/{{secretName}}-{{32 digits random string}}```
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// testsecret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The status. Valid values:
	//
	// - **success**: The operation is successful.
	//
	// - **fail**: The operation failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) GetMessage() *string {
	return s.Message
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

func (s *CreateSecretResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateSecretResponseBody) SetMessage(v string) *CreateSecretResponseBody {
	s.Message = &v
	return s
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

func (s *CreateSecretResponseBody) SetStatus(v string) *CreateSecretResponseBody {
	s.Status = &v
	return s
}

func (s *CreateSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
