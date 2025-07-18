// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSecretResponseBody
	GetRequestId() *string
	SetSecretArn(v string) *DeleteSecretResponseBody
	GetSecretArn() *string
	SetStatus(v string) *DeleteSecretResponseBody
	GetStatus() *string
}

type DeleteSecretResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// The status of the operation. Valid values:
	//
	// 	- **fail**
	//
	// 	- **success**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretResponseBody) GetSecretArn() *string {
	return s.SecretArn
}

func (s *DeleteSecretResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteSecretResponseBody) SetMessage(v string) *DeleteSecretResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSecretArn(v string) *DeleteSecretResponseBody {
	s.SecretArn = &v
	return s
}

func (s *DeleteSecretResponseBody) SetStatus(v string) *DeleteSecretResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
