// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMonitorGroupInstancesResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitorGroupInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeMonitorGroupInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMonitorGroupInstancesResponseBody
	GetRequestId() *string
	SetResources(v *DescribeMonitorGroupInstancesResponseBodyResources) *DescribeMonitorGroupInstancesResponseBody
	GetResources() *DescribeMonitorGroupInstancesResponseBodyResources
	SetSuccess(v bool) *DescribeMonitorGroupInstancesResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeMonitorGroupInstancesResponseBody
	GetTotal() *int32
}

type DescribeMonitorGroupInstancesResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 97F2A410-9412-499C-9AD1-76EF7EC02DF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resources in the application group.
	Resources *DescribeMonitorGroupInstancesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMonitorGroupInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetResources() *DescribeMonitorGroupInstancesResponseBodyResources {
	return s.Resources
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorGroupInstancesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetCode(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetMessage(v string) *DescribeMonitorGroupInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetPageNumber(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetPageSize(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetRequestId(v string) *DescribeMonitorGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetResources(v *DescribeMonitorGroupInstancesResponseBodyResources) *DescribeMonitorGroupInstancesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) SetTotal(v int32) *DescribeMonitorGroupInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupInstancesResponseBodyResources struct {
	Resource []*DescribeMonitorGroupInstancesResponseBodyResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupInstancesResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBodyResources) GetResource() []*DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	return s.Resource
}

func (s *DescribeMonitorGroupInstancesResponseBodyResources) SetResource(v []*DescribeMonitorGroupInstancesResponseBodyResourcesResource) *DescribeMonitorGroupInstancesResponseBodyResources {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResources) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupInstancesResponseBodyResourcesResource struct {
	// The abbreviation of the service name.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-2ze3w55tr2r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// hostIP
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupInstancesResponseBodyResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupInstancesResponseBodyResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) GetId() *int64 {
	return s.Id
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetCategory(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetId(v int64) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.Id = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetInstanceId(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetInstanceName(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.InstanceName = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) SetRegionId(v string) *DescribeMonitorGroupInstancesResponseBodyResourcesResource {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupInstancesResponseBodyResourcesResource) Validate() error {
	return dara.Validate(s)
}
