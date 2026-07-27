// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *ModifyInstanceConfigRequest
	GetBranchName() *string
	SetClientToken(v string) *ModifyInstanceConfigRequest
	GetClientToken() *string
	SetConfigName(v string) *ModifyInstanceConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *ModifyInstanceConfigRequest
	GetConfigValue() *string
	SetInstanceName(v string) *ModifyInstanceConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceConfigRequest
	GetRegionId() *string
}

type ModifyInstanceConfigRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The idempotency parameter.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the configuration item to modify. This parameter is used together with ConfigValue.
	//
	// example:
	//
	// eip、nat
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The value of the configuration item to modify. This parameter is used together with ConfigName.
	//
	// example:
	//
	// on、off
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *ModifyInstanceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyInstanceConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyInstanceConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceConfigRequest) SetBranchName(v string) *ModifyInstanceConfigRequest {
	s.BranchName = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetClientToken(v string) *ModifyInstanceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigName(v string) *ModifyInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetConfigValue(v string) *ModifyInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetInstanceName(v string) *ModifyInstanceConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetRegionId(v string) *ModifyInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
