// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageUpscaleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageUpscaleResponseBody
	GetCode() *string
	SetData(v *ImageUpscaleResponseBodyData) *ImageUpscaleResponseBody
	GetData() *ImageUpscaleResponseBodyData
	SetMessage(v string) *ImageUpscaleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageUpscaleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageUpscaleResponseBody
	GetSuccess() *bool
}

type ImageUpscaleResponseBody struct {
	// The status code, such as success.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result struct information.
	Data *ImageUpscaleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message or error prompt.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageUpscaleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleResponseBody) GoString() string {
	return s.String()
}

func (s *ImageUpscaleResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageUpscaleResponseBody) GetData() *ImageUpscaleResponseBodyData {
	return s.Data
}

func (s *ImageUpscaleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageUpscaleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageUpscaleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageUpscaleResponseBody) SetCode(v string) *ImageUpscaleResponseBody {
	s.Code = &v
	return s
}

func (s *ImageUpscaleResponseBody) SetData(v *ImageUpscaleResponseBodyData) *ImageUpscaleResponseBody {
	s.Data = v
	return s
}

func (s *ImageUpscaleResponseBody) SetMessage(v string) *ImageUpscaleResponseBody {
	s.Message = &v
	return s
}

func (s *ImageUpscaleResponseBody) SetRequestId(v string) *ImageUpscaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageUpscaleResponseBody) SetSuccess(v bool) *ImageUpscaleResponseBody {
	s.Success = &v
	return s
}

func (s *ImageUpscaleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageUpscaleResponseBodyData struct {
	// The high-definition upscaling result data.
	Data *ImageUpscaleResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Returned upon failure, such as InvalidParameter or DownstreamUnavailable.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. Returned upon failure.
	//
	// example:
	//
	// Success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the operation is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The downstream task ID. This parameter is empty for synchronous operations.
	//
	// example:
	//
	// ed9d8504-6141-9fbb-8345-4fa36433483f
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information.
	UsageMap *ImageUpscaleResponseBodyDataUsageMap `json:"UsageMap,omitempty" xml:"UsageMap,omitempty" type:"Struct"`
}

func (s ImageUpscaleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageUpscaleResponseBodyData) GetData() *ImageUpscaleResponseBodyDataData {
	return s.Data
}

func (s *ImageUpscaleResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ImageUpscaleResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ImageUpscaleResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ImageUpscaleResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageUpscaleResponseBodyData) GetUsageMap() *ImageUpscaleResponseBodyDataUsageMap {
	return s.UsageMap
}

func (s *ImageUpscaleResponseBodyData) SetData(v *ImageUpscaleResponseBodyDataData) *ImageUpscaleResponseBodyData {
	s.Data = v
	return s
}

func (s *ImageUpscaleResponseBodyData) SetErrorCode(v string) *ImageUpscaleResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *ImageUpscaleResponseBodyData) SetErrorMessage(v string) *ImageUpscaleResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *ImageUpscaleResponseBodyData) SetSuccess(v bool) *ImageUpscaleResponseBodyData {
	s.Success = &v
	return s
}

func (s *ImageUpscaleResponseBodyData) SetTaskId(v string) *ImageUpscaleResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageUpscaleResponseBodyData) SetUsageMap(v *ImageUpscaleResponseBodyDataUsageMap) *ImageUpscaleResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageUpscaleResponseBodyData) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.UsageMap != nil {
		if err := s.UsageMap.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageUpscaleResponseBodyDataData struct {
	// The actual height of the generated image.
	//
	// example:
	//
	// 4096
	ImageHeight *int32 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// The result information of the high-definition upscaling for a single image.
	//
	// example:
	//
	// https://example.com/output.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The actual width of the generated image.
	//
	// example:
	//
	// 2048
	ImageWidth *int32 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
}

func (s ImageUpscaleResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ImageUpscaleResponseBodyDataData) GetImageHeight() *int32 {
	return s.ImageHeight
}

func (s *ImageUpscaleResponseBodyDataData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageUpscaleResponseBodyDataData) GetImageWidth() *int32 {
	return s.ImageWidth
}

func (s *ImageUpscaleResponseBodyDataData) SetImageHeight(v int32) *ImageUpscaleResponseBodyDataData {
	s.ImageHeight = &v
	return s
}

func (s *ImageUpscaleResponseBodyDataData) SetImageUrl(v string) *ImageUpscaleResponseBodyDataData {
	s.ImageUrl = &v
	return s
}

func (s *ImageUpscaleResponseBodyDataData) SetImageWidth(v int32) *ImageUpscaleResponseBodyDataData {
	s.ImageWidth = &v
	return s
}

func (s *ImageUpscaleResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type ImageUpscaleResponseBodyDataUsageMap struct {
	// The number of processed images.
	//
	// example:
	//
	// 1
	ProcessedImageCount *int32 `json:"ProcessedImageCount,omitempty" xml:"ProcessedImageCount,omitempty"`
}

func (s ImageUpscaleResponseBodyDataUsageMap) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleResponseBodyDataUsageMap) GoString() string {
	return s.String()
}

func (s *ImageUpscaleResponseBodyDataUsageMap) GetProcessedImageCount() *int32 {
	return s.ProcessedImageCount
}

func (s *ImageUpscaleResponseBodyDataUsageMap) SetProcessedImageCount(v int32) *ImageUpscaleResponseBodyDataUsageMap {
	s.ProcessedImageCount = &v
	return s
}

func (s *ImageUpscaleResponseBodyDataUsageMap) Validate() error {
	return dara.Validate(s)
}
