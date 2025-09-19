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
	// Specifies whether to enable the feature. Valid values:
	//
	// 	- **on**: yes
	//
	// 	- **off**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to configure assets for the feature. Default value: **false**. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// >  This parameter takes effect only when you set **Config*	- to **on**.
	//
	// example:
	//
	// true
	NoTargetAsOn *bool `json:"NoTargetAsOn,omitempty" xml:"NoTargetAsOn,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 222.178.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- **auto_breaking**: Anti-Virus
	//
	// 	- **ransomware_breaking**: Anti-ransomware (Bait Capture)
	//
	// 	- **webshell_cloud_breaking**: Webshell Protection
	//
	// 	- **alinet**: Behavior prevention
	//
	// 	- **k8s_log_analysis**: K8s Threat Detection
	//
	// 	- **alisecguard**: Defense mode for Client Protection
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
