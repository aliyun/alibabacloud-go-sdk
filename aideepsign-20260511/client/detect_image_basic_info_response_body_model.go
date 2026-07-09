// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetectImageBasicInfoResponseBody
	GetCode() *string
	SetDpi(v string) *DetectImageBasicInfoResponseBody
	GetDpi() *string
	SetHttpStatusCode(v int32) *DetectImageBasicInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DetectImageBasicInfoResponseBody
	GetMessage() *string
	SetName(v string) *DetectImageBasicInfoResponseBody
	GetName() *string
	SetRequestId(v string) *DetectImageBasicInfoResponseBody
	GetRequestId() *string
	SetSize(v string) *DetectImageBasicInfoResponseBody
	GetSize() *string
	SetSuccess(v bool) *DetectImageBasicInfoResponseBody
	GetSuccess() *bool
	SetType(v string) *DetectImageBasicInfoResponseBody
	GetType() *string
}

type DetectImageBasicInfoResponseBody struct {
	// The business error code. "OK" is returned if the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The image resolution (width × height), such as 1920 × 1080. This value is empty if the resolution cannot be identified.
	//
	// example:
	//
	// 2048 	- 2048
	Dpi *string `json:"Dpi,omitempty" xml:"Dpi,omitempty"`
	// The HTTP status code. 200 is returned if the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. "success" is returned if the request was successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The file name.
	//
	// example:
	//
	// photo.jpg
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The file size in a human-readable format, such as 1.5 MB or 256 KB.
	//
	// example:
	//
	// 2.3 MB
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The image format, such as JPEG, PNG, GIF, or WEBP. UNKNOWN is returned if the format cannot be identified.
	//
	// example:
	//
	// PNG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectImageBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageBasicInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetectImageBasicInfoResponseBody) GetDpi() *string {
	return s.Dpi
}

func (s *DetectImageBasicInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DetectImageBasicInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectImageBasicInfoResponseBody) GetName() *string {
	return s.Name
}

func (s *DetectImageBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageBasicInfoResponseBody) GetSize() *string {
	return s.Size
}

func (s *DetectImageBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetectImageBasicInfoResponseBody) GetType() *string {
	return s.Type
}

func (s *DetectImageBasicInfoResponseBody) SetCode(v string) *DetectImageBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetDpi(v string) *DetectImageBasicInfoResponseBody {
	s.Dpi = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetHttpStatusCode(v int32) *DetectImageBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetMessage(v string) *DetectImageBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetName(v string) *DetectImageBasicInfoResponseBody {
	s.Name = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetRequestId(v string) *DetectImageBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetSize(v string) *DetectImageBasicInfoResponseBody {
	s.Size = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetSuccess(v bool) *DetectImageBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) SetType(v string) *DetectImageBasicInfoResponseBody {
	s.Type = &v
	return s
}

func (s *DetectImageBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
