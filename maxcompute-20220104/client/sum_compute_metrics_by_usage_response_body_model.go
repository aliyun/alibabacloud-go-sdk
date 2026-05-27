// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumComputeMetricsByUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SumComputeMetricsByUsageResponseBodyData) *SumComputeMetricsByUsageResponseBody
	GetData() []*SumComputeMetricsByUsageResponseBodyData
	SetHttpCode(v int32) *SumComputeMetricsByUsageResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumComputeMetricsByUsageResponseBody
	GetRequestId() *string
}

type SumComputeMetricsByUsageResponseBody struct {
	Data []*SumComputeMetricsByUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc0598d17544456742466519e6611
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumComputeMetricsByUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByUsageResponseBody) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByUsageResponseBody) GetData() []*SumComputeMetricsByUsageResponseBodyData {
	return s.Data
}

func (s *SumComputeMetricsByUsageResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumComputeMetricsByUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumComputeMetricsByUsageResponseBody) SetData(v []*SumComputeMetricsByUsageResponseBodyData) *SumComputeMetricsByUsageResponseBody {
	s.Data = v
	return s
}

func (s *SumComputeMetricsByUsageResponseBody) SetHttpCode(v int32) *SumComputeMetricsByUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBody) SetRequestId(v string) *SumComputeMetricsByUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBody) Validate() error {
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

type SumComputeMetricsByUsageResponseBodyData struct {
	DailyComputeMetrics []*SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics `json:"dailyComputeMetrics,omitempty" xml:"dailyComputeMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// ComputationSql
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SumComputeMetricsByUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByUsageResponseBodyData) GetDailyComputeMetrics() []*SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics {
	return s.DailyComputeMetrics
}

func (s *SumComputeMetricsByUsageResponseBodyData) GetType() *string {
	return s.Type
}

func (s *SumComputeMetricsByUsageResponseBodyData) SetDailyComputeMetrics(v []*SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) *SumComputeMetricsByUsageResponseBodyData {
	s.DailyComputeMetrics = v
	return s
}

func (s *SumComputeMetricsByUsageResponseBodyData) SetType(v string) *SumComputeMetricsByUsageResponseBodyData {
	s.Type = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBodyData) Validate() error {
	if s.DailyComputeMetrics != nil {
		for _, item := range s.DailyComputeMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics struct {
	// example:
	//
	// 20260413
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// GBCplx
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 1.149683987
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) String() string {
	return dara.Prettify(s)
}

func (s SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) GoString() string {
	return s.String()
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) GetDateTime() *string {
	return s.DateTime
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) GetUnit() *string {
	return s.Unit
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) GetUsage() *string {
	return s.Usage
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) SetDateTime(v string) *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics {
	s.DateTime = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) SetUnit(v string) *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics {
	s.Unit = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) SetUsage(v string) *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics {
	s.Usage = &v
	return s
}

func (s *SumComputeMetricsByUsageResponseBodyDataDailyComputeMetrics) Validate() error {
	return dara.Validate(s)
}
