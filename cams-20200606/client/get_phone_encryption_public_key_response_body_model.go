// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneEncryptionPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetPhoneEncryptionPublicKeyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetPhoneEncryptionPublicKeyResponseBody
	GetCode() *string
	SetData(v *GetPhoneEncryptionPublicKeyResponseBodyData) *GetPhoneEncryptionPublicKeyResponseBody
	GetData() *GetPhoneEncryptionPublicKeyResponseBodyData
	SetMessage(v string) *GetPhoneEncryptionPublicKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPhoneEncryptionPublicKeyResponseBody
	GetRequestId() *string
}

type GetPhoneEncryptionPublicKeyResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPhoneEncryptionPublicKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error description information.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) GetData() *GetPhoneEncryptionPublicKeyResponseBodyData {
	return s.Data
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetAccessDeniedDetail(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetCode(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetData(v *GetPhoneEncryptionPublicKeyResponseBodyData) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetMessage(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) SetRequestId(v string) *GetPhoneEncryptionPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPhoneEncryptionPublicKeyResponseBodyData struct {
	// The public key.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----
	//
	// AAA
	//
	// BBB
	//
	// CCC
	//
	// DDD
	//
	// EEE
	//
	// FFF
	//
	// GGG
	//
	// -----END PUBLIC KEY-----
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	// The validity state of the public key. Valid values:
	//
	// 	- MISMATCH: The public key is invalid.
	//
	// 	- VALID: The public key is valid.
	//
	// example:
	//
	// VALID
	EncryptionPublicKeyStatus *string `json:"EncryptionPublicKeyStatus,omitempty" xml:"EncryptionPublicKeyStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 86138000**
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s GetPhoneEncryptionPublicKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneEncryptionPublicKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) GetEncryptionPublicKey() *string {
	return s.EncryptionPublicKey
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) GetEncryptionPublicKeyStatus() *string {
	return s.EncryptionPublicKeyStatus
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetEncryptionPublicKey(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.EncryptionPublicKey = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetEncryptionPublicKeyStatus(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.EncryptionPublicKeyStatus = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) SetPhoneNumber(v string) *GetPhoneEncryptionPublicKeyResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetPhoneEncryptionPublicKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
