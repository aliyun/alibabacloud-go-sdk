// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationCategory(v string) *DescribeDBClusterPriceRequest
	GetCreationCategory() *string
	SetDBClusterId(v string) *DescribeDBClusterPriceRequest
	GetDBClusterId() *string
	SetDBNodeClass(v string) *DescribeDBClusterPriceRequest
	GetDBNodeClass() *string
	SetDBNodeIds(v []*string) *DescribeDBClusterPriceRequest
	GetDBNodeIds() []*string
	SetDBNodeNum(v int32) *DescribeDBClusterPriceRequest
	GetDBNodeNum() *int32
	SetDBNodes(v []*DescribeDBClusterPriceRequestDBNodes) *DescribeDBClusterPriceRequest
	GetDBNodes() []*DescribeDBClusterPriceRequestDBNodes
	SetDBType(v string) *DescribeDBClusterPriceRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterPriceRequest
	GetDBVersion() *string
	SetHotStandbyCluster(v string) *DescribeDBClusterPriceRequest
	GetHotStandbyCluster() *string
	SetModifyType(v string) *DescribeDBClusterPriceRequest
	GetModifyType() *string
	SetOrderType(v string) *DescribeDBClusterPriceRequest
	GetOrderType() *string
	SetPayType(v string) *DescribeDBClusterPriceRequest
	GetPayType() *string
	SetPeriod(v string) *DescribeDBClusterPriceRequest
	GetPeriod() *string
	SetProvisionedIops(v string) *DescribeDBClusterPriceRequest
	GetProvisionedIops() *string
	SetRegionId(v string) *DescribeDBClusterPriceRequest
	GetRegionId() *string
	SetServerlessType(v string) *DescribeDBClusterPriceRequest
	GetServerlessType() *string
	SetStorageChargeType(v string) *DescribeDBClusterPriceRequest
	GetStorageChargeType() *string
	SetStorageSpace(v string) *DescribeDBClusterPriceRequest
	GetStorageSpace() *string
	SetStorageType(v string) *DescribeDBClusterPriceRequest
	GetStorageType() *string
	SetUsedTime(v string) *DescribeDBClusterPriceRequest
	GetUsedTime() *string
	SetZoneId(v string) *DescribeDBClusterPriceRequest
	GetZoneId() *string
}

