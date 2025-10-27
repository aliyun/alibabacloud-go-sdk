// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *DeleteFlowAliasRequest
	GetFlowName() *string
	SetName(v string) *DeleteFlowAliasRequest
	GetName() *string
}

type DeleteFlowAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alias_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFlowAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowAliasRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowAliasRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DeleteFlowAliasRequest) GetName() *string {
	return s.Name
}

func (s *DeleteFlowAliasRequest) SetFlowName(v string) *DeleteFlowAliasRequest {
	s.FlowName = &v
	return s
}

func (s *DeleteFlowAliasRequest) SetName(v string) *DeleteFlowAliasRequest {
	s.Name = &v
	return s
}

func (s *DeleteFlowAliasRequest) Validate() error {
	return dara.Validate(s)
}
