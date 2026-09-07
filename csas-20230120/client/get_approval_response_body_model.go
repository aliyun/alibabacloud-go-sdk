// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApproval(v []*GetApprovalResponseBodyApproval) *GetApprovalResponseBody
	GetApproval() []*GetApprovalResponseBodyApproval
	SetRequestId(v string) *GetApprovalResponseBody
	GetRequestId() *string
}

type GetApprovalResponseBody struct {
	// The approval details list, which typically contains one record.
	Approval []*GetApprovalResponseBodyApproval `json:"Approval,omitempty" xml:"Approval,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBody) GetApproval() []*GetApprovalResponseBodyApproval {
	return s.Approval
}

func (s *GetApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApprovalResponseBody) SetApproval(v []*GetApprovalResponseBodyApproval) *GetApprovalResponseBody {
	s.Approval = v
	return s
}

func (s *GetApprovalResponseBody) SetRequestId(v string) *GetApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApprovalResponseBody) Validate() error {
	if s.Approval != nil {
		for _, item := range s.Approval {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalResponseBodyApproval struct {
	// The details of the approval instance.
	//
	// example:
	//
	// {"applicationId":"pa-application-eb75f0c80c28****","applicationName":"App***","associatedPolicyName":"Private access***"}
	ApprovalDetail *string `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty"`
	// The approval instance ID.
	//
	// example:
	//
	// approval-3564b140642f****
	ApprovalId *string `json:"ApprovalId,omitempty" xml:"ApprovalId,omitempty"`
	// The approval progress list. For backend reports without approval nodes, an empty array is returned.
	ApprovalProgresses []*GetApprovalResponseBodyApprovalApprovalProgresses `json:"ApprovalProgresses,omitempty" xml:"ApprovalProgresses,omitempty" type:"Repeated"`
	// The approval type. Valid values:
	//
	// 	- 0: built-in approval.
	//
	// 	- 1: DingTalk approval.
	//
	// 	- 2: WeCom approval.
	//
	// 	- 3: Lark approval.
	//
	// example:
	//
	// 0
	ApprovalType *int32 `json:"ApprovalType,omitempty" xml:"ApprovalType,omitempty"`
	// The backend report details. This value is returned only when ReportType is set to BackendReport.
	BackendReportDetail *GetApprovalResponseBodyApprovalBackendReportDetail `json:"BackendReportDetail,omitempty" xml:"BackendReportDetail,omitempty" type:"Struct"`
	// The creation time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2026-08-18 17:48:44
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creation time as a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1787046524
	CreateTimeUnix *int64 `json:"CreateTimeUnix,omitempty" xml:"CreateTimeUnix,omitempty"`
	// The department path of the report initiator.
	//
	// example:
	//
	// CN=cn***,OU=ou***
	CreatorDepartment *string `json:"CreatorDepartment,omitempty" xml:"CreatorDepartment,omitempty"`
	// The device ID of the terminal that created the approval instance.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	CreatorDevTag *string `json:"CreatorDevTag,omitempty" xml:"CreatorDevTag,omitempty"`
	// The ID of the user who created the approval instance. For backend reports, this is the actual effective user, not the administrator.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// The username of the user who created the approval instance.
	//
	// example:
	//
	// user***
	CreatorUsername *string `json:"CreatorUsername,omitempty" xml:"CreatorUsername,omitempty"`
	// The effective status of the report. This value is an empty string when the approval status is not Approved. Valid values:
	//
	// 	- Enabled: valid.
	//
	// 	- Expired: expired.
	//
	// example:
	//
	// Enabled
	EffectStatus *string `json:"EffectStatus,omitempty" xml:"EffectStatus,omitempty"`
	// The expiration time of the approval instance. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1757952000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The policy type associated with the approval instance. Valid values:
	//
	// - **DomainBlacklist**: Domain name blacklist.
	//
	// - **DomainWhitelist**: Domain name whitelist.
	//
	// - **SoftwareBlock**: Software blocking.
	//
	// - **DeviceRegistration**: Excess registration.
	//
	// - **AppUninstall**: Client uninstallation.
	//
	// - **DlpSend**: File outbound transfer.
	//
	// - **PeripheralBlock**: Peripheral control.
	//
	// - **EndpointHardening**: Endpoint hardening.
	//
	// - **oftwareHardening**: Software hardening.
	//
	// - **AiAgentBlock**: AI Agent control.
	//
	// - **PrivateAccessBlock**: Private access.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The ID of the process associated with the approval instance.
	//
	// example:
	//
	// approval-process-fcc351b8a95b****
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the process associated with the approval instance.
	//
	// example:
	//
	// Approval***
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The reason for creating the approval instance.
	//
	// example:
	//
	// Temporary access for a project
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The report type. Valid values:
	//
	// 	- ApprovalReport: approval report.
	//
	// 	- BackendReport: backend report.
	//
	// example:
	//
	// BackendReport
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The content of the template associated with the approval instance.
	//
	// example:
	//
	// {"form": {"labelCol": 6,"wrapperCol": 12}}
	SchemaContent *string `json:"SchemaContent,omitempty" xml:"SchemaContent,omitempty"`
	// The ID of the template associated with the approval instance.
	//
	// example:
	//
	// approval-schema-090134f1ebff****
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
	// The name of the template associated with the approval instance.
	//
	// example:
	//
	// Template***
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The instance status. Valid values:
	//
	// - **Pending**: Pending approval.
	//
	// - **Approved**: Approved.
	//
	// - **Rejected**: Denied.
	//
	// - **Revoked**: Revoked.
	//
	// - **Expired**: Expired.
	//
	// - **Deleted**: Deleted.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The validity duration type. Valid values:
	//
	// - **FixedTime**: Expires at a specified time.
	//
	// - **Permanent**: Permanently valid.
	//
	// example:
	//
	// Permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s GetApprovalResponseBodyApproval) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApproval) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApproval) GetApprovalDetail() *string {
	return s.ApprovalDetail
}

