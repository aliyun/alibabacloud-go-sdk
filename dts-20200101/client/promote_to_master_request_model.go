// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromoteToMasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PromoteToMasterRequest
	GetInstanceId() *string
	SetMasterDbInstanceId(v string) *PromoteToMasterRequest
	GetMasterDbInstanceId() *string
	SetRegionId(v string) *PromoteToMasterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *PromoteToMasterRequest
	GetResourceGroupId() *string
	SetSlaveDbInstanceId(v string) *PromoteToMasterRequest
	GetSlaveDbInstanceId() *string
}

type PromoteToMasterRequest struct {
	// example:
	//
	// gad-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rm-bp1756****
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// rm-bp1756****
	SlaveDbInstanceId *string `json:"SlaveDbInstanceId,omitempty" xml:"SlaveDbInstanceId,omitempty"`
}

func (s PromoteToMasterRequest) String() string {
	return dara.Prettify(s)
}

func (s PromoteToMasterRequest) GoString() string {
	return s.String()
}

func (s *PromoteToMasterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PromoteToMasterRequest) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *PromoteToMasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PromoteToMasterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *PromoteToMasterRequest) GetSlaveDbInstanceId() *string {
	return s.SlaveDbInstanceId
}

func (s *PromoteToMasterRequest) SetInstanceId(v string) *PromoteToMasterRequest {
	s.InstanceId = &v
	return s
}

func (s *PromoteToMasterRequest) SetMasterDbInstanceId(v string) *PromoteToMasterRequest {
	s.MasterDbInstanceId = &v
	return s
}

func (s *PromoteToMasterRequest) SetRegionId(v string) *PromoteToMasterRequest {
	s.RegionId = &v
	return s
}

func (s *PromoteToMasterRequest) SetResourceGroupId(v string) *PromoteToMasterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *PromoteToMasterRequest) SetSlaveDbInstanceId(v string) *PromoteToMasterRequest {
	s.SlaveDbInstanceId = &v
	return s
}

func (s *PromoteToMasterRequest) Validate() error {
	return dara.Validate(s)
}
