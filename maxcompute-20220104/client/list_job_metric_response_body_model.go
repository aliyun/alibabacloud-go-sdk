// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListJobMetricResponseBodyData) *ListJobMetricResponseBody
	GetData() *ListJobMetricResponseBodyData
	SetErrorCode(v string) *ListJobMetricResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListJobMetricResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListJobMetricResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListJobMetricResponseBody
	GetRequestId() *string
}

type ListJobMetricResponseBody struct {
	Data      *ListJobMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                        `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                         `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobMetricResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBody) GetData() *ListJobMetricResponseBodyData {
	return s.Data
}

func (s *ListJobMetricResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobMetricResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListJobMetricResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListJobMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobMetricResponseBody) SetData(v *ListJobMetricResponseBodyData) *ListJobMetricResponseBody {
	s.Data = v
	return s
}

func (s *ListJobMetricResponseBody) SetErrorCode(v string) *ListJobMetricResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobMetricResponseBody) SetErrorMsg(v string) *ListJobMetricResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListJobMetricResponseBody) SetHttpCode(v int32) *ListJobMetricResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobMetricResponseBody) SetRequestId(v string) *ListJobMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobMetricResponseBodyData struct {
	Category *string                                 `json:"category,omitempty" xml:"category,omitempty"`
	Metrics  []*ListJobMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	Name     *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Period   *int64                                  `json:"period,omitempty" xml:"period,omitempty"`
}

func (s ListJobMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *ListJobMetricResponseBodyData) GetMetrics() []*ListJobMetricResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListJobMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListJobMetricResponseBodyData) GetPeriod() *int64 {
	return s.Period
}

func (s *ListJobMetricResponseBodyData) SetCategory(v string) *ListJobMetricResponseBodyData {
	s.Category = &v
	return s
}

func (s *ListJobMetricResponseBodyData) SetMetrics(v []*ListJobMetricResponseBodyDataMetrics) *ListJobMetricResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListJobMetricResponseBodyData) SetName(v string) *ListJobMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListJobMetricResponseBodyData) SetPeriod(v int64) *ListJobMetricResponseBodyData {
	s.Period = &v
	return s
}

func (s *ListJobMetricResponseBodyData) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobMetricResponseBodyDataMetrics struct {
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	Values [][]*float64       `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListJobMetricResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListJobMetricResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListJobMetricResponseBodyDataMetrics) GetMetric() map[string]*string {
	return s.Metric
}

func (s *ListJobMetricResponseBodyDataMetrics) GetValues() [][]*float64 {
	return s.Values
}

func (s *ListJobMetricResponseBodyDataMetrics) SetMetric(v map[string]*string) *ListJobMetricResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *ListJobMetricResponseBodyDataMetrics) SetValues(v [][]*float64) *ListJobMetricResponseBodyDataMetrics {
	s.Values = v
	return s
}

func (s *ListJobMetricResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
