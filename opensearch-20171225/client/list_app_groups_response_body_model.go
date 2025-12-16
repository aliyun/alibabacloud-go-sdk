// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAppGroupsResponseBody
	GetRequestId() *string
	SetResult(v []*ListAppGroupsResponseBodyResult) *ListAppGroupsResponseBody
	GetResult() []*ListAppGroupsResponseBodyResult
	SetTotalCount(v int32) *ListAppGroupsResponseBody
	GetTotalCount() *int32
}

type ListAppGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	//
	// For more information, see [AppGroup](https://help.aliyun.com/document_detail/170000.html).
	//
	// example:
	//
	// []
	Result []*ListAppGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppGroupsResponseBody) GetResult() []*ListAppGroupsResponseBodyResult {
	return s.Result
}

func (s *ListAppGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAppGroupsResponseBody) SetRequestId(v string) *ListAppGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupsResponseBody) SetResult(v []*ListAppGroupsResponseBodyResult) *ListAppGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppGroupsResponseBody) SetTotalCount(v int32) *ListAppGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAppGroupsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppGroupsResponseBodyResult struct {
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
	// 1575442875
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 110116134
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// "xxx"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry of the application.
	//
	// example:
	//
	// ""
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
	// "xxx"
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
	// 110116134
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// "xxx"
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
	// Indicates whether the instance is automatically locked after it expires.
	//
	// example:
	//
	// 0
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The application name.
	//
	// example:
	//
	// os_function_test_v1
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
	// "xxx"
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application. For more information, see [Quota](https://help.aliyun.com/document_detail/170001.html).
	//
	// example:
	//
	// {}
	Quota *ListAppGroupsResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The state of the application. Valid values:
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
	// 0
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The application tags.
	Tags []*ListAppGroupsResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1578916076
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListAppGroupsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListAppGroupsResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *ListAppGroupsResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListAppGroupsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListAppGroupsResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ListAppGroupsResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListAppGroupsResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListAppGroupsResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *ListAppGroupsResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *ListAppGroupsResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *ListAppGroupsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListAppGroupsResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAppGroupsResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *ListAppGroupsResponseBodyResult) GetLockedByExpiration() *int32 {
	return s.LockedByExpiration
}

func (s *ListAppGroupsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListAppGroupsResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *ListAppGroupsResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListAppGroupsResponseBodyResult) GetQuota() *ListAppGroupsResponseBodyResultQuota {
	return s.Quota
}

func (s *ListAppGroupsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListAppGroupsResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *ListAppGroupsResponseBodyResult) GetTags() []*ListAppGroupsResponseBodyResultTags {
	return s.Tags
}

func (s *ListAppGroupsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListAppGroupsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListAppGroupsResponseBodyResult) SetChargeType(v string) *ListAppGroupsResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetChargingWay(v int32) *ListAppGroupsResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCommodityCode(v string) *ListAppGroupsResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCreated(v int32) *ListAppGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCurrentVersion(v string) *ListAppGroupsResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDescription(v string) *ListAppGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDomain(v string) *ListAppGroupsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetEngineType(v string) *ListAppGroupsResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetExpireOn(v string) *ListAppGroupsResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ListAppGroupsResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetId(v string) *ListAppGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetInstanceId(v string) *ListAppGroupsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockMode(v string) *ListAppGroupsResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockedByExpiration(v int32) *ListAppGroupsResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetName(v string) *ListAppGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProduced(v int32) *ListAppGroupsResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProjectId(v string) *ListAppGroupsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetQuota(v *ListAppGroupsResponseBodyResultQuota) *ListAppGroupsResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetStatus(v string) *ListAppGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetSwitchedTime(v int32) *ListAppGroupsResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetTags(v []*ListAppGroupsResponseBodyResultTags) *ListAppGroupsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetType(v string) *ListAppGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetUpdated(v int32) *ListAppGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppGroupsResponseBodyResultQuota struct {
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

func (s ListAppGroupsResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *ListAppGroupsResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *ListAppGroupsResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *ListAppGroupsResponseBodyResultQuota) SetComputeResource(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetDocSize(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetSpec(v string) *ListAppGroupsResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type ListAppGroupsResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bar
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppGroupsResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *ListAppGroupsResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *ListAppGroupsResponseBodyResultTags) SetKey(v string) *ListAppGroupsResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultTags) SetValue(v string) *ListAppGroupsResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
