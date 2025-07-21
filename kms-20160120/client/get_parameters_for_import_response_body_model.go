// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersForImportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImportToken(v string) *GetParametersForImportResponseBody
	GetImportToken() *string
	SetKeyId(v string) *GetParametersForImportResponseBody
	GetKeyId() *string
	SetPublicKey(v string) *GetParametersForImportResponseBody
	GetPublicKey() *string
	SetRequestId(v string) *GetParametersForImportResponseBody
	GetRequestId() *string
	SetTokenExpireTime(v string) *GetParametersForImportResponseBody
	GetTokenExpireTime() *string
}

type GetParametersForImportResponseBody struct {
	// The token that is used to import key material.
	//
	// The token is valid for 24 hours. The value of this parameter is required when you call the [ImportKeyMaterial](https://help.aliyun.com/document_detail/68622.html) operation.
	//
	// example:
	//
	// Base64String
	ImportToken *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty"`
	// The globally unique ID of the CMK.
	//
	// The value of this parameter is required when you call the [ImportKeyMaterial](https://help.aliyun.com/document_detail/68622.html) operation.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The public key that is used to encrypt key material.
	//
	// The public key is Base64-encoded.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlls4uIBxD0GG84C+lGBO6Dhpf1J3XimC6cPmPNaKKJMOzoX4tD+C+r7aZv8lZ3vnPfxuxvy/YwG+whUxTEEFUdqJTOIzhPfYucupqKM92crVHIuG+xtMVeHKjyTr+UrtKCsQikqHT+19yDRN/RMoo2HUx0gmEnRyXd8t3JyUXun9FdoxKA08GrsV7nodb9ZsoBLhnev7tTLcXvLyKW6XG1ZQCQm6dPnbnwLeDXR7uK0Lqn9PM28mBIdaiQUQxj2XbM1CoJA+JiyVX3Ptdb+4rqukb4Rb05B80Bs9xV/cf7FIku08l7xGhrGiQFq+DFXwQWtwihXHZxz3LhldU+4ZPwID****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8cdf51fd-bcd6-d79a-0ef4-e52c9b5466dc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the token expires.
	//
	// example:
	//
	// 2018-01-25T00:01:02Z
	TokenExpireTime *string `json:"TokenExpireTime,omitempty" xml:"TokenExpireTime,omitempty"`
}

func (s GetParametersForImportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetParametersForImportResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponseBody) GetImportToken() *string {
	return s.ImportToken
}

func (s *GetParametersForImportResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GetParametersForImportResponseBody) GetPublicKey() *string {
	return s.PublicKey
}

func (s *GetParametersForImportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetParametersForImportResponseBody) GetTokenExpireTime() *string {
	return s.TokenExpireTime
}

func (s *GetParametersForImportResponseBody) SetImportToken(v string) *GetParametersForImportResponseBody {
	s.ImportToken = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetKeyId(v string) *GetParametersForImportResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetPublicKey(v string) *GetParametersForImportResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetRequestId(v string) *GetParametersForImportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersForImportResponseBody) SetTokenExpireTime(v string) *GetParametersForImportResponseBody {
	s.TokenExpireTime = &v
	return s
}

func (s *GetParametersForImportResponseBody) Validate() error {
	return dara.Validate(s)
}
