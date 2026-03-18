// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryStorageMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryStorageMetricResponseBodyData) *QueryStorageMetricResponseBody
	GetData() *QueryStorageMetricResponseBodyData
	SetErrorCode(v string) *QueryStorageMetricResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryStorageMetricResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *QueryStorageMetricResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *QueryStorageMetricResponseBody
	GetRequestId() *string
}

type QueryStorageMetricResponseBody struct {
	// The data returned.
	Data *QueryStorageMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0A3B1FD2006A24C8D8BE65CDAC028298
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: Informational - The request was received and is being processed.
	//
	// - 2xx: Success - The request was successfully received, understood, and accepted by the server.
	//
	// - 3xx: Redirection - The request was redirected. Further action is needed to complete the request.
	//
	// - 4xx: Client error - The request contains incorrect parameters or syntax, or cannot be fulfilled.
	//
	// - 5xx: Server error - The server failed to fulfill the request for other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0bb16654558425251398e27a9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryStorageMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryStorageMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryStorageMetricResponseBody) GetData() *QueryStorageMetricResponseBodyData {
	return s.Data
}

func (s *QueryStorageMetricResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryStorageMetricResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryStorageMetricResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryStorageMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryStorageMetricResponseBody) SetData(v *QueryStorageMetricResponseBodyData) *QueryStorageMetricResponseBody {
	s.Data = v
	return s
}

func (s *QueryStorageMetricResponseBody) SetErrorCode(v string) *QueryStorageMetricResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryStorageMetricResponseBody) SetErrorMsg(v string) *QueryStorageMetricResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryStorageMetricResponseBody) SetHttpCode(v int32) *QueryStorageMetricResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryStorageMetricResponseBody) SetRequestId(v string) *QueryStorageMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryStorageMetricResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryStorageMetricResponseBodyData struct {
	// The category of the metric.
	//
	// example:
	//
	// storage
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The metric values.
	Metrics []*QueryStorageMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// example:
	//
	// summary
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The step size of the monitoring data.
	//
	// example:
	//
	// 3600
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryStorageMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryStorageMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryStorageMetricResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *QueryStorageMetricResponseBodyData) GetMetrics() []*QueryStorageMetricResponseBodyDataMetrics {
	return s.Metrics
}

func (s *QueryStorageMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryStorageMetricResponseBodyData) GetPeriod() *int64 {
	return s.Period
}

func (s *QueryStorageMetricResponseBodyData) SetCategory(v string) *QueryStorageMetricResponseBodyData {
	s.Category = &v
	return s
}

func (s *QueryStorageMetricResponseBodyData) SetMetrics(v []*QueryStorageMetricResponseBodyDataMetrics) *QueryStorageMetricResponseBodyData {
	s.Metrics = v
	return s
}

func (s *QueryStorageMetricResponseBodyData) SetName(v string) *QueryStorageMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryStorageMetricResponseBodyData) SetPeriod(v int64) *QueryStorageMetricResponseBodyData {
	s.Period = &v
	return s
}

func (s *QueryStorageMetricResponseBodyData) Validate() error {
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

type QueryStorageMetricResponseBodyDataMetrics struct {
	// The metadata of the metric.
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	// The time series data.
	Values [][]*float64 `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryStorageMetricResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryStorageMetricResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *QueryStorageMetricResponseBodyDataMetrics) GetMetric() map[string]*string {
	return s.Metric
}

func (s *QueryStorageMetricResponseBodyDataMetrics) GetValues() [][]*float64 {
	return s.Values
}

func (s *QueryStorageMetricResponseBodyDataMetrics) SetMetric(v map[string]*string) *QueryStorageMetricResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *QueryStorageMetricResponseBodyDataMetrics) SetValues(v [][]*float64) *QueryStorageMetricResponseBodyDataMetrics {
	s.Values = v
	return s
}

func (s *QueryStorageMetricResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
