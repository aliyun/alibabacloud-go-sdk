// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateScenarioRequest
	GetDescription() *string
	SetScenarioId(v string) *UpdateScenarioRequest
	GetScenarioId() *string
	SetScenarioName(v string) *UpdateScenarioRequest
	GetScenarioName() *string
	SetTid(v int64) *UpdateScenarioRequest
	GetTid() *int64
}

type UpdateScenarioRequest struct {
	// The description of the business scenario.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the business scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
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
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScenarioRequest) GoString() string {
	return s.String()
}

func (s *UpdateScenarioRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateScenarioRequest) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *UpdateScenarioRequest) GetScenarioName() *string {
	return s.ScenarioName
}

func (s *UpdateScenarioRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateScenarioRequest) SetDescription(v string) *UpdateScenarioRequest {
	s.Description = &v
	return s
}

func (s *UpdateScenarioRequest) SetScenarioId(v string) *UpdateScenarioRequest {
	s.ScenarioId = &v
	return s
}

func (s *UpdateScenarioRequest) SetScenarioName(v string) *UpdateScenarioRequest {
	s.ScenarioName = &v
	return s
}

func (s *UpdateScenarioRequest) SetTid(v int64) *UpdateScenarioRequest {
	s.Tid = &v
	return s
}

func (s *UpdateScenarioRequest) Validate() error {
	return dara.Validate(s)
}