func (s *GetApprovalResponseBodyApproval) GetApprovalId() *string {
	return s.ApprovalId
}

func (s *GetApprovalResponseBodyApproval) GetApprovalProgresses() []*GetApprovalResponseBodyApprovalApprovalProgresses {
	return s.ApprovalProgresses
}

func (s *GetApprovalResponseBodyApproval) GetApprovalType() *int32 {
	return s.ApprovalType
}

func (s *GetApprovalResponseBodyApproval) GetBackendReportDetail() *GetApprovalResponseBodyApprovalBackendReportDetail {
	return s.BackendReportDetail
}

func (s *GetApprovalResponseBodyApproval) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApprovalResponseBodyApproval) GetCreateTimeUnix() *int64 {
	return s.CreateTimeUnix
}

func (s *GetApprovalResponseBodyApproval) GetCreatorDepartment() *string {
	return s.CreatorDepartment
}

func (s *GetApprovalResponseBodyApproval) GetCreatorDevTag() *string {
	return s.CreatorDevTag
}

func (s *GetApprovalResponseBodyApproval) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *GetApprovalResponseBodyApproval) GetCreatorUsername() *string {
	return s.CreatorUsername
}

func (s *GetApprovalResponseBodyApproval) GetEffectStatus() *string {
	return s.EffectStatus
}

func (s *GetApprovalResponseBodyApproval) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetApprovalResponseBodyApproval) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetApprovalResponseBodyApproval) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetApprovalResponseBodyApproval) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetApprovalResponseBodyApproval) GetReason() *string {
	return s.Reason
}

func (s *GetApprovalResponseBodyApproval) GetReportType() *string {
	return s.ReportType
}

func (s *GetApprovalResponseBodyApproval) GetSchemaContent() *string {
	return s.SchemaContent
}

func (s *GetApprovalResponseBodyApproval) GetSchemaId() *string {
	return s.SchemaId
}

func (s *GetApprovalResponseBodyApproval) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GetApprovalResponseBodyApproval) GetStatus() *string {
	return s.Status
}

