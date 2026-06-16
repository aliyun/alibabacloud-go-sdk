// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageCroppingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageCroppingResponseBody
	GetCode() *string
	SetData(v *ImageCroppingResponseBodyData) *ImageCroppingResponseBody
	GetData() *ImageCroppingResponseBodyData
	SetMessage(v string) *ImageCroppingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageCroppingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageCroppingResponseBody
	GetSuccess() *bool
}

type ImageCroppingResponseBody struct {
	// Response code
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Crop result
	Data *ImageCroppingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful: true indicates success, false indicates failure
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageCroppingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageCroppingResponseBody) GoString() string {
	return s.String()
}

func (s *ImageCroppingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageCroppingResponseBody) GetData() *ImageCroppingResponseBodyData {
	return s.Data
}

func (s *ImageCroppingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageCroppingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageCroppingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageCroppingResponseBody) SetCode(v string) *ImageCroppingResponseBody {
	s.Code = &v
	return s
}

func (s *ImageCroppingResponseBody) SetData(v *ImageCroppingResponseBodyData) *ImageCroppingResponseBody {
	s.Data = v
	return s
}

func (s *ImageCroppingResponseBody) SetMessage(v string) *ImageCroppingResponseBody {
	s.Message = &v
	return s
}

func (s *ImageCroppingResponseBody) SetRequestId(v string) *ImageCroppingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageCroppingResponseBody) SetSuccess(v bool) *ImageCroppingResponseBody {
	s.Success = &v
	return s
}

func (s *ImageCroppingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageCroppingResponseBodyData struct {
	// Image height
	//
	// example:
	//
	// 800
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// URL of the cropped image
	//
	// example:
	//
	// https://example.com/cropped.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Usage information
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
	// Image width
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageCroppingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageCroppingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageCroppingResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *ImageCroppingResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageCroppingResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageCroppingResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *ImageCroppingResponseBodyData) SetHeight(v int32) *ImageCroppingResponseBodyData {
	s.Height = &v
	return s
}

func (s *ImageCroppingResponseBodyData) SetImageUrl(v string) *ImageCroppingResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageCroppingResponseBodyData) SetUsageMap(v map[string]*int64) *ImageCroppingResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageCroppingResponseBodyData) SetWidth(v int32) *ImageCroppingResponseBodyData {
	s.Width = &v
	return s
}

func (s *ImageCroppingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
