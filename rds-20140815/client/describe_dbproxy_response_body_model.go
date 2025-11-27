// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBProxyAVZones(v *DescribeDBProxyResponseBodyDBProxyAVZones) *DescribeDBProxyResponseBody
	GetDBProxyAVZones() *DescribeDBProxyResponseBodyDBProxyAVZones
	SetDBProxyConnectStringItems(v *DescribeDBProxyResponseBodyDBProxyConnectStringItems) *DescribeDBProxyResponseBody
	GetDBProxyConnectStringItems() *DescribeDBProxyResponseBodyDBProxyConnectStringItems
	SetDBProxyEngineType(v string) *DescribeDBProxyResponseBody
	GetDBProxyEngineType() *string
	SetDBProxyInstanceCurrentMinorVersion(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceCurrentMinorVersion() *string
	SetDBProxyInstanceLatestMinorVersion(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceLatestMinorVersion() *string
	SetDBProxyInstanceMinorVersions(v *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) *DescribeDBProxyResponseBody
	GetDBProxyInstanceMinorVersions() *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions
	SetDBProxyInstanceName(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceName() *string
	SetDBProxyInstanceNum(v int32) *DescribeDBProxyResponseBody
	GetDBProxyInstanceNum() *int32
	SetDBProxyInstanceSize(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceSize() *string
	SetDBProxyInstanceStatus(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceStatus() *string
	SetDBProxyInstanceType(v string) *DescribeDBProxyResponseBody
	GetDBProxyInstanceType() *string
	SetDBProxyKindCode(v string) *DescribeDBProxyResponseBody
	GetDBProxyKindCode() *string
	SetDBProxyNodes(v *DescribeDBProxyResponseBodyDBProxyNodes) *DescribeDBProxyResponseBody
	GetDBProxyNodes() *DescribeDBProxyResponseBodyDBProxyNodes
	SetDBProxyPersistentConnectionStatus(v string) *DescribeDBProxyResponseBody
	GetDBProxyPersistentConnectionStatus() *string
	SetDBProxyServiceStatus(v string) *DescribeDBProxyResponseBody
	GetDBProxyServiceStatus() *string
	SetDbProxyEndpointItems(v *DescribeDBProxyResponseBodyDbProxyEndpointItems) *DescribeDBProxyResponseBody
	GetDbProxyEndpointItems() *DescribeDBProxyResponseBodyDbProxyEndpointItems
	SetRequestId(v string) *DescribeDBProxyResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeDBProxyResponseBody
	GetResourceGroupId() *string
}

type DescribeDBProxyResponseBody struct {
	// The list of zones that are available for the database proxy.
	DBProxyAVZones *DescribeDBProxyResponseBodyDBProxyAVZones `json:"DBProxyAVZones,omitempty" xml:"DBProxyAVZones,omitempty" type:"Struct"`
	// An array consisting of the information about the database proxy endpoint that is created for the instance.
	DBProxyConnectStringItems *DescribeDBProxyResponseBodyDBProxyConnectStringItems `json:"DBProxyConnectStringItems,omitempty" xml:"DBProxyConnectStringItems,omitempty" type:"Struct"`
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The version of the proxy instance.
	//
	// example:
	//
	// 1.13.11
	DBProxyInstanceCurrentMinorVersion *string `json:"DBProxyInstanceCurrentMinorVersion,omitempty" xml:"DBProxyInstanceCurrentMinorVersion,omitempty"`
	// The latest version that is available for the proxy instance.
	//
	// example:
	//
	// 1.13.12
	DBProxyInstanceLatestMinorVersion *string                                                  `json:"DBProxyInstanceLatestMinorVersion,omitempty" xml:"DBProxyInstanceLatestMinorVersion,omitempty"`
	DBProxyInstanceMinorVersions      *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions `json:"DBProxyInstanceMinorVersions,omitempty" xml:"DBProxyInstanceMinorVersions,omitempty" type:"Struct"`
	// The name of the proxy instance.
	//
	// example:
	//
	// gos787jog2wk0ye1****
	DBProxyInstanceName *string `json:"DBProxyInstanceName,omitempty" xml:"DBProxyInstanceName,omitempty"`
	// The number of proxies that are enabled on the instance.
	//
	// example:
	//
	// 1
	DBProxyInstanceNum *int32 `json:"DBProxyInstanceNum,omitempty" xml:"DBProxyInstanceNum,omitempty"`
	// This parameter is available only for ApsaraDB RDS for PostgreSQL instances. The specifications of the proxy instance that is enabled.
	//
	// Format: `Number of cores/Memory capacity`.
	//
	// For example, a value of 4/8 indicates that the proxy instance has 4 cores and 8 GB of memory.
	//
	// example:
	//
	// 4/8
	DBProxyInstanceSize *string `json:"DBProxyInstanceSize,omitempty" xml:"DBProxyInstanceSize,omitempty"`
	// The status of the proxy instance.
	//
	// 	- DBInstanceClassChanging: The specifications of the proxy instance are being changed.
	//
	// 	- Creating: The proxy instance is being created.
	//
	// 	- Running: The proxy instance is running.
	//
	// 	- Deleting: The proxy instance is being deleted.
	//
	// example:
	//
	// Running
	DBProxyInstanceStatus *string `json:"DBProxyInstanceStatus,omitempty" xml:"DBProxyInstanceStatus,omitempty"`
	// The type of the database proxy that is enabled on the instance. Valid values:
	//
	// 	- 1: shared database proxy
	//
	// 	- 2: dedicated database proxy
	//
	// 	- 3: general-purpose database proxy
	//
	// >  ApsaraDB RDS for PostgreSQL does not support shared database proxies.
	//
	// example:
	//
	// 2
	DBProxyInstanceType *string `json:"DBProxyInstanceType,omitempty" xml:"DBProxyInstanceType,omitempty"`
	// An internal parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// 18
	DBProxyKindCode *string `json:"DBProxyKindCode,omitempty" xml:"DBProxyKindCode,omitempty"`
	// The proxy nodes.
	DBProxyNodes *DescribeDBProxyResponseBodyDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Struct"`
	// The status of persistence connections. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// 	- **Unsupported**
	//
	// example:
	//
	// Disabled
	DBProxyPersistentConnectionStatus *string `json:"DBProxyPersistentConnectionStatus,omitempty" xml:"DBProxyPersistentConnectionStatus,omitempty"`
	// The status of the database proxy.
	//
	// 	- Shutdown: disabled
	//
	// 	- Startup: enabled
	//
	// example:
	//
	// Startup
	DBProxyServiceStatus *string `json:"DBProxyServiceStatus,omitempty" xml:"DBProxyServiceStatus,omitempty"`
	// The proxy terminals of the instance.
	DbProxyEndpointItems *DescribeDBProxyResponseBodyDbProxyEndpointItems `json:"DbProxyEndpointItems,omitempty" xml:"DbProxyEndpointItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 909A69EE-71C8-4417-A0B9-FF085407E1E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBody) GetDBProxyAVZones() *DescribeDBProxyResponseBodyDBProxyAVZones {
	return s.DBProxyAVZones
}

func (s *DescribeDBProxyResponseBody) GetDBProxyConnectStringItems() *DescribeDBProxyResponseBodyDBProxyConnectStringItems {
	return s.DBProxyConnectStringItems
}

func (s *DescribeDBProxyResponseBody) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceCurrentMinorVersion() *string {
	return s.DBProxyInstanceCurrentMinorVersion
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceLatestMinorVersion() *string {
	return s.DBProxyInstanceLatestMinorVersion
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceMinorVersions() *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions {
	return s.DBProxyInstanceMinorVersions
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceName() *string {
	return s.DBProxyInstanceName
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceNum() *int32 {
	return s.DBProxyInstanceNum
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceSize() *string {
	return s.DBProxyInstanceSize
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceStatus() *string {
	return s.DBProxyInstanceStatus
}

func (s *DescribeDBProxyResponseBody) GetDBProxyInstanceType() *string {
	return s.DBProxyInstanceType
}

func (s *DescribeDBProxyResponseBody) GetDBProxyKindCode() *string {
	return s.DBProxyKindCode
}

func (s *DescribeDBProxyResponseBody) GetDBProxyNodes() *DescribeDBProxyResponseBodyDBProxyNodes {
	return s.DBProxyNodes
}

func (s *DescribeDBProxyResponseBody) GetDBProxyPersistentConnectionStatus() *string {
	return s.DBProxyPersistentConnectionStatus
}

func (s *DescribeDBProxyResponseBody) GetDBProxyServiceStatus() *string {
	return s.DBProxyServiceStatus
}

func (s *DescribeDBProxyResponseBody) GetDbProxyEndpointItems() *DescribeDBProxyResponseBodyDbProxyEndpointItems {
	return s.DbProxyEndpointItems
}

func (s *DescribeDBProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBProxyResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBProxyResponseBody) SetDBProxyAVZones(v *DescribeDBProxyResponseBodyDBProxyAVZones) *DescribeDBProxyResponseBody {
	s.DBProxyAVZones = v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyConnectStringItems(v *DescribeDBProxyResponseBodyDBProxyConnectStringItems) *DescribeDBProxyResponseBody {
	s.DBProxyConnectStringItems = v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyEngineType(v string) *DescribeDBProxyResponseBody {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceCurrentMinorVersion(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceCurrentMinorVersion = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceLatestMinorVersion(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceLatestMinorVersion = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceMinorVersions(v *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceMinorVersions = v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceName(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceName = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceNum(v int32) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceNum = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceSize(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceSize = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceStatus(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceStatus = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyInstanceType(v string) *DescribeDBProxyResponseBody {
	s.DBProxyInstanceType = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyKindCode(v string) *DescribeDBProxyResponseBody {
	s.DBProxyKindCode = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyNodes(v *DescribeDBProxyResponseBodyDBProxyNodes) *DescribeDBProxyResponseBody {
	s.DBProxyNodes = v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyPersistentConnectionStatus(v string) *DescribeDBProxyResponseBody {
	s.DBProxyPersistentConnectionStatus = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDBProxyServiceStatus(v string) *DescribeDBProxyResponseBody {
	s.DBProxyServiceStatus = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetDbProxyEndpointItems(v *DescribeDBProxyResponseBodyDbProxyEndpointItems) *DescribeDBProxyResponseBody {
	s.DbProxyEndpointItems = v
	return s
}

func (s *DescribeDBProxyResponseBody) SetRequestId(v string) *DescribeDBProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBProxyResponseBody) SetResourceGroupId(v string) *DescribeDBProxyResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBProxyResponseBody) Validate() error {
	if s.DBProxyAVZones != nil {
		if err := s.DBProxyAVZones.Validate(); err != nil {
			return err
		}
	}
	if s.DBProxyConnectStringItems != nil {
		if err := s.DBProxyConnectStringItems.Validate(); err != nil {
			return err
		}
	}
	if s.DBProxyInstanceMinorVersions != nil {
		if err := s.DBProxyInstanceMinorVersions.Validate(); err != nil {
			return err
		}
	}
	if s.DBProxyNodes != nil {
		if err := s.DBProxyNodes.Validate(); err != nil {
			return err
		}
	}
	if s.DbProxyEndpointItems != nil {
		if err := s.DbProxyEndpointItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBProxyResponseBodyDBProxyAVZones struct {
	DBProxyAVZones []*string `json:"DBProxyAVZones,omitempty" xml:"DBProxyAVZones,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyResponseBodyDBProxyAVZones) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyAVZones) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyAVZones) GetDBProxyAVZones() []*string {
	return s.DBProxyAVZones
}

func (s *DescribeDBProxyResponseBodyDBProxyAVZones) SetDBProxyAVZones(v []*string) *DescribeDBProxyResponseBodyDBProxyAVZones {
	s.DBProxyAVZones = v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyAVZones) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyResponseBodyDBProxyConnectStringItems struct {
	DBProxyConnectStringItems []*DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems `json:"DBProxyConnectStringItems,omitempty" xml:"DBProxyConnectStringItems,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyResponseBodyDBProxyConnectStringItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyConnectStringItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItems) GetDBProxyConnectStringItems() []*DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	return s.DBProxyConnectStringItems
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItems) SetDBProxyConnectStringItems(v []*DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) *DescribeDBProxyResponseBodyDBProxyConnectStringItems {
	s.DBProxyConnectStringItems = v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItems) Validate() error {
	if s.DBProxyConnectStringItems != nil {
		for _, item := range s.DBProxyConnectStringItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems struct {
	// The database proxy endpoint.
	//
	// example:
	//
	// gos787jog2wk0ye1****-rw4rm.rwlb.rds.aliyuncs.com
	DBProxyConnectString *string `json:"DBProxyConnectString,omitempty" xml:"DBProxyConnectString,omitempty"`
	// The network type of the database proxy endpoint. A database proxy endpoint is formerly referred to as a proxy terminal. Valid values:
	//
	// 	- OuterString: Internet
	//
	// 	- InnerString: internal network
	//
	// example:
	//
	// InnerString
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// The network type of the database proxy. Valid values:
	//
	// 	- 0: Internet
	//
	// 	- 1: classic network
	//
	// 	- 2: virtual private cloud (VPC)
	//
	// example:
	//
	// 2
	DBProxyConnectStringNetWorkType *string `json:"DBProxyConnectStringNetWorkType,omitempty" xml:"DBProxyConnectStringNetWorkType,omitempty"`
	// The port that is associated with the database proxy endpoint.
	//
	// example:
	//
	// 3306
	DBProxyConnectStringPort *string `json:"DBProxyConnectStringPort,omitempty" xml:"DBProxyConnectStringPort,omitempty"`
	// The ID of the backend database proxy endpoint.
	//
	// example:
	//
	// 20****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// The name of the database proxy endpoint. The name can be replaced by the ID of the database proxy endpoint.
	//
	// example:
	//
	// gos787jog2wk0ye1****
	DBProxyEndpointName *string `json:"DBProxyEndpointName,omitempty" xml:"DBProxyEndpointName,omitempty"`
	// The VPC of the database proxy.
	//
	// example:
	//
	// vpc-uf6oobt****
	DBProxyVpcId *string `json:"DBProxyVpcId,omitempty" xml:"DBProxyVpcId,omitempty"`
	// The ID of the database proxy instance.
	//
	// example:
	//
	// rm-bp145737x5****131161274792****
	DBProxyVpcInstanceId *string `json:"DBProxyVpcInstanceId,omitempty" xml:"DBProxyVpcInstanceId,omitempty"`
	// The vSwitch of the database proxy.
	//
	// example:
	//
	// vsw-uf6l0pic17****
	DBProxyVswitchId *string `json:"DBProxyVswitchId,omitempty" xml:"DBProxyVswitchId,omitempty"`
}

func (s DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyConnectString() *string {
	return s.DBProxyConnectString
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyConnectStringNetWorkType() *string {
	return s.DBProxyConnectStringNetWorkType
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyConnectStringPort() *string {
	return s.DBProxyConnectStringPort
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyEndpointName() *string {
	return s.DBProxyEndpointName
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyVpcId() *string {
	return s.DBProxyVpcId
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyVpcInstanceId() *string {
	return s.DBProxyVpcInstanceId
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) GetDBProxyVswitchId() *string {
	return s.DBProxyVswitchId
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyConnectString(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyConnectString = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyConnectStringNetType(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyConnectStringNetWorkType(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyConnectStringNetWorkType = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyConnectStringPort(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyConnectStringPort = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyEndpointId(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyEndpointId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyEndpointName(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyEndpointName = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyVpcId(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyVpcId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyVpcInstanceId(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyVpcInstanceId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) SetDBProxyVswitchId(v string) *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems {
	s.DBProxyVswitchId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyConnectStringItemsDBProxyConnectStringItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions struct {
	DBProxyInstanceMinorVersions []*string `json:"DBProxyInstanceMinorVersions,omitempty" xml:"DBProxyInstanceMinorVersions,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) GetDBProxyInstanceMinorVersions() []*string {
	return s.DBProxyInstanceMinorVersions
}

func (s *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) SetDBProxyInstanceMinorVersions(v []*string) *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions {
	s.DBProxyInstanceMinorVersions = v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyInstanceMinorVersions) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyResponseBodyDBProxyNodes struct {
	DBProxyNodes []*DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyResponseBodyDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyNodes) GetDBProxyNodes() []*DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes {
	return s.DBProxyNodes
}

func (s *DescribeDBProxyResponseBodyDBProxyNodes) SetDBProxyNodes(v []*DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) *DescribeDBProxyResponseBodyDBProxyNodes {
	s.DBProxyNodes = v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyNodes) Validate() error {
	if s.DBProxyNodes != nil {
		for _, item := range s.DBProxyNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes struct {
	// The number of CPU cores of the node.
	//
	// example:
	//
	// 2
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	// The ID of the proxy node.
	//
	// example:
	//
	// pn-xxxxxxx01
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The ID of the zone in which the node is deployed.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) SetCpuCores(v string) *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) SetNodeId(v string) *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) SetZoneId(v string) *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDBProxyNodesDBProxyNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyResponseBodyDbProxyEndpointItems struct {
	DbProxyEndpointItems []*DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems `json:"DbProxyEndpointItems,omitempty" xml:"DbProxyEndpointItems,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyResponseBodyDbProxyEndpointItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDbProxyEndpointItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItems) GetDbProxyEndpointItems() []*DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems {
	return s.DbProxyEndpointItems
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItems) SetDbProxyEndpointItems(v []*DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) *DescribeDBProxyResponseBodyDbProxyEndpointItems {
	s.DbProxyEndpointItems = v
	return s
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItems) Validate() error {
	if s.DbProxyEndpointItems != nil {
		for _, item := range s.DbProxyEndpointItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems struct {
	// The description of the database proxy endpoint.
	//
	// example:
	//
	// proxy-test
	DbProxyEndpointAliases *string `json:"DbProxyEndpointAliases,omitempty" xml:"DbProxyEndpointAliases,omitempty"`
	// The ID of the database proxy endpoint.
	//
	// example:
	//
	// gos787jog2wk0ye1****
	DbProxyEndpointName *string `json:"DbProxyEndpointName,omitempty" xml:"DbProxyEndpointName,omitempty"`
	// The type of the database proxy endpoint. Valid values:
	//
	// 	- Custom: custom database proxy endpoint
	//
	// 	- RWSplit: default database proxy endpoint
	//
	// example:
	//
	// RWSplit
	DbProxyEndpointType *string `json:"DbProxyEndpointType,omitempty" xml:"DbProxyEndpointType,omitempty"`
	// The read and write attributes of the database proxy endpoint.
	//
	// 	- ReadOnly
	//
	// 	- ReadWrite
	//
	// example:
	//
	// ReadWrite
	DbProxyReadWriteMode *string `json:"DbProxyReadWriteMode,omitempty" xml:"DbProxyReadWriteMode,omitempty"`
}

func (s DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) GetDbProxyEndpointAliases() *string {
	return s.DbProxyEndpointAliases
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) GetDbProxyEndpointName() *string {
	return s.DbProxyEndpointName
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) GetDbProxyEndpointType() *string {
	return s.DbProxyEndpointType
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) GetDbProxyReadWriteMode() *string {
	return s.DbProxyReadWriteMode
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) SetDbProxyEndpointAliases(v string) *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems {
	s.DbProxyEndpointAliases = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) SetDbProxyEndpointName(v string) *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems {
	s.DbProxyEndpointName = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) SetDbProxyEndpointType(v string) *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems {
	s.DbProxyEndpointType = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) SetDbProxyReadWriteMode(v string) *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems {
	s.DbProxyReadWriteMode = &v
	return s
}

func (s *DescribeDBProxyResponseBodyDbProxyEndpointItemsDbProxyEndpointItems) Validate() error {
	return dara.Validate(s)
}
