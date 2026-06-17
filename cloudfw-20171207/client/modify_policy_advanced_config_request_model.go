// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolicyAdvancedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEips(v []*string) *ModifyPolicyAdvancedConfigRequest
	GetEips() []*string
	SetInternetSwitch(v string) *ModifyPolicyAdvancedConfigRequest
	GetInternetSwitch() *string
	SetLang(v string) *ModifyPolicyAdvancedConfigRequest
	GetLang() *string
	SetSourceIp(v string) *ModifyPolicyAdvancedConfigRequest
	GetSourceIp() *string
}

type ModifyPolicyAdvancedConfigRequest struct {
	// A list of IP addresses. The IP addresses must use the same protocol version. You can specify up to 100 IP addresses.
	Eips []*string `json:"Eips,omitempty" xml:"Eips,omitempty" type:"Repeated"`
	// Specifies whether to enable or disable the strict mode for access control policies. Valid values:
	//
	// - **on**: Enables strict mode.
	//
	// - **off**: Disables strict mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the visitor.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyPolicyAdvancedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigRequest) GetEips() []*string {
	return s.Eips
}

func (s *ModifyPolicyAdvancedConfigRequest) GetInternetSwitch() *string {
	return s.InternetSwitch
}

func (s *ModifyPolicyAdvancedConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyPolicyAdvancedConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifyPolicyAdvancedConfigRequest) SetEips(v []*string) *ModifyPolicyAdvancedConfigRequest {
	s.Eips = v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetInternetSwitch(v string) *ModifyPolicyAdvancedConfigRequest {
	s.InternetSwitch = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetLang(v string) *ModifyPolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetSourceIp(v string) *ModifyPolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) Validate() error {
	return dara.Validate(s)
}
