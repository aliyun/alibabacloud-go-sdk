// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetFlowRequest
	GetFlowId() *string
	SetFlowVersion(v string) *GetFlowRequest
	GetFlowVersion() *string
}

type GetFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow-xxxxx
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 6
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s GetFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *GetFlowRequest) SetFlowId(v string) *GetFlowRequest {
	s.FlowId = &v
	return s
}

func (s *GetFlowRequest) SetFlowVersion(v string) *GetFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *GetFlowRequest) Validate() error {
	return dara.Validate(s)
}
