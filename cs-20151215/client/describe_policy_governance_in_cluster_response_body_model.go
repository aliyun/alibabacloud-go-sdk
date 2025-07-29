// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyGovernanceInClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdmitLog(v *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) *DescribePolicyGovernanceInClusterResponseBody
	GetAdmitLog() *DescribePolicyGovernanceInClusterResponseBodyAdmitLog
	SetOnState(v []*DescribePolicyGovernanceInClusterResponseBodyOnState) *DescribePolicyGovernanceInClusterResponseBody
	GetOnState() []*DescribePolicyGovernanceInClusterResponseBodyOnState
	SetTotalViolations(v *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) *DescribePolicyGovernanceInClusterResponseBody
	GetTotalViolations() *DescribePolicyGovernanceInClusterResponseBodyTotalViolations
	SetViolations(v *DescribePolicyGovernanceInClusterResponseBodyViolations) *DescribePolicyGovernanceInClusterResponseBody
	GetViolations() *DescribePolicyGovernanceInClusterResponseBodyViolations
}

type DescribePolicyGovernanceInClusterResponseBody struct {
	// The audit logs of the policies in the cluster.
	AdmitLog *DescribePolicyGovernanceInClusterResponseBodyAdmitLog `json:"admit_log,omitempty" xml:"admit_log,omitempty" type:"Struct"`
	// Details about the policies of different severity levels that are enabled for the cluster.
	OnState []*DescribePolicyGovernanceInClusterResponseBodyOnState `json:"on_state,omitempty" xml:"on_state,omitempty" type:"Repeated"`
	// Details about the blocking and alerting events that are triggered by policies of different severity levels.
	TotalViolations *DescribePolicyGovernanceInClusterResponseBodyTotalViolations `json:"totalViolations,omitempty" xml:"totalViolations,omitempty" type:"Struct"`
	// Details about the blocking and alerting events that are triggered by different policies.
	Violations *DescribePolicyGovernanceInClusterResponseBodyViolations `json:"violations,omitempty" xml:"violations,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetAdmitLog() *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	return s.AdmitLog
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetOnState() []*DescribePolicyGovernanceInClusterResponseBodyOnState {
	return s.OnState
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetTotalViolations() *DescribePolicyGovernanceInClusterResponseBodyTotalViolations {
	return s.TotalViolations
}

func (s *DescribePolicyGovernanceInClusterResponseBody) GetViolations() *DescribePolicyGovernanceInClusterResponseBodyViolations {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetAdmitLog(v *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) *DescribePolicyGovernanceInClusterResponseBody {
	s.AdmitLog = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetOnState(v []*DescribePolicyGovernanceInClusterResponseBodyOnState) *DescribePolicyGovernanceInClusterResponseBody {
	s.OnState = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetTotalViolations(v *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) *DescribePolicyGovernanceInClusterResponseBody {
	s.TotalViolations = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) SetViolations(v *DescribePolicyGovernanceInClusterResponseBodyViolations) *DescribePolicyGovernanceInClusterResponseBody {
	s.Violations = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyAdmitLog struct {
	// The number of audit log entries.
	//
	// example:
	//
	// 100
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// The audit log content.
	Log *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog `json:"log,omitempty" xml:"log,omitempty" type:"Struct"`
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

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetLog() *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	return s.Log
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) GetProgress() *string {
	return s.Progress
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetCount(v int64) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Count = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetLog(v *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Log = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) SetProgress(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLog {
	s.Progress = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLog) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog struct {
	// The cluster ID.
	//
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The policy type.
	//
	// example:
	//
	// ACKAllowedRepos
	ConstraintKind *string `json:"constraint_kind,omitempty" xml:"constraint_kind,omitempty"`
	// The message that appears when an event is generated by a policy.
	//
	// example:
	//
	// d4hdhs*****
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Pod
	ResourceKind *string `json:"resource_kind,omitempty" xml:"resource_kind,omitempty"`
	// The resource name.
	//
	// example:
	//
	// nginx-deployment-basic2-84ccb74bfc-df22p
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// The namespace to which the resource belongs.
	//
	// example:
	//
	// default
	ResourceNamespace *string `json:"resource_namespace,omitempty" xml:"resource_namespace,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetConstraintKind() *string {
	return s.ConstraintKind
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetMsg() *string {
	return s.Msg
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) GetResourceNamespace() *string {
	return s.ResourceNamespace
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetClusterId(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetConstraintKind(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.ConstraintKind = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetMsg(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.Msg = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetResourceKind(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.ResourceKind = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetResourceName(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.ResourceName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) SetResourceNamespace(v string) *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog {
	s.ResourceNamespace = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyAdmitLogLog) Validate() error {
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

type DescribePolicyGovernanceInClusterResponseBodyTotalViolations struct {
	// Details about the blocking events that are triggered by the policies of each severity level.
	Deny *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny `json:"deny,omitempty" xml:"deny,omitempty" type:"Struct"`
	// Details about the alerting events that are triggered by the policies of each severity level.
	Warn *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn `json:"warn,omitempty" xml:"warn,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolations) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) GetDeny() *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny {
	return s.Deny
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) GetWarn() *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn {
	return s.Warn
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) SetDeny(v *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyTotalViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) SetWarn(v *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyTotalViolations {
	s.Warn = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolations) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny struct {
	// The severity level of the policy.
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The number of blocking events that are triggered.
	//
	// example:
	//
	// 0
	Violations *int64 `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsDeny) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn struct {
	// The severity level of the policy.
	//
	// example:
	//
	// low
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The number of alerting events that are triggered.
	//
	// example:
	//
	// 5
	Violations *int64 `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyTotalViolationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolations struct {
	// Details about the blocking events that are triggered by each policy.
	Deny *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny `json:"deny,omitempty" xml:"deny,omitempty" type:"Struct"`
	// Details about the alerting events that are triggered by the policies of each severity level.
	Warn *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn `json:"warn,omitempty" xml:"warn,omitempty" type:"Struct"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolations) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolations) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolations) GetDeny() *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny {
	return s.Deny
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolations) GetWarn() *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn {
	return s.Warn
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolations) SetDeny(v *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) *DescribePolicyGovernanceInClusterResponseBodyViolations {
	s.Deny = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolations) SetWarn(v *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) *DescribePolicyGovernanceInClusterResponseBodyViolations {
	s.Warn = v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolations) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolationsDeny struct {
	// The policy description.
	//
	// example:
	//
	// Requires container images to begin with a repo string from a specified list.
	PolicyDescription *string `json:"policyDescription,omitempty" xml:"policyDescription,omitempty"`
	// The policy name.
	//
	// example:
	//
	// policy-gatekeeper-ackallowedrepos
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The total number of blocking events that are triggered by the policy.
	//
	// example:
	//
	// 11
	Violations *int64 `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsDeny) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyGovernanceInClusterResponseBodyViolationsWarn struct {
	// The policy description.
	//
	// example:
	//
	// Controls Linux capabilities.
	PolicyDescription *string `json:"policyDescription,omitempty" xml:"policyDescription,omitempty"`
	// The policy name.
	//
	// example:
	//
	// policy-gatekeeper-ackpspcapabilities
	PolicyName *string `json:"policyName,omitempty" xml:"policyName,omitempty"`
	// The severity level of the policy.
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The total number of alerting events that are triggered by the policy.
	//
	// example:
	//
	// 81
	Violations *int64 `json:"violations,omitempty" xml:"violations,omitempty"`
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) GoString() string {
	return s.String()
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) GetSeverity() *string {
	return s.Severity
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) GetViolations() *int64 {
	return s.Violations
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) SetPolicyDescription(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) SetPolicyName(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn {
	s.PolicyName = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) SetSeverity(v string) *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn {
	s.Severity = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) SetViolations(v int64) *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn {
	s.Violations = &v
	return s
}

func (s *DescribePolicyGovernanceInClusterResponseBodyViolationsWarn) Validate() error {
	return dara.Validate(s)
}
