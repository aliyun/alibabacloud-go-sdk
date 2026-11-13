// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMetricDataResponseBody
	GetCode() *string
	SetMessage(v string) *GetMetricDataResponseBody
	GetMessage() *string
	SetMetricData(v *MetricData) *GetMetricDataResponseBody
	GetMetricData() *MetricData
	SetRequestId(v string) *GetMetricDataResponseBody
	GetRequestId() *string
}

type GetMetricDataResponseBody struct {
	// 成功时为 200。失败时为对应的错误码。
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 本次请求的结果说明。
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 各指标的查询结果及对应的时间和分组。
	//
	// example:
	//
	// {"period":60,"startTime":"2026-11-05T00:00:00Z","endTime":"2026-11-05T00:02:00Z","groupBy":["teamID"],"results":[{"metricName":"cpuUsedCores","unit":"Core","series":[{"dimensions":{"teamID":"70d1c834-0333-58d8-97ac-5336eb9179eb"},"points":[{"timestamp":"2026-11-05T00:00:00Z","value":0.25},{"timestamp":"2026-11-05T00:01:00Z","value":0.5}]}]}]}
	MetricData *MetricData `json:"metricData,omitempty" xml:"metricData,omitempty"`
	// 本次请求的唯一标识。
	//
	// example:
	//
	// 12345678-1234-5678-9abc-123456789abc
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMetricDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMetricDataResponseBody) GetMetricData() *MetricData {
	return s.MetricData
}

func (s *GetMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetricDataResponseBody) SetCode(v string) *GetMetricDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetMetricDataResponseBody) SetMessage(v string) *GetMetricDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetMetricDataResponseBody) SetMetricData(v *MetricData) *GetMetricDataResponseBody {
	s.MetricData = v
	return s
}

func (s *GetMetricDataResponseBody) SetRequestId(v string) *GetMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricDataResponseBody) Validate() error {
	if s.MetricData != nil {
		if err := s.MetricData.Validate(); err != nil {
			return err
		}
	}
	return nil
}
