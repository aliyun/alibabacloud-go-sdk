// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchSecurityProxyRequest
	GetLang() *string
	SetProxyId(v string) *SwitchSecurityProxyRequest
	GetProxyId() *string
	SetSwitch(v string) *SwitchSecurityProxyRequest
	GetSwitch() *string
}

type SwitchSecurityProxyRequest struct {
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
	// The ID of the NAT firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// proxy-natbf******
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The security protection switch. Valid values:
	//
	// - open: Enables security protection.
	//
	// - close: Disables security protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	Switch *string `json:"Switch,omitempty" xml:"Switch,omitempty"`
}

func (s SwitchSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *SwitchSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchSecurityProxyRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *SwitchSecurityProxyRequest) GetSwitch() *string {
	return s.Switch
}

func (s *SwitchSecurityProxyRequest) SetLang(v string) *SwitchSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *SwitchSecurityProxyRequest) SetProxyId(v string) *SwitchSecurityProxyRequest {
	s.ProxyId = &v
	return s
}

func (s *SwitchSecurityProxyRequest) SetSwitch(v string) *SwitchSecurityProxyRequest {
	s.Switch = &v
	return s
}

func (s *SwitchSecurityProxyRequest) Validate() error {
	return dara.Validate(s)
}
