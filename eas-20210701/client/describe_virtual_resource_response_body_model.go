// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeVirtualResourceResponseBody
	GetCreateTime() *string
	SetDisableSpotProtectionPeriod(v bool) *DescribeVirtualResourceResponseBody
	GetDisableSpotProtectionPeriod() *bool
	SetFeatures(v []*string) *DescribeVirtualResourceResponseBody
	GetFeatures() []*string
	SetRequestId(v string) *DescribeVirtualResourceResponseBody
	GetRequestId() *string
	SetResources(v []*DescribeVirtualResourceResponseBodyResources) *DescribeVirtualResourceResponseBody
	GetResources() []*DescribeVirtualResourceResponseBodyResources
	SetServiceCount(v int32) *DescribeVirtualResourceResponseBody
	GetServiceCount() *int32
	SetUpdateTime(v string) *DescribeVirtualResourceResponseBody
	GetUpdateTime() *string
	SetVirtualResourceId(v string) *DescribeVirtualResourceResponseBody
	GetVirtualResourceId() *string
	SetVirtualResourceName(v string) *DescribeVirtualResourceResponseBody
	GetVirtualResourceName() *string
}

type DescribeVirtualResourceResponseBody struct {
	// The time when the virtual resource group was created.
	//
	// example:
	//
	// 2024-10-16T17:52:49Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the retention period of preemptible instances was disabled.
	//
	// example:
	//
	// true
	DisableSpotProtectionPeriod *bool     `json:"DisableSpotProtectionPeriod,omitempty" xml:"DisableSpotProtectionPeriod,omitempty"`
	Features                    []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources in the virtual resource group.
	Resources []*DescribeVirtualResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The number of deployed services.
	//
	// example:
	//
	// 1
	ServiceCount *int32 `json:"ServiceCount,omitempty" xml:"ServiceCount,omitempty"`
	// The time when the virtual resource group was last updated.
	//
	// example:
	//
	// 2024-10-16T19:52:49Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual resource group.
	//
	// example:
	//
	// eas-vr-npovr28onap1xxxxxx
	VirtualResourceId *string `json:"VirtualResourceId,omitempty" xml:"VirtualResourceId,omitempty"`
	// The name of the virtual resource group.
	//
	// example:
	//
	// MyVirtualResource
	VirtualResourceName *string `json:"VirtualResourceName,omitempty" xml:"VirtualResourceName,omitempty"`
}

func (s DescribeVirtualResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualResourceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeVirtualResourceResponseBody) GetDisableSpotProtectionPeriod() *bool {
	return s.DisableSpotProtectionPeriod
}

func (s *DescribeVirtualResourceResponseBody) GetFeatures() []*string {
	return s.Features
}

func (s *DescribeVirtualResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualResourceResponseBody) GetResources() []*DescribeVirtualResourceResponseBodyResources {
	return s.Resources
}

func (s *DescribeVirtualResourceResponseBody) GetServiceCount() *int32 {
	return s.ServiceCount
}

func (s *DescribeVirtualResourceResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeVirtualResourceResponseBody) GetVirtualResourceId() *string {
	return s.VirtualResourceId
}

func (s *DescribeVirtualResourceResponseBody) GetVirtualResourceName() *string {
	return s.VirtualResourceName
}

func (s *DescribeVirtualResourceResponseBody) SetCreateTime(v string) *DescribeVirtualResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetDisableSpotProtectionPeriod(v bool) *DescribeVirtualResourceResponseBody {
	s.DisableSpotProtectionPeriod = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetFeatures(v []*string) *DescribeVirtualResourceResponseBody {
	s.Features = v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetRequestId(v string) *DescribeVirtualResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetResources(v []*DescribeVirtualResourceResponseBodyResources) *DescribeVirtualResourceResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetServiceCount(v int32) *DescribeVirtualResourceResponseBody {
	s.ServiceCount = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetUpdateTime(v string) *DescribeVirtualResourceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetVirtualResourceId(v string) *DescribeVirtualResourceResponseBody {
	s.VirtualResourceId = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) SetVirtualResourceName(v string) *DescribeVirtualResourceResponseBody {
	s.VirtualResourceName = &v
	return s
}

func (s *DescribeVirtualResourceResponseBody) Validate() error {
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

type DescribeVirtualResourceResponseBodyResources struct {
	// The instance type of the public resource group.
	//
	// example:
	//
	// ecs.s6-c1m2.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The priority of resource scheduling. A greater number specifies a higher priority.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The instance type of the public resource group.
	//
	// example:
	//
	// quota185lqxxxxxx
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The region where the resource resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the dedicated resource group.
	//
	// example:
	//
	// eas-r-g55ieatgg3buxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The maximum price of preemptible instances in a public resource group.
	//
	// example:
	//
	// 10.05
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s DescribeVirtualResourceResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualResourceResponseBodyResources) GoString() string {
	return s.String()
}

func (s *DescribeVirtualResourceResponseBodyResources) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeVirtualResourceResponseBodyResources) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeVirtualResourceResponseBodyResources) GetQuotaId() *string {
	return s.QuotaId
}

func (s *DescribeVirtualResourceResponseBodyResources) GetRegion() *string {
	return s.Region
}

func (s *DescribeVirtualResourceResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeVirtualResourceResponseBodyResources) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeVirtualResourceResponseBodyResources) SetInstanceType(v string) *DescribeVirtualResourceResponseBodyResources {
	s.InstanceType = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) SetPriority(v int32) *DescribeVirtualResourceResponseBodyResources {
	s.Priority = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) SetQuotaId(v string) *DescribeVirtualResourceResponseBodyResources {
	s.QuotaId = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) SetRegion(v string) *DescribeVirtualResourceResponseBodyResources {
	s.Region = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) SetResourceId(v string) *DescribeVirtualResourceResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) SetSpotPriceLimit(v float32) *DescribeVirtualResourceResponseBodyResources {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeVirtualResourceResponseBodyResources) Validate() error {
	return dara.Validate(s)
}
