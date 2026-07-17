// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppGroupResponseBody
	GetRequestId() *string
	SetResult(v *ModifyAppGroupResponseBodyResult) *ModifyAppGroupResponseBody
	GetResult() *ModifyAppGroupResponseBodyResult
}

type ModifyAppGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
	Result *ModifyAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppGroupResponseBody) GetResult() *ModifyAppGroupResponseBodyResult {
	return s.Result
}

func (s *ModifyAppGroupResponseBody) SetRequestId(v string) *ModifyAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupResponseBody) SetResult(v *ModifyAppGroupResponseBodyResult) *ModifyAppGroupResponseBody {
	s.Result = v
	return s
}

func (s *ModifyAppGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppGroupResponseBodyResult struct {
	// The billing method. Valid values:
	//
	// - POSTPAY: pay-as-you-go.
	//
	// - PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billable item. Valid values:
	//
	// - 1: computing resources.
	//
	// - 2: QPS.
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
	// The UNIX timestamp when the application was created.
	//
	// example:
	//
	// 159013954
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
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry. Valid values:
	//
	// - GENERAL: general.
	//
	// - ECOMMERCE: e-commerce.
	//
	// - IT_CONTENT: IT content.
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
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
	// The approval status of the quota. Valid values:
	//
	// - 0: normal.
	//
	// - 1: being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 10030288
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// - Unlock: The instance is not locked.
	//
	// - LockByExpiration: The instance is automatically locked after it expires.
	//
	// - ManualLock: The instance is manually locked.
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
	// Indicates whether the application is created. Valid values:
	//
	// - 0: The application is being created.
	//
	// - 1: The application is created.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test project.
	//
	// example:
	//
	// 1
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The quota information of the application.
	//
	// example:
	//
	// {}
	Quota *ModifyAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmoiyerh6nzly
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the application. Valid values:
	//
	// - producing: The application is being created.
	//
	// - review_pending: The application is under review.
	//
	// - config_pending: The application requires configuration.
	//
	// - normal: The application is running.
	//
	// - frozen: The application is frozen.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The UNIX timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// - standard: a Standard Edition application.
	//
	// - advance: an Advanced Edition application of an old version. New applications do not support this type.
	//
	// - enhanced: an Enhanced Edition application.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The UNIX timestamp when the application was last modified.
	//
	// example:
	//
	// 1590978265
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyAppGroupResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *ModifyAppGroupResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ModifyAppGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ModifyAppGroupResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ModifyAppGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppGroupResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ModifyAppGroupResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *ModifyAppGroupResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *ModifyAppGroupResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *ModifyAppGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ModifyAppGroupResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAppGroupResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *ModifyAppGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ModifyAppGroupResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *ModifyAppGroupResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifyAppGroupResponseBodyResult) GetQuota() *ModifyAppGroupResponseBodyResultQuota {
	return s.Quota
}

func (s *ModifyAppGroupResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAppGroupResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ModifyAppGroupResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *ModifyAppGroupResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ModifyAppGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ModifyAppGroupResponseBodyResult) SetChargeType(v string) *ModifyAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCreated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDescription(v string) *ModifyAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDomain(v string) *ModifyAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetEngineType(v string) *ModifyAppGroupResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetId(v string) *ModifyAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetLockMode(v string) *ModifyAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetName(v string) *ModifyAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProduced(v int32) *ModifyAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProjectId(v string) *ModifyAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetQuota(v *ModifyAppGroupResponseBodyResultQuota) *ModifyAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetResourceGroupId(v string) *ModifyAppGroupResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetStatus(v string) *ModifyAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetType(v string) *ModifyAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppGroupResponseBodyResultQuota struct {
	// The computing resources in logical compute units (LCUs).
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
	// The specifications. Valid values:
	//
	// - opensearch.share.junior: Basic
	//
	// - opensearch.share.common: Shared General-purpose
	//
	// - opensearch.share.compute: Shared Compute-optimized
	//
	// - opensearch.share.storage: Shared Storage-optimized
	//
	// - opensearch.private.common: Exclusive General-purpose
	//
	// - opensearch.private.compute: Exclusive Compute-optimized
	//
	// - opensearch.private.storage: Exclusive Storage-optimized
	//
	// example:
	//
	// opensearch.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *ModifyAppGroupResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *ModifyAppGroupResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}
