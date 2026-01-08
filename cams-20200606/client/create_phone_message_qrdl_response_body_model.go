// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhoneMessageQrdlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreatePhoneMessageQrdlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreatePhoneMessageQrdlResponseBody
	GetCode() *string
	SetData(v *CreatePhoneMessageQrdlResponseBodyData) *CreatePhoneMessageQrdlResponseBody
	GetData() *CreatePhoneMessageQrdlResponseBodyData
	SetMessage(v string) *CreatePhoneMessageQrdlResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePhoneMessageQrdlResponseBody
	GetRequestId() *string
}

type CreatePhoneMessageQrdlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreatePhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// none
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePhoneMessageQrdlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreatePhoneMessageQrdlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePhoneMessageQrdlResponseBody) GetData() *CreatePhoneMessageQrdlResponseBodyData {
	return s.Data
}

func (s *CreatePhoneMessageQrdlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePhoneMessageQrdlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePhoneMessageQrdlResponseBody) SetAccessDeniedDetail(v string) *CreatePhoneMessageQrdlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetCode(v string) *CreatePhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetData(v *CreatePhoneMessageQrdlResponseBodyData) *CreatePhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetMessage(v string) *CreatePhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) SetRequestId(v string) *CreatePhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePhoneMessageQrdlResponseBodyData struct {
	// The URL of the deep link.
	//
	// example:
	//
	// https://wa.qrdl/
	DeepLinkUrl *string `json:"DeepLinkUrl,omitempty" xml:"DeepLinkUrl,omitempty"`
	// The format of the generated image.
	//
	// example:
	//
	// PNG
	GenerateQrImage *string `json:"GenerateQrImage,omitempty" xml:"GenerateQrImage,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The message content.
	//
	// example:
	//
	// Hello
	PrefilledMessage *string `json:"PrefilledMessage,omitempty" xml:"PrefilledMessage,omitempty"`
	// The URL of the QR code.
	//
	// example:
	//
	// http://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// The mode of the quick-response (QR) code.
	//
	// example:
	//
	// D9II3***
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s CreatePhoneMessageQrdlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetDeepLinkUrl() *string {
	return s.DeepLinkUrl
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetGenerateQrImage() *string {
	return s.GenerateQrImage
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetPrefilledMessage() *string {
	return s.PrefilledMessage
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetQrImageUrl() *string {
	return s.QrImageUrl
}

func (s *CreatePhoneMessageQrdlResponseBodyData) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *CreatePhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

func (s *CreatePhoneMessageQrdlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
