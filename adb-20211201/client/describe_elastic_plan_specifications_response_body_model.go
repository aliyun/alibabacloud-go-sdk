// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanSpecificationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeElasticPlanSpecificationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeElasticPlanSpecificationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeElasticPlanSpecificationsResponseBody
	GetRequestId() *string
	SetSpecifications(v []*string) *DescribeElasticPlanSpecificationsResponseBody
	GetSpecifications() []*string
	SetTotalCount(v int32) *DescribeElasticPlanSpecificationsResponseBody
	GetTotalCount() *int32
}

type DescribeElasticPlanSpecificationsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried resource specifications.
	Specifications []*string `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlanSpecificationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeElasticPlanSpecificationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeElasticPlanSpecificationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlanSpecificationsResponseBody) GetSpecifications() []*string {
	return s.Specifications
}

func (s *DescribeElasticPlanSpecificationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetPageNumber(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetPageSize(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetRequestId(v string) *DescribeElasticPlanSpecificationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetSpecifications(v []*string) *DescribeElasticPlanSpecificationsResponseBody {
	s.Specifications = v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) SetTotalCount(v int32) *DescribeElasticPlanSpecificationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsResponseBody) Validate() error {
	return dara.Validate(s)
}
