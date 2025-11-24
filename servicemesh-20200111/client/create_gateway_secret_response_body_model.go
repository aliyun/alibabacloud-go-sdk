// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewaySecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGatewaySecretResponseBody
	GetRequestId() *string
	SetSecretCreateRecord(v map[string]*SecretCreateRecordValue) *CreateGatewaySecretResponseBody
	GetSecretCreateRecord() map[string]*SecretCreateRecordValue
}

type CreateGatewaySecretResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The record of creating the secret.
	SecretCreateRecord map[string]*SecretCreateRecordValue `json:"SecretCreateRecord,omitempty" xml:"SecretCreateRecord,omitempty"`
}

func (s CreateGatewaySecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewaySecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGatewaySecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGatewaySecretResponseBody) GetSecretCreateRecord() map[string]*SecretCreateRecordValue {
	return s.SecretCreateRecord
}

func (s *CreateGatewaySecretResponseBody) SetRequestId(v string) *CreateGatewaySecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGatewaySecretResponseBody) SetSecretCreateRecord(v map[string]*SecretCreateRecordValue) *CreateGatewaySecretResponseBody {
	s.SecretCreateRecord = v
	return s
}

func (s *CreateGatewaySecretResponseBody) Validate() error {
	return dara.Validate(s)
}
