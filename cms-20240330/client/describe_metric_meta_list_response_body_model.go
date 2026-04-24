// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricMetaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeMetricMetaListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMetricMetaListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMetricMetaListResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeMetricMetaListResponseBodyResources) *DescribeMetricMetaListResponseBody
	GetResources() []*DescribeMetricMetaListResponseBodyResources
	SetTotalCount(v int64) *DescribeMetricMetaListResponseBody
	GetTotalCount() *int64
}

type DescribeMetricMetaListResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Resources []*DescribeMetricMetaListResponseBodyResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
	// example:
	//
	// 6370
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeMetricMetaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetricMetaListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetricMetaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricMetaListResponseBody) GetResources() []*DescribeMetricMetaListResponseBodyResources {
	return s.Resources
}

func (s *DescribeMetricMetaListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMetricMetaListResponseBody) SetPageNumber(v int32) *DescribeMetricMetaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetPageSize(v int32) *DescribeMetricMetaListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetRequestId(v string) *DescribeMetricMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetResources(v []*DescribeMetricMetaListResponseBodyResources) *DescribeMetricMetaListResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMetricMetaListResponseBody) SetTotalCount(v int64) *DescribeMetricMetaListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetricMetaListResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricMetaListResponseBodyResources struct {
	Description          *string                                                            `json:"description,omitempty" xml:"description,omitempty"`
	DimensionDescription []*DescribeMetricMetaListResponseBodyResourcesDimensionDescription `json:"dimensionDescription,omitempty" xml:"dimensionDescription,omitempty" type:"Repeated"`
	Dimensions           []*string                                                          `json:"dimensions,omitempty" xml:"dimensions,omitempty" type:"Repeated"`
	Labels               map[string]*string                                                 `json:"labels,omitempty" xml:"labels,omitempty"`
	MetricName           *string                                                            `json:"metricName,omitempty" xml:"metricName,omitempty"`
	Namespace            *string                                                            `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Periods              *string                                                            `json:"periods,omitempty" xml:"periods,omitempty"`
	Statistics           *string                                                            `json:"statistics,omitempty" xml:"statistics,omitempty"`
	Type                 *string                                                            `json:"type,omitempty" xml:"type,omitempty"`
	Unit                 *string                                                            `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s DescribeMetricMetaListResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResources) GetDescription() *string {
	return s.Description
}

func (s *DescribeMetricMetaListResponseBodyResources) GetDimensionDescription() []*DescribeMetricMetaListResponseBodyResourcesDimensionDescription {
	return s.DimensionDescription
}

func (s *DescribeMetricMetaListResponseBodyResources) GetDimensions() []*string {
	return s.Dimensions
}

func (s *DescribeMetricMetaListResponseBodyResources) GetLabels() map[string]*string {
	return s.Labels
}

func (s *DescribeMetricMetaListResponseBodyResources) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricMetaListResponseBodyResources) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricMetaListResponseBodyResources) GetPeriods() *string {
	return s.Periods
}

func (s *DescribeMetricMetaListResponseBodyResources) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricMetaListResponseBodyResources) GetType() *string {
	return s.Type
}

func (s *DescribeMetricMetaListResponseBodyResources) GetUnit() *string {
	return s.Unit
}

func (s *DescribeMetricMetaListResponseBodyResources) SetDescription(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Description = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetDimensionDescription(v []*DescribeMetricMetaListResponseBodyResourcesDimensionDescription) *DescribeMetricMetaListResponseBodyResources {
	s.DimensionDescription = v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetDimensions(v []*string) *DescribeMetricMetaListResponseBodyResources {
	s.Dimensions = v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetLabels(v map[string]*string) *DescribeMetricMetaListResponseBodyResources {
	s.Labels = v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetMetricName(v string) *DescribeMetricMetaListResponseBodyResources {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetNamespace(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetPeriods(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Periods = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetStatistics(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetType(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Type = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) SetUnit(v string) *DescribeMetricMetaListResponseBodyResources {
	s.Unit = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResources) Validate() error {
	if s.DimensionDescription != nil {
		for _, item := range s.DimensionDescription {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricMetaListResponseBodyResourcesDimensionDescription struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeMetricMetaListResponseBodyResourcesDimensionDescription) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListResponseBodyResourcesDimensionDescription) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListResponseBodyResourcesDimensionDescription) GetName() *string {
	return s.Name
}

func (s *DescribeMetricMetaListResponseBodyResourcesDimensionDescription) SetName(v string) *DescribeMetricMetaListResponseBodyResourcesDimensionDescription {
	s.Name = &v
	return s
}

func (s *DescribeMetricMetaListResponseBodyResourcesDimensionDescription) Validate() error {
	return dara.Validate(s)
}
