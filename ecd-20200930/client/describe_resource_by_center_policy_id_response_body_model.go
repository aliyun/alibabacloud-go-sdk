// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceByCenterPolicyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v string) *DescribeResourceByCenterPolicyIdResponseBody
	GetCount() *string
	SetNextToken(v string) *DescribeResourceByCenterPolicyIdResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeResourceByCenterPolicyIdResponseBody
	GetRequestId() *string
	SetResourceModelList(v []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) *DescribeResourceByCenterPolicyIdResponseBody
	GetResourceModelList() []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelList
}

type DescribeResourceByCenterPolicyIdResponseBody struct {
	// The total number of resources.
	//
	// example:
	//
	// 2
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 48174475-5EB2-5F99-A9E9-6F892D645****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	ResourceModelList []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelList `json:"ResourceModelList,omitempty" xml:"ResourceModelList,omitempty" type:"Repeated"`
}

func (s DescribeResourceByCenterPolicyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByCenterPolicyIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) GetCount() *string {
	return s.Count
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) GetResourceModelList() []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	return s.ResourceModelList
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) SetCount(v string) *DescribeResourceByCenterPolicyIdResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) SetNextToken(v string) *DescribeResourceByCenterPolicyIdResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) SetRequestId(v string) *DescribeResourceByCenterPolicyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) SetResourceModelList(v []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) *DescribeResourceByCenterPolicyIdResponseBody {
	s.ResourceModelList = v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBody) Validate() error {
	if s.ResourceModelList != nil {
		for _, item := range s.ResourceModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceByCenterPolicyIdResponseBodyResourceModelList struct {
	// The cloud applications.
	AppModelList []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList `json:"AppModelList,omitempty" xml:"AppModelList,omitempty" type:"Repeated"`
	// The number of vCPUs.
	//
	// example:
	//
	// 64
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The cloud computer type. You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the IDs of the cloud computer types supported by Alibaba Cloud Workspace.
	//
	// example:
	//
	// eds.enterprise_office.8c32g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 0.125
	GpuCount *float64 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU type.
	//
	// example:
	//
	// 2GiB
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 10240
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The OS type.
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The billing method.
	//
	// example:
	//
	// postPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The service type.
	//
	// example:
	//
	// desktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-d7pasxsd3b9nhq**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The number of associated resource groups
	//
	// example:
	//
	// 10
	ResourceGroupRelCount *int32 `json:"ResourceGroupRelCount,omitempty" xml:"ResourceGroupRelCount,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// ecd-7o96aa08fr****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shenzhen
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource status.
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GoString() string {
	return s.String()
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetAppModelList() []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList {
	return s.AppModelList
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetGpuCount() *float64 {
	return s.GpuCount
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetOsType() *string {
	return s.OsType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetPayType() *string {
	return s.PayType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceGroupRelCount() *int32 {
	return s.ResourceGroupRelCount
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetAppModelList(v []*DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.AppModelList = v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetCpu(v int32) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.Cpu = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetDesktopType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.DesktopType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetGpuCount(v float64) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.GpuCount = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetGpuSpec(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.GpuSpec = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetMemory(v int64) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.Memory = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetOsType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.OsType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetPayType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.PayType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetProductType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ProductType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetProtocolType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ProtocolType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceGroupId(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceGroupName(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceGroupRelCount(v int32) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceGroupRelCount = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceId(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceName(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceRegionId(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetResourceType(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) SetStatus(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList {
	s.Status = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelList) Validate() error {
	if s.AppModelList != nil {
		for _, item := range s.AppModelList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList struct {
	// The application ID. This parameter is only applicable to cloud applications.
	//
	// example:
	//
	// 18
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// alipic-powergem
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) GoString() string {
	return s.String()
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) GetAppId() *string {
	return s.AppId
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) GetAppName() *string {
	return s.AppName
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) SetAppId(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList {
	s.AppId = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) SetAppName(v string) *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList {
	s.AppName = &v
	return s
}

func (s *DescribeResourceByCenterPolicyIdResponseBodyResourceModelListAppModelList) Validate() error {
	return dara.Validate(s)
}
