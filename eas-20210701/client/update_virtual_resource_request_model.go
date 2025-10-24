// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSpotProtectionPeriod(v bool) *UpdateVirtualResourceRequest
	GetDisableSpotProtectionPeriod() *bool
	SetResources(v []*UpdateVirtualResourceRequestResources) *UpdateVirtualResourceRequest
	GetResources() []*UpdateVirtualResourceRequestResources
	SetVirtualResourceName(v string) *UpdateVirtualResourceRequest
	GetVirtualResourceName() *string
}

type UpdateVirtualResourceRequest struct {
	// Specifies whether to disable the retention period of preemptible instances.
	//
	// example:
	//
	// true
	DisableSpotProtectionPeriod *bool `json:"DisableSpotProtectionPeriod,omitempty" xml:"DisableSpotProtectionPeriod,omitempty"`
	// The resources in the virtual resource group.
	//
	// >  If you specify this parameter, previous data is overwritten.
	Resources []*UpdateVirtualResourceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The new name of the virtual resource group.
	//
	// example:
	//
	// NewMyVirtualResource
	VirtualResourceName *string `json:"VirtualResourceName,omitempty" xml:"VirtualResourceName,omitempty"`
}

func (s UpdateVirtualResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualResourceRequest) GetDisableSpotProtectionPeriod() *bool {
	return s.DisableSpotProtectionPeriod
}

func (s *UpdateVirtualResourceRequest) GetResources() []*UpdateVirtualResourceRequestResources {
	return s.Resources
}

func (s *UpdateVirtualResourceRequest) GetVirtualResourceName() *string {
	return s.VirtualResourceName
}

func (s *UpdateVirtualResourceRequest) SetDisableSpotProtectionPeriod(v bool) *UpdateVirtualResourceRequest {
	s.DisableSpotProtectionPeriod = &v
	return s
}

func (s *UpdateVirtualResourceRequest) SetResources(v []*UpdateVirtualResourceRequestResources) *UpdateVirtualResourceRequest {
	s.Resources = v
	return s
}

func (s *UpdateVirtualResourceRequest) SetVirtualResourceName(v string) *UpdateVirtualResourceRequest {
	s.VirtualResourceName = &v
	return s
}

func (s *UpdateVirtualResourceRequest) Validate() error {
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

type UpdateVirtualResourceRequestResources struct {
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
	// quota185lqf994k6
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
	// eas-r-g55ieatgg3butwrn7a
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

func (s UpdateVirtualResourceRequestResources) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualResourceRequestResources) GoString() string {
	return s.String()
}

func (s *UpdateVirtualResourceRequestResources) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateVirtualResourceRequestResources) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateVirtualResourceRequestResources) GetQuotaId() *string {
	return s.QuotaId
}

func (s *UpdateVirtualResourceRequestResources) GetRegion() *string {
	return s.Region
}

func (s *UpdateVirtualResourceRequestResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateVirtualResourceRequestResources) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *UpdateVirtualResourceRequestResources) SetInstanceType(v string) *UpdateVirtualResourceRequestResources {
	s.InstanceType = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) SetPriority(v int32) *UpdateVirtualResourceRequestResources {
	s.Priority = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) SetQuotaId(v string) *UpdateVirtualResourceRequestResources {
	s.QuotaId = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) SetRegion(v string) *UpdateVirtualResourceRequestResources {
	s.Region = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) SetResourceId(v string) *UpdateVirtualResourceRequestResources {
	s.ResourceId = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) SetSpotPriceLimit(v float32) *UpdateVirtualResourceRequestResources {
	s.SpotPriceLimit = &v
	return s
}

func (s *UpdateVirtualResourceRequestResources) Validate() error {
	return dara.Validate(s)
}
