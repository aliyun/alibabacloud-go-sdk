// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousOverallConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *OperateSuspiciousOverallConfigRequest
	GetConfig() *string
	SetLang(v string) *OperateSuspiciousOverallConfigRequest
	GetLang() *string
	SetNoTargetAsOn(v bool) *OperateSuspiciousOverallConfigRequest
	GetNoTargetAsOn() *bool
	SetSourceIp(v string) *OperateSuspiciousOverallConfigRequest
	GetSourceIp() *string
	SetType(v string) *OperateSuspiciousOverallConfigRequest
	GetType() *string
}

type OperateSuspiciousOverallConfigRequest struct {
	// The switch status. Valid values:
	//
	// - **on**: Enable
	//
	// - **off**: Disable
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether asset configuration is required. Default value: **false**. Valid values:
	//
	// - **true**: Required
	//
	// - **false**: Not required
	//
	// > This value takes effect only when **config*	- is set to **on**.
	//
	// example:
	//
	// true
	NoTargetAsOn *bool `json:"NoTargetAsOn,omitempty" xml:"NoTargetAsOn,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 222.178.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The switch type. Valid values:
	//
	// - **auto_breaking**: Anti-virus
	//
	// - **ransomware_breaking**: Anti-ransomware (bait capture)
	//
	// - **webshell_cloud_breaking**: Website backdoor connection defense
	//
	// - **alinet**: Malicious network behavior defense
	//
	// - **k8s_log_analysis**: Container K8s threat detection
	//
	// - **alisecguard**: Client self-protection defense mode
	//
	// This parameter is required.
	//
	// example:
	//
	// k8s_log_analysis
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateSuspiciousOverallConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousOverallConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousOverallConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *OperateSuspiciousOverallConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateSuspiciousOverallConfigRequest) GetNoTargetAsOn() *bool {
	return s.NoTargetAsOn
}

func (s *OperateSuspiciousOverallConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *OperateSuspiciousOverallConfigRequest) GetType() *string {
	return s.Type
}

func (s *OperateSuspiciousOverallConfigRequest) SetConfig(v string) *OperateSuspiciousOverallConfigRequest {
	s.Config = &v
	return s
}

func (s *OperateSuspiciousOverallConfigRequest) SetLang(v string) *OperateSuspiciousOverallConfigRequest {
	s.Lang = &v
	return s
}

func (s *OperateSuspiciousOverallConfigRequest) SetNoTargetAsOn(v bool) *OperateSuspiciousOverallConfigRequest {
	s.NoTargetAsOn = &v
	return s
}

func (s *OperateSuspiciousOverallConfigRequest) SetSourceIp(v string) *OperateSuspiciousOverallConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OperateSuspiciousOverallConfigRequest) SetType(v string) *OperateSuspiciousOverallConfigRequest {
	s.Type = &v
	return s
}

func (s *OperateSuspiciousOverallConfigRequest) Validate() error {
	return dara.Validate(s)
}
