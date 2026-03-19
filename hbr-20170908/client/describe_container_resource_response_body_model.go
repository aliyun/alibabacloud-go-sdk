// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeContainerResourceResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeContainerResourceResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeContainerResourceResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeContainerResourceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeContainerResourceResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeContainerResourceResponseBodyResources) *DescribeContainerResourceResponseBody
	GetResources() []*DescribeContainerResourceResponseBodyResources
	SetSuccess(v bool) *DescribeContainerResourceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeContainerResourceResponseBody
	GetTotalCount() *int64
}

type DescribeContainerResourceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// D98A2895-745B-5530-A8C1-9A86F0A82354
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*DescribeContainerResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeContainerResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeContainerResourceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeContainerResourceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerResourceResponseBody) GetResources() []*DescribeContainerResourceResponseBodyResources {
	return s.Resources
}

func (s *DescribeContainerResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeContainerResourceResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeContainerResourceResponseBody) SetCode(v string) *DescribeContainerResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetMessage(v string) *DescribeContainerResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetPageNumber(v int32) *DescribeContainerResourceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetPageSize(v int32) *DescribeContainerResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetRequestId(v string) *DescribeContainerResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetResources(v []*DescribeContainerResourceResponseBodyResources) *DescribeContainerResourceResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetSuccess(v bool) *DescribeContainerResourceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) SetTotalCount(v int64) *DescribeContainerResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeContainerResourceResponseBody) Validate() error {
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

type DescribeContainerResourceResponseBodyResources struct {
	// example:
	//
	// cc-0005**********hhjw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// c21b653f********695e892e718c419a4
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" xml:"ClusterIdentifier,omitempty"`
	// example:
	//
	// a9ab843d-****-****-8e46-1d67a82128a7
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// {"pv_name":"nas-a9ab843d-****-****-8e46-1d67a82128a7","pv_size":"1000Gi","storage_class":"opk8s-nas","pvc_name":"**-pvc","namespace":"default"}
	ResourceInfo *string `json:"ResourceInfo,omitempty" xml:"ResourceInfo,omitempty"`
	// example:
	//
	// PV
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeContainerResourceResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerResourceResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeContainerResourceResponseBodyResources) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerResourceResponseBodyResources) GetClusterIdentifier() *string {
	return s.ClusterIdentifier
}

func (s *DescribeContainerResourceResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeContainerResourceResponseBodyResources) GetResourceInfo() *string {
	return s.ResourceInfo
}

func (s *DescribeContainerResourceResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeContainerResourceResponseBodyResources) SetClusterId(v string) *DescribeContainerResourceResponseBodyResources {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerResourceResponseBodyResources) SetClusterIdentifier(v string) *DescribeContainerResourceResponseBodyResources {
	s.ClusterIdentifier = &v
	return s
}

func (s *DescribeContainerResourceResponseBodyResources) SetResourceId(v string) *DescribeContainerResourceResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeContainerResourceResponseBodyResources) SetResourceInfo(v string) *DescribeContainerResourceResponseBodyResources {
	s.ResourceInfo = &v
	return s
}

func (s *DescribeContainerResourceResponseBodyResources) SetResourceType(v string) *DescribeContainerResourceResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeContainerResourceResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
