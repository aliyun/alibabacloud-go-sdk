// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppGroupQuotaResponseBody
	GetRequestId() *string
	SetResult(v *ModifyAppGroupQuotaResponseBodyResult) *ModifyAppGroupQuotaResponseBody
	GetResult() *ModifyAppGroupQuotaResponseBodyResult
}

type ModifyAppGroupQuotaResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	Result *ModifyAppGroupQuotaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppGroupQuotaResponseBody) GetResult() *ModifyAppGroupQuotaResponseBodyResult {
	return s.Result
}

func (s *ModifyAppGroupQuotaResponseBody) SetRequestId(v string) *ModifyAppGroupQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBody) SetResult(v *ModifyAppGroupQuotaResponseBodyResult) *ModifyAppGroupQuotaResponseBody {
	s.Result = v
	return s
}

func (s *ModifyAppGroupQuotaResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppGroupQuotaResponseBodyResult struct {
	// The billing method.
	//
	// - POSTPAY: pay-as-you-go
	//
	// - PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model.
	//
	// - 1: by compute resources
	//
	// - 2: by queries per second (QPS)
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1590139542
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine type.
	//
	// example:
	//
	// ha3
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// Indicates whether a quota is pending for approval.
	//
	// - 0: No quota is pending for approval.
	//
	// - 1: A quota is pending for approval.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status.
	//
	// - Unlock: The application is not locked.
	//
	// - LockByExpiration: The application is automatically locked upon expiration.
	//
	// - ManualLock: The application is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the application is provisioned.
	//
	// - 0: The application is being provisioned.
	//
	// - 1: The application is provisioned.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test project.
	//
	// example:
	//
	// 1000
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The quota information of the application.
	Quota *ModifyAppGroupQuotaResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmoiyerh6nzly
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the application.
	//
	// - producing: The application is being provisioned.
	//
	// - review_pending: The application is pending for review.
	//
	// - config_pending: The application is pending for configuration.
	//
	// - normal: The application is running as normal.
	//
	// - frozen: The application is frozen.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the online version was switched.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application.
	//
	// - standard: Standard Edition
	//
	// - advance: an earlier version of Premium Edition. This type is not supported for new applications.
	//
	// - enhanced: a new version of Premium Edition.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1590978265
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetQuota() *ModifyAppGroupQuotaResponseBodyResultQuota {
	return s.Quota
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ModifyAppGroupQuotaResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargeType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCreated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetDescription(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetEngineType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetLockMode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetName(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProduced(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProjectId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetQuota(v *ModifyAppGroupQuotaResponseBodyResultQuota) *ModifyAppGroupQuotaResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetResourceGroupId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetStatus(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppGroupQuotaResponseBodyResultQuota struct {
	// The compute resources in LCU.
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity in GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications.
	//
	// - opensearch.share.junior: Entry-level
	//
	// - opensearch.share.common: Shared General-purpose
	//
	// - opensearch.share.compute: Shared Compute-optimized
	//
	// - opensearch.share.storage: Shared Storage-optimized
	//
	// - opensearch.private.common: Dedicated General-purpose
	//
	// - opensearch.private.compute: Dedicated Compute-optimized
	//
	// - opensearch.private.storage: Dedicated Storage-optimized
	//
	// example:
	//
	// opensearch.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}
