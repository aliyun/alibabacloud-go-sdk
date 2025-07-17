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
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@28f515462******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the business monitoring job.
	//
	// This parameter is required.
	//
	// example:
	//
	// pro-content
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjaikou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scenario where the business monitoring job is used. Valid values:
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
	// The code of the business monitoring job. Set this parameter when you know the code of the business monitoring job you want to query.
	//
	// example:
	//
	// a9f8****
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
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
