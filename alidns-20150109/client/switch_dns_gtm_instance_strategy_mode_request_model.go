// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDnsGtmInstanceStrategyModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchDnsGtmInstanceStrategyModeRequest
	GetInstanceId() *string
	SetLang(v string) *SwitchDnsGtmInstanceStrategyModeRequest
	GetLang() *string
	SetStrategyMode(v string) *SwitchDnsGtmInstanceStrategyModeRequest
	GetStrategyMode() *string
}

type SwitchDnsGtmInstanceStrategyModeRequest struct {
	// The ID of the GTM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The access policy type. Valid values:
	//
	// 	- GEO: geographical location-based
	//
	// 	- LATENCY: latency-based
	//
	// This parameter is required.
	//
	// example:
	//
	// GEO
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
}

func (s SwitchDnsGtmInstanceStrategyModeRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetInstanceId(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetLang(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.Lang = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) SetStrategyMode(v string) *SwitchDnsGtmInstanceStrategyModeRequest {
	s.StrategyMode = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeRequest) Validate() error {
	return dara.Validate(s)
}
