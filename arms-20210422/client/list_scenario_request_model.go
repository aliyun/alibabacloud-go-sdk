// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListScenarioRequest
	GetAppId() *string
	SetName(v string) *ListScenarioRequest
	GetName() *string
	SetRegionId(v string) *ListScenarioRequest
	GetRegionId() *string
	SetScenario(v string) *ListScenarioRequest
	GetScenario() *string
	SetSign(v string) *ListScenarioRequest
	GetSign() *string
}

type ListScenarioRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign     *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s ListScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScenarioRequest) GoString() string {
	return s.String()
}

func (s *ListScenarioRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListScenarioRequest) GetName() *string {
	return s.Name
}

func (s *ListScenarioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListScenarioRequest) GetScenario() *string {
	return s.Scenario
}

func (s *ListScenarioRequest) GetSign() *string {
	return s.Sign
}

func (s *ListScenarioRequest) SetAppId(v string) *ListScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ListScenarioRequest) SetName(v string) *ListScenarioRequest {
	s.Name = &v
	return s
}

func (s *ListScenarioRequest) SetRegionId(v string) *ListScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ListScenarioRequest) SetScenario(v string) *ListScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ListScenarioRequest) SetSign(v string) *ListScenarioRequest {
	s.Sign = &v
	return s
}

func (s *ListScenarioRequest) Validate() error {
	return dara.Validate(s)
}
