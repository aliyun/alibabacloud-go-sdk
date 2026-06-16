// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSizeChartDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SizeChartDetectResponseBody
	GetCode() *string
	SetData(v *SizeChartDetectResponseBodyData) *SizeChartDetectResponseBody
	GetData() *SizeChartDetectResponseBodyData
	SetMessage(v string) *SizeChartDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *SizeChartDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SizeChartDetectResponseBody
	GetSuccess() *bool
}

type SizeChartDetectResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The size chart detection result.
	Data *SizeChartDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
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
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SizeChartDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SizeChartDetectResponseBody) GoString() string {
	return s.String()
}

func (s *SizeChartDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *SizeChartDetectResponseBody) GetData() *SizeChartDetectResponseBodyData {
	return s.Data
}

func (s *SizeChartDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SizeChartDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SizeChartDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SizeChartDetectResponseBody) SetCode(v string) *SizeChartDetectResponseBody {
	s.Code = &v
	return s
}

func (s *SizeChartDetectResponseBody) SetData(v *SizeChartDetectResponseBodyData) *SizeChartDetectResponseBody {
	s.Data = v
	return s
}

func (s *SizeChartDetectResponseBody) SetMessage(v string) *SizeChartDetectResponseBody {
	s.Message = &v
	return s
}

func (s *SizeChartDetectResponseBody) SetRequestId(v string) *SizeChartDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SizeChartDetectResponseBody) SetSuccess(v bool) *SizeChartDetectResponseBody {
	s.Success = &v
	return s
}

func (s *SizeChartDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SizeChartDetectResponseBodyData struct {
	// Indicates whether the image is a size chart.
	//
	// example:
	//
	// false
	IsSizeChart *bool `json:"IsSizeChart,omitempty" xml:"IsSizeChart,omitempty"`
	// The usage information. The key is the usage name, and the value is the count.
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s SizeChartDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SizeChartDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *SizeChartDetectResponseBodyData) GetIsSizeChart() *bool {
	return s.IsSizeChart
}

func (s *SizeChartDetectResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *SizeChartDetectResponseBodyData) SetIsSizeChart(v bool) *SizeChartDetectResponseBodyData {
	s.IsSizeChart = &v
	return s
}

func (s *SizeChartDetectResponseBodyData) SetUsageMap(v map[string]*int64) *SizeChartDetectResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *SizeChartDetectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
