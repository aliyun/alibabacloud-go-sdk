// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageMattingResponseBody
	GetCode() *string
	SetData(v *ImageMattingResponseBodyData) *ImageMattingResponseBody
	GetData() *ImageMattingResponseBodyData
	SetMessage(v string) *ImageMattingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageMattingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageMattingResponseBody
	GetSuccess() *bool
}

type ImageMattingResponseBody struct {
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The matting result.
	Data *ImageMattingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
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
	// Indicates whether the call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageMattingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingResponseBody) GoString() string {
	return s.String()
}

func (s *ImageMattingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageMattingResponseBody) GetData() *ImageMattingResponseBodyData {
	return s.Data
}

func (s *ImageMattingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageMattingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageMattingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageMattingResponseBody) SetCode(v string) *ImageMattingResponseBody {
	s.Code = &v
	return s
}

func (s *ImageMattingResponseBody) SetData(v *ImageMattingResponseBodyData) *ImageMattingResponseBody {
	s.Data = v
	return s
}

func (s *ImageMattingResponseBody) SetMessage(v string) *ImageMattingResponseBody {
	s.Message = &v
	return s
}

func (s *ImageMattingResponseBody) SetRequestId(v string) *ImageMattingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageMattingResponseBody) SetSuccess(v bool) *ImageMattingResponseBody {
	s.Success = &v
	return s
}

func (s *ImageMattingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageMattingResponseBodyData struct {
	// The height of the result image in pixels.
	//
	// example:
	//
	// 800
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The URL of the matting result image.
	//
	// example:
	//
	// http://dashscope-7c2c.oss-cn-shanghai.aliyuncs.com/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The quality score of the matting result.
	//
	// example:
	//
	// 0.11
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The usage information.
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
	// The width of the result image in pixels.
	//
	// example:
	//
	// 800
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageMattingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageMattingResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *ImageMattingResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageMattingResponseBodyData) GetScore() *float64 {
	return s.Score
}

func (s *ImageMattingResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageMattingResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *ImageMattingResponseBodyData) SetHeight(v int32) *ImageMattingResponseBodyData {
	s.Height = &v
	return s
}

func (s *ImageMattingResponseBodyData) SetImageUrl(v string) *ImageMattingResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageMattingResponseBodyData) SetScore(v float64) *ImageMattingResponseBodyData {
	s.Score = &v
	return s
}

func (s *ImageMattingResponseBodyData) SetUsageMap(v map[string]*int64) *ImageMattingResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageMattingResponseBodyData) SetWidth(v int32) *ImageMattingResponseBodyData {
	s.Width = &v
	return s
}

func (s *ImageMattingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
