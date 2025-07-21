// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateGatewayRequest
	GetChargeType() *string
	SetGatewayType(v string) *CreateGatewayRequest
	GetGatewayType() *string
	SetLogConfig(v *CreateGatewayRequestLogConfig) *CreateGatewayRequest
	GetLogConfig() *CreateGatewayRequestLogConfig
	SetName(v string) *CreateGatewayRequest
	GetName() *string
	SetNetworkAccessConfig(v *CreateGatewayRequestNetworkAccessConfig) *CreateGatewayRequest
	GetNetworkAccessConfig() *CreateGatewayRequestNetworkAccessConfig
	SetResourceGroupId(v string) *CreateGatewayRequest
	GetResourceGroupId() *string
	SetSpec(v string) *CreateGatewayRequest
	GetSpec() *string
	SetTag(v []*CreateGatewayRequestTag) *CreateGatewayRequest
	GetTag() []*CreateGatewayRequestTag
	SetVpcId(v string) *CreateGatewayRequest
	GetVpcId() *string
	SetZoneConfig(v *CreateGatewayRequestZoneConfig) *CreateGatewayRequest
	GetZoneConfig() *CreateGatewayRequestZoneConfig
}

type CreateGatewayRequest struct {
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// example:
	//
	// API
	GatewayType *string                        `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	LogConfig   *CreateGatewayRequestLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	// example:
	//
	// test-ceshi
	Name                *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	NetworkAccessConfig *CreateGatewayRequestNetworkAccessConfig `json:"networkAccessConfig,omitempty" xml:"networkAccessConfig,omitempty" type:"Struct"`
	// example:
	//
	// rg-ahr5uil8raz0rq3b
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// apigw.dev.x2
	Spec *string                    `json:"spec,omitempty" xml:"spec,omitempty"`
	Tag  []*CreateGatewayRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-zm0x16tomfiat1mk9f6rs
	VpcId      *string                         `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneConfig *CreateGatewayRequestZoneConfig `json:"zoneConfig,omitempty" xml:"zoneConfig,omitempty" type:"Struct"`
}

func (s CreateGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateGatewayRequest) GetLogConfig() *CreateGatewayRequestLogConfig {
	return s.LogConfig
}

func (s *CreateGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateGatewayRequest) GetNetworkAccessConfig() *CreateGatewayRequestNetworkAccessConfig {
	return s.NetworkAccessConfig
}

func (s *CreateGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGatewayRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateGatewayRequest) GetTag() []*CreateGatewayRequestTag {
	return s.Tag
}

func (s *CreateGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayRequest) GetZoneConfig() *CreateGatewayRequestZoneConfig {
	return s.ZoneConfig
}

func (s *CreateGatewayRequest) SetChargeType(v string) *CreateGatewayRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateGatewayRequest) SetGatewayType(v string) *CreateGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateGatewayRequest) SetLogConfig(v *CreateGatewayRequestLogConfig) *CreateGatewayRequest {
	s.LogConfig = v
	return s
}

func (s *CreateGatewayRequest) SetName(v string) *CreateGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateGatewayRequest) SetNetworkAccessConfig(v *CreateGatewayRequestNetworkAccessConfig) *CreateGatewayRequest {
	s.NetworkAccessConfig = v
	return s
}

func (s *CreateGatewayRequest) SetResourceGroupId(v string) *CreateGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGatewayRequest) SetSpec(v string) *CreateGatewayRequest {
	s.Spec = &v
	return s
}

func (s *CreateGatewayRequest) SetTag(v []*CreateGatewayRequestTag) *CreateGatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateGatewayRequest) SetVpcId(v string) *CreateGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayRequest) SetZoneConfig(v *CreateGatewayRequestZoneConfig) *CreateGatewayRequest {
	s.ZoneConfig = v
	return s
}

func (s *CreateGatewayRequest) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestLogConfig struct {
	Sls *CreateGatewayRequestLogConfigSls `json:"sls,omitempty" xml:"sls,omitempty" type:"Struct"`
}

func (s CreateGatewayRequestLogConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestLogConfig) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestLogConfig) GetSls() *CreateGatewayRequestLogConfigSls {
	return s.Sls
}

func (s *CreateGatewayRequestLogConfig) SetSls(v *CreateGatewayRequestLogConfigSls) *CreateGatewayRequestLogConfig {
	s.Sls = v
	return s
}

func (s *CreateGatewayRequestLogConfig) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestLogConfigSls struct {
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateGatewayRequestLogConfigSls) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestLogConfigSls) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestLogConfigSls) GetEnable() *bool {
	return s.Enable
}

func (s *CreateGatewayRequestLogConfigSls) SetEnable(v bool) *CreateGatewayRequestLogConfigSls {
	s.Enable = &v
	return s
}

func (s *CreateGatewayRequestLogConfigSls) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestNetworkAccessConfig struct {
	// example:
	//
	// Internet
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateGatewayRequestNetworkAccessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestNetworkAccessConfig) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestNetworkAccessConfig) GetType() *string {
	return s.Type
}

func (s *CreateGatewayRequestNetworkAccessConfig) SetType(v string) *CreateGatewayRequestNetworkAccessConfig {
	s.Type = &v
	return s
}

func (s *CreateGatewayRequestNetworkAccessConfig) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestTag struct {
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateGatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateGatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateGatewayRequestTag) SetKey(v string) *CreateGatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateGatewayRequestTag) SetValue(v string) *CreateGatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateGatewayRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestZoneConfig struct {
	// example:
	//
	// Manual
	SelectOption *string `json:"selectOption,omitempty" xml:"selectOption,omitempty"`
	// example:
	//
	// vsw-xxx
	VSwitchId *string                                `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	Zones     []*CreateGatewayRequestZoneConfigZones `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s CreateGatewayRequestZoneConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestZoneConfig) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestZoneConfig) GetSelectOption() *string {
	return s.SelectOption
}

func (s *CreateGatewayRequestZoneConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGatewayRequestZoneConfig) GetZones() []*CreateGatewayRequestZoneConfigZones {
	return s.Zones
}

func (s *CreateGatewayRequestZoneConfig) SetSelectOption(v string) *CreateGatewayRequestZoneConfig {
	s.SelectOption = &v
	return s
}

func (s *CreateGatewayRequestZoneConfig) SetVSwitchId(v string) *CreateGatewayRequestZoneConfig {
	s.VSwitchId = &v
	return s
}

func (s *CreateGatewayRequestZoneConfig) SetZones(v []*CreateGatewayRequestZoneConfigZones) *CreateGatewayRequestZoneConfig {
	s.Zones = v
	return s
}

func (s *CreateGatewayRequestZoneConfig) Validate() error {
	return dara.Validate(s)
}

type CreateGatewayRequestZoneConfigZones struct {
	// example:
	//
	// vsw-xx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateGatewayRequestZoneConfigZones) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequestZoneConfigZones) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequestZoneConfigZones) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGatewayRequestZoneConfigZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateGatewayRequestZoneConfigZones) SetVSwitchId(v string) *CreateGatewayRequestZoneConfigZones {
	s.VSwitchId = &v
	return s
}

func (s *CreateGatewayRequestZoneConfigZones) SetZoneId(v string) *CreateGatewayRequestZoneConfigZones {
	s.ZoneId = &v
	return s
}

func (s *CreateGatewayRequestZoneConfigZones) Validate() error {
	return dara.Validate(s)
}
