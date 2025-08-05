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
	// The data returned.
	Data *ListJobMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// HTTP status code.
	//
	// - 1xx: Informational response - Request received, processing continues.
	//
	// - 2xx: Success - The request has been successfully received, understood, and accepted by the server.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client error - The request contains bad syntax or cannot be fulfilled.
	//
	// - 5xx: Server error - The server failed to fulfill an apparently valid request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0aa16667684362147582e038f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	return dara.Validate(s)
}

type ListJobMetricResponseBodyData struct {
	// The category of the metrics.
	//
	// example:
	//
	// job
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Metric details.
	Metrics []*ListJobMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// The name of observation metric.
	//
	// example:
	//
	// num
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The monitoring statistical period.Unit:Second(s).
	//
	// example:
	//
	// 3600
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
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
	return dara.Validate(s)
}

type ListJobMetricResponseBodyDataMetrics struct {
	// Metric related information.
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	// Metric values information.
	Values [][]*float64 `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
