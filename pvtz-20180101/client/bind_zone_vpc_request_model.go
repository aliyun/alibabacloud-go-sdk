// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindZoneVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *BindZoneVpcRequest
	GetClientToken() *string
	SetLang(v string) *BindZoneVpcRequest
	GetLang() *string
	SetUserClientIp(v string) *BindZoneVpcRequest
	GetUserClientIp() *string
	SetVpcs(v []*BindZoneVpcRequestVpcs) *BindZoneVpcRequest
	GetVpcs() []*BindZoneVpcRequestVpcs
	SetZoneId(v string) *BindZoneVpcRequest
	GetZoneId() *string
}

type BindZoneVpcRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The VPCs.
	//
	// >  If Vpcs is left empty, all VPCs that are associated with the zone are disassociated from the zone.
	Vpcs []*BindZoneVpcRequestVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
	// The zone ID. This ID uniquely identifies the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34d4031b63c527358b710a61346a****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s BindZoneVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s BindZoneVpcRequest) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BindZoneVpcRequest) GetLang() *string {
	return s.Lang
}

func (s *BindZoneVpcRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *BindZoneVpcRequest) GetVpcs() []*BindZoneVpcRequestVpcs {
	return s.Vpcs
}

func (s *BindZoneVpcRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *BindZoneVpcRequest) SetClientToken(v string) *BindZoneVpcRequest {
	s.ClientToken = &v
	return s
}

func (s *BindZoneVpcRequest) SetLang(v string) *BindZoneVpcRequest {
	s.Lang = &v
	return s
}

func (s *BindZoneVpcRequest) SetUserClientIp(v string) *BindZoneVpcRequest {
	s.UserClientIp = &v
	return s
}

func (s *BindZoneVpcRequest) SetVpcs(v []*BindZoneVpcRequestVpcs) *BindZoneVpcRequest {
	s.Vpcs = v
	return s
}

func (s *BindZoneVpcRequest) SetZoneId(v string) *BindZoneVpcRequest {
	s.ZoneId = &v
	return s
}

func (s *BindZoneVpcRequest) Validate() error {
	return dara.Validate(s)
}

type BindZoneVpcRequestVpcs struct {
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC ID. If the zone is already associated with VPCs and you do not specify this parameter, the associated VPCs are disassociated from the zone.
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC type. Valid values:
	//
	// 	- **STANDARD**: standard VPC
	//
	// 	- **EDS**: Elastic Desktop Service (EDS) workspace VPC
	//
	// example:
	//
	// STANDARD
	VpcType *string `json:"VpcType,omitempty" xml:"VpcType,omitempty"`
}

func (s BindZoneVpcRequestVpcs) String() string {
	return dara.Prettify(s)
}

func (s BindZoneVpcRequestVpcs) GoString() string {
	return s.String()
}

func (s *BindZoneVpcRequestVpcs) GetRegionId() *string {
	return s.RegionId
}

func (s *BindZoneVpcRequestVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *BindZoneVpcRequestVpcs) GetVpcType() *string {
	return s.VpcType
}

func (s *BindZoneVpcRequestVpcs) SetRegionId(v string) *BindZoneVpcRequestVpcs {
	s.RegionId = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) SetVpcId(v string) *BindZoneVpcRequestVpcs {
	s.VpcId = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) SetVpcType(v string) *BindZoneVpcRequestVpcs {
	s.VpcType = &v
	return s
}

func (s *BindZoneVpcRequestVpcs) Validate() error {
	return dara.Validate(s)
}
