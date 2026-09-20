// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTryOnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AiTryOnResponseBody
	GetCode() *string
	SetData(v *AiTryOnResponseBodyData) *AiTryOnResponseBody
	GetData() *AiTryOnResponseBodyData
	SetMessage(v string) *AiTryOnResponseBody
	GetMessage() *string
	SetRequestId(v string) *AiTryOnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AiTryOnResponseBody
	GetSuccess() *bool
}

type AiTryOnResponseBody struct {
	// example:
	//
	// 200
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AiTryOnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2728332e-72c1-9c0d-8869-5781b2cd25d4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AiTryOnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnResponseBody) GoString() string {
	return s.String()
}

func (s *AiTryOnResponseBody) GetCode() *string {
	return s.Code
}

func (s *AiTryOnResponseBody) GetData() *AiTryOnResponseBodyData {
	return s.Data
}

func (s *AiTryOnResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AiTryOnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AiTryOnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AiTryOnResponseBody) SetCode(v string) *AiTryOnResponseBody {
	s.Code = &v
	return s
}

func (s *AiTryOnResponseBody) SetData(v *AiTryOnResponseBodyData) *AiTryOnResponseBody {
	s.Data = v
	return s
}

func (s *AiTryOnResponseBody) SetMessage(v string) *AiTryOnResponseBody {
	s.Message = &v
	return s
}

func (s *AiTryOnResponseBody) SetRequestId(v string) *AiTryOnResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiTryOnResponseBody) SetSuccess(v bool) *AiTryOnResponseBody {
	s.Success = &v
	return s
}

func (s *AiTryOnResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiTryOnResponseBodyData struct {
	// example:
	//
	// 1360
	ImageHeight *string `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// example:
	//
	// https://example.com/virtual-try-on-result.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 768
	ImageWidth *string `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// example:
	//
	// {"ProcessedImageCount":1,"Resolution":"1K"}
	UsageMap *AiTryOnResponseBodyDataUsageMap `json:"UsageMap,omitempty" xml:"UsageMap,omitempty" type:"Struct"`
}

func (s AiTryOnResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnResponseBodyData) GoString() string {
	return s.String()
}

func (s *AiTryOnResponseBodyData) GetImageHeight() *string {
	return s.ImageHeight
}

func (s *AiTryOnResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *AiTryOnResponseBodyData) GetImageWidth() *string {
	return s.ImageWidth
}

func (s *AiTryOnResponseBodyData) GetUsageMap() *AiTryOnResponseBodyDataUsageMap {
	return s.UsageMap
}

func (s *AiTryOnResponseBodyData) SetImageHeight(v string) *AiTryOnResponseBodyData {
	s.ImageHeight = &v
	return s
}

func (s *AiTryOnResponseBodyData) SetImageUrl(v string) *AiTryOnResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *AiTryOnResponseBodyData) SetImageWidth(v string) *AiTryOnResponseBodyData {
	s.ImageWidth = &v
	return s
}

func (s *AiTryOnResponseBodyData) SetUsageMap(v *AiTryOnResponseBodyDataUsageMap) *AiTryOnResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *AiTryOnResponseBodyData) Validate() error {
	if s.UsageMap != nil {
		if err := s.UsageMap.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiTryOnResponseBodyDataUsageMap struct {
	// example:
	//
	// 1
	ProcessedImageCount *int64 `json:"ProcessedImageCount,omitempty" xml:"ProcessedImageCount,omitempty"`
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s AiTryOnResponseBodyDataUsageMap) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnResponseBodyDataUsageMap) GoString() string {
	return s.String()
}

func (s *AiTryOnResponseBodyDataUsageMap) GetProcessedImageCount() *int64 {
	return s.ProcessedImageCount
}

func (s *AiTryOnResponseBodyDataUsageMap) GetResolution() *string {
	return s.Resolution
}

func (s *AiTryOnResponseBodyDataUsageMap) SetProcessedImageCount(v int64) *AiTryOnResponseBodyDataUsageMap {
	s.ProcessedImageCount = &v
	return s
}

func (s *AiTryOnResponseBodyDataUsageMap) SetResolution(v string) *AiTryOnResponseBodyDataUsageMap {
	s.Resolution = &v
	return s
}

func (s *AiTryOnResponseBodyDataUsageMap) Validate() error {
	return dara.Validate(s)
}
