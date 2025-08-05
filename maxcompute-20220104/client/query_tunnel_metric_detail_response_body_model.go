// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTunnelMetricDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryTunnelMetricDetailResponseBodyData) *QueryTunnelMetricDetailResponseBody
	GetData() *QueryTunnelMetricDetailResponseBodyData
	SetErrorCode(v string) *QueryTunnelMetricDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *QueryTunnelMetricDetailResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *QueryTunnelMetricDetailResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *QueryTunnelMetricDetailResponseBody
	GetRequestId() *string
}

type QueryTunnelMetricDetailResponseBody struct {
	Data *QueryTunnelMetricDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// OBJECT_NOT_EXIST
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
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryTunnelMetricDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricDetailResponseBody) GetData() *QueryTunnelMetricDetailResponseBodyData {
	return s.Data
}

func (s *QueryTunnelMetricDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryTunnelMetricDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *QueryTunnelMetricDetailResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *QueryTunnelMetricDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTunnelMetricDetailResponseBody) SetData(v *QueryTunnelMetricDetailResponseBodyData) *QueryTunnelMetricDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryTunnelMetricDetailResponseBody) SetErrorCode(v string) *QueryTunnelMetricDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryTunnelMetricDetailResponseBody) SetErrorMsg(v string) *QueryTunnelMetricDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *QueryTunnelMetricDetailResponseBody) SetHttpCode(v int32) *QueryTunnelMetricDetailResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryTunnelMetricDetailResponseBody) SetRequestId(v string) *QueryTunnelMetricDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTunnelMetricDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTunnelMetricDetailResponseBodyData struct {
	Metrics []*QueryTunnelMetricDetailResponseBodyDataMetrics `json:"metrics,omitempty" xml:"metrics,omitempty" type:"Repeated"`
	// example:
	//
	// tableA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s QueryTunnelMetricDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricDetailResponseBodyData) GetMetrics() []*QueryTunnelMetricDetailResponseBodyDataMetrics {
	return s.Metrics
}

func (s *QueryTunnelMetricDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryTunnelMetricDetailResponseBodyData) SetMetrics(v []*QueryTunnelMetricDetailResponseBodyDataMetrics) *QueryTunnelMetricDetailResponseBodyData {
	s.Metrics = v
	return s
}

func (s *QueryTunnelMetricDetailResponseBodyData) SetName(v string) *QueryTunnelMetricDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryTunnelMetricDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTunnelMetricDetailResponseBodyDataMetrics struct {
	Metric map[string]*string `json:"metric,omitempty" xml:"metric,omitempty"`
	// example:
	//
	// "avgValue":"11.5"
	Value map[string]interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s QueryTunnelMetricDetailResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s QueryTunnelMetricDetailResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *QueryTunnelMetricDetailResponseBodyDataMetrics) GetMetric() map[string]*string {
	return s.Metric
}

func (s *QueryTunnelMetricDetailResponseBodyDataMetrics) GetValue() map[string]interface{} {
	return s.Value
}

func (s *QueryTunnelMetricDetailResponseBodyDataMetrics) SetMetric(v map[string]*string) *QueryTunnelMetricDetailResponseBodyDataMetrics {
	s.Metric = v
	return s
}

func (s *QueryTunnelMetricDetailResponseBodyDataMetrics) SetValue(v map[string]interface{}) *QueryTunnelMetricDetailResponseBodyDataMetrics {
	s.Value = v
	return s
}

func (s *QueryTunnelMetricDetailResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
