// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneMessageQrdlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListPhoneMessageQrdlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListPhoneMessageQrdlResponseBody
	GetCode() *string
	SetData(v []*ListPhoneMessageQrdlResponseBodyData) *ListPhoneMessageQrdlResponseBody
	GetData() []*ListPhoneMessageQrdlResponseBodyData
	SetMessage(v string) *ListPhoneMessageQrdlResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPhoneMessageQrdlResponseBody
	GetRequestId() *string
}

type ListPhoneMessageQrdlResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListPhoneMessageQrdlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPhoneMessageQrdlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneMessageQrdlResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListPhoneMessageQrdlResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPhoneMessageQrdlResponseBody) GetData() []*ListPhoneMessageQrdlResponseBodyData {
	return s.Data
}

func (s *ListPhoneMessageQrdlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPhoneMessageQrdlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPhoneMessageQrdlResponseBody) SetAccessDeniedDetail(v string) *ListPhoneMessageQrdlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetCode(v string) *ListPhoneMessageQrdlResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetData(v []*ListPhoneMessageQrdlResponseBodyData) *ListPhoneMessageQrdlResponseBody {
	s.Data = v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetMessage(v string) *ListPhoneMessageQrdlResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) SetRequestId(v string) *ListPhoneMessageQrdlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPhoneMessageQrdlResponseBodyData struct {
	// The URL of the deep link.
	//
	// example:
	//
	// https://wa.msg/
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
	// https://img.png
	QrImageUrl *string `json:"QrImageUrl,omitempty" xml:"QrImageUrl,omitempty"`
	// The mode of the quick-response (QR) code.
	//
	// example:
	//
	// IUIED999
	QrdlCode *string `json:"QrdlCode,omitempty" xml:"QrdlCode,omitempty"`
}

func (s ListPhoneMessageQrdlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneMessageQrdlResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetDeepLinkUrl() *string {
	return s.DeepLinkUrl
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetGenerateQrImage() *string {
	return s.GenerateQrImage
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetPrefilledMessage() *string {
	return s.PrefilledMessage
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetQrImageUrl() *string {
	return s.QrImageUrl
}

func (s *ListPhoneMessageQrdlResponseBodyData) GetQrdlCode() *string {
	return s.QrdlCode
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetDeepLinkUrl(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.DeepLinkUrl = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetGenerateQrImage(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.GenerateQrImage = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetPhoneNumber(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetPrefilledMessage(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.PrefilledMessage = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetQrImageUrl(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.QrImageUrl = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) SetQrdlCode(v string) *ListPhoneMessageQrdlResponseBodyData {
	s.QrdlCode = &v
	return s
}

func (s *ListPhoneMessageQrdlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
