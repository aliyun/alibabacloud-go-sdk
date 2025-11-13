// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClustersNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) *DescribeDBInstanceNetInfoResponseBody
	GetDBClustersNetInfos() []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos
	SetDBInstanceNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody
	GetDBInstanceNetInfos() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos
	SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceNetInfoResponseBody struct {
	// The network information about the backend (BE) clusters.
	DBClustersNetInfos []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos `json:"DBClustersNetInfos,omitempty" xml:"DBClustersNetInfos,omitempty" type:"Repeated"`
	// The network information about the instance.
	DBInstanceNetInfos []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ADF42B18-43FD-5100-83A9-BE81AB70C863
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetDBClustersNetInfos() []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	return s.DBClustersNetInfos
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetDBInstanceNetInfos() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	return s.DBInstanceNetInfos
}

func (s *DescribeDBInstanceNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBClustersNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBClustersNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetDBInstanceNetInfos(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponseBody {
	s.DBInstanceNetInfos = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) SetRequestId(v string) *DescribeDBInstanceNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBody) Validate() error {
	if s.DBClustersNetInfos != nil {
		for _, item := range s.DBClustersNetInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DBInstanceNetInfos != nil {
		for _, item := range s.DBInstanceNetInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-****-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection string of the BE cluster.
	//
	// example:
	//
	// selectdb-cn-****-fe.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address of the BE cluster.
	//
	// example:
	//
	// 8.131.***.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The network type of the BE cluster.
	//
	// example:
	//
	// VPC/PUBLIC
	NetType  *string                                                            `json:"NetType,omitempty" xml:"NetType,omitempty"`
	PortList []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// Indicates whether the network information is visible to users.
	//
	// example:
	//
	// true/false
	UserVisible *bool `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// selectdb-cn-****-fe-20230816101006
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetIp() *string {
	return s.Ip
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetPortList() []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList {
	return s.PortList
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetUserVisible() *bool {
	return s.UserVisible
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetClusterId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.ClusterId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetIp(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.Ip = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetNetType(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.NetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetPortList(v []*DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.PortList = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetUserVisible(v bool) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.UserVisible = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVpcId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) SetVswitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfos) Validate() error {
	if s.PortList != nil {
		for _, item := range s.PortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList struct {
	// The port that is used to connect to the BE cluster.
	//
	// example:
	//
	// MySQLPort/HttpPort
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the port.
	//
	// example:
	//
	// 9030/8080
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) SetPort(v int32) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) SetProtocol(v string) *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList {
	s.Protocol = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBClustersNetInfosPortList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos struct {
	// The cluster ID.
	//
	// example:
	//
	// selectdb-cn-****-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The connection string of the instance.
	//
	// example:
	//
	// selectdb-cn-h033cnd****-fe.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: indicates a virtual private cloud (VPC)-connected instance.
	//
	// 	- **PUBLIC**: indicates an Internet-connected instance.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The ports.
	PortList []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// Indicates whether the network information is visible to users. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	UserVisible *bool `json:"UserVisible,omitempty" xml:"UserVisible,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-wz90scxq6ods388ft****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPC-connected instance.
	//
	// example:
	//
	// selectdb-cn-h033cnd****-fe-20230816101006
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6mlqti065rer6m0****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetIp() *string {
	return s.Ip
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetPortList() []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList {
	return s.PortList
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetUserVisible() *bool {
	return s.UserVisible
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetClusterId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.ClusterId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetIp(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.Ip = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetNetType(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.NetType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetPortList(v []*DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.PortList = v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetUserVisible(v bool) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.UserVisible = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVpcId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) SetVswitchId(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos {
	s.VswitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfos) Validate() error {
	if s.PortList != nil {
		for _, item := range s.PortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList struct {
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// MySQLPort
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol of the port. Valid values:
	//
	// 	- **HttpPort**: HTTP port.
	//
	// 	- **MySQLPort**: MySQL port.
	//
	// example:
	//
	// 9030
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) SetPort(v int32) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) SetProtocol(v string) *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList {
	s.Protocol = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseBodyDBInstanceNetInfosPortList) Validate() error {
	return dara.Validate(s)
}
