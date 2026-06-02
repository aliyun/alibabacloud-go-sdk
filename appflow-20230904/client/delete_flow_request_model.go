// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *DeleteFlowRequest
	GetFlowId() *string
	SetFlowVersion(v int32) *DeleteFlowRequest
	GetFlowVersion() *int32
}

type DeleteFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow-xxxxxxxxx
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 8
	FlowVersion *int32 `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeleteFlowRequest) GetFlowVersion() *int32 {
	return s.FlowVersion
}

func (s *DeleteFlowRequest) SetFlowId(v string) *DeleteFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DeleteFlowRequest) SetFlowVersion(v int32) *DeleteFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *DeleteFlowRequest) Validate() error {
	return dara.Validate(s)
}
