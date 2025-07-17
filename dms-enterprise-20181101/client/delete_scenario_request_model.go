// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenarioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScenarioId(v int64) *DeleteScenarioRequest
	GetScenarioId() *int64
	SetTid(v int64) *DeleteScenarioRequest
	GetTid() *int64
}

type DeleteScenarioRequest struct {
	// The ID of the business scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	ScenarioId *int64 `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteScenarioRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenarioRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenarioRequest) GetScenarioId() *int64 {
	return s.ScenarioId
}

func (s *DeleteScenarioRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteScenarioRequest) SetScenarioId(v int64) *DeleteScenarioRequest {
	s.ScenarioId = &v
	return s
}

func (s *DeleteScenarioRequest) SetTid(v int64) *DeleteScenarioRequest {
	s.Tid = &v
	return s
}

func (s *DeleteScenarioRequest) Validate() error {
	return dara.Validate(s)
}
