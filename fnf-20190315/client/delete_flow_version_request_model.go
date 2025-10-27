// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *DeleteFlowVersionRequest
	GetFlowName() *string
	SetFlowVersion(v string) *DeleteFlowVersionRequest
	GetFlowVersion() *string
}

type DeleteFlowVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example-flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s DeleteFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DeleteFlowVersionRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *DeleteFlowVersionRequest) SetFlowName(v string) *DeleteFlowVersionRequest {
	s.FlowName = &v
	return s
}

func (s *DeleteFlowVersionRequest) SetFlowVersion(v string) *DeleteFlowVersionRequest {
	s.FlowVersion = &v
	return s
}

func (s *DeleteFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
