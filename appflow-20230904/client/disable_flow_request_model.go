// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *DisableFlowRequest
	GetFlowId() *string
	SetFlowVersion(v int32) *DisableFlowRequest
	GetFlowVersion() *int32
}

type DisableFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// flow-xxxxxx
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// 1
	FlowVersion *int32 `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s DisableFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableFlowRequest) GoString() string {
	return s.String()
}

func (s *DisableFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DisableFlowRequest) GetFlowVersion() *int32 {
	return s.FlowVersion
}

func (s *DisableFlowRequest) SetFlowId(v string) *DisableFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DisableFlowRequest) SetFlowVersion(v int32) *DisableFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *DisableFlowRequest) Validate() error {
	return dara.Validate(s)
}
