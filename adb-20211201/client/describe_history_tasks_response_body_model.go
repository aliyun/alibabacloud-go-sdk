// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeHistoryTasksResponseBodyAccessDeniedDetail) *DescribeHistoryTasksResponseBody
	GetAccessDeniedDetail() *DescribeHistoryTasksResponseBodyAccessDeniedDetail
	SetCode(v string) *DescribeHistoryTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeHistoryTasksResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*DescribeHistoryTasksResponseBodyItems) *DescribeHistoryTasksResponseBody
	GetItems() []*DescribeHistoryTasksResponseBodyItems
	SetMessage(v string) *DescribeHistoryTasksResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeHistoryTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHistoryTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHistoryTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeHistoryTasksResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeHistoryTasksResponseBody
	GetTotalCount() *int32
}

type DescribeHistoryTasksResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *DescribeHistoryTasksResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The API status code or POP error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The task list.
	Items []*DescribeHistoryTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The additional information about the call result. Valid values:
	//
	// - If the request is normal, **Success*	- is returned.
	//
	// - If the request is abnormal, a specific error code is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the query result.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E51713A3-BB45-5616-AB15-170B704F8F01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponseBody) GetAccessDeniedDetail() *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeHistoryTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHistoryTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeHistoryTasksResponseBody) GetItems() []*DescribeHistoryTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeHistoryTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHistoryTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHistoryTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHistoryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHistoryTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHistoryTasksResponseBody) SetAccessDeniedDetail(v *DescribeHistoryTasksResponseBodyAccessDeniedDetail) *DescribeHistoryTasksResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetCode(v string) *DescribeHistoryTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetHttpStatusCode(v int32) *DescribeHistoryTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetItems(v []*DescribeHistoryTasksResponseBodyItems) *DescribeHistoryTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetMessage(v string) *DescribeHistoryTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetPageNumber(v int32) *DescribeHistoryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetPageSize(v int32) *DescribeHistoryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetRequestId(v string) *DescribeHistoryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetSuccess(v bool) *DescribeHistoryTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) SetTotalCount(v int32) *DescribeHistoryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHistoryTasksResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHistoryTasksResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// test
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// test
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The Alibaba Cloud account ID of the authentication principal.
	//
	// example:
	//
	// 141345906006****
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authorization principal.
	//
	// example:
	//
	// test
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFmy+hZZsvoWTQ2NTJFQUFBLUIyNjktMzk5NS05MEZCLUMxMzgwNDUwMDA2Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the access policy denial. Valid values:
	//
	// - **ImplicitDeny**: The resource owner has not configured a relevant access policy for the current user, and unauthorized operations are denied by default.
	//
	// - **ExplicitDeny**: The RAM policy configured by the resource owner explicitly denies the current user from accessing the corresponding resource.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The type of the policy that causes the permission denial. Valid values:
	//
	// - **ControlPolicy**: Control policy.
	//
	// - **SessionPolicy**: Session policy, which is an additional permission policy added to a temporary token.
	//
	// - **AssumeRolePolicy**: Assume role policy, which is the trust policy of a RAM role.
	//
	// - **AccountLevelIdentityBasedPolicy**: Account-level identity-access policy, which includes custom policies and system policies.
	//
	// - **ResourceGroupLevelIdentityBasedPolicy**: Resource group-level identity-access policy, which includes custom policies and system policies.
	//
	// example:
	//
	// ControlPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeHistoryTasksResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeHistoryTasksResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeHistoryTasksResponseBodyItems struct {
	// The allowed operation information. When you use this parameter, the system matches the operation action based on the currentStepName and status in this information. If no action is matched, the current task status does not support the operation.
	//
	// example:
	//
	// {\\"steps\\":[{\\"action_info\\":{\\"Waiting\\":[\\"modifySwitchTime\\"]},\\"step_name\\":\\"exec_task\\"}]}
	ActionInfo *string `json:"ActionInfo,omitempty" xml:"ActionInfo,omitempty"`
	// The ID of the requesting user. If callerSource is User, it represents the user UID.
	//
	// example:
	//
	// 141345906006****
	CallerSource *string `json:"CallerSource,omitempty" xml:"CallerSource,omitempty"`
	// The request source. Valid values:
	//
	// - **System**: System.
	//
	// - **User**: User.
	//
	// example:
	//
	// User
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The name of the current step. If this parameter is empty, the task has not started.
	//
	// example:
	//
	// exec_task
	CurrentStepName *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	// The database type.
	//
	// example:
	//
	// analyticdb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The end time of the task. The time is displayed in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2022-02-03T12:06:17Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// Instance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The product.
	//
	// example:
	//
	// ads
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The current progress.
	//
	// example:
	//
	// 79.0
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The reason why the current task is initiated.
	//
	// example:
	//
	// test
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The estimated remaining execution time. Unit: seconds.
	//
	// example:
	//
	// 1000
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The start time of the task. The time is displayed in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2022-02-03T11:31:03Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - **Scheduled**: Pending execution.
	//
	// - **Running**: Running.
	//
	// - **Succeed**: Succeeded.
	//
	// - **Failed**: Failed.
	//
	// - **Cancelling**: Being canceled.
	//
	// - **Canceled**: Canceled.
	//
	// - **Waiting**: Waiting for the scheduled time.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task details.
	//
	// example:
	//
	// {\\"callerUid\\":\\"test\\"}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-83br18hloy3faf****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// example:
	//
	// autotest_dispatch_cases
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The ID of the user to which the resource belongs.
	//
	// example:
	//
	// 141345906006****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHistoryTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksResponseBodyItems) GetActionInfo() *string {
	return s.ActionInfo
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCallerSource() *string {
	return s.CallerSource
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCallerUid() *string {
	return s.CallerUid
}

func (s *DescribeHistoryTasksResponseBodyItems) GetCurrentStepName() *string {
	return s.CurrentStepName
}

func (s *DescribeHistoryTasksResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeHistoryTasksResponseBodyItems) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetProduct() *string {
	return s.Product
}

func (s *DescribeHistoryTasksResponseBodyItems) GetProgress() *float32 {
	return s.Progress
}

func (s *DescribeHistoryTasksResponseBodyItems) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeHistoryTasksResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHistoryTasksResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskDetail() *string {
	return s.TaskDetail
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeHistoryTasksResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeHistoryTasksResponseBodyItems) GetUid() *string {
	return s.Uid
}

func (s *DescribeHistoryTasksResponseBodyItems) SetActionInfo(v string) *DescribeHistoryTasksResponseBodyItems {
	s.ActionInfo = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCallerSource(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CallerSource = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCallerUid(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CallerUid = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetCurrentStepName(v string) *DescribeHistoryTasksResponseBodyItems {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetDbType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetEndTime(v string) *DescribeHistoryTasksResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceName(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceName = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetInstanceType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.InstanceType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetProduct(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Product = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetProgress(v float32) *DescribeHistoryTasksResponseBodyItems {
	s.Progress = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetReasonCode(v string) *DescribeHistoryTasksResponseBodyItems {
	s.ReasonCode = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetRegionId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetRemainTime(v int32) *DescribeHistoryTasksResponseBodyItems {
	s.RemainTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetStartTime(v string) *DescribeHistoryTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetStatus(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskDetail(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskDetail = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskId(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetTaskType(v string) *DescribeHistoryTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) SetUid(v string) *DescribeHistoryTasksResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *DescribeHistoryTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
