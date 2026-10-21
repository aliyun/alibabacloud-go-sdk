// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentInvestigationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentInvestigation(v *GetIncidentInvestigationResponseBodyIncidentInvestigation) *GetIncidentInvestigationResponseBody
	GetIncidentInvestigation() *GetIncidentInvestigationResponseBodyIncidentInvestigation
	SetRequestId(v string) *GetIncidentInvestigationResponseBody
	GetRequestId() *string
}

type GetIncidentInvestigationResponseBody struct {
	// The AI investigation details of the incident. This field is empty if no matching investigation record is found.
	IncidentInvestigation *GetIncidentInvestigationResponseBodyIncidentInvestigation `json:"IncidentInvestigation,omitempty" xml:"IncidentInvestigation,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot this call.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIncidentInvestigationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentInvestigationResponseBody) GoString() string {
	return s.String()
}

func (s *GetIncidentInvestigationResponseBody) GetIncidentInvestigation() *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	return s.IncidentInvestigation
}

func (s *GetIncidentInvestigationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIncidentInvestigationResponseBody) SetIncidentInvestigation(v *GetIncidentInvestigationResponseBodyIncidentInvestigation) *GetIncidentInvestigationResponseBody {
	s.IncidentInvestigation = v
	return s
}

func (s *GetIncidentInvestigationResponseBody) SetRequestId(v string) *GetIncidentInvestigationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIncidentInvestigationResponseBody) Validate() error {
	if s.IncidentInvestigation != nil {
		if err := s.IncidentInvestigation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIncidentInvestigationResponseBodyIncidentInvestigation struct {
	// The name of the alert that triggered this AI investigation.
	//
	// example:
	//
	// Webshell Alert
	IncidentInvestigationAlertName *string `json:"IncidentInvestigationAlertName,omitempty" xml:"IncidentInvestigationAlertName,omitempty"`
	// The judgment conclusion of the AI investigation, corresponding to ai_judgement_conclusion in the report.
	//
	// example:
	//
	// Insufficient evidence to confirm whether the attack chain is complete. Supplement correlated logs for further analysis
	IncidentInvestigationConclusion *string `json:"IncidentInvestigationConclusion,omitempty" xml:"IncidentInvestigationConclusion,omitempty"`
	// The display ID of the AI investigation report. You can use this value as the IncidentInvestigationDisplayId parameter of this operation to query the report.
	//
	// example:
	//
	// 755616bb5d63046c1e4f62ea872a****
	IncidentInvestigationDisplayId *string `json:"IncidentInvestigationDisplayId,omitempty" xml:"IncidentInvestigationDisplayId,omitempty"`
	// The end time of the AI investigation, in seconds (UNIX timestamp). If the execution end time is not recorded, the last update time of the investigation record is used. This field is empty if neither time exists.
	//
	// example:
	//
	// 1786195753
	IncidentInvestigationEndTime *int64 `json:"IncidentInvestigationEndTime,omitempty" xml:"IncidentInvestigationEndTime,omitempty"`
	// The unique ID of the AI investigation report. You can use this value as the IncidentInvestigationId parameter of this operation to query the report.
	//
	// example:
	//
	// 951a6a0b-14c7-4332-ace6-bff21c******
	IncidentInvestigationId *string `json:"IncidentInvestigationId,omitempty" xml:"IncidentInvestigationId,omitempty"`
	// The output content of the AI investigation, in JSON-formatted string. This field can contain fields such as investigation_report (the investigation process and evidence body, which may use Markdown format). This field is empty if the corresponding output is not generated. Internal fields may change across report versions.
	//
	// example:
	//
	// {"investigation_report":"### Investigation Results\\nSuspicious file detected, but execution evidence is missing."}
	IncidentInvestigationOutput *string `json:"IncidentInvestigationOutput,omitempty" xml:"IncidentInvestigationOutput,omitempty"`
	// The AI investigation report content, in JSON-formatted string. This field can contain fields such as judgement_category (judgment category), ai_judgement_conclusion (judgment conclusion), ai_summary (summary), attack process, and attack graph. The specific content may change across report versions.
	//
	// example:
	//
	// {"judgement_category":"insufficient_information_to_evaluate","ai_judgement_conclusion":"Insufficient evidence to confirm whether the attack chain is complete. Supplement correlated logs for further analysis.","ai_summary":"Suspicious file detected. Insufficient evidence to confirm whether the attack was successful."}
	IncidentInvestigationReport *string `json:"IncidentInvestigationReport,omitempty" xml:"IncidentInvestigationReport,omitempty"`
	// The start time of the AI investigation, in seconds (UNIX timestamp). If the execution start time is not recorded, the creation time of the investigation record is used. This field is empty if neither time exists.
	//
	// example:
	//
	// 1786195533
	IncidentInvestigationStartTime *int64 `json:"IncidentInvestigationStartTime,omitempty" xml:"IncidentInvestigationStartTime,omitempty"`
	// The judgment category of the AI investigation report, corresponding to judgement_category in the report. This field does not indicate the execution status of the asynchronous task. Valid values:
	//
	// - real_attack: Real attack.
	//
	// - blocked: Risk blocked.
	//
	// - false_positive: Suspected false positive.
	//
	// - insufficient_information_to_evaluate: Unable to determine.
	//
	// example:
	//
	// insufficient_information_to_evaluate
	IncidentInvestigationStatus *string `json:"IncidentInvestigationStatus,omitempty" xml:"IncidentInvestigationStatus,omitempty"`
	// The summary of the AI investigation report, corresponding to ai_summary in the investigation results.
	//
	// example:
	//
	// Suspicious file detected. Insufficient evidence to confirm whether the attack was successful
	IncidentInvestigationSummary *string `json:"IncidentInvestigationSummary,omitempty" xml:"IncidentInvestigationSummary,omitempty"`
	// The globally unique UUID of the incident.
	//
	// example:
	//
	// 9f00c254391400f6898c4d4fea******
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
}

func (s GetIncidentInvestigationResponseBodyIncidentInvestigation) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentInvestigationResponseBodyIncidentInvestigation) GoString() string {
	return s.String()
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationAlertName() *string {
	return s.IncidentInvestigationAlertName
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationConclusion() *string {
	return s.IncidentInvestigationConclusion
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationDisplayId() *string {
	return s.IncidentInvestigationDisplayId
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationEndTime() *int64 {
	return s.IncidentInvestigationEndTime
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationId() *string {
	return s.IncidentInvestigationId
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationOutput() *string {
	return s.IncidentInvestigationOutput
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationReport() *string {
	return s.IncidentInvestigationReport
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationStartTime() *int64 {
	return s.IncidentInvestigationStartTime
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationStatus() *string {
	return s.IncidentInvestigationStatus
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentInvestigationSummary() *string {
	return s.IncidentInvestigationSummary
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationAlertName(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationAlertName = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationConclusion(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationConclusion = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationDisplayId(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationDisplayId = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationEndTime(v int64) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationEndTime = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationId(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationId = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationOutput(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationOutput = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationReport(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationReport = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationStartTime(v int64) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationStartTime = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationStatus(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationStatus = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentInvestigationSummary(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentInvestigationSummary = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) SetIncidentUuid(v string) *GetIncidentInvestigationResponseBodyIncidentInvestigation {
	s.IncidentUuid = &v
	return s
}

func (s *GetIncidentInvestigationResponseBodyIncidentInvestigation) Validate() error {
	return dara.Validate(s)
}
