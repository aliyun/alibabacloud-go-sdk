// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryTunnelMetricResponseBodyData) *QueryTunnelMetricResponseBody
	GetData() *QueryTunnelMetricResponseBodyData
	SetErrorCode(v string) *QueryTunnelMetricResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryTunnelMetricResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *QueryTunnelMetricResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *QueryTunnelMetricResponseBody
	GetRequestId() *string
}

type QueryTunnelMetricResponseBody struct {
	Data *QueryTunnelMetricResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 0A3B1E82006A23A918C70905BF08AEC7
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc3b4b016674434996033675e71ee
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryTunnelMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricResponseBody) GetData() *QueryTunnelMetricResponseBodyData {
	return s.Data
}

func (s *QueryTunnelMetricResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryTunnelMetricResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTunnelMetricResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryTunnelMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTunnelMetricResponseBody) SetData(v *QueryTunnelMetricResponseBodyData) *QueryTunnelMetricResponseBody {
	s.Data = v
	return s
}

func (s *QueryTunnelMetricResponseBody) SetErrorCode(v string) *QueryTunnelMetricResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryTunnelMetricResponseBody) SetErrorMsg(v string) *QueryTunnelMetricResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTunnelMetricResponseBody) SetHttpCode(v int32) *QueryTunnelMetricResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryTunnelMetricResponseBody) SetRequestId(v string) *QueryTunnelMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTunnelMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTunnelMetricResponseBodyData struct {
	// example:
	//
	// tunnel
	Category *string                                     `json:"category,omitempty" xml:"category,omitempty"`
	Metrics  []*QueryTunnelMetricResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// example:
	//
	// slot_usage
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 60
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
}

func (s QueryTunnelMetricResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *QueryTunnelMetricResponseBodyData) GetMetrics() []*QueryTunnelMetricResponseBodyDataMetrics {
	return s.Metrics
}

func (s *QueryTunnelMetricResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryTunnelMetricResponseBodyData) GetPeriod() *int64 {
	return s.Period
}

func (s *QueryTunnelMetricResponseBodyData) SetCategory(v string) *QueryTunnelMetricResponseBodyData {
	s.Category = &v
	return s
}

func (s *QueryTunnelMetricResponseBodyData) SetMetrics(v []*QueryTunnelMetricResponseBodyDataMetrics) *QueryTunnelMetricResponseBodyData {
	s.Metrics = v
	return s
}

func (s *QueryTunnelMetricResponseBodyData) SetName(v string) *QueryTunnelMetricResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryTunnelMetricResponseBodyData) SetPeriod(v int64) *QueryTunnelMetricResponseBodyData {
	s.Period = &v
	return s
}

func (s *QueryTunnelMetricResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTunnelMetricResponseBodyDataMetrics struct {
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	Values [][]*float64       `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s QueryTunnelMetricResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricResponseBodyDataMetrics) GetMetric() map[string]*string {
	return s.Metric
}

func (s *QueryTunnelMetricResponseBodyDataMetrics) GetValues() [][]*float64 {
	return s.Values
}

func (s *QueryTunnelMetricResponseBodyDataMetrics) SetMetric(v map[string]*string) *QueryTunnelMetricResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *QueryTunnelMetricResponseBodyDataMetrics) SetValues(v [][]*float64) *QueryTunnelMetricResponseBodyDataMetrics {
	s.Values = v
	return s
}

func (s *QueryTunnelMetricResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
