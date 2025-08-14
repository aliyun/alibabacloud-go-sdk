// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *UpdateMediaConnectFlowStatusRequest
	GetFlowId() *string
	SetStatus(v string) *UpdateMediaConnectFlowStatusRequest
	GetStatus() *string
}

type UpdateMediaConnectFlowStatusRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The flow state. Valid values:
	//
	// 	- online: starts the flow.
	//
	// 	- offline: stops the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateMediaConnectFlowStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowStatusRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *UpdateMediaConnectFlowStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateMediaConnectFlowStatusRequest) SetFlowId(v string) *UpdateMediaConnectFlowStatusRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusRequest) SetStatus(v string) *UpdateMediaConnectFlowStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusRequest) Validate() error {
	return dara.Validate(s)
}
