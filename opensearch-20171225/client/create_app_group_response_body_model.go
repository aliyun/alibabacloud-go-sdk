// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppGroupResponseBody
	GetRequestId() *string
	SetResult(v *CreateAppGroupResponseBodyResult) *CreateAppGroupResponseBody
	GetResult() *CreateAppGroupResponseBodyResult
}

type CreateAppGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 766CF6DB-CA02-3E12-7CBA-6AC21FC978FD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// None
	Result *CreateAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppGroupResponseBody) GetResult() *CreateAppGroupResponseBodyResult {
	return s.Result
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetResult(v *CreateAppGroupResponseBodyResult) *CreateAppGroupResponseBody {
	s.Result = v
	return s
}

func (s *CreateAppGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppGroupResponseBodyResult struct {
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
	// The type of billing. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: queries per second (QPS).
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
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
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
	// -
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval state of the quotas. Valid values:
	//
	// 	- 0: The application is in service.
	//
	// 	- 1: The quotas are being reviewed.
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
	// The ID of the instance.
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
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the application is created. Valid values:
	//
	// 	- 0: The application is being created.
	//
	// 	- 1: The application is created.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// -
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *CreateAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The status of the application. Valid values:
	//
	// 	- producing: The application is being created.
	//
	// 	- review_pending: The application is being reviewed.
	//
	// 	- config_pending: The application is to be configured.
	//
	// 	- normal: The application is in service.
	//
	// 	- frozen: The application is frozen.
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
	// 	- advance: an advanced edition which is of an old version. New version is not supported for this edition.
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

func (s CreateAppGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateAppGroupResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *CreateAppGroupResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateAppGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *CreateAppGroupResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *CreateAppGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *CreateAppGroupResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *CreateAppGroupResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateAppGroupResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *CreateAppGroupResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *CreateAppGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateAppGroupResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAppGroupResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *CreateAppGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateAppGroupResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *CreateAppGroupResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAppGroupResponseBodyResult) GetQuota() *CreateAppGroupResponseBodyResultQuota {
	return s.Quota
}

func (s *CreateAppGroupResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CreateAppGroupResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *CreateAppGroupResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *CreateAppGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *CreateAppGroupResponseBodyResult) SetChargeType(v string) *CreateAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetChargingWay(v int32) *CreateAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCommodityCode(v string) *CreateAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCreated(v int32) *CreateAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCurrentVersion(v string) *CreateAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDescription(v string) *CreateAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDomain(v string) *CreateAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetEngineType(v string) *CreateAppGroupResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetExpireOn(v string) *CreateAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *CreateAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetId(v string) *CreateAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetInstanceId(v string) *CreateAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetLockMode(v string) *CreateAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetName(v string) *CreateAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProduced(v int32) *CreateAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProjectId(v string) *CreateAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetQuota(v *CreateAppGroupResponseBodyResultQuota) *CreateAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetStatus(v string) *CreateAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetSwitchedTime(v int32) *CreateAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetType(v string) *CreateAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetUpdated(v int32) *CreateAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppGroupResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
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

func (s CreateAppGroupResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *CreateAppGroupResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *CreateAppGroupResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *CreateAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetDocSize(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetSpec(v string) *CreateAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}
