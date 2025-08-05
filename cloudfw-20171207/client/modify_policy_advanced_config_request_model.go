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
	// The IP addresses. The versions of the IP addresses must be the same. You can specify a maximum of 100 IP addresses.
	Eips []*string `json:"Eips,omitempty" xml:"Eips,omitempty" type:"Repeated"`
	// Specifies whether to enable the strict mode for the access control policy. Valid values:
	//
	// 	- **on**: enables the strict mode.
	//
	// 	- **off**: disables the strict mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
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
