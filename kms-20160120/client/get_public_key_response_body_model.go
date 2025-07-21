// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GetPublicKeyResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *GetPublicKeyResponseBody
	GetKeyVersionId() *string
	SetPublicKey(v string) *GetPublicKeyResponseBody
	GetPublicKey() *string
	SetRequestId(v string) *GetPublicKeyResponseBody
	GetRequestId() *string
}

type GetPublicKeyResponseBody struct {
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to the alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The public key returned in the PEM format.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs5Yu9AEgATN2/e3nUz1K\\nEy6ng8MSPutcse2/VECG/NUF9C6D4IsJ64ShzY3dcn34WYzTOe916eMJFxyrNrSw\\nHtc4UOR5AvaoRrfpgu2uq+i70/ZXrWL+pGb1hgZV8cWheIHMxwrR3IiQlM5qN7EF\\n9BdyWtyBfUGsp0Bn1VqlPc5G0x0a9xU2z9YtP994yDenNVIoIQ6Cov1lIEuwXAb2\\n7boC41ePXwD0JWt41sP+rgCmpjBx00puIG+IlnoReEgI1ZGYmK98GgA/XzmNjZiD\\nyvXJZAcM33Ue85+PkR5iHTtSEbi4QAoqpJabprUzz3Fin2j1dRrcacxGb7p31A9c\\nJQIDAQAB\\n-----END PUBLIC KEY-----\\n
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GetPublicKeyResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *GetPublicKeyResponseBody) GetPublicKey() *string {
	return s.PublicKey
}

func (s *GetPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublicKeyResponseBody) SetKeyId(v string) *GetPublicKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetKeyVersionId(v string) *GetPublicKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetPublicKey(v string) *GetPublicKeyResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetRequestId(v string) *GetPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
