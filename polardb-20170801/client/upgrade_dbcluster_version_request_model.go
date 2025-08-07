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
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to immediately run the kernel upgrade task. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// >  This parameter is not required when you call the operation.
	//
	// example:
	//
	// false
	FromTimeService *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The latest start time to run the task that updates the kernel version of the cluster. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// > 	- The value of this parameter must be at least 30 minutes later than the value of PlannedStartTime.
	//
	// >	- If you specify `PlannedStartTime` but do not specify PlannedEndTime, the latest start time of the task is `PlannedEndTime + 30 minutes`. For example, if you set `PlannedStartTime` to `2021-01-14T09:00:00Z` and do not specify PlannedEndTime, the latest start time of the task is set to `2021-01-14T09:30:00Z`.
	//
	// example:
	//
	// 2021-01-14T09:30:00Z
	PlannedEndTime *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	// The earliest start time to run the task that updates the kernel version of the cluster. Specify the time in the `YYYY-MM-DDThh:mm:ssZ` format. The time must be in UTC.
	//
	// >
	//
	// 	- The earliest start time of the task can be a point in time within the next 72 hours. For example, if the current time is `2021-01-14T09:00:00Z`, you can specify a point in time from `2021-01-14T09:00:00Z` to `2021-01-17T09:00:00Z`.
	//
	// 	- If you do not specify this parameter, the kernel update task runs immediately after you submit the request.
	//
	// example:
	//
	// 2021-01-14T09:00:00Z
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the db version to which you want to upgrade the cluster. You can call the [DescribeDBClusterVersion](https://help.aliyun.com/document_detail/2319145.html) operation to query the version code.
	//
	// example:
	//
	// 20230707
	TargetDBRevisionVersionCode *string `json:"TargetDBRevisionVersionCode,omitempty" xml:"TargetDBRevisionVersionCode,omitempty"`
	// The code of the proxy version to which you want to upgrade the cluster. You can call the [DescribeDBClusterVersion](https://help.aliyun.com/document_detail/2319145.html) operation to query the version code.
	//
	// example:
	//
	// 20240702
	TargetProxyRevisionVersionCode *string `json:"TargetProxyRevisionVersionCode,omitempty" xml:"TargetProxyRevisionVersionCode,omitempty"`
	// The upgrade tag. The value is fixed as **INNOVATE**.
	//
	// > 	- This parameter is applicable only when you upgrade PolarDB for MySQL 8.0.1 to PolarDB for MySQL 8.0.2.
	//
	// >	- If you specify this parameter, you must set `UpgradePolicy` to **COLD**.
	//
	// example:
	//
	// INNOVATE
	UpgradeLabel *string `json:"UpgradeLabel,omitempty" xml:"UpgradeLabel,omitempty"`
	// The engine version upgrade policy. Valid values:
	//
	// 	- **HOT**: hot upgrade.
	//
	// 	- **COLD**: cold upgrade. Only PolarDB for MySQL 8.0 Cluster Edition supports this upgrade method.
	//
	// example:
	//
	// HOT
	UpgradePolicy *string `json:"UpgradePolicy,omitempty" xml:"UpgradePolicy,omitempty"`
	// The update type. Valid values:
	//
	// 	- **PROXY**: specifies to upgrade PloarProxy.
	//
	// 	- **DB**: specifies to upgrade the kernel version.
	//
	// 	- **ALL**: specifies to upgrade both PloarProxy and kernel version.
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
