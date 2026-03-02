// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceMetricPageResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*MetricData) *ServiceMetricPageResult
	GetData() []*MetricData
	SetIntTotal(v int32) *ServiceMetricPageResult
	GetIntTotal() *int32
	SetRequestId(v string) *ServiceMetricPageResult
	GetRequestId() *string
	SetTotal(v int64) *ServiceMetricPageResult
	GetTotal() *int64
}

type ServiceMetricPageResult struct {
	Data      []*MetricData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	IntTotal  *int32        `json:"intTotal,omitempty" xml:"intTotal,omitempty"`
	RequestId *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ServiceMetricPageResult) String() string {
	return dara.Prettify(s)
}

func (s ServiceMetricPageResult) GoString() string {
	return s.String()
}

func (s *ServiceMetricPageResult) GetData() []*MetricData {
	return s.Data
}

func (s *ServiceMetricPageResult) GetIntTotal() *int32 {
	return s.IntTotal
}

func (s *ServiceMetricPageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ServiceMetricPageResult) GetTotal() *int64 {
	return s.Total
}

func (s *ServiceMetricPageResult) SetData(v []*MetricData) *ServiceMetricPageResult {
	s.Data = v
	return s
}

func (s *ServiceMetricPageResult) SetIntTotal(v int32) *ServiceMetricPageResult {
	s.IntTotal = &v
	return s
}

func (s *ServiceMetricPageResult) SetRequestId(v string) *ServiceMetricPageResult {
	s.RequestId = &v
	return s
}

func (s *ServiceMetricPageResult) SetTotal(v int64) *ServiceMetricPageResult {
	s.Total = &v
	return s
}

func (s *ServiceMetricPageResult) Validate() error {
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
