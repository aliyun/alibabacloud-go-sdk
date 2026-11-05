// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSuspiciousTargetConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *OperateSuspiciousTargetConfigRequest
	GetLang() *string
	SetSourceIp(v string) *OperateSuspiciousTargetConfigRequest
	GetSourceIp() *string
	SetTargetOperations(v string) *OperateSuspiciousTargetConfigRequest
	GetTargetOperations() *string
	SetTargetType(v string) *OperateSuspiciousTargetConfigRequest
	GetTargetType() *string
	SetType(v string) *OperateSuspiciousTargetConfigRequest
	GetType() *string
}

type OperateSuspiciousTargetConfigRequest struct {
	// Specifies the language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Specifies the proactive defense parameters for servers. The following parameters are included:
	//
	// - **targetType**: the dimension of the defense configuration. Currently, only the UUID dimension is supported. Fixed value: **uuid**.
	//
	// - **target**: the UUID of the server for which you want to configure defense settings.
	//
	// - **flag**: specifies whether to enable or disable proactive defense for the server. Valid values: **add*	- (enable) and **del*	- (disable).
	//
	// This parameter is required.
	//
	// example:
	//
	// "[{"targetType":"uuid","target":"0585f81a-dd84-4ddf-9971-f59d12345678","flag":"add"},{"targetType":"uuid","target":"01acfd9d-e6a4-4e61-b9eb-aae012345678","flag":"add"},{"targetType":"uuid","target":"04a0e735-ad32-4835-b635-045812345678","flag":"add"}]"
	TargetOperations *string `json:"TargetOperations,omitempty" xml:"TargetOperations,omitempty"`
	// Specifies the dimension for managing proactive defense. Currently, only the server UUID dimension is supported.
	//
	// Valid value: **uuid**.
	//
	// This parameter is required.
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// Specifies the type of proactive defense. Valid values:
	//
	// - **auto_breaking**: automatic interception
	//
	// - **webshell_cloud_breaking**: web shell connection defense
	//
	// - **alinet**: malicious network behavior defense
	//
	// - **ransomware_breaking**: anti-ransomware (bait capture)
	//
	// - **alisecguard**: client self-protection
	//
	// This parameter is required.
	//
	// example:
	//
	// auto_breaking
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateSuspiciousTargetConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateSuspiciousTargetConfigRequest) GoString() string {
	return s.String()
}

func (s *OperateSuspiciousTargetConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *OperateSuspiciousTargetConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *OperateSuspiciousTargetConfigRequest) GetTargetOperations() *string {
	return s.TargetOperations
}

func (s *OperateSuspiciousTargetConfigRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *OperateSuspiciousTargetConfigRequest) GetType() *string {
	return s.Type
}

func (s *OperateSuspiciousTargetConfigRequest) SetLang(v string) *OperateSuspiciousTargetConfigRequest {
	s.Lang = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetSourceIp(v string) *OperateSuspiciousTargetConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetTargetOperations(v string) *OperateSuspiciousTargetConfigRequest {
	s.TargetOperations = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetTargetType(v string) *OperateSuspiciousTargetConfigRequest {
	s.TargetType = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) SetType(v string) *OperateSuspiciousTargetConfigRequest {
	s.Type = &v
	return s
}

func (s *OperateSuspiciousTargetConfigRequest) Validate() error {
	return dara.Validate(s)
}
