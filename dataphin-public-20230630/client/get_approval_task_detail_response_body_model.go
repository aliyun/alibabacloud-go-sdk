// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApprovalTaskDetailResponseBody
	GetCode() *string
	SetDetailInfo(v *GetApprovalTaskDetailResponseBodyDetailInfo) *GetApprovalTaskDetailResponseBody
	GetDetailInfo() *GetApprovalTaskDetailResponseBodyDetailInfo
	SetHttpStatusCode(v int32) *GetApprovalTaskDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApprovalTaskDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApprovalTaskDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApprovalTaskDetailResponseBody
	GetSuccess() *bool
}

type GetApprovalTaskDetailResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The approval task details.
	DetailInfo *GetApprovalTaskDetailResponseBodyDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend response exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetApprovalTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalTaskDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApprovalTaskDetailResponseBody) GetDetailInfo() *GetApprovalTaskDetailResponseBodyDetailInfo {
	return s.DetailInfo
}

func (s *GetApprovalTaskDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApprovalTaskDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApprovalTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalTaskDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApprovalTaskDetailResponseBody) SetCode(v string) *GetApprovalTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) SetDetailInfo(v *GetApprovalTaskDetailResponseBodyDetailInfo) *GetApprovalTaskDetailResponseBody {
	s.DetailInfo = v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) SetHttpStatusCode(v int32) *GetApprovalTaskDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) SetMessage(v string) *GetApprovalTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) SetRequestId(v string) *GetApprovalTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) SetSuccess(v bool) *GetApprovalTaskDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBody) Validate() error {
	if s.DetailInfo != nil {
		if err := s.DetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalTaskDetailResponseBodyDetailInfo struct {
	// The ID of the applicant.
	//
	// example:
	//
	// user001
	ApplicantId *string `json:"ApplicantId,omitempty" xml:"ApplicantId,omitempty"`
	// The name of the applicant.
	//
	// example:
	//
	// Zhang San
	ApplicantName *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	// The approval type. Valid values:
	//
	// - APPROVE: Permission approval.
	//
	// - MANAGE: Management.
	//
	// - OTHERS: Others.
	//
	// - ATOMIC: Atomic metric approval.
	//
	// - BIZ_OBJECT: Business object approval.
	//
	// - BIZ_PROCESS: Business activity approval.
	//
	// - PUBLISH_APPROVE: Publish approval.
	//
	// - BASELINE_APPROVE: Baseline approval.
	//
	// - CODE_REVIEW: Asset approval.
	//
	// - OBJECT_CODE_REVIEW: Code review.
	//
	// - STANDARD_APPROVAL: Standard online approval.
	//
	// - BATCH_STANDARD_APPROVAL: Batch standard online approval.
	//
	// - STANDARD_OFFLINE_APPROVAL: Standard offline approval.
	//
	// - BATCH_STANDARD_OFFLINE_APPROVAL: Batch standard offline approval.
	//
	// - PRIVILEGE_TRANSFER_APPROVAL: Privilege transfer approval.
	//
	// - QD_FEATURE_ONLINE: Tag listing.
	//
	// - QD_FEATURE_OFFLINE: Tag delisting.
	//
	// - QD_CLUSTER_ONLINE: Group online.
	//
	// - QD_CLUSTER_OFFLINE: Group offline.
	//
	// - QD_MEMBER_ADD_APP: Member added to application.
	//
	// - QD_FEATURE_ADD_APP: Tag added to application.
	//
	// - QD_CLUSTER_ADD_APP: Group added to application.
	//
	// - QD_FEATURE_ADD_PROJECT: Tag added to project.
	//
	// - QD_CLUSTER_ADD_PROJECT: Group added to project.
	//
	// - TASK_DATA_DOWNLOAD: Data download.
	//
	// - CUSTOM_OPERATE: Custom operation.
	//
	// - PRIVACY_COMPUTING: Privacy-preserving computation.
	//
	// - MDC_TOPIC_DIR_PUBLISH: Asset topic directory publishing.
	//
	// - ASSET_PUBLISH: Asset listing approval.
	//
	// - ASSET_UN_PUBLISH: Asset delisting approval.
	//
	// - APPLICATION_CREATE: Application creation approval.
	//
	// example:
	//
	// DATA_SOURCE
	ApprovalType *string `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The list of approval process steps.
	ApproveProcessSteps []*GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps `json:"ApproveProcessSteps,omitempty" xml:"ApproveProcessSteps,omitempty" type:"Repeated"`
	// The JSON string of the task details.
	//
	// example:
	//
	// {}
	BizContent *string `json:"BizContent,omitempty" xml:"BizContent,omitempty"`
	// The completion time. Format: yyyy-MM-dd HH:mm:ss. This field is empty if the task is not completed.
	//
	// example:
	//
	// 2025-07-30 12:00:00
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The approval flow ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source system.
	//
	// example:
	//
	// Dataphin
	SourceSystem *string `json:"SourceSystem,omitempty" xml:"SourceSystem,omitempty"`
	// The approval status. Valid values:
	//
	// - APPROVING: Approving.
	//
	// - APPROVED: Approved.
	//
	// - REJECTED: Rejected.
	//
	// - REVOKED: Revoked.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The submission time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-07-30 10:00:00
	SubmittedAt *string `json:"SubmittedAt,omitempty" xml:"SubmittedAt,omitempty"`
	// The task name.
	//
	// example:
	//
	// Data source application
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetApprovalTaskDetailResponseBodyDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalTaskDetailResponseBodyDetailInfo) GoString() string {
	return s.String()
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetApplicantId() *string {
	return s.ApplicantId
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetApplicantName() *string {
	return s.ApplicantName
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetApproveProcessSteps() []*GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	return s.ApproveProcessSteps
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetBizContent() *string {
	return s.BizContent
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetId() *int64 {
	return s.Id
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetSourceSystem() *string {
	return s.SourceSystem
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetStatus() *string {
	return s.Status
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetSubmittedAt() *string {
	return s.SubmittedAt
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) GetTitle() *string {
	return s.Title
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetApplicantId(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.ApplicantId = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetApplicantName(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.ApplicantName = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetApprovalType(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.ApprovalType = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetApproveProcessSteps(v []*GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.ApproveProcessSteps = v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetBizContent(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.BizContent = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetFinishedAt(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.FinishedAt = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetId(v int64) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.Id = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetSourceSystem(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.SourceSystem = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetStatus(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.Status = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetSubmittedAt(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.SubmittedAt = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) SetTitle(v string) *GetApprovalTaskDetailResponseBodyDetailInfo {
	s.Title = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfo) Validate() error {
	if s.ApproveProcessSteps != nil {
		for _, item := range s.ApproveProcessSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps struct {
	// The list of candidate approver IDs.
	//
	// example:
	//
	// user001,user002
	ApproveUserIds *string `json:"ApproveUserIds,omitempty" xml:"ApproveUserIds,omitempty"`
	// The list of candidate approver names.
	//
	// example:
	//
	// Zhang San, Li Si
	ApproveUserNames *string `json:"ApproveUserNames,omitempty" xml:"ApproveUserNames,omitempty"`
	// The operation time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-07-30 11:00:00
	OperateTime *string `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// user001
	OperateUserId *string `json:"OperateUserId,omitempty" xml:"OperateUserId,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// Zhang San
	OperateUserName *string `json:"OperateUserName,omitempty" xml:"OperateUserName,omitempty"`
	// The operation.
	//
	// example:
	//
	// SUBMIT
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The step name.
	//
	// example:
	//
	// Submit application
	ProcessStepName *string `json:"ProcessStepName,omitempty" xml:"ProcessStepName,omitempty"`
	// The approval comment.
	//
	// example:
	//
	// Approved
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The approval result.
	//
	// example:
	//
	// APPROVED
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GoString() string {
	return s.String()
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetApproveUserIds() *string {
	return s.ApproveUserIds
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetApproveUserNames() *string {
	return s.ApproveUserNames
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetOperateTime() *string {
	return s.OperateTime
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetOperateUserId() *string {
	return s.OperateUserId
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetOperateUserName() *string {
	return s.OperateUserName
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetOperation() *string {
	return s.Operation
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetProcessStepName() *string {
	return s.ProcessStepName
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetReason() *string {
	return s.Reason
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) GetResult() *string {
	return s.Result
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetApproveUserIds(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.ApproveUserIds = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetApproveUserNames(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.ApproveUserNames = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetOperateTime(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.OperateTime = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetOperateUserId(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.OperateUserId = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetOperateUserName(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.OperateUserName = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetOperation(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.Operation = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetProcessStepName(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.ProcessStepName = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetReason(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.Reason = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) SetResult(v string) *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps {
	s.Result = &v
	return s
}

func (s *GetApprovalTaskDetailResponseBodyDetailInfoApproveProcessSteps) Validate() error {
	return dara.Validate(s)
}
