// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBClusterVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpgradeDBClusterVersionRequest
	GetDBClusterId() *string
	SetFromTimeService(v bool) *UpgradeDBClusterVersionRequest
	GetFromTimeService() *bool
	SetOwnerAccount(v string) *UpgradeDBClusterVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeDBClusterVersionRequest
	GetOwnerId() *int64
	SetPlannedEndTime(v string) *UpgradeDBClusterVersionRequest
	GetPlannedEndTime() *string
	SetPlannedStartTime(v string) *UpgradeDBClusterVersionRequest
	GetPlannedStartTime() *string
	SetResourceOwnerAccount(v string) *UpgradeDBClusterVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBClusterVersionRequest
	GetResourceOwnerId() *int64
	SetTargetDBRevisionVersionCode(v string) *UpgradeDBClusterVersionRequest
	GetTargetDBRevisionVersionCode() *string
	SetTargetProxyRevisionVersionCode(v string) *UpgradeDBClusterVersionRequest
	GetTargetProxyRevisionVersionCode() *string
	SetUpgradeLabel(v string) *UpgradeDBClusterVersionRequest
	GetUpgradeLabel() *string
	SetUpgradePolicy(v string) *UpgradeDBClusterVersionRequest
	GetUpgradePolicy() *string
	SetUpgradeType(v string) *UpgradeDBClusterVersionRequest
	GetUpgradeType() *string
}

type UpgradeDBClusterVersionRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters under your account, such as the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to execute the kernel version upgrade immediately or on a schedule. Valid values:
	//
	// - **false*	- (default): The upgrade is executed on a schedule.
	//
	// - **true**: The upgrade is executed immediately.
	//
	// > You do not need to specify this parameter when you invoke this operation.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest time to start the scheduled task. The time is in the `YYYY-MM-DDThh:mm:ssZ` format and displayed in UTC.
	//
	// > 	- The latest time must be at least 30 minutes later than the start time.
	//
	// > 	- If you specify PlannedStartTime but do not specify this parameter, the latest time to start the task defaults to the start time plus 30 minutes. For example, if PlannedStartTime is set to `2021-01-14T09:00:00Z` and this parameter is left empty, the task starts at `2021-01-14T09:30:00Z` at the latest.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest time to start the scheduled Milvus version upgrade task. The time is in the `YYYY-MM-DDThh:mm:ssZ` format and displayed in UTC.
	//
	// > 	- The start time must be within the next 72 hours. For example, if the current time is `2021-01-14T09:00:00Z`, the allowed start time ranges from `2021-01-14T09:00:00Z` to `2021-01-17T09:00:00Z`.
	//
	// > 	- If you leave this parameter empty, the Milvus version upgrade task is executed immediately by default.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The VersionCode of the target database version. You can obtain the value from the [DescribeDBClusterVersion](https://help.aliyun.com/document_detail/2319145.html) operation.
	//
	// example:
	//
	// 20230707
	TargetDBRevisionVersionCode *string `json:"TargetDBRevisionVersionCode,omitempty" xml:"TargetDBRevisionVersionCode,omitempty"`
	// The VersionCode of the target PolarProxy version. You can obtain the value from the [DescribeDBClusterVersion](https://help.aliyun.com/document_detail/2319145.html) operation.
	//
	// example:
	//
	// 20240702
	TargetProxyRevisionVersionCode *string `json:"TargetProxyRevisionVersionCode,omitempty" xml:"TargetProxyRevisionVersionCode,omitempty"`
	// The kernel version upgrade label. Set the value to **INNOVATE**.
	//
	// > - This parameter applies only to upgrades from PolarDB for MySQL 8.0.1 to PolarDB for MySQL 8.0.2.
	//
	// > - When you specify this parameter, you must set UpgradePolicy to **COLD**.
	//
	// example:
	//
	// INNOVATE
	UpgradeLabel *string `json:"UpgradeLabel,omitempty" xml:"UpgradeLabel,omitempty"`
	// The kernel version upgrade policy. Valid values:
	//
	// - **HOT**: Hot upgrade.
	//
	// - **COLD**: Cold upgrade. Currently, only PolarDB for MySQL 8.0 Cluster Edition supports this upgrade method.
	//
	// example:
	//
	// HOT
	UpgradePolicy *string `json:"UpgradePolicy,omitempty" xml:"UpgradePolicy,omitempty"`
	// The upgrade type. Valid values:
	//
	// - **PROXY**: Upgrades only PolarProxy.
	//
	// - **DB**: Upgrades only the database engine.
	//
	// - **ALL*	- (default): Upgrades both PolarProxy and the database engine.
	//
	// example:
	//
	// PROXY
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBClusterVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBClusterVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpgradeDBClusterVersionRequest) GetFromTimeService() *bool {
	return s.FromTimeService
}

func (s *UpgradeDBClusterVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeDBClusterVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBClusterVersionRequest) GetPlannedEndTime() *string {
	return s.PlannedEndTime
}

func (s *UpgradeDBClusterVersionRequest) GetPlannedStartTime() *string {
	return s.PlannedStartTime
}

func (s *UpgradeDBClusterVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBClusterVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBClusterVersionRequest) GetTargetDBRevisionVersionCode() *string {
	return s.TargetDBRevisionVersionCode
}

func (s *UpgradeDBClusterVersionRequest) GetTargetProxyRevisionVersionCode() *string {
	return s.TargetProxyRevisionVersionCode
}

func (s *UpgradeDBClusterVersionRequest) GetUpgradeLabel() *string {
	return s.UpgradeLabel
}

func (s *UpgradeDBClusterVersionRequest) GetUpgradePolicy() *string {
	return s.UpgradePolicy
}

func (s *UpgradeDBClusterVersionRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *UpgradeDBClusterVersionRequest) SetDBClusterId(v string) *UpgradeDBClusterVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetFromTimeService(v bool) *UpgradeDBClusterVersionRequest {
	s.FromTimeService = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetOwnerAccount(v string) *UpgradeDBClusterVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetOwnerId(v int64) *UpgradeDBClusterVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetPlannedEndTime(v string) *UpgradeDBClusterVersionRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetPlannedStartTime(v string) *UpgradeDBClusterVersionRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBClusterVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBClusterVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetTargetDBRevisionVersionCode(v string) *UpgradeDBClusterVersionRequest {
	s.TargetDBRevisionVersionCode = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetTargetProxyRevisionVersionCode(v string) *UpgradeDBClusterVersionRequest {
	s.TargetProxyRevisionVersionCode = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetUpgradeLabel(v string) *UpgradeDBClusterVersionRequest {
	s.UpgradeLabel = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetUpgradePolicy(v string) *UpgradeDBClusterVersionRequest {
	s.UpgradePolicy = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetUpgradeType(v string) *UpgradeDBClusterVersionRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) Validate() error {
	return dara.Validate(s)
}
