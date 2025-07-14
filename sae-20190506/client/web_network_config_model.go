// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetInternetAccess(v bool) *WebNetworkConfig
	GetInternetAccess() *bool
	SetSecurityGroupId(v string) *WebNetworkConfig
	GetSecurityGroupId() *string
	SetVSwitchIds(v []*string) *WebNetworkConfig
	GetVSwitchIds() []*string
}

type WebNetworkConfig struct {
	// example:
	//
	// true
	InternetAccess *bool `json:"InternetAccess,omitempty" xml:"InternetAccess,omitempty"`
	// example:
	//
	// sg-bp18hj1wtxgy3b0***
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchIds      []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s WebNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s WebNetworkConfig) GoString() string {
	return s.String()
}

func (s *WebNetworkConfig) GetInternetAccess() *bool {
	return s.InternetAccess
}

func (s *WebNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *WebNetworkConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *WebNetworkConfig) SetInternetAccess(v bool) *WebNetworkConfig {
	s.InternetAccess = &v
	return s
}

func (s *WebNetworkConfig) SetSecurityGroupId(v string) *WebNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *WebNetworkConfig) SetVSwitchIds(v []*string) *WebNetworkConfig {
	s.VSwitchIds = v
	return s
}

func (s *WebNetworkConfig) Validate() error {
	return dara.Validate(s)
}
