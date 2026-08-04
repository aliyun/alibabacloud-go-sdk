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
	SetSchemaId(v string) *ListApprovalsRequest
	GetSchemaId() *string
	SetSchemaName(v string) *ListApprovalsRequest
	GetSchemaName() *string
	SetStatuses(v []*string) *ListApprovalsRequest
	GetStatuses() []*string
}

type ListApprovalsRequest struct {
	// Collection of approval instance IDs.
	ApprovalIds []*string `json:"ApprovalIds,omitempty" xml:"ApprovalIds,omitempty" type:"Repeated"`
	// End time when the approval instance was created, in seconds since the Unix epoch.
	//
	// example:
	//
	// 1736750500
	CreateEndTime *int64 `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	// Start time when the approval instance was created, in seconds since the Unix epoch.
	//
	// example:
	//
	// 1730000000
	CreateStartTime *int64 `json:"CreateStartTime,omitempty" xml:"CreateStartTime,omitempty"`
	// Department of the user who created the approval instance.
	//
	// example:
	//
	// 测试部
	CreatorDepartment *string `json:"CreatorDepartment,omitempty" xml:"CreatorDepartment,omitempty"`
	// ID of the device used to create the approval instance.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// ID of the user who created the approval instance.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// Username of the user who created the approval instance.
	//
	// example:
	//
	// 王先生
	CreatorUsername *string `json:"CreatorUsername,omitempty" xml:"CreatorUsername,omitempty"`
	// Page number for the current page in a paged query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// ID of the user who performed an operation on the approval instance.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	OperatorUserId *string `json:"OperatorUserId,omitempty" xml:"OperatorUserId,omitempty"`
	// Username of the user who performed an operation on the approval instance.
	//
	// example:
	//
	// 李小姐
	OperatorUsername *string `json:"OperatorUsername,omitempty" xml:"OperatorUsername,omitempty"`
	// Number of entries per page in a paged query. Valid values: 1 to 500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Policy type. Valid values:
	//
	// - **DomainBlacklist**: Domain blacklist.
	//
	// - **DomainWhitelist**: Domain whitelist.
	//
	// - **SoftwareBlock**: Software blocking.
	//
	// - **AppUninstall**: App uninstallation.
	//
	// - **DlpSend**: File outbound transfer.
	//
	// - **PeripheralBlock**: Peripheral control.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// ID of the associated approval process.
	//
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Name of the associated approval process.
	//
	// example:
	//
	// 测试
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// ID of the associated approval template.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	// Name of the associated approval template.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// Collection of approval instance statuses.
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
