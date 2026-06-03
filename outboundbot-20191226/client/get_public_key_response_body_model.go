// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPublicKeyResponseBody
	GetCode() *string
	SetData(v string) *GetPublicKeyResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *GetPublicKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPublicKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPublicKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPublicKeyResponseBody
	GetSuccess() *bool
}

type GetPublicKeyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwkftR4j5P9ng+Si/2ydc\\\\/K03NlhpzI4nW3JoNZIZR83P\\nMeyoULt+ivvFI7R++BU413QfX7l5FZnuUrII\\nNNBfFX84m1tmsdythDQmS2soG2sBiGKMv6O5mlBvXi+GA0/GqQ2juEv5DAb0GfOk\\nw8syQDkpNZflUSTnh10qbnDQxIGeisv1S4/Eo00djX48y5N8qXEcz9CUgwQpKQ0s\\nWQIDAQAB\\n-----END PUBLIC KEY-----\\n
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPublicKeyResponseBody) GetData() *string {
	return s.Data
}

func (s *GetPublicKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPublicKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPublicKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPublicKeyResponseBody) SetCode(v string) *GetPublicKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetData(v string) *GetPublicKeyResponseBody {
	s.Data = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetHttpStatusCode(v int32) *GetPublicKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetMessage(v string) *GetPublicKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetRequestId(v string) *GetPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicKeyResponseBody) SetSuccess(v bool) *GetPublicKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
