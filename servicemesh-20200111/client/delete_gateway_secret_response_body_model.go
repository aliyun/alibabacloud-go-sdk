// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewaySecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGatewaySecretResponseBody
	GetRequestId() *string
	SetSecretDeleteRecord(v map[string]*SecretDeleteRecordValue) *DeleteGatewaySecretResponseBody
	GetSecretDeleteRecord() map[string]*SecretDeleteRecordValue
}

type DeleteGatewaySecretResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The records of deleting the secret in all clusters.
	SecretDeleteRecord map[string]*SecretDeleteRecordValue `json:"SecretDeleteRecord,omitempty" xml:"SecretDeleteRecord,omitempty"`
}

func (s DeleteGatewaySecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewaySecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewaySecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewaySecretResponseBody) GetSecretDeleteRecord() map[string]*SecretDeleteRecordValue {
	return s.SecretDeleteRecord
}

func (s *DeleteGatewaySecretResponseBody) SetRequestId(v string) *DeleteGatewaySecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewaySecretResponseBody) SetSecretDeleteRecord(v map[string]*SecretDeleteRecordValue) *DeleteGatewaySecretResponseBody {
	s.SecretDeleteRecord = v
	return s
}

func (s *DeleteGatewaySecretResponseBody) Validate() error {
	return dara.Validate(s)
}
