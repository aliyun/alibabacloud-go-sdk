// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSpotProtectionPeriod(v bool) *CreateVirtualResourceRequest
	GetDisableSpotProtectionPeriod() *bool
	SetResources(v []*CreateVirtualResourceRequestResources) *CreateVirtualResourceRequest
	GetResources() []*CreateVirtualResourceRequestResources
	SetVirtualResourceName(v string) *CreateVirtualResourceRequest
	GetVirtualResourceName() *string
}

type CreateVirtualResourceRequest struct {
	// Specifies whether to disable the retention period of preemptible instances.
	//
	// example:
	//
	// true
	DisableSpotProtectionPeriod *bool `json:"DisableSpotProtectionPeriod,omitempty" xml:"DisableSpotProtectionPeriod,omitempty"`
	// The resources in the virtual resource group.
	Resources []*CreateVirtualResourceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The name of the virtual resource group. Default value: the ID of the virtual resource group.
	//
	// example:
	//
	// MyVirtualResource
	VirtualResourceName *string `json:"VirtualResourceName,omitempty" xml:"VirtualResourceName,omitempty"`
}

func (s CreateVirtualResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualResourceRequest) GetDisableSpotProtectionPeriod() *bool {
	return s.DisableSpotProtectionPeriod
}

func (s *CreateVirtualResourceRequest) GetResources() []*CreateVirtualResourceRequestResources {
	return s.Resources
}

func (s *CreateVirtualResourceRequest) GetVirtualResourceName() *string {
	return s.VirtualResourceName
}

func (s *CreateVirtualResourceRequest) SetDisableSpotProtectionPeriod(v bool) *CreateVirtualResourceRequest {
	s.DisableSpotProtectionPeriod = &v
	return s
}

func (s *CreateVirtualResourceRequest) SetResources(v []*CreateVirtualResourceRequestResources) *CreateVirtualResourceRequest {
	s.Resources = v
	return s
}

func (s *CreateVirtualResourceRequest) SetVirtualResourceName(v string) *CreateVirtualResourceRequest {
	s.VirtualResourceName = &v
	return s
}

func (s *CreateVirtualResourceRequest) Validate() error {
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

type CreateVirtualResourceRequestResources struct {
	// The instance type of the public resource group.
	//
	// >  You must specify one and only one of the InstanceType, ResourceId, and QuotaId parameters.
	//
	// example:
	//
	// ecs.s6-c1m2.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The priority of resource scheduling. A greater number indicates a higher priority.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the Lingjun resource quota.
	//
	// >  You must specify one and only one of the InstanceType, ResourceId, and QuotaId parameters.
	//
	// example:
	//
	// quota185lqxxxxxx
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The region in which the resource resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the dedicated resource group. For information about how to obtain the ID of a dedicated resource group, see [ListResources](https://help.aliyun.com/document_detail/412133.html).
	//
	// >  You must specify one and only one of the InstanceType, ResourceId, and QuotaId parameters.
	//
	// example:
	//
	// eas-r-g55ieatgg3buxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The maximum price of preemptible instances in a public resource group.
	//
	// >  If you leave this parameter empty, preemptible instances are not used.
	//
	// example:
	//
	// 10.05
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
}

func (s CreateVirtualResourceRequestResources) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualResourceRequestResources) GoString() string {
	return s.String()
}

func (s *CreateVirtualResourceRequestResources) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateVirtualResourceRequestResources) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateVirtualResourceRequestResources) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CreateVirtualResourceRequestResources) GetRegion() *string {
	return s.Region
}

func (s *CreateVirtualResourceRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateVirtualResourceRequestResources) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateVirtualResourceRequestResources) SetInstanceType(v string) *CreateVirtualResourceRequestResources {
	s.InstanceType = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) SetPriority(v int32) *CreateVirtualResourceRequestResources {
	s.Priority = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) SetQuotaId(v string) *CreateVirtualResourceRequestResources {
	s.QuotaId = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) SetRegion(v string) *CreateVirtualResourceRequestResources {
	s.Region = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) SetResourceId(v string) *CreateVirtualResourceRequestResources {
	s.ResourceId = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) SetSpotPriceLimit(v float32) *CreateVirtualResourceRequestResources {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateVirtualResourceRequestResources) Validate() error {
	return dara.Validate(s)
}
