// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGovernanceInClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetViolation(v *DescribePolicyGovernanceInClusterResponseBodyViolation) *DescribePolicyGovernanceInClusterResponseBody
	GetViolation() *DescribePolicyGovernanceInClusterResponseBodyViolation
	SetAdmitLog(v *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) *DescribePolicyGovernanceInClusterResponseBody
	GetAdmitLog() *DescribePolicyGovernanceInClusterResponseBodyAdmitLog
	SetOnState(v []*DescribePolicyGovernanceInClusterResponseBodyOnState) *DescribePolicyGovernanceInClusterResponseBody
	GetOnState() []*DescribePolicyGovernanceInClusterResponseBodyOnState
}

type DescribePolicyGovernanceInClusterResponseBody struct {
	Violation *DescribePolicyGovernanceInClusterResponseBodyViolation `json:"Violation,omitempty" xml:"Violation,omitempty" type:"Struct"`
	// The audit logs of the policies in the cluster.
	AdmitLog *DescribePolicyGovernanceInClusterResponseBodyAdmitLog `json:"admit_log,omitempty" xml:"admit_log,omitempty" type:"Struct"`
	// Details about the policies of different severity levels that are enabled for the cluster.
	OnState []*DescribePolicyGovernanceInClusterResponseBodyOnState `json:"on_state,omitempty" xml:"on_state,omitempty" type:"Repeated"`
}

func (s DescribePolicyGovernanceInClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetViolation() *DescribePolicyGovernanceInClusterResponseBodyViolation {
	return s.Violation
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetAdmitLog() *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	return s.AdmitLog
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetOnState() []*DescribePolicyGovernanceInClusterResponseBodyOnState {
	return s.OnState
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetViolation(v *DescribePolicyGovernanceInClusterResponseBodyViolation) *DescribePolicyGovernanceInClusterResponseBody {
	s.Violation = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetAdmitLog(v *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) *DescribePolicyGovernanceInClusterResponseBody {
	s.AdmitLog = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetOnState(v []*DescribePolicyGovernanceInClusterResponseBodyOnState) *DescribePolicyGovernanceInClusterResponseBody {
	s.OnState = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) Validate() error {
	if s.Violation != nil {
		if err := s.Violation.Validate(); err != nil {
			return err
		}
	}
	if s.AdmitLog != nil {
		if err := s.AdmitLog.Validate(); err != nil {
			return err
		}
	}
	if s.OnState != nil {
		for _, item := range s.OnState {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyGovernanceInClusterResponseBodyViolation struct {
	TotalViolations *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations `json:"totalViolations,omitempty" xml:"totalViolations,omitempty" type:"Struct"`
	Violations      *DescribePolicyGovernanceInClusterResponseBodyViolationViolations      `json:"violations,omitempty" xml:"violations,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolation) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolation) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolation) GetTotalViolations() *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations {
	return s.TotalViolations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolation) GetViolations() *DescribePolicyGovernanceInClusterResponseBodyViolationViolations {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolation) SetTotalViolations(v *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) *DescribePolicyGovernanceInClusterResponseBodyViolation {
	s.TotalViolations = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolation) SetViolations(v *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) *DescribePolicyGovernanceInClusterResponseBodyViolation {
	s.Violations = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolation) Validate() error {
	if s.TotalViolations != nil {
		if err := s.TotalViolations.Validate(); err != nil {
			return err
		}
	}
	if s.Violations != nil {
		if err := s.Violations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations struct {
	Deny []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny `json:"deny,omitempty" xml:"deny,omitempty" type:"Repeated"`
	Warn []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn `json:"warn,omitempty" xml:"warn,omitempty" type:"Repeated"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) GetDeny() []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny {
	return s.Deny
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) GetWarn() []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn {
	return s.Warn
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) SetDeny(v []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) SetWarn(v []*DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations {
	s.Warn = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolations) Validate() error {
	if s.Deny != nil {
		for _, item := range s.Deny {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Warn != nil {
		for _, item := range s.Warn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny struct {
	Severity   *string `json:"severity,omitempty" xml:"severity,omitempty"`
	Violations *string `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) GetViolations() *string {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) SetViolations(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsDeny) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn struct {
	Severity   *string `json:"severity,omitempty" xml:"severity,omitempty"`
	Violations *int64  `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationTotalViolationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolationViolations struct {
	Deny []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny `json:"deny,omitempty" xml:"deny,omitempty" type:"Repeated"`
	Warn []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn `json:"warn,omitempty" xml:"warn,omitempty" type:"Repeated"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolations) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) GetDeny() []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny {
	return s.Deny
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) GetWarn() []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn {
	return s.Warn
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) SetDeny(v []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyViolationViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) SetWarn(v []*DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyViolationViolations {
	s.Warn = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolations) Validate() error {
	if s.Deny != nil {
		for _, item := range s.Deny {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Warn != nil {
		for _, item := range s.Warn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny struct {
	PolicyDescription *string `json:"policyDescription,omitempty" xml:"policyDescription,omitempty"`
	PolicyName        *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	Severity          *string `json:"severity,omitempty" xml:"severity,omitempty"`
	Violations        *int64  `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsDeny) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn struct {
	PolicyDescription *string `json:"policyDescription,omitempty" xml:"policyDescription,omitempty"`
	PolicyName        *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	Severity          *string `json:"severity,omitempty" xml:"severity,omitempty"`
	Violations        *int64  `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationViolationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyAdmitLog struct {
	// The number of audit log entries.
	//
	// example:
	//
	// 100
	Count      *int64                                                       `json:"count,omitempty" xml:"count,omitempty"`
	LogProject *string                                                      `json:"log_project,omitempty" xml:"log_project,omitempty"`
	LogStore   *string                                                      `json:"log_store,omitempty" xml:"log_store,omitempty"`
	Logs       []*DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// The status of the query. Valid values:
	//
	// 	- `Complete`: The query succeeded and the complete query result is returned.
	//
	// 	- `Incomplete`: The query succeeded but the query result is incomplete. To obtain the complete query result, you must repeat the request.
	//
	// example:
	//
	// Complete
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLog) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetCount() *int64 {
	return s.Count
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetLogProject() *string {
	return s.LogProject
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetLogs() []*DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	return s.Logs
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetProgress() *string {
	return s.Progress
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetCount(v int64) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Count = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetLogProject(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.LogProject = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetLogStore(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.LogStore = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetLogs(v []*DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Logs = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetProgress(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Progress = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs struct {
	ClusterId            *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	ConstraintAction     *string `json:"constraint_action,omitempty" xml:"constraint_action,omitempty"`
	ConstraintApiVersion *string `json:"constraint_api_version,omitempty" xml:"constraint_api_version,omitempty"`
	ConstraintCategory   *string `json:"constraint_category,omitempty" xml:"constraint_category,omitempty"`
	ConstraintGroup      *string `json:"constraint_group,omitempty" xml:"constraint_group,omitempty"`
	ConstraintKind       *string `json:"constraint_kind,omitempty" xml:"constraint_kind,omitempty"`
	ConstraintName       *string `json:"constraint_name,omitempty" xml:"constraint_name,omitempty"`
	EventMsg             *string `json:"event_msg,omitempty" xml:"event_msg,omitempty"`
	EventType            *string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	RequestUid           *string `json:"request_uid,omitempty" xml:"request_uid,omitempty"`
	RequestUserinfo      *string `json:"request_userinfo,omitempty" xml:"request_userinfo,omitempty"`
	RequestUsername      *string `json:"request_username,omitempty" xml:"request_username,omitempty"`
	ResourceKind         *string `json:"resource_kind,omitempty" xml:"resource_kind,omitempty"`
	ResourceName         *string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	Time                 *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintAction() *string {
	return s.ConstraintAction
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintApiVersion() *string {
	return s.ConstraintApiVersion
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintCategory() *string {
	return s.ConstraintCategory
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintGroup() *string {
	return s.ConstraintGroup
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintKind() *string {
	return s.ConstraintKind
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetConstraintName() *string {
	return s.ConstraintName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetEventMsg() *string {
	return s.EventMsg
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetEventType() *string {
	return s.EventType
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetRequestUid() *string {
	return s.RequestUid
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetRequestUserinfo() *string {
	return s.RequestUserinfo
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetRequestUsername() *string {
	return s.RequestUsername
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) GetTime() *string {
	return s.Time
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetClusterId(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintAction(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintAction = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintApiVersion(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintApiVersion = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintCategory(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintCategory = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintGroup(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintGroup = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintKind(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintKind = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetConstraintName(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ConstraintName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetEventMsg(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.EventMsg = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetEventType(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.EventType = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetRequestUid(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.RequestUid = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetRequestUserinfo(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.RequestUserinfo = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetRequestUsername(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.RequestUsername = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetResourceKind(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ResourceKind = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetResourceName(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.ResourceName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) SetTime(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs {
	s.Time = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLogs) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyOnState struct {
	// The number of policies that are enabled.
	//
	// example:
	//
	// 3
	EnabledCount *int32 `json:"enabled_count,omitempty" xml:"enabled_count,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The total number of policies of the severity level.
	//
	// example:
	//
	// 8
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyOnState) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyOnState) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) GetEnabledCount() *int32 {
	return s.EnabledCount
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) SetEnabledCount(v int32) *DescribePolicyGovernanceInClusterResponseBodyOnState {
	s.EnabledCount = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyOnState {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) SetTotal(v int32) *DescribePolicyGovernanceInClusterResponseBodyOnState {
	s.Total = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyOnState) Validate() error {
	return dara.Validate(s)
}
