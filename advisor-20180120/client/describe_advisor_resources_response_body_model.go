// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAdvisorResourcesResponseBodyData) *DescribeAdvisorResourcesResponseBody
	GetData() *DescribeAdvisorResourcesResponseBodyData
	SetRequestId(v string) *DescribeAdvisorResourcesResponseBody
	GetRequestId() *string
}

type DescribeAdvisorResourcesResponseBody struct {
	Data *DescribeAdvisorResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBody) GetData() *DescribeAdvisorResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeAdvisorResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvisorResourcesResponseBody) SetData(v *DescribeAdvisorResourcesResponseBodyData) *DescribeAdvisorResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorResourcesResponseBody) SetRequestId(v string) *DescribeAdvisorResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvisorResourcesResponseBodyData struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   *DescribeAdvisorResourcesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAdvisorResourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvisorResourcesResponseBodyData) GetResult() *DescribeAdvisorResourcesResponseBodyDataResult {
	return s.Result
}

func (s *DescribeAdvisorResourcesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetPageNo(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetPageSize(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetResult(v *DescribeAdvisorResourcesResponseBodyDataResult) *DescribeAdvisorResourcesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) SetTotal(v int32) *DescribeAdvisorResourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvisorResourcesResponseBodyDataResult struct {
	Resource []*DescribeAdvisorResourcesResponseBodyDataResultResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeAdvisorResourcesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyDataResult) GetResource() []*DescribeAdvisorResourcesResponseBodyDataResultResource {
	return s.Resource
}

func (s *DescribeAdvisorResourcesResponseBodyDataResult) SetResource(v []*DescribeAdvisorResourcesResponseBodyDataResultResource) *DescribeAdvisorResourcesResponseBodyDataResult {
	s.Resource = v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResult) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvisorResourcesResponseBodyDataResultResource struct {
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s DescribeAdvisorResourcesResponseBodyDataResultResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesResponseBodyDataResultResource) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) GetData() *string {
	return s.Data
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetData(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.Data = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetProduct(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetRegionId(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.RegionId = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetResourceId(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) SetResourceName(v string) *DescribeAdvisorResourcesResponseBodyDataResultResource {
	s.ResourceName = &v
	return s
}

func (s *DescribeAdvisorResourcesResponseBodyDataResultResource) Validate() error {
	return dara.Validate(s)
}
