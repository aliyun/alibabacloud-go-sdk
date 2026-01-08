// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneMessageQrdlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdatePhoneMessageQrdlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdatePhoneMessageQrdlResponseBody
	GetCode() *string
	SetData(v *UpdatePhoneMessageQrdlResponseBodyData) *UpdatePhoneMessageQrdlResponseBody
	GetData() *UpdatePhoneMessageQrdlResponseBodyData
	SetMessage(v string) *UpdatePhoneMessageQrdlResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePhoneMessageQrdlResponseBody
	GetRequestId() *string
}

type UpdatePhoneMessageQrdlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The result returns OK as normal.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UpdatePhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1612C226-E271-4CFE-9F18-4066D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdatePhoneMessageQrdlResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePhoneMessageQrdlResponseBody) GetData() *UpdatePhoneMessageQrdlResponseBodyData {
	return s.Data
}

func (s *UpdatePhoneMessageQrdlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePhoneMessageQrdlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetAccessDeniedDetail(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetCode(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetData(v *UpdatePhoneMessageQrdlResponseBodyData) *UpdatePhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetMessage(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) SetRequestId(v string) *UpdatePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePhoneMessageQrdlResponseBodyData struct {
	// Deep link address.
	//
	// example:
	//
	// https://wa.msg/
	DeepLinkUrl *string `json:"DeepLinkUrl,omitempty" xml:"DeepLinkUrl,omitempty"`
	// Generate image types.
	//
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// Number.
	//
	// example:
	//
	// 8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Message content.
	//
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// QR code address.
	//
	// example:
	//
	// https://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// QR code encoding.
	//
	// example:
	//
	// DEDEE998
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s UpdatePhoneMessageQrdlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetDeepLinkUrl() *string {
	return s.DeepLinkUrl
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetGenerateQrImage() *string {
	return s.GenerateQrImage
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetPrefilledMessage() *string {
	return s.PrefilledMessage
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetQrImageUrl() *string {
	return s.QrImageUrl
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *UpdatePhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

func (s *UpdatePhoneMessageQrdlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
