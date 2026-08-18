// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalIds(v []*string) *ListApprovalsRequest
	GetApprovalIds() []*string
	SetCreateEndTime(v int64) *ListApprovalsRequest
	GetCreateEndTime() *int64
	SetCreateStartTime(v int64) *ListApprovalsRequest
	GetCreateStartTime() *int64
	SetCreatorDepartment(v string) *ListApprovalsRequest
	GetCreatorDepartment() *string
	SetCreatorDevTag(v string) *ListApprovalsRequest
	GetCreatorDevTag() *string
	SetCreatorUserId(v string) *ListApprovalsRequest
	GetCreatorUserId() *string
	SetCreatorUsername(v string) *ListApprovalsRequest
	GetCreatorUsername() *string
	SetCurrentPage(v int64) *ListApprovalsRequest
	GetCurrentPage() *int64
	SetEffectStatuses(v []*string) *ListApprovalsRequest
	GetEffectStatuses() []*string
	SetOperatorUserId(v string) *ListApprovalsRequest
	GetOperatorUserId() *string
	SetOperatorUsername(v string) *ListApprovalsRequest
	GetOperatorUsername() *string
	SetPageSize(v int64) *ListApprovalsRequest
	GetPageSize() *int64
	SetPolicyType(v string) *ListApprovalsRequest
	GetPolicyType() *string
	SetProcessId(v string) *ListApprovalsRequest
	GetProcessId() *string
	SetProcessName(v string) *ListApprovalsRequest
	GetProcessName() *string
	SetReportTypes(v []*string) *ListApprovalsRequest
	GetReportTypes() []*string
	SetSchemaId(v string) *ListApprovalsRequest
	GetSchemaId() *string
	SetSchemaName(v string) *ListApprovalsRequest
	GetSchemaName() *string
	SetStatuses(v []*string) *ListApprovalsRequest
	GetStatuses() []*string
}

type ListApprovalsRequest struct {
	// The collection of approval instance IDs.
	ApprovalIds []*string `json:"ApprovalIds,omitempty" xml:"ApprovalIds,omitempty" type:"Repeated"`
	// The end time for approval instance creation, in seconds-level timestamp.
	//
	// example:
	//
	// 1736750500
	CreateEndTime *int64 `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	// The start time for approval instance creation, in seconds-level timestamp.
	//
	// example:
	//
	// 1730000000
	CreateStartTime *int64 `json:"CreateStartTime,omitempty" xml:"CreateStartTime,omitempty"`
	// The department of the approval instance creator.
	//
	// example:
	//
	// QA Department
	CreatorDepartment *string `json:"CreatorDepartment,omitempty" xml:"CreatorDepartment,omitempty"`
	// The terminal device ID of the approval instance creator.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// The ID of the approval instance creator.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// The username of the approval instance creator.
	//
	// example:
	//
	// Mr. Wang
	CreatorUsername *string `json:"CreatorUsername,omitempty" xml:"CreatorUsername,omitempty"`
	// The page number of the current page in a paging query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of report effective statuses. Valid values: Enabled, Expired.
	EffectStatuses []*string `json:"EffectStatuses,omitempty" xml:"EffectStatuses,omitempty" type:"Repeated"`
	// The ID of the approval instance operator.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// The username of the approval instance operator.
	//
	// example:
	//
	// Ms. Li
	OperatorUsername *string `json:"OperatorUsername,omitempty" xml:"OperatorUsername,omitempty"`
	// The number of entries per page in a paging query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The adaptation policy type. Valid values:
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The associated approval process ID.
	//
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The associated approval process name.
	//
	// example:
	//
	// Test
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The list of report types. If not specified, only ApprovalReport is queried.
	ReportTypes []*string `json:"ReportTypes,omitempty" xml:"ReportTypes,omitempty" type:"Repeated"`
	// The associated approval template ID.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	// The associated approval template name.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The collection of approval instance statuses.
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListApprovalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalsRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalsRequest) GetApprovalIds() []*string {
	return s.ApprovalIds
}

func (s *ListApprovalsRequest) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListApprovalsRequest) GetCreateStartTime() *int64 {
	return s.CreateStartTime
}

func (s *ListApprovalsRequest) GetCreatorDepartment() *string {
	return s.CreatorDepartment
}

func (s *ListApprovalsRequest) GetCreatorDevTag() *string {
	return s.CreatorDevTag
}

func (s *ListApprovalsRequest) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *ListApprovalsRequest) GetCreatorUsername() *string {
	return s.CreatorUsername
}

func (s *ListApprovalsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListApprovalsRequest) GetEffectStatuses() []*string {
	return s.EffectStatuses
}

func (s *ListApprovalsRequest) GetOperatorUserId() *string {
	return s.OperatorUserId
}

func (s *ListApprovalsRequest) GetOperatorUsername() *string {
	return s.OperatorUsername
}

func (s *ListApprovalsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApprovalsRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalsRequest) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListApprovalsRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListApprovalsRequest) GetReportTypes() []*string {
	return s.ReportTypes
}

func (s *ListApprovalsRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ListApprovalsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListApprovalsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListApprovalsRequest) SetApprovalIds(v []*string) *ListApprovalsRequest {
	s.ApprovalIds = v
	return s
}

func (s *ListApprovalsRequest) SetCreateEndTime(v int64) *ListApprovalsRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListApprovalsRequest) SetCreateStartTime(v int64) *ListApprovalsRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListApprovalsRequest) SetCreatorDepartment(v string) *ListApprovalsRequest {
	s.CreatorDepartment = &v
	return s
}

func (s *ListApprovalsRequest) SetCreatorDevTag(v string) *ListApprovalsRequest {
	s.CreatorDevTag = &v
	return s
}

func (s *ListApprovalsRequest) SetCreatorUserId(v string) *ListApprovalsRequest {
	s.CreatorUserId = &v
	return s
}

func (s *ListApprovalsRequest) SetCreatorUsername(v string) *ListApprovalsRequest {
	s.CreatorUsername = &v
	return s
}

func (s *ListApprovalsRequest) SetCurrentPage(v int64) *ListApprovalsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApprovalsRequest) SetEffectStatuses(v []*string) *ListApprovalsRequest {
	s.EffectStatuses = v
	return s
}

func (s *ListApprovalsRequest) SetOperatorUserId(v string) *ListApprovalsRequest {
	s.OperatorUserId = &v
	return s
}

func (s *ListApprovalsRequest) SetOperatorUsername(v string) *ListApprovalsRequest {
	s.OperatorUsername = &v
	return s
}

func (s *ListApprovalsRequest) SetPageSize(v int64) *ListApprovalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApprovalsRequest) SetPolicyType(v string) *ListApprovalsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalsRequest) SetProcessId(v string) *ListApprovalsRequest {
	s.ProcessId = &v
	return s
}

func (s *ListApprovalsRequest) SetProcessName(v string) *ListApprovalsRequest {
	s.ProcessName = &v
	return s
}

func (s *ListApprovalsRequest) SetReportTypes(v []*string) *ListApprovalsRequest {
	s.ReportTypes = v
	return s
}

func (s *ListApprovalsRequest) SetSchemaId(v string) *ListApprovalsRequest {
	s.SchemaId = &v
	return s
}

func (s *ListApprovalsRequest) SetSchemaName(v string) *ListApprovalsRequest {
	s.SchemaName = &v
	return s
}

func (s *ListApprovalsRequest) SetStatuses(v []*string) *ListApprovalsRequest {
	s.Statuses = v
	return s
}

func (s *ListApprovalsRequest) Validate() error {
	return dara.Validate(s)
}
