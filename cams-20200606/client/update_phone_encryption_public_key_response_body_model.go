// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneEncryptionPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdatePhoneEncryptionPublicKeyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdatePhoneEncryptionPublicKeyResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePhoneEncryptionPublicKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePhoneEncryptionPublicKeyResponseBody
	GetRequestId() *string
}

type UpdatePhoneEncryptionPublicKeyResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneEncryptionPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneEncryptionPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetAccessDeniedDetail(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetCode(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetMessage(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) SetRequestId(v string) *UpdatePhoneEncryptionPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePhoneEncryptionPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
