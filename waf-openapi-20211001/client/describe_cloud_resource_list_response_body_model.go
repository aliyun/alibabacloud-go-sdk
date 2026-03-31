// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceList(v []*DescribeCloudResourceListResponseBodyCloudResourceList) *DescribeCloudResourceListResponseBody
	GetCloudResourceList() []*DescribeCloudResourceListResponseBodyCloudResourceList
	SetMaxResults(v int32) *DescribeCloudResourceListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudResourceListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCloudResourceListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudResourceListResponseBody
	GetTotalCount() *int32
}

type DescribeCloudResourceListResponseBody struct {
	CloudResourceList []*DescribeCloudResourceListResponseBodyCloudResourceList `json:"CloudResourceList,omitempty" xml:"CloudResourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAGBgV9tolsLfijC4wam2htS*****D/46H3X2wIS
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 118
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceListResponseBody) GetCloudResourceList() []*DescribeCloudResourceListResponseBodyCloudResourceList {
	return s.CloudResourceList
}

func (s *DescribeCloudResourceListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudResourceListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudResourceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudResourceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudResourceListResponseBody) SetCloudResourceList(v []*DescribeCloudResourceListResponseBodyCloudResourceList) *DescribeCloudResourceListResponseBody {
	s.CloudResourceList = v
	return s
}

func (s *DescribeCloudResourceListResponseBody) SetMaxResults(v int32) *DescribeCloudResourceListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudResourceListResponseBody) SetNextToken(v string) *DescribeCloudResourceListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudResourceListResponseBody) SetRequestId(v string) *DescribeCloudResourceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourceListResponseBody) SetTotalCount(v int32) *DescribeCloudResourceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudResourceListResponseBody) Validate() error {
	if s.CloudResourceList != nil {
		for _, item := range s.CloudResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudResourceListResponseBodyCloudResourceList struct {
	// example:
	//
	// i-8vbdlsd********81e22-80-ecs
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// i-8vbdlsd********81e22
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// rg-aek2uo2****lbka
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// ecs
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeCloudResourceListResponseBodyCloudResourceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceListResponseBodyCloudResourceList) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetCloudResourceId(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.CloudResourceId = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetPort(v int32) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.Port = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetProtocol(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.Protocol = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetResourceInstanceId(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetResourceManagerResourceGroupId(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetResourceProduct(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) SetResourceRegionId(v string) *DescribeCloudResourceListResponseBodyCloudResourceList {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeCloudResourceListResponseBodyCloudResourceList) Validate() error {
	return dara.Validate(s)
}