type DescribeDBClusterPriceRequest struct {
	// The cluster edition. Valid values:
	//
	// - Normal: Cluster Edition.
	//
	// - Basic: Single Node Edition.
	//
	// - ArchiveNormal: Archive Database.
	//
	// - SENormal: Standard Edition.
	//
	// - NormalMultimaster: Multi-master Cluster.
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The cluster ID. Required for non-BUY scenarios.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The node specifications. Required for the BUY scenario. Example format: polar.mysql.x4.large.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The list of node IDs to delete. Used when ModifyType is set to DELETE.
	//
	// example:
	//
	// pi-**************
	DBNodeIds []*string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// The number of nodes. Valid for the BUY scenario. This value includes the read/write node. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 2
	DBNodeNum *int32 `json:"DBNodeNum,omitempty" xml:"DBNodeNum,omitempty"`
	// The list of heterogeneous specification change nodes. Used for specification change scenarios to specify the target specifications for each node.
	DBNodes []*DescribeDBClusterPriceRequestDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The database engine type. Required for the BUY scenario. Valid values: MySQL, PostgreSQL, and Oracle.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The database engine version. Required for the BUY scenario. Valid values for MySQL: 5.6, 5.7, and 8.0. Valid values for PostgreSQL: 11 and 14. Valid values for Oracle: 11 and 14.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Specifies whether to enable the hot standby cluster. Valid values:
	//
	// - ON: Enable.
	//
	// - OFF: Disable.
	//
	// Valid for the BUY and specification change scenarios.
	//
	// example:
	//
	// ON
	HotStandbyCluster *string `json:"HotStandbyCluster,omitempty" xml:"HotStandbyCluster,omitempty"`
	// The specification change direction. Valid values:
	//
	// - ADD: add nodes.
	//
	// - DELETE: remove nodes.
	//
	// - UPGRADE: upgrade specifications.
	//
	// - DOWNGRADE: downgrade specifications.
	//
	// - HOT_STANDBY: hot standby change.
	//
	// - STORAGE: storage space change.
	//
	// - STORAGE_TYPE: storage type change.
	//
	// example:
	//
	// UPGRADE
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The order type. Valid values:
	//
	// - BUY: new purchase.
	//
	// - CONVERT: billing method conversion.
	//
	// - RENEW: renewal.
	//
	// - UPGRADE: upgrade specifications or add nodes.
	//
	// - DOWNGRADE: downgrade specifications or remove nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The billing method. Required for the BUY and CONVERT scenarios. Valid values:
	//
	// - Prepaid: subscription.
	//
	// - Postpaid: pay-as-you-go.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The subscription cycle. Valid values:
	//
	// - Month: monthly.
	//
	// - Year: yearly.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The provisioned IOPS. Used for the Standard Edition (SENormal) scenario.
	//
	// example:
	//
	// 1000
	ProvisionedIops *string `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The serverless type. Valid values: AgileServerless.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The storage billing type. Valid values:
	//
	// - Prepaid: subscription.
	//
	// - Postpaid: pay-as-you-go.
	//
	// example:
	//
	// Prepaid
	StorageChargeType *string `json:"StorageChargeType,omitempty" xml:"StorageChargeType,omitempty"`
	// The storage space, in GB. Used for prepaid storage or storage specification change scenarios.
	//
	// example:
	//
	// 50
	StorageSpace *string `json:"StorageSpace,omitempty" xml:"StorageSpace,omitempty"`
	// The storage type. Valid values:
	//
	// - PSL5: high performance.
	//
	// - PSL4: standard.
	//
	// - ESSDPL0
	//
	// - ESSDPL1
	//
	// - ESSDPL2
	//
	// - ESSDPL3
	//
	// - ESSDAUTOPL
	//
	// example:
	//
	// PSL5
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The subscription duration. Used together with Period. Valid for the BUY, CONVERT, and RENEW scenarios when the billing method is Prepaid.
	//
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The zone ID. We recommend that you specify this parameter for the BUY scenario.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *DescribeDBClusterPriceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterPriceRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterPriceRequest) GetDBNodeIds() []*string {
	return s.DBNodeIds
}

func (s *DescribeDBClusterPriceRequest) GetDBNodeNum() *int32 {
	return s.DBNodeNum
}

func (s *DescribeDBClusterPriceRequest) GetDBNodes() []*DescribeDBClusterPriceRequestDBNodes {
	return s.DBNodes
}

func (s *DescribeDBClusterPriceRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterPriceRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterPriceRequest) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *DescribeDBClusterPriceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *DescribeDBClusterPriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeDBClusterPriceRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterPriceRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeDBClusterPriceRequest) GetProvisionedIops() *string {
	return s.ProvisionedIops
}

func (s *DescribeDBClusterPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterPriceRequest) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClusterPriceRequest) GetStorageChargeType() *string {
	return s.StorageChargeType
}

func (s *DescribeDBClusterPriceRequest) GetStorageSpace() *string {
	return s.StorageSpace
}

func (s *DescribeDBClusterPriceRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClusterPriceRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *DescribeDBClusterPriceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterPriceRequest) SetCreationCategory(v string) *DescribeDBClusterPriceRequest {
	s.CreationCategory = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBClusterId(v string) *DescribeDBClusterPriceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBNodeClass(v string) *DescribeDBClusterPriceRequest {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBNodeIds(v []*string) *DescribeDBClusterPriceRequest {
	s.DBNodeIds = v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBNodeNum(v int32) *DescribeDBClusterPriceRequest {
	s.DBNodeNum = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBNodes(v []*DescribeDBClusterPriceRequestDBNodes) *DescribeDBClusterPriceRequest {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBType(v string) *DescribeDBClusterPriceRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetDBVersion(v string) *DescribeDBClusterPriceRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetHotStandbyCluster(v string) *DescribeDBClusterPriceRequest {
	s.HotStandbyCluster = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetModifyType(v string) *DescribeDBClusterPriceRequest {
	s.ModifyType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetOrderType(v string) *DescribeDBClusterPriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetPayType(v string) *DescribeDBClusterPriceRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetPeriod(v string) *DescribeDBClusterPriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetProvisionedIops(v string) *DescribeDBClusterPriceRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetRegionId(v string) *DescribeDBClusterPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetServerlessType(v string) *DescribeDBClusterPriceRequest {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetStorageChargeType(v string) *DescribeDBClusterPriceRequest {
	s.StorageChargeType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetStorageSpace(v string) *DescribeDBClusterPriceRequest {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetStorageType(v string) *DescribeDBClusterPriceRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetUsedTime(v string) *DescribeDBClusterPriceRequest {
	s.UsedTime = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) SetZoneId(v string) *DescribeDBClusterPriceRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterPriceRequest) Validate() error {
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPriceRequestDBNodes struct {
	// The target node specifications.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pi-**************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s DescribeDBClusterPriceRequestDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceRequestDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceRequestDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterPriceRequestDBNodes) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBClusterPriceRequestDBNodes) SetDBNodeClass(v string) *DescribeDBClusterPriceRequestDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterPriceRequestDBNodes) SetDBNodeId(v string) *DescribeDBClusterPriceRequestDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClusterPriceRequestDBNodes) Validate() error {
	return dara.Validate(s)
}
