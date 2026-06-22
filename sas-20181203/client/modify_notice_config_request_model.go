// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNoticeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ModifyNoticeConfigRequest
	GetBizType() *string
	SetFocusLevel(v string) *ModifyNoticeConfigRequest
	GetFocusLevel() *string
	SetProject(v string) *ModifyNoticeConfigRequest
	GetProject() *string
	SetRoute(v int32) *ModifyNoticeConfigRequest
	GetRoute() *int32
	SetSourceIp(v string) *ModifyNoticeConfigRequest
	GetSourceIp() *string
	SetTimeLimit(v int32) *ModifyNoticeConfigRequest
	GetTimeLimit() *int32
}

type ModifyNoticeConfigRequest struct {
	// The notification configuration type. By default, notifications are sent by SMS, email, or internal message. Valid values:
	//
	// - **cms**: CloudMonitor push.
	//
	// example:
	//
	// cms
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The focus level. Separate multiple levels with commas (,).
	//
	// When **Project*	- is **yundun_soar_incident_generate*	- or **yundun_soar_incident_update**, valid values:
	//
	// - **CRITICAL**: Critical.
	//
	// - **HIGH**: High-risk.
	//
	// - **MEDIUM**: Medium-risk.
	//
	// - **LOW**: Low-risk.
	//
	// - **INFO**: Informational.
	//
	// When **Project*	- is **yundun_sas_antiransomware_task**, valid values:
	//
	// - **Success**: Execution succeeded.
	//
	// - **Failed**: Execution failed.
	//
	// example:
	//
	// CRITICAL,HIGH
	FocusLevel *string `json:"FocusLevel,omitempty" xml:"FocusLevel,omitempty"`
	// ### Project identity
	//
	// #### When the BizType field is empty: valid values
	//
	// - **yundun_security_Weekreport**: Security weekly report (email only)
	//
	// - **sas_healthcheck**: Baseline check
	//
	// - **yundun_defennce_antiRansomware_overflow**: Anti-ransomware storage space exceeded
	//
	// - **yundun_sas_cloudsiem_log**: Cloud Threat Detection and Response (CTDR) log excess notification
	//
	// - **sas_suspicious**: Security alert
	//
	// - **yundun_aegis_AV_true**: Precise defense
	//
	// - **yundun_sas_ak_leakage AccessKey**: AccessKey leak intelligence
	//
	// - **yundun_sas_config_alert**: Cloud platform configuration check
	//
	// - **yundun_sas_vul_Emergency**: Emergency vulnerability intelligence
	//
	// - **yundun_webguard_event**: Web tamper-proofing
	//
	// - **yundun_sas_cloud_native_firewall**: Container firewall anomaly alert notification (email only)
	//
	// - **yundun_sas_cloud_native_firewall_Defense**: Container firewall active defense notification (email only)
	//
	// - **yundun_IP_Blocking**: Malicious IP blocking alerting notification
	//
	// - **yundun_sas_anti_virus_config**: Virus scan notification
	//
	// - **yundun_sas_log**: Log storage exceeded
	//
	// - **yundun_honeypot_alarm**: Cloud honeypot alerting
	//
	// - **aliyun_rasp_alarm**: Application protection alerting
	//
	// - **yundun_soar_incident_generate**: New security incident
	//
	// - **yundun_soar_incident_update**: Updated security incident
	//
	// > **yundun_security_Weekreport*	- sends a weekly report to notify you of unresolved vulnerabilities.
	//
	// ---
	//
	// #### When the BizType field is `cms`: valid values
	//
	// - **Vul_event**: Vulnerability result details
	//
	// - **Hc_summary**: Baseline check result summary
	//
	// - **Cspm_summary**: Cloud Security Posture Management (CSPM) result summary
	//
	// - **Vul_summary**: Vulnerability result summary
	//
	// - **Agentless_event**: Agentless detection result details
	//
	// - **Filedetect_event**: Malicious file SDK result details
	//
	// - **Rasp_event**: Application protection result details.
	//
	// example:
	//
	// yundun_webguard_event
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// ### Notification method
	//
	// #### When BizType is empty
	//
	// - 0: disabled
	//
	// - 1: SMS enabled
	//
	// - 2: email enabled
	//
	// - 4: internal message enabled
	//
	// - 3: SMS and email enabled
	//
	// - 5: SMS and internal message enabled
	//
	// - 6: email and internal message enabled
	//
	// - 7: SMS, email, and internal message all enabled
	//
	// #### When BizType is `cms`
	//
	// - 0: CloudMonitor push disabled
	//
	// - 1: CloudMonitor push enabled.
	//
	// example:
	//
	// 4
	Route *int32 `json:"Route,omitempty" xml:"Route,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 42.178.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// ### Notification time limit
	//
	// #### When the BizType field is empty: valid values
	//
	// - **0**: No limit.
	//
	// - **1**: Notifications are sent only between 08:00 and 22:00.
	//
	// #### When the BizType field is `cms`
	//
	// Specifies the push frequency limit, in seconds. The minimum value is **60**.
	//
	// example:
	//
	// 1
	TimeLimit *int32 `json:"TimeLimit,omitempty" xml:"TimeLimit,omitempty"`
}

func (s ModifyNoticeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNoticeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyNoticeConfigRequest) GetBizType() *string {
	return s.BizType
}

func (s *ModifyNoticeConfigRequest) GetFocusLevel() *string {
	return s.FocusLevel
}

func (s *ModifyNoticeConfigRequest) GetProject() *string {
	return s.Project
}

func (s *ModifyNoticeConfigRequest) GetRoute() *int32 {
	return s.Route
}

func (s *ModifyNoticeConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyNoticeConfigRequest) GetTimeLimit() *int32 {
	return s.TimeLimit
}

func (s *ModifyNoticeConfigRequest) SetBizType(v string) *ModifyNoticeConfigRequest {
	s.BizType = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetFocusLevel(v string) *ModifyNoticeConfigRequest {
	s.FocusLevel = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetProject(v string) *ModifyNoticeConfigRequest {
	s.Project = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetRoute(v int32) *ModifyNoticeConfigRequest {
	s.Route = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetSourceIp(v string) *ModifyNoticeConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyNoticeConfigRequest) SetTimeLimit(v int32) *ModifyNoticeConfigRequest {
	s.TimeLimit = &v
	return s
}

func (s *ModifyNoticeConfigRequest) Validate() error {
	return dara.Validate(s)
}
