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
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@28f515462f******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The configuration of the business monitoring job. The value is a JSON string. For more information about this parameter, see the following additional information about the **Config*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"rpcType":"0","nameMatchType":"EQUALS","service":"/api/pop/test","operator":"and","filterItems":[{"type":"HttpHeaders","key":"uid","opt":"==","value":"123456789"}],"group":{"type":"HttpRequestParameters","key":"name"}}
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the business monitoring job.
	//
	// This parameter is required.
	//
	// example:
	//
	// ScenarioName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjaikou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario where you want to use the business monitoring job. Valid values:
	//
	// 	- `USER-DEFINED`: user-defined. This is the default value.
	//
	// 	- `EDAS-ROLLOUT`: application release in Enterprise Distributed Application Service (EDAS)
	//
	// 	- `OAM-ROLLOUT`: application release based on Open Application Model (OAM)
	//
	// 	- `MSC-CANARY`: canary release based on Microservice Engine (MSE)
	//
	// example:
	//
	// USER-DEFINED
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The code of the business monitoring job. This parameter is not required when you create a business monitoring job. However, this parameter is required when you update a business monitoring job.
	//
	// example:
	//
	// a9f8****
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// Specifies whether to record business parameters to the trace marked with the coloring sign.
	//
	// 	- `true`
	//
	// 	- `false`: This is the default value.
	//
	// example:
	//
	// false
	SnDump *bool `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	// Specifies whether traffic in the trace marked with the coloring sign is all collected.
	//
	// 	- `true`
	//
	// 	- `false`: This is the default value.
	//
	// example:
	//
	// false
	SnForce *bool `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	// Specifies whether to count traffic based on the coloring sign.
	//
	// 	- `true`
	//
	// 	- `false`: This is the default value.
	//
	// example:
	//
	// false
	SnStat *bool `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	// Specifies whether the coloring sign is transparently passed down to downstream application nodes in the trace.
	//
	// 	- `true`
	//
	// 	- `false`: This is the default value.
	//
	// example:
	//
	// false
	SnTransfer *bool `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	// Specifies whether the operation is an update operation.
	//
	// 	- `true`: update
	//
	// 	- `false`: insert
	//
	// This parameter is required.
	//
	// example:
	//
	// false
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
