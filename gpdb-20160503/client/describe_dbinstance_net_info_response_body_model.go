// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody
	GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos
	SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody
	GetInstanceNetworkType() *string
	SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The connection information of the instance.
	DBInstanceNetInfos *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Struct"`
	// The network type of the instance. Valid values:
	//
	// 	- Classic: classic network.
	//
	// 	- VPC: VPC.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396F2CAD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetDBInstanceNetInfos() *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	return s.DBInstanceNetInfos
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) Validate() error {
	if s.DBInstanceNetInfos != nil {
		if err := s.DBInstanceNetInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetDBInstanceNetInfo() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	return s.DBInstanceNetInfo
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) Validate() error {
	if s.DBInstanceNetInfo != nil {
		for _, item := range s.DBInstanceNetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo struct {
	// The type of the endpoint.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The endpoint that is used to connect to the instance.
	//
	// example:
	//
	// gp-xxxxxxx.gpdb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 127.0.0.1
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The type of the IP address.
	//
	// 	- Valid values for instances in the classic network: Inner and Public.
	//
	// 	- Valid values for instances in a virtual private cloud (VPC): Private and Public.
	//
	// example:
	//
	// Inner
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The port number.
	//
	// example:
	//
	// 3432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The VPC ID of the instance.
	//
	// example:
	//
	// vpc-xxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the instance that is deployed in a VPC.
	//
	// example:
	//
	// vpc-xxxxxxx
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetAddressType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.AddressType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosDBInstanceNetInfo) Validate() error {
	return dara.Validate(s)
}
