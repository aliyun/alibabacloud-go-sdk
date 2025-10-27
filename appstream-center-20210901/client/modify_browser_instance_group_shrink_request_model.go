// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBrowserInstanceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserConfigShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetBrowserConfigShrink() *string
	SetBrowserInstanceGroupId(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetBrowserInstanceGroupId() *string
	SetCloudBrowserName(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetCloudBrowserName() *string
	SetNetworkShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetNetworkShrink() *string
	SetPolicyShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetPolicyShrink() *string
	SetTimersShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest
	GetTimersShrink() *string
}

type ModifyBrowserInstanceGroupShrinkRequest struct {
	BrowserConfigShrink *string `json:"BrowserConfig,omitempty" xml:"BrowserConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// big-0bz55ixxxxx9xi9w9
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// example:
	//
	// BrowserTest
	CloudBrowserName *string `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	NetworkShrink    *string `json:"Network,omitempty" xml:"Network,omitempty"`
	PolicyShrink     *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	TimersShrink     *string `json:"Timers,omitempty" xml:"Timers,omitempty"`
}

func (s ModifyBrowserInstanceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBrowserInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetBrowserConfigShrink() *string {
	return s.BrowserConfigShrink
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetCloudBrowserName() *string {
	return s.CloudBrowserName
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetNetworkShrink() *string {
	return s.NetworkShrink
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) GetTimersShrink() *string {
	return s.TimersShrink
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetBrowserConfigShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.BrowserConfigShrink = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetBrowserInstanceGroupId(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetCloudBrowserName(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.CloudBrowserName = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetNetworkShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.NetworkShrink = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetPolicyShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) SetTimersShrink(v string) *ModifyBrowserInstanceGroupShrinkRequest {
	s.TimersShrink = &v
	return s
}

func (s *ModifyBrowserInstanceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
