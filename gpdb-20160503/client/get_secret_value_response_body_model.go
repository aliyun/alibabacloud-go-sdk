// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecretValueResponseBody
	GetCode() *string
	SetDBInstanceId(v string) *GetSecretValueResponseBody
	GetDBInstanceId() *string
	SetDescription(v string) *GetSecretValueResponseBody
	GetDescription() *string
	SetMessage(v string) *GetSecretValueResponseBody
	GetMessage() *string
	SetPassword(v string) *GetSecretValueResponseBody
	GetPassword() *string
	SetRequestId(v string) *GetSecretValueResponseBody
	GetRequestId() *string
	SetSecretArn(v string) *GetSecretValueResponseBody
	GetSecretArn() *string
	SetSecretName(v string) *GetSecretValueResponseBody
	GetSecretName() *string
	SetStatus(v string) *GetSecretValueResponseBody
	GetStatus() *string
	SetUsername(v string) *GetSecretValueResponseBody
	GetUsername() *string
}

type GetSecretValueResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Secret.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the access credential.
	//
	// example:
	//
	// test secret
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// pwd123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARN of the access credential for the created Data API account. Format: `acs:gpdb:{{region}}:{{accountId}}:secret/{{secretName}}-{{32 digits random string}`.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the access credential.
	//
	// example:
	//
	// testsecret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testacc
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecretValueResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GetSecretValueResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSecretValueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecretValueResponseBody) GetPassword() *string {
	return s.Password
}

func (s *GetSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretValueResponseBody) GetSecretArn() *string {
	return s.SecretArn
}

func (s *GetSecretValueResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *GetSecretValueResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSecretValueResponseBody) GetUsername() *string {
	return s.Username
}

func (s *GetSecretValueResponseBody) SetCode(v string) *GetSecretValueResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecretValueResponseBody) SetDBInstanceId(v string) *GetSecretValueResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetDescription(v string) *GetSecretValueResponseBody {
	s.Description = &v
	return s
}

func (s *GetSecretValueResponseBody) SetMessage(v string) *GetSecretValueResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecretValueResponseBody) SetPassword(v string) *GetSecretValueResponseBody {
	s.Password = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRequestId(v string) *GetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretArn(v string) *GetSecretValueResponseBody {
	s.SecretArn = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretName(v string) *GetSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponseBody) SetStatus(v string) *GetSecretValueResponseBody {
	s.Status = &v
	return s
}

func (s *GetSecretValueResponseBody) SetUsername(v string) *GetSecretValueResponseBody {
	s.Username = &v
	return s
}

func (s *GetSecretValueResponseBody) Validate() error {
	return dara.Validate(s)
}
