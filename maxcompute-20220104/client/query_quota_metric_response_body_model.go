// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryQuotaMetricResponseBodyData) *QueryQuotaMetricResponseBody
	GetData() *QueryQuotaMetricResponseBodyData
	SetErrorCode(v string) *QueryQuotaMetricResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryQuotaMetricResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *QueryQuotaMetricResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *QueryQuotaMetricResponseBody
	GetRequestId() *string
}

type QueryQuotaMetricResponseBody struct {
	Data *QueryQuotaMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// plan \\"***\\" does not exist
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0b87b7b316643495896551555e855b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryQuotaMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricResponseBody) GetData() *QueryQuotaMetricResponseBodyData {
	return s.Data
}

func (s *QueryQuotaMetricResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryQuotaMetricResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryQuotaMetricResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryQuotaMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryQuotaMetricResponseBody) SetData(v *QueryQuotaMetricResponseBodyData) *QueryQuotaMetricResponseBody {
	s.Data = v
	return s
}

func (s *QueryQuotaMetricResponseBody) SetErrorCode(v string) *QueryQuotaMetricResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryQuotaMetricResponseBody) SetErrorMsg(v string) *QueryQuotaMetricResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryQuotaMetricResponseBody) SetHttpCode(v int32) *QueryQuotaMetricResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryQuotaMetricResponseBody) SetRequestId(v string) *QueryQuotaMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryQuotaMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryQuotaMetricResponseBodyData struct {
	Category *string                                    `json:"category,omitempty" xml:"category,omitempty"`
	Metrics  []*QueryQuotaMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// example:
	//
	// cpu
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 60
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryQuotaMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *QueryQuotaMetricResponseBodyData) GetMetrics() []*QueryQuotaMetricResponseBodyDataMetrics {
	return s.Metrics
}

func (s *QueryQuotaMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryQuotaMetricResponseBodyData) GetPeriod() *int64 {
	return s.Period
}

func (s *QueryQuotaMetricResponseBodyData) SetCategory(v string) *QueryQuotaMetricResponseBodyData {
	s.Category = &v
	return s
}

func (s *QueryQuotaMetricResponseBodyData) SetMetrics(v []*QueryQuotaMetricResponseBodyDataMetrics) *QueryQuotaMetricResponseBodyData {
	s.Metrics = v
	return s
}

func (s *QueryQuotaMetricResponseBodyData) SetName(v string) *QueryQuotaMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryQuotaMetricResponseBodyData) SetPeriod(v int64) *QueryQuotaMetricResponseBodyData {
	s.Period = &v
	return s
}

func (s *QueryQuotaMetricResponseBodyData) Validate() error {
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

type QueryQuotaMetricResponseBodyDataMetrics struct {
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	Values [][]*float64       `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryQuotaMetricResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricResponseBodyDataMetrics) GetMetric() map[string]*string {
	return s.Metric
}

func (s *QueryQuotaMetricResponseBodyDataMetrics) GetValues() [][]*float64 {
	return s.Values
}

func (s *QueryQuotaMetricResponseBodyDataMetrics) SetMetric(v map[string]*string) *QueryQuotaMetricResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *QueryQuotaMetricResponseBodyDataMetrics) SetValues(v [][]*float64) *QueryQuotaMetricResponseBodyDataMetrics {
	s.Values = v
	return s
}

func (s *QueryQuotaMetricResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
