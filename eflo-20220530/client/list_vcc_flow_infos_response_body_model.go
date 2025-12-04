// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVccFlowInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListVccFlowInfosResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListVccFlowInfosResponseBody
	GetCode() *int32
	SetContent(v *ListVccFlowInfosResponseBodyContent) *ListVccFlowInfosResponseBody
	GetContent() *ListVccFlowInfosResponseBodyContent
	SetMessage(v string) *ListVccFlowInfosResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVccFlowInfosResponseBody
	GetRequestId() *string
}

type ListVccFlowInfosResponseBody struct {
	// 访问被拒绝的详细原因。
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *ListVccFlowInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// Response
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// BDBCC783-84CA-5733-8EEA-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVccFlowInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVccFlowInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListVccFlowInfosResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListVccFlowInfosResponseBody) GetContent() *ListVccFlowInfosResponseBodyContent {
	return s.Content
}

func (s *ListVccFlowInfosResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVccFlowInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVccFlowInfosResponseBody) SetAccessDeniedDetail(v string) *ListVccFlowInfosResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetCode(v int32) *ListVccFlowInfosResponseBody {
	s.Code = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetContent(v *ListVccFlowInfosResponseBodyContent) *ListVccFlowInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetMessage(v string) *ListVccFlowInfosResponseBody {
	s.Message = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) SetRequestId(v string) *ListVccFlowInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVccFlowInfosResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVccFlowInfosResponseBodyContent struct {
	// Lingjun Connection Traffic Information
	Data []*ListVccFlowInfosResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListVccFlowInfosResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListVccFlowInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBodyContent) GetData() []*ListVccFlowInfosResponseBodyContentData {
	return s.Data
}

func (s *ListVccFlowInfosResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListVccFlowInfosResponseBodyContent) SetData(v []*ListVccFlowInfosResponseBodyContentData) *ListVccFlowInfosResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListVccFlowInfosResponseBodyContent) SetTotal(v int64) *ListVccFlowInfosResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContent) Validate() error {
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

type ListVccFlowInfosResponseBodyContentData struct {
	// The direction.
	//
	// example:
	//
	// OUT
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The metric. Valid values:
	//
	// example:
	//
	// passBytesRate
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Time
	//
	// example:
	//
	// 1689749749000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// Value
	//
	// example:
	//
	// 123
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
	// Lingjun Connection ID
	//
	// example:
	//
	// vcc-cn-zvp2w******
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s ListVccFlowInfosResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListVccFlowInfosResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListVccFlowInfosResponseBodyContentData) GetDirection() *string {
	return s.Direction
}

func (s *ListVccFlowInfosResponseBodyContentData) GetMetricName() *string {
	return s.MetricName
}

func (s *ListVccFlowInfosResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVccFlowInfosResponseBodyContentData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListVccFlowInfosResponseBodyContentData) GetValue() *float64 {
	return s.Value
}

func (s *ListVccFlowInfosResponseBodyContentData) GetVccId() *string {
	return s.VccId
}

func (s *ListVccFlowInfosResponseBodyContentData) SetDirection(v string) *ListVccFlowInfosResponseBodyContentData {
	s.Direction = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetMetricName(v string) *ListVccFlowInfosResponseBodyContentData {
	s.MetricName = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetRegionId(v string) *ListVccFlowInfosResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetTimestamp(v int64) *ListVccFlowInfosResponseBodyContentData {
	s.Timestamp = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetValue(v float64) *ListVccFlowInfosResponseBodyContentData {
	s.Value = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) SetVccId(v string) *ListVccFlowInfosResponseBodyContentData {
	s.VccId = &v
	return s
}

func (s *ListVccFlowInfosResponseBodyContentData) Validate() error {
	return dara.Validate(s)
}
