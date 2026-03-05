// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody
	GetInstanceNetworkType() *string
	SetNetInfoItems(v *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) *DescribeDBInstanceNetInfoResponseBody
	GetNetInfoItems() *DescribeDBInstanceNetInfoResponseBodyNetInfoItems
	SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The network type. Valid values:
	//
	// 	- **CLASSIC**: The instance runs in a classic network.
	//
	// 	- **VPC**: The instance runs in a virtual private cloud (VPC).
	//
	// example:
	//
	// CLASSIC
	InstanceNetworkType *string                                            `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	NetInfoItems        *DescribeDBInstanceNetInfoResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FC77D4E1-2A7C-4F0B-A4CC-CE0B9C314B9B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetNetInfoItems() *DescribeDBInstanceNetInfoResponseBodyNetInfoItems {
	return s.NetInfoItems
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetNetInfoItems(v *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) *DescribeDBInstanceNetInfoResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) Validate() error {
	if s.NetInfoItems != nil {
		if err := s.NetInfoItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItems struct {
	InstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo `json:"InstanceNetInfo,omitempty" xml:"InstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) GetInstanceNetInfo() []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	return s.InstanceNetInfo
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) SetInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyNetInfoItems {
	s.InstanceNetInfo = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItems) Validate() error {
	if s.InstanceNetInfo != nil {
		for _, item := range s.InstanceNetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo struct {
	ConnectionString  *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	DirectConnection  *int32  `json:"DirectConnection,omitempty" xml:"DirectConnection,omitempty"`
	ExpiredTime       *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IPAddress         *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	IPType            *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	IsSlaveProxy      *int32  `json:"IsSlaveProxy,omitempty" xml:"IsSlaveProxy,omitempty"`
	Port              *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Upgradeable       *string `json:"Upgradeable,omitempty" xml:"Upgradeable,omitempty"`
	VPCId             *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VPCInstanceId     *string `json:"VPCInstanceId,omitempty" xml:"VPCInstanceId,omitempty"`
	VSwitchId         *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetDirectConnection() *int32 {
	return s.DirectConnection
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetIsSlaveProxy() *int32 {
	return s.IsSlaveProxy
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetUpgradeable() *string {
	return s.Upgradeable
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVPCInstanceId() *string {
	return s.VPCInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDBInstanceNetType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetDirectConnection(v int32) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.DirectConnection = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetExpiredTime(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetIsSlaveProxy(v int32) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.IsSlaveProxy = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetUpgradeable(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.Upgradeable = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVPCInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VPCInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyNetInfoItemsInstanceNetInfo) Validate() error {
	return dara.Validate(s)
}
