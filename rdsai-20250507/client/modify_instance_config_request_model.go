// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The ID of the RDS Supabase instance.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// eip、nat
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The name of the configuration item that you want to modify. Configure this parameter together with the ConfigValue parameter.
	//
	// example:
	//
	// on、off
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **ModifyInstanceConfig**.
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
