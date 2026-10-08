// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCausalConsistReadTimeout(v string) *DescribeDBProxyEndpointResponseBody
	GetCausalConsistReadTimeout() *string
	SetDBProxyConnectString(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyConnectString() *string
	SetDBProxyConnectStringNetType(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyConnectStringNetType() *string
	SetDBProxyConnectStringPort(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyConnectStringPort() *string
	SetDBProxyEndpointCostThresholdForDuckdb(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyEndpointCostThresholdForDuckdb() *string
	SetDBProxyEndpointId(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyEndpointId() *string
	SetDBProxyEndpointMinSlaveCount(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyEndpointMinSlaveCount() *string
	SetDBProxyEngineType(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyEngineType() *string
	SetDBProxyFeatures(v string) *DescribeDBProxyEndpointResponseBody
	GetDBProxyFeatures() *string
	SetDBProxyNodes(v *DescribeDBProxyEndpointResponseBodyDBProxyNodes) *DescribeDBProxyEndpointResponseBody
	GetDBProxyNodes() *DescribeDBProxyEndpointResponseBodyDBProxyNodes
	SetDbProxyEndpointAliases(v string) *DescribeDBProxyEndpointResponseBody
	GetDbProxyEndpointAliases() *string
	SetDbProxyEndpointReadWriteMode(v string) *DescribeDBProxyEndpointResponseBody
	GetDbProxyEndpointReadWriteMode() *string
	SetDbProxyEndpointVpcId(v string) *DescribeDBProxyEndpointResponseBody
	GetDbProxyEndpointVpcId() *string
	SetDbProxyEndpointVswitchId(v string) *DescribeDBProxyEndpointResponseBody
	GetDbProxyEndpointVswitchId() *string
	SetDbProxyEndpointZoneId(v string) *DescribeDBProxyEndpointResponseBody
	GetDbProxyEndpointZoneId() *string
	SetEndpointConnectItems(v *DescribeDBProxyEndpointResponseBodyEndpointConnectItems) *DescribeDBProxyEndpointResponseBody
	GetEndpointConnectItems() *DescribeDBProxyEndpointResponseBodyEndpointConnectItems
	SetReadOnlyInstanceDistributionType(v string) *DescribeDBProxyEndpointResponseBody
	GetReadOnlyInstanceDistributionType() *string
	SetReadOnlyInstanceMaxDelayTime(v string) *DescribeDBProxyEndpointResponseBody
	GetReadOnlyInstanceMaxDelayTime() *string
	SetReadOnlyInstanceWeight(v string) *DescribeDBProxyEndpointResponseBody
	GetReadOnlyInstanceWeight() *string
	SetRequestId(v string) *DescribeDBProxyEndpointResponseBody
	GetRequestId() *string
}

type DescribeDBProxyEndpointResponseBody struct {
	// The timeout period for consistency reads. Unit: milliseconds. Default value: **10**. Valid values: **0 to 60000**.
	//
	// example:
	//
	// 10
	CausalConsistReadTimeout *string `json:"CausalConsistReadTimeout,omitempty" xml:"CausalConsistReadTimeout,omitempty"`
	// The proxy endpoint.
	//
	// example:
	//
	// testproxy****.rwlb.rds.aliyuncs.com
	DBProxyConnectString *string `json:"DBProxyConnectString,omitempty" xml:"DBProxyConnectString,omitempty"`
	// The network type of the proxy endpoint. Valid values:
	//
	// 	- **InnerString**: internal endpoint.
	//
	// 	- **OuterString**: public endpoint.
	//
	// example:
	//
	// InnerString
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// The port of the proxy endpoint.
	//
	// example:
	//
	// 3306
	DBProxyConnectStringPort              *string `json:"DBProxyConnectStringPort,omitempty" xml:"DBProxyConnectStringPort,omitempty"`
	DBProxyEndpointCostThresholdForDuckdb *string `json:"DBProxyEndpointCostThresholdForDuckdb,omitempty" xml:"DBProxyEndpointCostThresholdForDuckdb,omitempty"`
	// The ID of the proxy endpoint.
	//
	// example:
	//
	// keaxncrjluwu0gue****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// The minimum number of reserved instances.
	//
	// example:
	//
	// 2
	DBProxyEndpointMinSlaveCount *string `json:"DBProxyEndpointMinSlaveCount,omitempty" xml:"DBProxyEndpointMinSlaveCount,omitempty"`
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The settings of the proxy endpoint in JSON format. The following parameters are included:
	//
	// 	- **TransactionReadSqlRouteOptimizeStatus**: the transaction splitting setting. The value is **0*	- (disabled) or **1*	- (enabled).
	//
	// 	- **ConnectionPersist**: the connection pool setting. The value is **0*	- (disabled), **1*	- (session-level connection pool), or **2*	- (transaction-level connection pooling).
	//
	// 	- **ReadWriteSpliting**: the read/write splitting setting. The value is **0*	- (disabled) or **1*	- (enabled).
	//
	// 	- **AZProximityAccess**: the nearest access feature. The value is **0*	- (disabled) or **1*	- (enabled).
	//
	// 	- **CausalConsistRead**: the read consistency setting. The value is **0*	- (eventual consistency), **1*	- (session consistency), or **2*	- (global consistency).
	//
	// 	- **HtapFilter**: the automatic request distribution among row store and column store nodes setting. The value is **0*	- (disabled) or **1*	- (enabled).
	//
	// 	- **PinPreparedStmt**: visible only for ApsaraDB RDS for PostgreSQL. This is an internal parameter.
	//
	// > ApsaraDB RDS for PostgreSQL supports modification of only **ReadWriteSpliting**. **TransactionReadSqlRouteOptimizeStatus*	- and **PinPreparedStmt*	- are set to 1 by default.
	//
	// example:
	//
	// TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:0;ReadWriteSpliting:1
	DBProxyFeatures *string                                          `json:"DBProxyFeatures,omitempty" xml:"DBProxyFeatures,omitempty"`
	DBProxyNodes    *DescribeDBProxyEndpointResponseBodyDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Struct"`
	// The description of the proxy endpoint.
	//
	// example:
	//
	// proxyterminal-test
	DbProxyEndpointAliases *string `json:"DbProxyEndpointAliases,omitempty" xml:"DbProxyEndpointAliases,omitempty"`
	// The read/write type of the proxy endpoint. Valid values:
	//
	// 	- **ReadWrite**: read/write splitting mode.
	//
	// 	- **ReadOnly**: read-only mode.
	//
	// example:
	//
	// ReadWrite
	DbProxyEndpointReadWriteMode *string `json:"DbProxyEndpointReadWriteMode,omitempty" xml:"DbProxyEndpointReadWriteMode,omitempty"`
	// The VPC ID of the proxy endpoint.
	//
	// example:
	//
	// vpc-****
	DbProxyEndpointVpcId *string `json:"DbProxyEndpointVpcId,omitempty" xml:"DbProxyEndpointVpcId,omitempty"`
	// The vSwitch ID of the proxy endpoint.
	//
	// example:
	//
	// vsw-****
	DbProxyEndpointVswitchId *string `json:"DbProxyEndpointVswitchId,omitempty" xml:"DbProxyEndpointVswitchId,omitempty"`
	// The zone information of the proxy endpoint.
	//
	// example:
	//
	// cn-hangzhou-c
	DbProxyEndpointZoneId *string                                                  `json:"DbProxyEndpointZoneId,omitempty" xml:"DbProxyEndpointZoneId,omitempty"`
	EndpointConnectItems  *DescribeDBProxyEndpointResponseBodyEndpointConnectItems `json:"EndpointConnectItems,omitempty" xml:"EndpointConnectItems,omitempty" type:"Struct"`
	// The read weight distribution mode. For more information, see [Read weight distribution](https://help.aliyun.com/document_detail/96076.html). Valid values:
	//
	// 	- **Standard**: automatically distributes weights based on instance specifications.
	//
	// 	- **Custom**: uses custom weight distribution.
	//
	// example:
	//
	// Standard
	ReadOnlyInstanceDistributionType *string `json:"ReadOnlyInstanceDistributionType,omitempty" xml:"ReadOnlyInstanceDistributionType,omitempty"`
	// The latency threshold for read/write splitting. When the latency of a read-only instance exceeds this threshold, read traffic is not routed to the instance. Unit: seconds.
	//
	// example:
	//
	// 30
	ReadOnlyInstanceMaxDelayTime *string `json:"ReadOnlyInstanceMaxDelayTime,omitempty" xml:"ReadOnlyInstanceMaxDelayTime,omitempty"`
	// The read weight distribution information, which specifies the read request weights of the primary instance and read-only instances. The value is in JSON format and includes the following parameters:
	//
	// 	- **DBInstanceId**: the instance ID.
	//
	// 	- **DBInstanceType**: the instance type. The value is **Master*	- (primary instance) or **ReadOnly*	- (read-only instance).
	//
	// 	- **NodeID**: the node ID of the primary node or secondary node of the primary instance in the Cluster Edition.
	//
	// 	- **NodeType**: the node type in the Cluster Edition. The value is **Primary*	- (primary node of the primary instance) or **Secondary*	- (secondary node of the primary instance).
	//
	// 	- **Weight**: the read request weight. The value increases in increments of **100**. Maximum value: **10000**.
	//
	// example:
	//
	// [{\\"Availability\\":\\"Available\\",\\"DBInstanceId\\":\\"rm-2z****\\"，\\"DBInstanceType\\":\\"Master\\",\\"NodeId\\":\\"rn-t2****\\",\\"NodeType\\":\\"Primary\\",\\"Weight\\":0}, {\\"Availability\\":\\"Available\\",\\"DBInstanceId\\":\\"rm-2z****\\"，\\"DBInstanceType\\":\\"Master\\",\\"NodeId\\":\\"rn-z9****\\",\\"NodeType\\":\\"Secondary\\",\\"Weight\\":400}, {\\"Availability\\":\\"Available\\",,\\"DBInstanceId\\":\\"rm-2z****\\"，\\"DBInstanceType\\":\\"Master\\",\\"NodeId\\":\\"rn-1c****\\",\\"NodeType\\":\\"Secondary\\",\\"Weight\\":400}]]
	ReadOnlyInstanceWeight *string `json:"ReadOnlyInstanceWeight,omitempty" xml:"ReadOnlyInstanceWeight,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBProxyEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponseBody) GetCausalConsistReadTimeout() *string {
	return s.CausalConsistReadTimeout
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyConnectString() *string {
	return s.DBProxyConnectString
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyConnectStringPort() *string {
	return s.DBProxyConnectStringPort
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyEndpointCostThresholdForDuckdb() *string {
	return s.DBProxyEndpointCostThresholdForDuckdb
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyEndpointMinSlaveCount() *string {
	return s.DBProxyEndpointMinSlaveCount
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyFeatures() *string {
	return s.DBProxyFeatures
}

func (s *DescribeDBProxyEndpointResponseBody) GetDBProxyNodes() *DescribeDBProxyEndpointResponseBodyDBProxyNodes {
	return s.DBProxyNodes
}

func (s *DescribeDBProxyEndpointResponseBody) GetDbProxyEndpointAliases() *string {
	return s.DbProxyEndpointAliases
}

func (s *DescribeDBProxyEndpointResponseBody) GetDbProxyEndpointReadWriteMode() *string {
	return s.DbProxyEndpointReadWriteMode
}

func (s *DescribeDBProxyEndpointResponseBody) GetDbProxyEndpointVpcId() *string {
	return s.DbProxyEndpointVpcId
}

func (s *DescribeDBProxyEndpointResponseBody) GetDbProxyEndpointVswitchId() *string {
	return s.DbProxyEndpointVswitchId
}

func (s *DescribeDBProxyEndpointResponseBody) GetDbProxyEndpointZoneId() *string {
	return s.DbProxyEndpointZoneId
}

func (s *DescribeDBProxyEndpointResponseBody) GetEndpointConnectItems() *DescribeDBProxyEndpointResponseBodyEndpointConnectItems {
	return s.EndpointConnectItems
}

func (s *DescribeDBProxyEndpointResponseBody) GetReadOnlyInstanceDistributionType() *string {
	return s.ReadOnlyInstanceDistributionType
}

func (s *DescribeDBProxyEndpointResponseBody) GetReadOnlyInstanceMaxDelayTime() *string {
	return s.ReadOnlyInstanceMaxDelayTime
}

func (s *DescribeDBProxyEndpointResponseBody) GetReadOnlyInstanceWeight() *string {
	return s.ReadOnlyInstanceWeight
}

func (s *DescribeDBProxyEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBProxyEndpointResponseBody) SetCausalConsistReadTimeout(v string) *DescribeDBProxyEndpointResponseBody {
	s.CausalConsistReadTimeout = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyConnectString(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyConnectString = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyConnectStringNetType(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyConnectStringPort(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyConnectStringPort = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyEndpointCostThresholdForDuckdb(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyEndpointCostThresholdForDuckdb = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyEndpointId(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyEndpointId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyEndpointMinSlaveCount(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyEndpointMinSlaveCount = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyEngineType(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyFeatures(v string) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyFeatures = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDBProxyNodes(v *DescribeDBProxyEndpointResponseBodyDBProxyNodes) *DescribeDBProxyEndpointResponseBody {
	s.DBProxyNodes = v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDbProxyEndpointAliases(v string) *DescribeDBProxyEndpointResponseBody {
	s.DbProxyEndpointAliases = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDbProxyEndpointReadWriteMode(v string) *DescribeDBProxyEndpointResponseBody {
	s.DbProxyEndpointReadWriteMode = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDbProxyEndpointVpcId(v string) *DescribeDBProxyEndpointResponseBody {
	s.DbProxyEndpointVpcId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDbProxyEndpointVswitchId(v string) *DescribeDBProxyEndpointResponseBody {
	s.DbProxyEndpointVswitchId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetDbProxyEndpointZoneId(v string) *DescribeDBProxyEndpointResponseBody {
	s.DbProxyEndpointZoneId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetEndpointConnectItems(v *DescribeDBProxyEndpointResponseBodyEndpointConnectItems) *DescribeDBProxyEndpointResponseBody {
	s.EndpointConnectItems = v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetReadOnlyInstanceDistributionType(v string) *DescribeDBProxyEndpointResponseBody {
	s.ReadOnlyInstanceDistributionType = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetReadOnlyInstanceMaxDelayTime(v string) *DescribeDBProxyEndpointResponseBody {
	s.ReadOnlyInstanceMaxDelayTime = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetReadOnlyInstanceWeight(v string) *DescribeDBProxyEndpointResponseBody {
	s.ReadOnlyInstanceWeight = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) SetRequestId(v string) *DescribeDBProxyEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBody) Validate() error {
	if s.DBProxyNodes != nil {
		if err := s.DBProxyNodes.Validate(); err != nil {
			return err
		}
	}
	if s.EndpointConnectItems != nil {
		if err := s.EndpointConnectItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBProxyEndpointResponseBodyDBProxyNodes struct {
	DBProxyNodes []*DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes `json:"DBProxyNodes,omitempty" xml:"DBProxyNodes,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyEndpointResponseBodyDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponseBodyDBProxyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodes) GetDBProxyNodes() []*DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes {
	return s.DBProxyNodes
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodes) SetDBProxyNodes(v []*DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) *DescribeDBProxyEndpointResponseBodyDBProxyNodes {
	s.DBProxyNodes = v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodes) Validate() error {
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

type DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes struct {
	CpuCores *string `json:"cpuCores,omitempty" xml:"cpuCores,omitempty"`
	NodeId   *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	ZoneId   *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) GetCpuCores() *string {
	return s.CpuCores
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) SetCpuCores(v string) *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) SetNodeId(v string) *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) SetZoneId(v string) *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyDBProxyNodesDBProxyNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBProxyEndpointResponseBodyEndpointConnectItems struct {
	EndpointConnectItems []*DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems `json:"EndpointConnectItems,omitempty" xml:"EndpointConnectItems,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyEndpointResponseBodyEndpointConnectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponseBodyEndpointConnectItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItems) GetEndpointConnectItems() []*DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems {
	return s.EndpointConnectItems
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItems) SetEndpointConnectItems(v []*DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) *DescribeDBProxyEndpointResponseBodyEndpointConnectItems {
	s.EndpointConnectItems = v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItems) Validate() error {
	if s.EndpointConnectItems != nil {
		for _, item := range s.EndpointConnectItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems struct {
	DbProxyEndpointConnectString *string `json:"DbProxyEndpointConnectString,omitempty" xml:"DbProxyEndpointConnectString,omitempty"`
	DbProxyEndpointNetType       *string `json:"DbProxyEndpointNetType,omitempty" xml:"DbProxyEndpointNetType,omitempty"`
	DbProxyEndpointPort          *string `json:"DbProxyEndpointPort,omitempty" xml:"DbProxyEndpointPort,omitempty"`
}

func (s DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) GetDbProxyEndpointConnectString() *string {
	return s.DbProxyEndpointConnectString
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) GetDbProxyEndpointNetType() *string {
	return s.DbProxyEndpointNetType
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) GetDbProxyEndpointPort() *string {
	return s.DbProxyEndpointPort
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) SetDbProxyEndpointConnectString(v string) *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems {
	s.DbProxyEndpointConnectString = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) SetDbProxyEndpointNetType(v string) *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems {
	s.DbProxyEndpointNetType = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) SetDbProxyEndpointPort(v string) *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems {
	s.DbProxyEndpointPort = &v
	return s
}

func (s *DescribeDBProxyEndpointResponseBodyEndpointConnectItemsEndpointConnectItems) Validate() error {
	return dara.Validate(s)
}
