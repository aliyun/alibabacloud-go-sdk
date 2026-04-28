// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateGatewayRequest
	GetAutoRenew() *bool
	SetDBClusterClass(v string) *CreateGatewayRequest
	GetDBClusterClass() *string
	SetDBType(v string) *CreateGatewayRequest
	GetDBType() *string
	SetPayType(v string) *CreateGatewayRequest
	GetPayType() *string
	SetPeriod(v string) *CreateGatewayRequest
	GetPeriod() *string
	SetRegionId(v string) *CreateGatewayRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *CreateGatewayRequest
	GetSecurityGroupId() *string
	SetUsedTime(v string) *CreateGatewayRequest
	GetUsedTime() *string
	SetVPCId(v string) *CreateGatewayRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateGatewayRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateGatewayRequest
	GetZoneId() *string
}

type CreateGatewayRequest struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// polar.app.g2.medium
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// sg-bp**************
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 3
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-*******************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-beijing-l
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateGatewayRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *CreateGatewayRequest) GetDBType() *string {
	return s.DBType
}

func (s *CreateGatewayRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateGatewayRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGatewayRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateGatewayRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateGatewayRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateGatewayRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateGatewayRequest) SetAutoRenew(v bool) *CreateGatewayRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGatewayRequest) SetDBClusterClass(v string) *CreateGatewayRequest {
	s.DBClusterClass = &v
	return s
}

func (s *CreateGatewayRequest) SetDBType(v string) *CreateGatewayRequest {
	s.DBType = &v
	return s
}

func (s *CreateGatewayRequest) SetPayType(v string) *CreateGatewayRequest {
	s.PayType = &v
	return s
}

func (s *CreateGatewayRequest) SetPeriod(v string) *CreateGatewayRequest {
	s.Period = &v
	return s
}

func (s *CreateGatewayRequest) SetRegionId(v string) *CreateGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGatewayRequest) SetSecurityGroupId(v string) *CreateGatewayRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateGatewayRequest) SetUsedTime(v string) *CreateGatewayRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateGatewayRequest) SetVPCId(v string) *CreateGatewayRequest {
	s.VPCId = &v
	return s
}

func (s *CreateGatewayRequest) SetVSwitchId(v string) *CreateGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateGatewayRequest) SetZoneId(v string) *CreateGatewayRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateGatewayRequest) Validate() error {
	return dara.Validate(s)
}
