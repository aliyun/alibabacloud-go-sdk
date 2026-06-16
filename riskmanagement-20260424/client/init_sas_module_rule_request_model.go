// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitSasModuleRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBind(v int32) *InitSasModuleRuleRequest
	GetAutoBind() *int32
	SetInstances(v []*InitSasModuleRuleRequestInstances) *InitSasModuleRuleRequest
	GetInstances() []*InitSasModuleRuleRequestInstances
	SetIsTrial(v bool) *InitSasModuleRuleRequest
	GetIsTrial() *bool
	SetRegionId(v string) *InitSasModuleRuleRequest
	GetRegionId() *string
}

type InitSasModuleRuleRequest struct {
	AutoBind  *int32                               `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	Instances []*InitSasModuleRuleRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	IsTrial   *bool                                `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	RegionId  *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitSasModuleRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleRequest) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *InitSasModuleRuleRequest) GetInstances() []*InitSasModuleRuleRequestInstances {
	return s.Instances
}

func (s *InitSasModuleRuleRequest) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *InitSasModuleRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitSasModuleRuleRequest) SetAutoBind(v int32) *InitSasModuleRuleRequest {
	s.AutoBind = &v
	return s
}

func (s *InitSasModuleRuleRequest) SetInstances(v []*InitSasModuleRuleRequestInstances) *InitSasModuleRuleRequest {
	s.Instances = v
	return s
}

func (s *InitSasModuleRuleRequest) SetIsTrial(v bool) *InitSasModuleRuleRequest {
	s.IsTrial = &v
	return s
}

func (s *InitSasModuleRuleRequest) SetRegionId(v string) *InitSasModuleRuleRequest {
	s.RegionId = &v
	return s
}

func (s *InitSasModuleRuleRequest) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InitSasModuleRuleRequestInstances struct {
	Cores      *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s InitSasModuleRuleRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleRequestInstances) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleRequestInstances) GetCores() *string {
	return s.Cores
}

func (s *InitSasModuleRuleRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitSasModuleRuleRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *InitSasModuleRuleRequestInstances) GetUuid() *string {
	return s.Uuid
}

func (s *InitSasModuleRuleRequestInstances) SetCores(v string) *InitSasModuleRuleRequestInstances {
	s.Cores = &v
	return s
}

func (s *InitSasModuleRuleRequestInstances) SetInstanceId(v string) *InitSasModuleRuleRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *InitSasModuleRuleRequestInstances) SetRegionId(v string) *InitSasModuleRuleRequestInstances {
	s.RegionId = &v
	return s
}

func (s *InitSasModuleRuleRequestInstances) SetUuid(v string) *InitSasModuleRuleRequestInstances {
	s.Uuid = &v
	return s
}

func (s *InitSasModuleRuleRequestInstances) Validate() error {
	return dara.Validate(s)
}
