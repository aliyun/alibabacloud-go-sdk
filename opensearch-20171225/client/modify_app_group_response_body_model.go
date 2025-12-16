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
	// Response parameters
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
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billable item. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: QPS.
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
	// 1590139524
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
	// The type of the industry. Valid values:
	//
	// 	- GENERAL: general.
	//
	// 	- ECOMMERCE: e-commerce.
	//
	// 	- IT_CONTENT: IT content.
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
	// The time when the application expired.
	//
	// example:
	//
	// 1
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// 	- 0: normal.
	//
	// 	- 1: being approved.
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
	// The lock status. Valid values:
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
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// 	- 0: The order is in progress.
	//
	// 	- 1: The order is complete.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// 1
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
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
	// The state of the application. Valid values:
	//
	// 	- producing: being produced.
	//
	// 	- review_pending: being approved.
	//
	// 	- config_pending: to be configured.
	//
	// 	- normal: normal.
	//
	// 	- frozen: frozen.
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
	// The type of the application. Valid values:
	//
	// 	- standard: a standard edition application.
	//
	// 	- advance: an advanced edition application of an old version. New versions are not supported for this edition.
	//
	// 	- enhanced: an advanced edition application of a new version.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last modified.
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
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic.
	//
	// 	- opensearch.share.common: shared general-purpose.
	//
	// 	- opensearch.share.compute: shared computing.
	//
	// 	- opensearch.share.storage: shared storage.
	//
	// 	- opensearch.private.common: exclusive general-purpose.
	//
	// 	- opensearch.private.compute: exclusive computing.
	//
	// 	- opensearch.private.storage: exclusive storage.
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
