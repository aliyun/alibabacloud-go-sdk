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
	// The error code. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetPhoneEncryptionPublicKeyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DAC72B08-3327-33EF-BEDC-8EC3E83A6575
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPhoneEncryptionPublicKeyResponseBodyData struct {
	// The public key.
	//
	// example:
	//
	// -----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAt+PMrYw4zUDEk+VeTrp0\\n8LZaoVpiVFErX7iuoDjUs4F9vkxMQuIABjcXw\\/swzTMEopLORQV28uqN\\/2\\/x9hjU\\****\\/Zwa\\/Vk5Svp4\\niVY4e22FsJCCWUEMvayO8Q+3UGq0eXXQ+8SlUWEMq1VaJ4pwCLsMnmgybA+VmJxi\\nkwIDAQAB\\n-----END PUBLIC KEY-----"
	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" xml:"EncryptionPublicKey,omitempty"`
	// The status of the public key. Valid values:
	//
	// - MISMATCH: The public key is invalid.
	//
	// - VALID: The public key is valid.
	//
	// example:
	//
	// VALID
	EncryptionPublicKeyStatus *string `json:"EncryptionPublicKeyStatus,omitempty" xml:"EncryptionPublicKeyStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 861526377****
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
