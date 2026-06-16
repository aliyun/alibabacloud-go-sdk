// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageRemoveResponseBody
	GetCode() *string
	SetData(v *ImageRemoveResponseBodyData) *ImageRemoveResponseBody
	GetData() *ImageRemoveResponseBodyData
	SetMessage(v string) *ImageRemoveResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageRemoveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageRemoveResponseBody
	GetSuccess() *bool
}

type ImageRemoveResponseBody struct {
	// Error code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Intelligent removal result
	Data *ImageRemoveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2728332e-72c1-9c0d-8869-5781b2cd25d4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageRemoveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveResponseBody) GoString() string {
	return s.String()
}

func (s *ImageRemoveResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageRemoveResponseBody) GetData() *ImageRemoveResponseBodyData {
	return s.Data
}

func (s *ImageRemoveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageRemoveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageRemoveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageRemoveResponseBody) SetCode(v string) *ImageRemoveResponseBody {
	s.Code = &v
	return s
}

func (s *ImageRemoveResponseBody) SetData(v *ImageRemoveResponseBodyData) *ImageRemoveResponseBody {
	s.Data = v
	return s
}

func (s *ImageRemoveResponseBody) SetMessage(v string) *ImageRemoveResponseBody {
	s.Message = &v
	return s
}

func (s *ImageRemoveResponseBody) SetRequestId(v string) *ImageRemoveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageRemoveResponseBody) SetSuccess(v bool) *ImageRemoveResponseBody {
	s.Success = &v
	return s
}

func (s *ImageRemoveResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageRemoveResponseBodyData struct {
	// Height of the result image (pixels)
	//
	// example:
	//
	// 800
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// URL of the removal result image
	//
	// example:
	//
	// http://dashscope-7c2c.oss-cn-shanghai.aliyuncs.com/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Usage information
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
	// Width of the result image (pixels)
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageRemoveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageRemoveResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *ImageRemoveResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemoveResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageRemoveResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *ImageRemoveResponseBodyData) SetHeight(v int32) *ImageRemoveResponseBodyData {
	s.Height = &v
	return s
}

func (s *ImageRemoveResponseBodyData) SetImageUrl(v string) *ImageRemoveResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemoveResponseBodyData) SetUsageMap(v map[string]*int64) *ImageRemoveResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageRemoveResponseBodyData) SetWidth(v int32) *ImageRemoveResponseBodyData {
	s.Width = &v
	return s
}

func (s *ImageRemoveResponseBodyData) Validate() error {
	return dara.Validate(s)
}
