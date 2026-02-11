// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ApplyScenarioRequest
	GetAppId() *string
	SetConfig(v map[string]interface{}) *ApplyScenarioRequest
	GetConfig() map[string]interface{}
	SetName(v string) *ApplyScenarioRequest
	GetName() *string
	SetRegionId(v string) *ApplyScenarioRequest
	GetRegionId() *string
	SetScenario(v string) *ApplyScenarioRequest
	GetScenario() *string
	SetSign(v string) *ApplyScenarioRequest
	GetSign() *string
	SetSnDump(v bool) *ApplyScenarioRequest
	GetSnDump() *bool
	SetSnForce(v bool) *ApplyScenarioRequest
	GetSnForce() *bool
	SetSnStat(v bool) *ApplyScenarioRequest
	GetSnStat() *bool
	SetSnTransfer(v bool) *ApplyScenarioRequest
	GetSnTransfer() *bool
	SetUpdateOption(v bool) *ApplyScenarioRequest
	GetUpdateOption() *bool
}

type ApplyScenarioRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
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

func (s ApplyScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyScenarioRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioRequest) GetAppId() *string {
	return s.AppId
}

func (s *ApplyScenarioRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ApplyScenarioRequest) GetName() *string {
	return s.Name
}

func (s *ApplyScenarioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyScenarioRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ApplyScenarioRequest) GetSign() *string {
	return s.Sign
}

func (s *ApplyScenarioRequest) GetSnDump() *bool {
	return s.SnDump
}

func (s *ApplyScenarioRequest) GetSnForce() *bool {
	return s.SnForce
}

func (s *ApplyScenarioRequest) GetSnStat() *bool {
	return s.SnStat
}

func (s *ApplyScenarioRequest) GetSnTransfer() *bool {
	return s.SnTransfer
}

func (s *ApplyScenarioRequest) GetUpdateOption() *bool {
	return s.UpdateOption
}

func (s *ApplyScenarioRequest) SetAppId(v string) *ApplyScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioRequest) SetConfig(v map[string]interface{}) *ApplyScenarioRequest {
	s.Config = v
	return s
}

func (s *ApplyScenarioRequest) SetName(v string) *ApplyScenarioRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioRequest) SetRegionId(v string) *ApplyScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioRequest) SetScenario(v string) *ApplyScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioRequest) SetSign(v string) *ApplyScenarioRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnDump(v bool) *ApplyScenarioRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnForce(v bool) *ApplyScenarioRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnStat(v bool) *ApplyScenarioRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnTransfer(v bool) *ApplyScenarioRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioRequest) SetUpdateOption(v bool) *ApplyScenarioRequest {
	s.UpdateOption = &v
	return s
}

func (s *ApplyScenarioRequest) Validate() error {
	return dara.Validate(s)
}
