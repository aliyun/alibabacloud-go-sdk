// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalTasksByUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListApprovalTasksByUserRequestListQuery) *ListApprovalTasksByUserRequest
	GetListQuery() *ListApprovalTasksByUserRequestListQuery
	SetOpTenantId(v int64) *ListApprovalTasksByUserRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListApprovalTasksByUserRequest
	GetOpUserId() *string
}

type ListApprovalTasksByUserRequest struct {
	// The query conditions.
	//
	// This parameter is required.
	ListQuery *ListApprovalTasksByUserRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListApprovalTasksByUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserRequest) GetListQuery() *ListApprovalTasksByUserRequestListQuery {
	return s.ListQuery
}

func (s *ListApprovalTasksByUserRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListApprovalTasksByUserRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListApprovalTasksByUserRequest) SetListQuery(v *ListApprovalTasksByUserRequestListQuery) *ListApprovalTasksByUserRequest {
	s.ListQuery = v
	return s
}

func (s *ListApprovalTasksByUserRequest) SetOpTenantId(v int64) *ListApprovalTasksByUserRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListApprovalTasksByUserRequest) SetOpUserId(v string) *ListApprovalTasksByUserRequest {
	s.OpUserId = &v
	return s
}

func (s *ListApprovalTasksByUserRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApprovalTasksByUserRequestListQuery struct {
	// The approval task type. Valid values:
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
	// - BIZ_PROCESS: Business process approval.
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
	// - PRIVILEGE_TRANSFER_APPROVAL: Permission transfer approval.
	//
	// - QD_FEATURE_ONLINE: Label listing.
	//
	// - QD_FEATURE_OFFLINE: Label delisting.
	//
	// - QD_CLUSTER_ONLINE: Group online.
	//
	// - QD_CLUSTER_OFFLINE: Group offline.
	//
	// - QD_MEMBER_ADD_APP: Add member to application.
	//
	// - QD_FEATURE_ADD_APP: Add label to application.
	//
	// - QD_CLUSTER_ADD_APP: Add group to application.
	//
	// - QD_FEATURE_ADD_PROJECT: Add label to project.
	//
	// - QD_CLUSTER_ADD_PROJECT: Add group to project.
	//
	// - TASK_DATA_DOWNLOAD: Data download.
	//
	// - CUSTOM_OPERATE: Custom operation.
	//
	// - PRIVACY_COMPUTING: Privacy-preserving computation.
	//
	// - MDC_TOPIC_DIR_PUBLISH: Asset topic directory publish.
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
	// The keyword for fuzzy match on the task name.
	//
	// example:
	//
	// datasource
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of records per page. Default value: 20. Maximum value: 100. Values greater than 100 are automatically adjusted to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The relationship type between the current user and the approval task. This parameter is required. Valid values:
	//
	// - SUBMITTED: Submitted by me.
	//
	// - PENDING_APPROVAL: Pending my approval.
	//
	// - PROCESSED: Processed by me.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUBMITTED
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The approval status filter. Status filtering is not supported in the pending scenario. Valid values:
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
	// The start of the submission time range, in the format yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	SubmittedFrom *string `json:"SubmittedFrom,omitempty" xml:"SubmittedFrom,omitempty"`
	// The end of the submission time range, in the format yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-12-31 23:59:59
	SubmittedTo *string `json:"SubmittedTo,omitempty" xml:"SubmittedTo,omitempty"`
}

func (s ListApprovalTasksByUserRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalTasksByUserRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListApprovalTasksByUserRequestListQuery) GetApprovalType() *string {
	return s.ApprovalType
}

func (s *ListApprovalTasksByUserRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListApprovalTasksByUserRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListApprovalTasksByUserRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApprovalTasksByUserRequestListQuery) GetRelationType() *string {
	return s.RelationType
}

func (s *ListApprovalTasksByUserRequestListQuery) GetStatus() *string {
	return s.Status
}

func (s *ListApprovalTasksByUserRequestListQuery) GetSubmittedFrom() *string {
	return s.SubmittedFrom
}

func (s *ListApprovalTasksByUserRequestListQuery) GetSubmittedTo() *string {
	return s.SubmittedTo
}

func (s *ListApprovalTasksByUserRequestListQuery) SetApprovalType(v string) *ListApprovalTasksByUserRequestListQuery {
	s.ApprovalType = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetKeyword(v string) *ListApprovalTasksByUserRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetPage(v int32) *ListApprovalTasksByUserRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetPageSize(v int32) *ListApprovalTasksByUserRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetRelationType(v string) *ListApprovalTasksByUserRequestListQuery {
	s.RelationType = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetStatus(v string) *ListApprovalTasksByUserRequestListQuery {
	s.Status = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetSubmittedFrom(v string) *ListApprovalTasksByUserRequestListQuery {
	s.SubmittedFrom = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) SetSubmittedTo(v string) *ListApprovalTasksByUserRequestListQuery {
	s.SubmittedTo = &v
	return s
}

func (s *ListApprovalTasksByUserRequestListQuery) Validate() error {
	return dara.Validate(s)
}
