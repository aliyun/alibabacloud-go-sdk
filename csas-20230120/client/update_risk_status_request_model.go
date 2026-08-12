// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRiskConfirm(v string) *UpdateRiskStatusRequest
	GetRiskConfirm() *string
	SetRiskConfirmDesc(v string) *UpdateRiskStatusRequest
	GetRiskConfirmDesc() *string
	SetRiskId(v string) *UpdateRiskStatusRequest
	GetRiskId() *string
	SetRiskScene(v string) *UpdateRiskStatusRequest
	GetRiskScene() *string
	SetStatus(v string) *UpdateRiskStatusRequest
	GetStatus() *string
}

type UpdateRiskStatusRequest struct {
	// The manually confirmed risk conclusion. This parameter is required when `Status` is set to `Processed`. Do not specify this parameter when `Status` is set to `Unprocess` or `Processing`. Valid values:
	//
	// 	- `Risk`: Confirmed as risky.
	//
	// 	- `Ignore`: Confirmed as not risky.
	//
	// 	- `Invalid`: Confirmed as a false positive.
	//
	// example:
	//
	// Risk
	RiskConfirm *string `json:"RiskConfirm,omitempty" xml:"RiskConfirm,omitempty"`
	// The description of the risk event handling. The length must be 1 to 128 characters.
	//
	// example:
	//
	// After verification, this risk event is a real risk
	RiskConfirmDesc *string `json:"RiskConfirmDesc,omitempty" xml:"RiskConfirmDesc,omitempty"`
	// The risk event ID. You can obtain the value from the following operation:
	//
	// 	- `ListRiskItems`: Queries the list of risk events.
	//
	// example:
	//
	// 69ef648034cf53d7bac7a9c9c912****
	RiskId *string `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The risk scenario. This parameter is optional. If not specified, the system automatically populates it based on RiskId. Valid values:
	//
	// 	- account_share: Account sharing.
	//
	// 	- account_stolen: Account stolen.
	//
	// 	- device_share: Device sharing.
	//
	// 	- remote_logon: Remote logon.
	//
	// 	- sensitive_data_leakage: Sensitive data leakage.
	//
	// 	- lateral_scanning: Lateral scanning.
	//
	// 	- ai_skill_malware: Malicious skill.
	//
	// 	- ai_config_check: AI configuration check.
	//
	// 	- openclaw_vulnerability: OpenClaw vulnerability.
	//
	// example:
	//
	// account_stolen
	RiskScene *string `json:"RiskScene,omitempty" xml:"RiskScene,omitempty"`
	// The handling status of the risk event. Valid values:
	//
	// 	- `Unprocess`: Unprocessed.
	//
	// 	- `Processing`: Being processed.
	//
	// 	- `Processed`: Processed.
	//
	// example:
	//
	// Processed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateRiskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRiskStatusRequest) GetRiskConfirm() *string {
	return s.RiskConfirm
}

func (s *UpdateRiskStatusRequest) GetRiskConfirmDesc() *string {
	return s.RiskConfirmDesc
}

func (s *UpdateRiskStatusRequest) GetRiskId() *string {
	return s.RiskId
}

func (s *UpdateRiskStatusRequest) GetRiskScene() *string {
	return s.RiskScene
}

func (s *UpdateRiskStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateRiskStatusRequest) SetRiskConfirm(v string) *UpdateRiskStatusRequest {
	s.RiskConfirm = &v
	return s
}

func (s *UpdateRiskStatusRequest) SetRiskConfirmDesc(v string) *UpdateRiskStatusRequest {
	s.RiskConfirmDesc = &v
	return s
}

func (s *UpdateRiskStatusRequest) SetRiskId(v string) *UpdateRiskStatusRequest {
	s.RiskId = &v
	return s
}

func (s *UpdateRiskStatusRequest) SetRiskScene(v string) *UpdateRiskStatusRequest {
	s.RiskScene = &v
	return s
}

func (s *UpdateRiskStatusRequest) SetStatus(v string) *UpdateRiskStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateRiskStatusRequest) Validate() error {
	return dara.Validate(s)
}
