// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppGroupResponseBody
	GetRequestId() *string
	SetResult(v *DescribeAppGroupResponseBodyResult) *DescribeAppGroupResponseBody
	GetResult() *DescribeAppGroupResponseBodyResult
}

type DescribeAppGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	Result *DescribeAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppGroupResponseBody) GetResult() *DescribeAppGroupResponseBodyResult {
	return s.Result
}

func (s *DescribeAppGroupResponseBody) SetRequestId(v string) *DescribeAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupResponseBody) SetResult(v *DescribeAppGroupResponseBodyResult) *DescribeAppGroupResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppGroupResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppGroupResponseBodyResult struct {
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
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry of the application.
	//
	// example:
	//
	// ecommerce
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
	// The ID of the created rough sort expression.
	//
	// example:
	//
	// 0
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
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
	// The ID of the fine sort expression that is being created.
	//
	// example:
	//
	// 0
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete.
	//
	// example:
	//
	// -
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
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
	Quota *DescribeAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the resource group to which the network instance belongs.
	//
	// example:
	//
	// rg-acfmoiyerh6nzly
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The ID of the created fine sort expression.
	//
	// example:
	//
	// 0
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
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
	Tags []*DescribeAppGroupResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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

func (s DescribeAppGroupResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAppGroupResponseBodyResult) GetChargingWay() *int32 {
	return s.ChargingWay
}

func (s *DescribeAppGroupResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeAppGroupResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeAppGroupResponseBodyResult) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeAppGroupResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppGroupResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *DescribeAppGroupResponseBodyResult) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeAppGroupResponseBodyResult) GetExpireOn() *string {
	return s.ExpireOn
}

func (s *DescribeAppGroupResponseBodyResult) GetFirstRankAlgoDeploymentId() *int32 {
	return s.FirstRankAlgoDeploymentId
}

func (s *DescribeAppGroupResponseBodyResult) GetHasPendingQuotaReviewTask() *int32 {
	return s.HasPendingQuotaReviewTask
}

func (s *DescribeAppGroupResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeAppGroupResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAppGroupResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeAppGroupResponseBodyResult) GetLockedByExpiration() *int32 {
	return s.LockedByExpiration
}

func (s *DescribeAppGroupResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeAppGroupResponseBodyResult) GetPendingSecondRankAlgoDeploymentId() *int32 {
	return s.PendingSecondRankAlgoDeploymentId
}

func (s *DescribeAppGroupResponseBodyResult) GetProcessingOrderId() *string {
	return s.ProcessingOrderId
}

func (s *DescribeAppGroupResponseBodyResult) GetProduced() *int32 {
	return s.Produced
}

func (s *DescribeAppGroupResponseBodyResult) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAppGroupResponseBodyResult) GetQuota() *DescribeAppGroupResponseBodyResultQuota {
	return s.Quota
}

func (s *DescribeAppGroupResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAppGroupResponseBodyResult) GetSecondRankAlgoDeploymentId() *int32 {
	return s.SecondRankAlgoDeploymentId
}

func (s *DescribeAppGroupResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DescribeAppGroupResponseBodyResult) GetSwitchedTime() *int32 {
	return s.SwitchedTime
}

func (s *DescribeAppGroupResponseBodyResult) GetTags() []*DescribeAppGroupResponseBodyResultTags {
	return s.Tags
}

func (s *DescribeAppGroupResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeAppGroupResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeAppGroupResponseBodyResult) SetChargeType(v string) *DescribeAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetChargingWay(v int32) *DescribeAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCommodityCode(v string) *DescribeAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCreated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCurrentVersion(v string) *DescribeAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDescription(v string) *DescribeAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDomain(v string) *DescribeAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetEngineType(v string) *DescribeAppGroupResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetExpireOn(v string) *DescribeAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *DescribeAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetId(v string) *DescribeAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetInstanceId(v string) *DescribeAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockMode(v string) *DescribeAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *DescribeAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetName(v string) *DescribeAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProcessingOrderId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProduced(v int32) *DescribeAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProjectId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetQuota(v *DescribeAppGroupResponseBodyResultQuota) *DescribeAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetResourceGroupId(v string) *DescribeAppGroupResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetStatus(v string) *DescribeAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSwitchedTime(v int32) *DescribeAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetTags(v []*DescribeAppGroupResponseBodyResultTags) *DescribeAppGroupResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetType(v string) *DescribeAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetUpdated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) Validate() error {
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

type DescribeAppGroupResponseBodyResultQuota struct {
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

func (s DescribeAppGroupResponseBodyResultQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResultQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *DescribeAppGroupResponseBodyResultQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *DescribeAppGroupResponseBodyResultQuota) GetSpec() *string {
	return s.Spec
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetSpec(v string) *DescribeAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) Validate() error {
	return dara.Validate(s)
}

type DescribeAppGroupResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value
	//
	// example:
	//
	// bar
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeAppGroupResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAppGroupResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAppGroupResponseBodyResultTags) SetKey(v string) *DescribeAppGroupResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultTags) SetValue(v string) *DescribeAppGroupResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
