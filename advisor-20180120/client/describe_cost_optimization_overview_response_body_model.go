// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostOptimizationOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) *DescribeCostOptimizationOverviewResponseBody
	GetAccessDeniedDetail() *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail
	SetCode(v string) *DescribeCostOptimizationOverviewResponseBody
	GetCode() *string
	SetData(v *DescribeCostOptimizationOverviewResponseBodyData) *DescribeCostOptimizationOverviewResponseBody
	GetData() *DescribeCostOptimizationOverviewResponseBodyData
	SetMessage(v string) *DescribeCostOptimizationOverviewResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCostOptimizationOverviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCostOptimizationOverviewResponseBody
	GetSuccess() *bool
}

type DescribeCostOptimizationOverviewResponseBody struct {
	AccessDeniedDetail *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostOptimizationOverviewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Internal service issue. Detail:.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetAccessDeniedDetail() *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetData() *DescribeCostOptimizationOverviewResponseBodyData {
	return s.Data
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostOptimizationOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetAccessDeniedDetail(v *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) *DescribeCostOptimizationOverviewResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetCode(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetData(v *DescribeCostOptimizationOverviewResponseBodyData) *DescribeCostOptimizationOverviewResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetMessage(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetRequestId(v string) *DescribeCostOptimizationOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) SetSuccess(v bool) *DescribeCostOptimizationOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail struct {
	// AuthAction
	//
	// example:
	//
	// null
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// null
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// null
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// null
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// *****
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// null
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeCostOptimizationOverviewResponseBodyData struct {
	// example:
	//
	// 100
	BillingCycleDate *string `json:"BillingCycleDate,omitempty" xml:"BillingCycleDate,omitempty"`
	// example:
	//
	// 100
	CurrentBillingCost *string `json:"CurrentBillingCost,omitempty" xml:"CurrentBillingCost,omitempty"`
	// example:
	//
	// 100
	ExpectedSavingCost *string `json:"ExpectedSavingCost,omitempty" xml:"ExpectedSavingCost,omitempty"`
	// example:
	//
	// 2023-07-01 00:00:00
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 100
	OptCheckItemNum *string `json:"OptCheckItemNum,omitempty" xml:"OptCheckItemNum,omitempty"`
	// example:
	//
	// 100
	OptResourceNum         *string `json:"OptResourceNum,omitempty" xml:"OptResourceNum,omitempty"`
	ProcessedResourceCount *string `json:"ProcessedResourceCount,omitempty" xml:"ProcessedResourceCount,omitempty"`
	ProcessedSaveAmount    *string `json:"ProcessedSaveAmount,omitempty" xml:"ProcessedSaveAmount,omitempty"`
	// example:
	//
	// 95***135
	TaskId                   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WaitProcessResourceCount *string `json:"WaitProcessResourceCount,omitempty" xml:"WaitProcessResourceCount,omitempty"`
}

func (s DescribeCostOptimizationOverviewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostOptimizationOverviewResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetBillingCycleDate() *string {
	return s.BillingCycleDate
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetCurrentBillingCost() *string {
	return s.CurrentBillingCost
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetExpectedSavingCost() *string {
	return s.ExpectedSavingCost
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetOptCheckItemNum() *string {
	return s.OptCheckItemNum
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetOptResourceNum() *string {
	return s.OptResourceNum
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetProcessedResourceCount() *string {
	return s.ProcessedResourceCount
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetProcessedSaveAmount() *string {
	return s.ProcessedSaveAmount
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) GetWaitProcessResourceCount() *string {
	return s.WaitProcessResourceCount
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetBillingCycleDate(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.BillingCycleDate = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetCurrentBillingCost(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.CurrentBillingCost = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetExpectedSavingCost(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ExpectedSavingCost = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetGmtModified(v int64) *DescribeCostOptimizationOverviewResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetOptCheckItemNum(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.OptCheckItemNum = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetOptResourceNum(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.OptResourceNum = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetProcessedResourceCount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ProcessedResourceCount = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetProcessedSaveAmount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.ProcessedSaveAmount = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetTaskId(v int64) *DescribeCostOptimizationOverviewResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) SetWaitProcessResourceCount(v string) *DescribeCostOptimizationOverviewResponseBodyData {
	s.WaitProcessResourceCount = &v
	return s
}

func (s *DescribeCostOptimizationOverviewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
