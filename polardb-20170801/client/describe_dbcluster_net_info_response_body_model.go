// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody
	GetClusterNetworkType() *string
	SetDBClusterNetInfos(v *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) *DescribeDBClusterNetInfoResponseBody
	GetDBClusterNetInfos() *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos
	SetDBNodeNetInfos(v *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) *DescribeDBClusterNetInfoResponseBody
	GetDBNodeNetInfos() *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos
	SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBClusterNetInfoResponseBody struct {
	// example:
	//
	// VPC
	ClusterNetworkType *string                                                `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	DBClusterNetInfos  *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos `json:"DBClusterNetInfos,omitempty" xml:"DBClusterNetInfos,omitempty" type:"Struct"`
	DBNodeNetInfos     *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos    `json:"DBNodeNetInfos,omitempty" xml:"DBNodeNetInfos,omitempty" type:"Struct"`
	// example:
	//
	// 72D99256-ACF1-5F86-831F-8CB53E9C23ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBody) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *DescribeDBClusterNetInfoResponseBody) GetDBClusterNetInfos() *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos {
	return s.DBClusterNetInfos
}

func (s *DescribeDBClusterNetInfoResponseBody) GetDBNodeNetInfos() *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos {
	return s.DBNodeNetInfos
}

func (s *DescribeDBClusterNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetDBClusterNetInfos(v *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) *DescribeDBClusterNetInfoResponseBody {
	s.DBClusterNetInfos = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetDBNodeNetInfos(v *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) *DescribeDBClusterNetInfoResponseBody {
	s.DBNodeNetInfos = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) Validate() error {
	if s.DBClusterNetInfos != nil {
		if err := s.DBClusterNetInfos.Validate(); err != nil {
			return err
		}
	}
	if s.DBNodeNetInfos != nil {
		if err := s.DBNodeNetInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos struct {
	DBClusterNetInfo []*DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo `json:"DBClusterNetInfo,omitempty" xml:"DBClusterNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) GetDBClusterNetInfo() []*DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	return s.DBClusterNetInfo
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) SetDBClusterNetInfo(v []*DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos {
	s.DBClusterNetInfo = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfos) Validate() error {
	if s.DBClusterNetInfo != nil {
		for _, item := range s.DBClusterNetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo struct {
	// example:
	//
	// pc-****************.pg.polardb.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// MPP
	ConnectionStringType *string `json:"ConnectionStringType,omitempty" xml:"ConnectionStringType,omitempty"`
	// example:
	//
	// 121.*.*.173
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// example:
	//
	// Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// example:
	//
	// 1521
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// vpc-***************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-bp1li3eavsz8oaexq15dw
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetConnectionStringType() *string {
	return s.ConnectionStringType
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetConnectionStringType(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.ConnectionStringType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetIPType(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetPort(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetVPCId(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBClusterNetInfosDBClusterNetInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos struct {
	DBNodeNetInfo []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo `json:"DBNodeNetInfo,omitempty" xml:"DBNodeNetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) GetDBNodeNetInfo() []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo {
	return s.DBNodeNetInfo
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) SetDBNodeNetInfo(v []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos {
	s.DBNodeNetInfo = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfos) Validate() error {
	if s.DBNodeNetInfo != nil {
		for _, item := range s.DBNodeNetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo struct {
	// example:
	//
	// pc-*************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// Writer
	DBNodeRole *string                                                                  `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	NetInfos   *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos `json:"NetInfos,omitempty" xml:"NetInfos,omitempty" type:"Struct"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) GetNetInfos() *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos {
	return s.NetInfos
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) SetDBInstanceId(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) SetDBNodeRole(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) SetNetInfos(v *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo {
	s.NetInfos = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfo) Validate() error {
	if s.NetInfos != nil {
		if err := s.NetInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos struct {
	NetInfo []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo `json:"NetInfo,omitempty" xml:"NetInfo,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) GetNetInfo() []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	return s.NetInfo
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) SetNetInfo(v []*DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos {
	s.NetInfo = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfos) Validate() error {
	if s.NetInfo != nil {
		for _, item := range s.NetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo struct {
	// example:
	//
	// pc-****************.pg.polardb.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 47.*.*.203
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// example:
	//
	// Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// vpc-****************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetIPType() *string {
	return s.IPType
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetIPType(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetPort(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetVPCId(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyDBNodeNetInfosDBNodeNetInfoNetInfosNetInfo) Validate() error {
	return dara.Validate(s)
}
