// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScenarioShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ApplyScenarioShrinkRequest
	GetAppId() *string
	SetConfigShrink(v string) *ApplyScenarioShrinkRequest
	GetConfigShrink() *string
	SetName(v string) *ApplyScenarioShrinkRequest
	GetName() *string
	SetRegionId(v string) *ApplyScenarioShrinkRequest
	GetRegionId() *string
	SetScenario(v string) *ApplyScenarioShrinkRequest
	GetScenario() *string
	SetSign(v string) *ApplyScenarioShrinkRequest
	GetSign() *string
	SetSnDump(v bool) *ApplyScenarioShrinkRequest
	GetSnDump() *bool
	SetSnForce(v bool) *ApplyScenarioShrinkRequest
	GetSnForce() *bool
	SetSnStat(v bool) *ApplyScenarioShrinkRequest
	GetSnStat() *bool
	SetSnTransfer(v bool) *ApplyScenarioShrinkRequest
	GetSnTransfer() *bool
	SetUpdateOption(v bool) *ApplyScenarioShrinkRequest
	GetUpdateOption() *bool
}

type ApplyScenarioShrinkRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario   *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SnDump     *bool   `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce    *bool   `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	SnStat     *bool   `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnTransfer *bool   `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	// This parameter is required.
	UpdateOption *bool `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyScenarioShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ApplyScenarioShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *ApplyScenarioShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ApplyScenarioShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyScenarioShrinkRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ApplyScenarioShrinkRequest) GetSign() *string {
	return s.Sign
}

func (s *ApplyScenarioShrinkRequest) GetSnDump() *bool {
	return s.SnDump
}

func (s *ApplyScenarioShrinkRequest) GetSnForce() *bool {
	return s.SnForce
}

func (s *ApplyScenarioShrinkRequest) GetSnStat() *bool {
	return s.SnStat
}

func (s *ApplyScenarioShrinkRequest) GetSnTransfer() *bool {
	return s.SnTransfer
}

func (s *ApplyScenarioShrinkRequest) GetUpdateOption() *bool {
	return s.UpdateOption
}

func (s *ApplyScenarioShrinkRequest) SetAppId(v string) *ApplyScenarioShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetConfigShrink(v string) *ApplyScenarioShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetName(v string) *ApplyScenarioShrinkRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetRegionId(v string) *ApplyScenarioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetScenario(v string) *ApplyScenarioShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSign(v string) *ApplyScenarioShrinkRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnDump(v bool) *ApplyScenarioShrinkRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnForce(v bool) *ApplyScenarioShrinkRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnStat(v bool) *ApplyScenarioShrinkRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnTransfer(v bool) *ApplyScenarioShrinkRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetUpdateOption(v bool) *ApplyScenarioShrinkRequest {
	s.UpdateOption = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) Validate() error {
	return dara.Validate(s)
}
