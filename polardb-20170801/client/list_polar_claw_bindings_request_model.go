// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentList(v []*string) *ListPolarClawBindingsRequest
	GetAgentList() []*string
	SetApplicationId(v string) *ListPolarClawBindingsRequest
	GetApplicationId() *string
}

type ListPolarClawBindingsRequest struct {
	// example:
	//
	// work,research
	AgentList []*string `json:"AgentList,omitempty" xml:"AgentList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListPolarClawBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListPolarClawBindingsRequest) GetAgentList() []*string {
	return s.AgentList
}

func (s *ListPolarClawBindingsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolarClawBindingsRequest) SetAgentList(v []*string) *ListPolarClawBindingsRequest {
	s.AgentList = v
	return s
}

func (s *ListPolarClawBindingsRequest) SetApplicationId(v string) *ListPolarClawBindingsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPolarClawBindingsRequest) Validate() error {
	return dara.Validate(s)
}
