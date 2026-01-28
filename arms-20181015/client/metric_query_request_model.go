// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v []*string) *MetricQueryRequest
	GetDimensions() []*string
	SetEndTime(v int64) *MetricQueryRequest
	GetEndTime() *int64
	SetFilters(v []*MetricQueryRequestFilters) *MetricQueryRequest
	GetFilters() []*MetricQueryRequestFilters
	SetHackerUserId(v string) *MetricQueryRequest
	GetHackerUserId() *string
	SetIintervalInSec(v int32) *MetricQueryRequest
	GetIintervalInSec() *int32
	SetLimit(v int32) *MetricQueryRequest
	GetLimit() *int32
	SetMeasures(v []*string) *MetricQueryRequest
	GetMeasures() []*string
	SetMetric(v string) *MetricQueryRequest
	GetMetric() *string
	SetOrder(v string) *MetricQueryRequest
	GetOrder() *string
	SetOrderBy(v string) *MetricQueryRequest
	GetOrderBy() *string
	SetSecurityToken(v string) *MetricQueryRequest
	GetSecurityToken() *string
	SetStartTime(v int64) *MetricQueryRequest
	GetStartTime() *int64
}

type MetricQueryRequest struct {
	Dimensions []*string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// This parameter is required.
	EndTime        *int64                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters        []*MetricQueryRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	HackerUserId   *string                      `json:"HackerUserId,omitempty" xml:"HackerUserId,omitempty"`
	IintervalInSec *int32                       `json:"IintervalInSec,omitempty" xml:"IintervalInSec,omitempty"`
	Limit          *int32                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is required.
	Measures []*string `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	// This parameter is required.
	Metric        *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy       *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s MetricQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s MetricQueryRequest) GoString() string {
	return s.String()
}

func (s *MetricQueryRequest) GetDimensions() []*string {
	return s.Dimensions
}

func (s *MetricQueryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *MetricQueryRequest) GetFilters() []*MetricQueryRequestFilters {
	return s.Filters
}

func (s *MetricQueryRequest) GetHackerUserId() *string {
	return s.HackerUserId
}

func (s *MetricQueryRequest) GetIintervalInSec() *int32 {
	return s.IintervalInSec
}

func (s *MetricQueryRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *MetricQueryRequest) GetMeasures() []*string {
	return s.Measures
}

func (s *MetricQueryRequest) GetMetric() *string {
	return s.Metric
}

func (s *MetricQueryRequest) GetOrder() *string {
	return s.Order
}

func (s *MetricQueryRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *MetricQueryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *MetricQueryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *MetricQueryRequest) SetDimensions(v []*string) *MetricQueryRequest {
	s.Dimensions = v
	return s
}

func (s *MetricQueryRequest) SetEndTime(v int64) *MetricQueryRequest {
	s.EndTime = &v
	return s
}

func (s *MetricQueryRequest) SetFilters(v []*MetricQueryRequestFilters) *MetricQueryRequest {
	s.Filters = v
	return s
}

func (s *MetricQueryRequest) SetHackerUserId(v string) *MetricQueryRequest {
	s.HackerUserId = &v
	return s
}

func (s *MetricQueryRequest) SetIintervalInSec(v int32) *MetricQueryRequest {
	s.IintervalInSec = &v
	return s
}

func (s *MetricQueryRequest) SetLimit(v int32) *MetricQueryRequest {
	s.Limit = &v
	return s
}

func (s *MetricQueryRequest) SetMeasures(v []*string) *MetricQueryRequest {
	s.Measures = v
	return s
}

func (s *MetricQueryRequest) SetMetric(v string) *MetricQueryRequest {
	s.Metric = &v
	return s
}

func (s *MetricQueryRequest) SetOrder(v string) *MetricQueryRequest {
	s.Order = &v
	return s
}

func (s *MetricQueryRequest) SetOrderBy(v string) *MetricQueryRequest {
	s.OrderBy = &v
	return s
}

func (s *MetricQueryRequest) SetSecurityToken(v string) *MetricQueryRequest {
	s.SecurityToken = &v
	return s
}

func (s *MetricQueryRequest) SetStartTime(v int64) *MetricQueryRequest {
	s.StartTime = &v
	return s
}

func (s *MetricQueryRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MetricQueryRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s MetricQueryRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s MetricQueryRequestFilters) GoString() string {
	return s.String()
}

func (s *MetricQueryRequestFilters) GetKey() *string {
	return s.Key
}

func (s *MetricQueryRequestFilters) GetValue() *string {
	return s.Value
}

func (s *MetricQueryRequestFilters) SetKey(v string) *MetricQueryRequestFilters {
	s.Key = &v
	return s
}

func (s *MetricQueryRequestFilters) SetValue(v string) *MetricQueryRequestFilters {
	s.Value = &v
	return s
}

func (s *MetricQueryRequestFilters) Validate() error {
	return dara.Validate(s)
}
