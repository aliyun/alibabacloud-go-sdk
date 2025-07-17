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
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
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
