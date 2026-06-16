// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationStandardResponseBody
	GetCode() *string
	SetData(v *ImageTranslationStandardResponseBodyData) *ImageTranslationStandardResponseBody
	GetData() *ImageTranslationStandardResponseBodyData
	SetMessage(v string) *ImageTranslationStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationStandardResponseBody
	GetSuccess() *bool
}

type ImageTranslationStandardResponseBody struct {
	// The response code. A value of 200 indicates a successful call. For other response codes, see the error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The translation result data, including the translated image URL and usage information.
	Data *ImageTranslationStandardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. "Success" is returned for a successful call. A specific error message is returned for a failed call.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 1CEC4D94-905A-1ED1-A7B4-1BFEFFB3D850
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationStandardResponseBody) GetData() *ImageTranslationStandardResponseBodyData {
	return s.Data
}

func (s *ImageTranslationStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationStandardResponseBody) SetCode(v string) *ImageTranslationStandardResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetData(v *ImageTranslationStandardResponseBodyData) *ImageTranslationStandardResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetMessage(v string) *ImageTranslationStandardResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetRequestId(v string) *ImageTranslationStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetSuccess(v bool) *ImageTranslationStandardResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationStandardResponseBodyData struct {
	// The URL of the image generated after image translation.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The usage information, including the number of processed images.
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s ImageTranslationStandardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationStandardResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageTranslationStandardResponseBodyData) SetImageUrl(v string) *ImageTranslationStandardResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyData) SetUsageMap(v map[string]*int64) *ImageTranslationStandardResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageTranslationStandardResponseBodyData) Validate() error {
	return dara.Validate(s)
}
