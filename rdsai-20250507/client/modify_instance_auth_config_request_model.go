// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigList(v []*ModifyInstanceAuthConfigRequestConfigList) *ModifyInstanceAuthConfigRequest
	GetConfigList() []*ModifyInstanceAuthConfigRequestConfigList
	SetInstanceName(v string) *ModifyInstanceAuthConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceAuthConfigRequest
	GetRegionId() *string
}

type ModifyInstanceAuthConfigRequest struct {
	ConfigList []*ModifyInstanceAuthConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigRequest) GetConfigList() []*ModifyInstanceAuthConfigRequestConfigList {
	return s.ConfigList
}

func (s *ModifyInstanceAuthConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAuthConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAuthConfigRequest) SetConfigList(v []*ModifyInstanceAuthConfigRequestConfigList) *ModifyInstanceAuthConfigRequest {
	s.ConfigList = v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) SetInstanceName(v string) *ModifyInstanceAuthConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) SetRegionId(v string) *ModifyInstanceAuthConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequest) Validate() error {
	if s.ConfigList != nil {
		for _, item := range s.ConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceAuthConfigRequestConfigList struct {
	// example:
	//
	// GOTRUE_SITE_URL
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// http://8.152. XXX.XXX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceAuthConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigRequestConfigList) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceAuthConfigRequestConfigList) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceAuthConfigRequestConfigList) SetName(v string) *ModifyInstanceAuthConfigRequestConfigList {
	s.Name = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequestConfigList) SetValue(v string) *ModifyInstanceAuthConfigRequestConfigList {
	s.Value = &v
	return s
}

func (s *ModifyInstanceAuthConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
