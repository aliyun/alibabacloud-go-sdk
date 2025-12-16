// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceAppGroupCommodityCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceAppGroupCommodityCodeResponseBody
	GetRequestId() *string
	SetResult(v *ReplaceAppGroupCommodityCodeResponseBodyResult) *ReplaceAppGroupCommodityCodeResponseBody
	GetResult() *ReplaceAppGroupCommodityCodeResponseBodyResult
}

type ReplaceAppGroupCommodityCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AC5F78BA-66B9-545B-9CF1-8F542E682E1F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {}
	Result *ReplaceAppGroupCommodityCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReplaceAppGroupCommodityCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) GetResult() *ReplaceAppGroupCommodityCodeResponseBodyResult {
	return s.Result
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetRequestId(v string) *ReplaceAppGroupCommodityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetResult(v *ReplaceAppGroupCommodityCodeResponseBodyResult) *ReplaceAppGroupCommodityCodeResponseBody {
	s.Result = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplaceAppGroupCommodityCodeResponseBodyResult struct {
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing type. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: queries per second (QPS).
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// -
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the rough sort expression.
	//
	// example:
	//
	// 0
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval state of the quotas. Valid values:
	//
	// 	- 0: The approval status is normal.
	//
	// 	- 1: The quotas are being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// -
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// -
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock state. Valid values:
	//
	// 	- Unlock: The instance is unlocked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires. Valid values:
	//
	// 	- 0: The instance is not automatically locked after it expires.
	//
	// 	- 1: The instance is automatically locked after it expires.
	//
	// example:
	//
	// 0
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the order.
	//
	// example:
	//
	// -
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	//
	// example:
	//
	// 0
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is in progress.
	//
	// example:
	//
	// -
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is produced.
	//
	// example:
	//
	// 0
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// -
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The configuration information.
	Quota *ReplaceAppGroupCommodityCodeResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the fine sort expression.
	//
	// example:
	//
	// 0
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application.
	//
	// example:
	//
	// -
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was updated.
	//
	// example:
	//
	// 1581065904
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The versions.
	Versions []*string `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetFirstRankAlgoDeploymentId() *int32 {
	return s.FirstRankAlgoDeploymentId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetLockedByExpiration() *int32 {
	return s.LockedByExpiration
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetPendingSecondRankAlgoDeploymentId() *int32 {
	return s.PendingSecondRankAlgoDeploymentId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetProcessingOrderId() *string {
	return s.ProcessingOrderId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetQuota() *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	return s.Quota
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetSecondRankAlgoDeploymentId() *int32 {
	return s.SecondRankAlgoDeploymentId
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) GetVersions() []*string {
	return s.Versions
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargeType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargingWay(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCommodityCode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCreated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCurrentVersion(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetDescription(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetExpireOn(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetInstanceId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockMode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockedByExpiration(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetName(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProcessingOrderId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProduced(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProjectId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetQuota(v *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetStatus(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSwitchedTime(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetUpdated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetVersions(v []*string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Versions = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplaceAppGroupCommodityCodeResponseBodyResultQuota struct {
	// The number of computing resources configured.
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications configured.
	//
	// example:
	//
	// -
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetComputeResource(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetDocSize(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetSpec(v string) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}
