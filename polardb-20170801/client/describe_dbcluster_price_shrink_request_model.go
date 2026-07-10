// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationCategory(v string) *DescribeDBClusterPriceShrinkRequest
	GetCreationCategory() *string
	SetDBClusterId(v string) *DescribeDBClusterPriceShrinkRequest
	GetDBClusterId() *string
	SetDBNodeClass(v string) *DescribeDBClusterPriceShrinkRequest
	GetDBNodeClass() *string
	SetDBNodeIds(v []*string) *DescribeDBClusterPriceShrinkRequest
	GetDBNodeIds() []*string
	SetDBNodeNum(v int32) *DescribeDBClusterPriceShrinkRequest
	GetDBNodeNum() *int32
	SetDBNodesShrink(v string) *DescribeDBClusterPriceShrinkRequest
	GetDBNodesShrink() *string
	SetDBType(v string) *DescribeDBClusterPriceShrinkRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterPriceShrinkRequest
	GetDBVersion() *string
	SetHotStandbyCluster(v string) *DescribeDBClusterPriceShrinkRequest
	GetHotStandbyCluster() *string
	SetModifyType(v string) *DescribeDBClusterPriceShrinkRequest
	GetModifyType() *string
	SetOrderType(v string) *DescribeDBClusterPriceShrinkRequest
	GetOrderType() *string
	SetPayType(v string) *DescribeDBClusterPriceShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *DescribeDBClusterPriceShrinkRequest
	GetPeriod() *string
	SetProvisionedIops(v string) *DescribeDBClusterPriceShrinkRequest
	GetProvisionedIops() *string
	SetRegionId(v string) *DescribeDBClusterPriceShrinkRequest
	GetRegionId() *string
	SetServerlessType(v string) *DescribeDBClusterPriceShrinkRequest
	GetServerlessType() *string
	SetStorageChargeType(v string) *DescribeDBClusterPriceShrinkRequest
	GetStorageChargeType() *string
	SetStorageSpace(v string) *DescribeDBClusterPriceShrinkRequest
	GetStorageSpace() *string
	SetStorageType(v string) *DescribeDBClusterPriceShrinkRequest
	GetStorageType() *string
	SetUsedTime(v string) *DescribeDBClusterPriceShrinkRequest
	GetUsedTime() *string
	SetZoneId(v string) *DescribeDBClusterPriceShrinkRequest
	GetZoneId() *string
}

type DescribeDBClusterPriceShrinkRequest struct {
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
	DBNodesShrink *string `json:"DBNodes,omitempty" xml:"DBNodes,omitempty"`
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

func (s DescribeDBClusterPriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceShrinkRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBNodeIds() []*string {
	return s.DBNodeIds
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBNodeNum() *int32 {
	return s.DBNodeNum
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBNodesShrink() *string {
	return s.DBNodesShrink
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterPriceShrinkRequest) GetHotStandbyCluster() *string {
	return s.HotStandbyCluster
}

func (s *DescribeDBClusterPriceShrinkRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeDBClusterPriceShrinkRequest) GetProvisionedIops() *string {
	return s.ProvisionedIops
}

func (s *DescribeDBClusterPriceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterPriceShrinkRequest) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetStorageChargeType() *string {
	return s.StorageChargeType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetStorageSpace() *string {
	return s.StorageSpace
}

func (s *DescribeDBClusterPriceShrinkRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBClusterPriceShrinkRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *DescribeDBClusterPriceShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterPriceShrinkRequest) SetCreationCategory(v string) *DescribeDBClusterPriceShrinkRequest {
	s.CreationCategory = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBClusterId(v string) *DescribeDBClusterPriceShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBNodeClass(v string) *DescribeDBClusterPriceShrinkRequest {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBNodeIds(v []*string) *DescribeDBClusterPriceShrinkRequest {
	s.DBNodeIds = v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBNodeNum(v int32) *DescribeDBClusterPriceShrinkRequest {
	s.DBNodeNum = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBNodesShrink(v string) *DescribeDBClusterPriceShrinkRequest {
	s.DBNodesShrink = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetDBVersion(v string) *DescribeDBClusterPriceShrinkRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetHotStandbyCluster(v string) *DescribeDBClusterPriceShrinkRequest {
	s.HotStandbyCluster = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetModifyType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.ModifyType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetOrderType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetPayType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetPeriod(v string) *DescribeDBClusterPriceShrinkRequest {
	s.Period = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetProvisionedIops(v string) *DescribeDBClusterPriceShrinkRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetRegionId(v string) *DescribeDBClusterPriceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetServerlessType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.ServerlessType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetStorageChargeType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.StorageChargeType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetStorageSpace(v string) *DescribeDBClusterPriceShrinkRequest {
	s.StorageSpace = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetStorageType(v string) *DescribeDBClusterPriceShrinkRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetUsedTime(v string) *DescribeDBClusterPriceShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) SetZoneId(v string) *DescribeDBClusterPriceShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterPriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
