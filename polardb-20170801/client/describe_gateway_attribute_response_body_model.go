// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClassCode(v string) *DescribeGatewayAttributeResponseBody
	GetClassCode() *string
	SetCreateTime(v string) *DescribeGatewayAttributeResponseBody
	GetCreateTime() *string
	SetCurrentVersion(v string) *DescribeGatewayAttributeResponseBody
	GetCurrentVersion() *string
	SetDbType(v string) *DescribeGatewayAttributeResponseBody
	GetDbType() *string
	SetEndpoints(v []*DescribeGatewayAttributeResponseBodyEndpoints) *DescribeGatewayAttributeResponseBody
	GetEndpoints() []*DescribeGatewayAttributeResponseBodyEndpoints
	SetExpireTime(v string) *DescribeGatewayAttributeResponseBody
	GetExpireTime() *string
	SetExpired(v bool) *DescribeGatewayAttributeResponseBody
	GetExpired() *bool
	SetGwClusterId(v string) *DescribeGatewayAttributeResponseBody
	GetGwClusterId() *string
	SetGwDescription(v string) *DescribeGatewayAttributeResponseBody
	GetGwDescription() *string
	SetLatestVersion(v string) *DescribeGatewayAttributeResponseBody
	GetLatestVersion() *string
	SetModifyTime(v string) *DescribeGatewayAttributeResponseBody
	GetModifyTime() *string
	SetPayType(v string) *DescribeGatewayAttributeResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeGatewayAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeGatewayAttributeResponseBody
	GetRequestId() *string
	SetRunningVersion(v string) *DescribeGatewayAttributeResponseBody
	GetRunningVersion() *string
	SetSecurityIPArrays(v []*DescribeGatewayAttributeResponseBodySecurityIPArrays) *DescribeGatewayAttributeResponseBody
	GetSecurityIPArrays() []*DescribeGatewayAttributeResponseBodySecurityIPArrays
	SetStatus(v string) *DescribeGatewayAttributeResponseBody
	GetStatus() *string
	SetVSwitchId(v string) *DescribeGatewayAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeGatewayAttributeResponseBody
	GetVpcId() *string
}

