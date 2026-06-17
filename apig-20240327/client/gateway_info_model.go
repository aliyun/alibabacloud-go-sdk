// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayInfo interface {
	dara.Model
	String() string
	GoString() string
	SetEngineVersion(v string) *GatewayInfo
	GetEngineVersion() *string
	SetGatewayEdition(v string) *GatewayInfo
	GetGatewayEdition() *string
	SetGatewayId(v string) *GatewayInfo
	GetGatewayId() *string
	SetName(v string) *GatewayInfo
	GetName() *string
	SetVpcInfo(v *GatewayInfoVpcInfo) *GatewayInfo
	GetVpcInfo() *GatewayInfoVpcInfo
}

type GatewayInfo struct {
	// The gateway DPI engine database engine version.
	//
	// example:
	//
	// 2.0.7
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// MultiTenantServerless
	GatewayEdition *string `json:"gatewayEdition,omitempty" xml:"gatewayEdition,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-cq7og15lhtgi6qasrj60
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The gateway name.
	//
	// example:
	//
	// apitest-gw
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The VPC information.
	VpcInfo *GatewayInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s GatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s GatewayInfo) GoString() string {
	return s.String()
}

func (s *GatewayInfo) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GatewayInfo) GetGatewayEdition() *string {
	return s.GatewayEdition
}

func (s *GatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GatewayInfo) GetName() *string {
	return s.Name
}

func (s *GatewayInfo) GetVpcInfo() *GatewayInfoVpcInfo {
	return s.VpcInfo
}

func (s *GatewayInfo) SetEngineVersion(v string) *GatewayInfo {
	s.EngineVersion = &v
	return s
}

func (s *GatewayInfo) SetGatewayEdition(v string) *GatewayInfo {
	s.GatewayEdition = &v
	return s
}

func (s *GatewayInfo) SetGatewayId(v string) *GatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *GatewayInfo) SetName(v string) *GatewayInfo {
	s.Name = &v
	return s
}

func (s *GatewayInfo) SetVpcInfo(v *GatewayInfoVpcInfo) *GatewayInfo {
	s.VpcInfo = v
	return s
}

func (s *GatewayInfo) Validate() error {
	if s.VpcInfo != nil {
		if err := s.VpcInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GatewayInfoVpcInfo struct {
	// The VPC name.
	//
	// example:
	//
	// 商品中心预发环境专有网络
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf664nyle5kh***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GatewayInfoVpcInfo) String() string {
	return dara.Prettify(s)
}

func (s GatewayInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *GatewayInfoVpcInfo) GetName() *string {
	return s.Name
}

func (s *GatewayInfoVpcInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *GatewayInfoVpcInfo) SetName(v string) *GatewayInfoVpcInfo {
	s.Name = &v
	return s
}

func (s *GatewayInfoVpcInfo) SetVpcId(v string) *GatewayInfoVpcInfo {
	s.VpcId = &v
	return s
}

func (s *GatewayInfoVpcInfo) Validate() error {
	return dara.Validate(s)
}
