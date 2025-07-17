// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateScenarioRequest
	GetDescription() *string
	SetScenarioName(v string) *CreateScenarioRequest
	GetScenarioName() *string
	SetTid(v int64) *CreateScenarioRequest
	GetTid() *int64
}

type CreateScenarioRequest struct {
	// The description of the business scenario.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the business scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business scenario - test
	ScenarioName *string `json:"ScenarioName,omitempty" xml:"ScenarioName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScenarioRequest) GoString() string {
	return s.String()
}

func (s *CreateScenarioRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScenarioRequest) GetScenarioName() *string {
	return s.ScenarioName
}

func (s *CreateScenarioRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateScenarioRequest) SetDescription(v string) *CreateScenarioRequest {
	s.Description = &v
	return s
}

func (s *CreateScenarioRequest) SetScenarioName(v string) *CreateScenarioRequest {
	s.ScenarioName = &v
	return s
}

func (s *CreateScenarioRequest) SetTid(v int64) *CreateScenarioRequest {
	s.Tid = &v
	return s
}

func (s *CreateScenarioRequest) Validate() error {
	return dara.Validate(s)
}