type DescribeGatewayAttributeResponseBody struct {
	// The instance type code.
	//
	// example:
	//
	// polar.app.g2.medium
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The creation time.
	//
	// The time is displayed in ISO 8601 format in UTC (YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2020-02-24T11:57:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current version number.
	//
	// example:
	//
	// 2.0.0
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The database type.
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The list of endpoints.
	Endpoints []*DescribeGatewayAttributeResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time.
	//
	// This parameter is empty when the billing method is Postpaid.
	//
	// The time is displayed in ISO 8601 format in UTC (YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2027-04-22T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// False
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The description of the gateway.
	//
	// example:
	//
	// xxx
	GwDescription *string `json:"GwDescription,omitempty" xml:"GwDescription,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// 2.0.1
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The last modification time.
	//
	// The time is displayed in ISO 8601 format in UTC (YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2025-10-07T02:19:55Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The billing method. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3E5CD764-FCCA-5C9C-838E-20E0DE84B2AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The running version number.
	//
	// example:
	//
	// 2.0.0
	RunningVersion *string `json:"RunningVersion,omitempty" xml:"RunningVersion,omitempty"`
	// The application-level IP whitelist.
	SecurityIPArrays []*DescribeGatewayAttributeResponseBodySecurityIPArrays `json:"SecurityIPArrays,omitempty" xml:"SecurityIPArrays,omitempty" type:"Repeated"`
	// The application status. Valid values:
	//
	// - CREATE: creating
	//
	// - ACTIVATION: running
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-*************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAttributeResponseBody) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeGatewayAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGatewayAttributeResponseBody) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeGatewayAttributeResponseBody) GetDbType() *string {
	return s.DbType
}

func (s *DescribeGatewayAttributeResponseBody) GetEndpoints() []*DescribeGatewayAttributeResponseBodyEndpoints {
	return s.Endpoints
}

func (s *DescribeGatewayAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGatewayAttributeResponseBody) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeGatewayAttributeResponseBody) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayAttributeResponseBody) GetGwDescription() *string {
	return s.GwDescription
}

func (s *DescribeGatewayAttributeResponseBody) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *DescribeGatewayAttributeResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeGatewayAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGatewayAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGatewayAttributeResponseBody) GetRunningVersion() *string {
	return s.RunningVersion
}

func (s *DescribeGatewayAttributeResponseBody) GetSecurityIPArrays() []*DescribeGatewayAttributeResponseBodySecurityIPArrays {
	return s.SecurityIPArrays
}

func (s *DescribeGatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeGatewayAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeGatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeGatewayAttributeResponseBody) SetClassCode(v string) *DescribeGatewayAttributeResponseBody {
	s.ClassCode = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetCreateTime(v string) *DescribeGatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetCurrentVersion(v string) *DescribeGatewayAttributeResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetDbType(v string) *DescribeGatewayAttributeResponseBody {
	s.DbType = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetEndpoints(v []*DescribeGatewayAttributeResponseBodyEndpoints) *DescribeGatewayAttributeResponseBody {
	s.Endpoints = v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetExpireTime(v string) *DescribeGatewayAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetExpired(v bool) *DescribeGatewayAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetGwClusterId(v string) *DescribeGatewayAttributeResponseBody {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetGwDescription(v string) *DescribeGatewayAttributeResponseBody {
	s.GwDescription = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetLatestVersion(v string) *DescribeGatewayAttributeResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetModifyTime(v string) *DescribeGatewayAttributeResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetPayType(v string) *DescribeGatewayAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetRegionId(v string) *DescribeGatewayAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetRequestId(v string) *DescribeGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetRunningVersion(v string) *DescribeGatewayAttributeResponseBody {
	s.RunningVersion = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetSecurityIPArrays(v []*DescribeGatewayAttributeResponseBodySecurityIPArrays) *DescribeGatewayAttributeResponseBody {
	s.SecurityIPArrays = v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetStatus(v string) *DescribeGatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetVSwitchId(v string) *DescribeGatewayAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) SetVpcId(v string) *DescribeGatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBody) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityIPArrays != nil {
		for _, item := range s.SecurityIPArrays {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGatewayAttributeResponseBodyEndpoints struct {
	// The endpoint address.
	//
	// example:
	//
	// pg-xxxxx.polardbaigateway.pre.rds.aliyuncs.com
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// xxx
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// - Private: VPC
	//
	// - Public: Internet
	//
	// example:
	//
	// Public
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// 1874631
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
	// The ID of the VPC to which the vSwitch belongs.
	//
	// example:
	//
	// vpc-*************
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeGatewayAttributeResponseBodyEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayAttributeResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetAddress() *string {
	return s.Address
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetNetType() *string {
	return s.NetType
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetPort() *string {
	return s.Port
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetTunnelId() *string {
	return s.TunnelId
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetAddress(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.Address = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetEndpointId(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.EndpointId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetGwClusterId(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetNetType(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.NetType = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetPort(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetTunnelId(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.TunnelId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) SetVpcId(v string) *DescribeGatewayAttributeResponseBodyEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodyEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeGatewayAttributeResponseBodySecurityIPArrays struct {
	// The name of the IP whitelist. Default value: `default`.
	//
	// example:
	//
	// default
	SecurityIPArrayName *string `json:"SecurityIPArrayName,omitempty" xml:"SecurityIPArrayName,omitempty"`
	// The tag of the IP whitelist.
	//
	// example:
	//
	// mytag
	SecurityIPArrayTag *string `json:"SecurityIPArrayTag,omitempty" xml:"SecurityIPArrayTag,omitempty"`
	// The IP addresses in the whitelist, separated by commas.
	//
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeGatewayAttributeResponseBodySecurityIPArrays) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayAttributeResponseBodySecurityIPArrays) GoString() string {
	return s.String()
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) GetSecurityIPArrayName() *string {
	return s.SecurityIPArrayName
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) GetSecurityIPArrayTag() *string {
	return s.SecurityIPArrayTag
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) SetSecurityIPArrayName(v string) *DescribeGatewayAttributeResponseBodySecurityIPArrays {
	s.SecurityIPArrayName = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) SetSecurityIPArrayTag(v string) *DescribeGatewayAttributeResponseBodySecurityIPArrays {
	s.SecurityIPArrayTag = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) SetSecurityIPList(v string) *DescribeGatewayAttributeResponseBodySecurityIPArrays {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeGatewayAttributeResponseBodySecurityIPArrays) Validate() error {
	return dara.Validate(s)
}
