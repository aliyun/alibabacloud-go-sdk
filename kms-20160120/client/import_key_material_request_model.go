// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyMaterialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptedKeyMaterial(v string) *ImportKeyMaterialRequest
	GetEncryptedKeyMaterial() *string
	SetImportToken(v string) *ImportKeyMaterialRequest
	GetImportToken() *string
	SetKeyId(v string) *ImportKeyMaterialRequest
	GetKeyId() *string
	SetKeyMaterialExpireUnix(v int64) *ImportKeyMaterialRequest
	GetKeyMaterialExpireUnix() *int64
}

type ImportKeyMaterialRequest struct {
	// Use **GetParametersForImport*	- the Returned public key and the base64-encoded key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// bCPZx7I6v6KXsqEpr2OXKxuj2CCRtKdwp75Bw+BGncYqBdfjFBYRtOE6HRlT0oeiRDWzwnw9OA54OL36smDJrq4Lo9x0CyYDiuKnRkcKtMtlzW0din7Pd7IlZWWRdVueiw2qpzl7PkUWQGTdsdbzpfJJQ+qj/cRIrk/E83UGyeyytSpgnb+lu0xEYcPajRyWNsbi98N3pqqQzHXNNHO2NJqHlnQgglqTiBEjkGeKFhfKmTc3vjulIdVa3EaVIN6lwWfgx+UUYSrvbA77WDYKlDsZ4SbK2/T7za9Tp1qU7Ynqba7OKGVVj7PMbiaO80AxWZnjUMYCgEp5w7V+seOXqw==
	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" xml:"EncryptedKeyMaterial,omitempty"`
	// By calling **GetParametersForImport*	- the import token.
	//
	// This parameter is required.
	//
	// example:
	//
	// Base64String
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The ID of the CMK to be imported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234abcd-12ab-34cd-56ef-12345678****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The time when the key material expires.
	//
	// If this parameter is not specified or set this parameter to 0, the key material does not expire.
	//
	// >  The value cannot be earlier than the time when the API is called (based on the server time).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	KeyMaterialExpireUnix *int64 `json:"KeyMaterialExpireUnix,omitempty" xml:"KeyMaterialExpireUnix,omitempty"`
}

func (s ImportKeyMaterialRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialRequest) GetEncryptedKeyMaterial() *string {
	return s.EncryptedKeyMaterial
}

func (s *ImportKeyMaterialRequest) GetImportToken() *string {
	return s.ImportToken
}

func (s *ImportKeyMaterialRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *ImportKeyMaterialRequest) GetKeyMaterialExpireUnix() *int64 {
	return s.KeyMaterialExpireUnix
}

func (s *ImportKeyMaterialRequest) SetEncryptedKeyMaterial(v string) *ImportKeyMaterialRequest {
	s.EncryptedKeyMaterial = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetImportToken(v string) *ImportKeyMaterialRequest {
	s.ImportToken = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyId(v string) *ImportKeyMaterialRequest {
	s.KeyId = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyMaterialExpireUnix(v int64) *ImportKeyMaterialRequest {
	s.KeyMaterialExpireUnix = &v
	return s
}

func (s *ImportKeyMaterialRequest) Validate() error {
	return dara.Validate(s)
}