func (s *GetApprovalResponseBodyApproval) GetValidityType() *string {
	return s.ValidityType
}

func (s *GetApprovalResponseBodyApproval) SetApprovalDetail(v string) *GetApprovalResponseBodyApproval {
	s.ApprovalDetail = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetApprovalId(v string) *GetApprovalResponseBodyApproval {
	s.ApprovalId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetApprovalProgresses(v []*GetApprovalResponseBodyApprovalApprovalProgresses) *GetApprovalResponseBodyApproval {
	s.ApprovalProgresses = v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetApprovalType(v int32) *GetApprovalResponseBodyApproval {
	s.ApprovalType = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetBackendReportDetail(v *GetApprovalResponseBodyApprovalBackendReportDetail) *GetApprovalResponseBodyApproval {
	s.BackendReportDetail = v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreateTime(v string) *GetApprovalResponseBodyApproval {
	s.CreateTime = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreateTimeUnix(v int64) *GetApprovalResponseBodyApproval {
	s.CreateTimeUnix = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorDepartment(v string) *GetApprovalResponseBodyApproval {
	s.CreatorDepartment = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorDevTag(v string) *GetApprovalResponseBodyApproval {
	s.CreatorDevTag = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorUserId(v string) *GetApprovalResponseBodyApproval {
	s.CreatorUserId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetCreatorUsername(v string) *GetApprovalResponseBodyApproval {
	s.CreatorUsername = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetEffectStatus(v string) *GetApprovalResponseBodyApproval {
	s.EffectStatus = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetEndTimestamp(v int64) *GetApprovalResponseBodyApproval {
	s.EndTimestamp = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetPolicyType(v string) *GetApprovalResponseBodyApproval {
	s.PolicyType = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetProcessId(v string) *GetApprovalResponseBodyApproval {
	s.ProcessId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetProcessName(v string) *GetApprovalResponseBodyApproval {
	s.ProcessName = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetReason(v string) *GetApprovalResponseBodyApproval {
	s.Reason = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetReportType(v string) *GetApprovalResponseBodyApproval {
	s.ReportType = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaContent(v string) *GetApprovalResponseBodyApproval {
	s.SchemaContent = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaId(v string) *GetApprovalResponseBodyApproval {
	s.SchemaId = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetSchemaName(v string) *GetApprovalResponseBodyApproval {
	s.SchemaName = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetStatus(v string) *GetApprovalResponseBodyApproval {
	s.Status = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) SetValidityType(v string) *GetApprovalResponseBodyApproval {
	s.ValidityType = &v
	return s
}

func (s *GetApprovalResponseBodyApproval) Validate() error {
	if s.ApprovalProgresses != nil {
		for _, item := range s.ApprovalProgresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BackendReportDetail != nil {
		if err := s.BackendReportDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalResponseBodyApprovalApprovalProgresses struct {
	// The action performed on the approval progress node. Valid values:
	//
	// - **Approve**: Approved.
	//
	// - **Reject**: Rejected.
	//
	// - **Revoke**: Revoked.
	//
	// - **Comment**: Commented.
	//
	// example:
	//
	// Approve
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The comment on the approval progress node.
	//
	// example:
	//
	// Approved
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The executor ID of the approval progress node.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The list of operators for the approval progress node.
	Operators []*GetApprovalResponseBodyApprovalApprovalProgressesOperators `json:"Operators,omitempty" xml:"Operators,omitempty" type:"Repeated"`
	// The status of the approval progress node. Valid values:
	//
	// - **Pending**: Pending approval.
	//
	// - **Approved**: Approved.
	//
	// - **Rejected**: Rejected.
	//
	// - **Revoked**: Revoked.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution time of the approval progress node. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1736752000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetApprovalResponseBodyApprovalApprovalProgresses) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalApprovalProgresses) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetAction() *string {
	return s.Action
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetComment() *string {
	return s.Comment
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetExecutor() *string {
	return s.Executor
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetOperators() []*GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	return s.Operators
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetStatus() *string {
	return s.Status
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetAction(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Action = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetComment(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Comment = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetExecutor(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Executor = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetOperators(v []*GetApprovalResponseBodyApprovalApprovalProgressesOperators) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Operators = v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetStatus(v string) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Status = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) SetTimestamp(v int64) *GetApprovalResponseBodyApprovalApprovalProgresses {
	s.Timestamp = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgresses) Validate() error {
	if s.Operators != nil {
		for _, item := range s.Operators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApprovalResponseBodyApprovalApprovalProgressesOperators struct {
	// The ID of the operator for the approval progress node.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// The username of the operator for the approval progress node.
	//
	// example:
	//
	// user***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetApprovalResponseBodyApprovalApprovalProgressesOperators) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalApprovalProgressesOperators) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) GetUsername() *string {
	return s.Username
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) SetSaseUserId(v string) *GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	s.SaseUserId = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) SetUsername(v string) *GetApprovalResponseBodyApprovalApprovalProgressesOperators {
	s.Username = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalApprovalProgressesOperators) Validate() error {
	return dara.Validate(s)
}

type GetApprovalResponseBodyApprovalBackendReportDetail struct {
	// The associated policy name.
	//
	// example:
	//
	// Private access***
	AssociatedPolicyName *string `json:"AssociatedPolicyName,omitempty" xml:"AssociatedPolicyName,omitempty"`
	// The associated policy type, which is the same as PolicyType.
	//
	// example:
	//
	// PrivateAccessBlock
	AssociatedPolicyType *string `json:"AssociatedPolicyType,omitempty" xml:"AssociatedPolicyType,omitempty"`
	// The remark for the backend report, which is the same as the report reason.
	//
	// example:
	//
	// Temporary access for a project
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The report object. The fields vary based on PolicyType. Fields within the object use camelCase naming.
	//
	// example:
	//
	// {"applicationId":"pa-application-eb75f0c80c28****","applicationName":"App***"}
	ReportObject interface{} `json:"ReportObject,omitempty" xml:"ReportObject,omitempty"`
	// The actual effective user of the backend report.
	TargetUser *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser `json:"TargetUser,omitempty" xml:"TargetUser,omitempty" type:"Struct"`
}

func (s GetApprovalResponseBodyApprovalBackendReportDetail) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalBackendReportDetail) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) GetAssociatedPolicyName() *string {
	return s.AssociatedPolicyName
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) GetAssociatedPolicyType() *string {
	return s.AssociatedPolicyType
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) GetRemark() *string {
	return s.Remark
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) GetReportObject() interface{} {
	return s.ReportObject
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) GetTargetUser() *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser {
	return s.TargetUser
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) SetAssociatedPolicyName(v string) *GetApprovalResponseBodyApprovalBackendReportDetail {
	s.AssociatedPolicyName = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) SetAssociatedPolicyType(v string) *GetApprovalResponseBodyApprovalBackendReportDetail {
	s.AssociatedPolicyType = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) SetRemark(v string) *GetApprovalResponseBodyApprovalBackendReportDetail {
	s.Remark = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) SetReportObject(v interface{}) *GetApprovalResponseBodyApprovalBackendReportDetail {
	s.ReportObject = v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) SetTargetUser(v *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) *GetApprovalResponseBodyApprovalBackendReportDetail {
	s.TargetUser = v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetail) Validate() error {
	if s.TargetUser != nil {
		if err := s.TargetUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApprovalResponseBodyApprovalBackendReportDetailTargetUser struct {
	// The SASE user ID of the actual effective user.
	//
	// example:
	//
	// su_70a1ed06a900d337527984de27568352fdfed1b19442a886d2a697c0327f****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the actual effective user.
	//
	// example:
	//
	// user***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) GoString() string {
	return s.String()
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) GetUserId() *string {
	return s.UserId
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) GetUsername() *string {
	return s.Username
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) SetUserId(v string) *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser {
	s.UserId = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) SetUsername(v string) *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser {
	s.Username = &v
	return s
}

func (s *GetApprovalResponseBodyApprovalBackendReportDetailTargetUser) Validate() error {
	return dara.Validate(s)
}
