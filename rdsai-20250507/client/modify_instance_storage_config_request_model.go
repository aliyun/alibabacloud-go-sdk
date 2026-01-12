// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceStorageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceStorageConfigRequest
	GetClientToken() *string
	SetConfigList(v []*ModifyInstanceStorageConfigRequestConfigList) *ModifyInstanceStorageConfigRequest
	GetConfigList() []*ModifyInstanceStorageConfigRequestConfigList
	SetInstanceName(v string) *ModifyInstanceStorageConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceStorageConfigRequest
	GetRegionId() *string
}

type ModifyInstanceStorageConfigRequest struct {
	// The value of the configuration item.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the RDS Supabase instance.
	ConfigList []*ModifyInstanceStorageConfigRequestConfigList `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **ModifyInstanceStorageConfig**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceStorageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceStorageConfigRequest) GetConfigList() []*ModifyInstanceStorageConfigRequestConfigList {
	return s.ConfigList
}

func (s *ModifyInstanceStorageConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceStorageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceStorageConfigRequest) SetClientToken(v string) *ModifyInstanceStorageConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetConfigList(v []*ModifyInstanceStorageConfigRequestConfigList) *ModifyInstanceStorageConfigRequest {
	s.ConfigList = v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetInstanceName(v string) *ModifyInstanceStorageConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) SetRegionId(v string) *ModifyInstanceStorageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequest) Validate() error {
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

type ModifyInstanceStorageConfigRequestConfigList struct {
	// example:
	//
	// TENANT_ID
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// test-prefix
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyInstanceStorageConfigRequestConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigRequestConfigList) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigRequestConfigList) GetName() *string {
	return s.Name
}

func (s *ModifyInstanceStorageConfigRequestConfigList) GetValue() *string {
	return s.Value
}

func (s *ModifyInstanceStorageConfigRequestConfigList) SetName(v string) *ModifyInstanceStorageConfigRequestConfigList {
	s.Name = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequestConfigList) SetValue(v string) *ModifyInstanceStorageConfigRequestConfigList {
	s.Value = &v
	return s
}

func (s *ModifyInstanceStorageConfigRequestConfigList) Validate() error {
	return dara.Validate(s)
}
