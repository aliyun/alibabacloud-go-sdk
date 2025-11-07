// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNisNetworkMetricsResponseBodyData) *GetNisNetworkMetricsResponseBody
	GetData() *GetNisNetworkMetricsResponseBodyData
	SetRequestId(v string) *GetNisNetworkMetricsResponseBody
	GetRequestId() *string
}

type GetNisNetworkMetricsResponseBody struct {
	Data *GetNisNetworkMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D5E98683-355B-5867-8D3D-A24755F6895B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNisNetworkMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBody) GetData() *GetNisNetworkMetricsResponseBodyData {
	return s.Data
}

func (s *GetNisNetworkMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNisNetworkMetricsResponseBody) SetData(v *GetNisNetworkMetricsResponseBodyData) *GetNisNetworkMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetNisNetworkMetricsResponseBody) SetRequestId(v string) *GetNisNetworkMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNisNetworkMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNisNetworkMetricsResponseBodyData struct {
	Metrics []*GetNisNetworkMetricsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// Bits/Second
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetNisNetworkMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBodyData) GetMetrics() []*GetNisNetworkMetricsResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetNisNetworkMetricsResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetNisNetworkMetricsResponseBodyData) SetMetrics(v []*GetNisNetworkMetricsResponseBodyDataMetrics) *GetNisNetworkMetricsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyData) SetUnit(v string) *GetNisNetworkMetricsResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyData) Validate() error {
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

type GetNisNetworkMetricsResponseBodyDataMetrics struct {
	// example:
	//
	// 1690684091100
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 88
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkMetricsResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) GetValue() *float64 {
	return s.Value
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) SetTimeStamp(v int64) *GetNisNetworkMetricsResponseBodyDataMetrics {
	s.TimeStamp = &v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) SetValue(v float64) *GetNisNetworkMetricsResponseBodyDataMetrics {
	s.Value = &v
	return s
}

func (s *GetNisNetworkMetricsResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}
