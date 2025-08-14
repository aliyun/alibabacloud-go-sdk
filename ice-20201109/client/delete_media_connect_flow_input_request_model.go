// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *DeleteMediaConnectFlowInputRequest
	GetFlowId() *string
	SetInputName(v string) *DeleteMediaConnectFlowInputRequest
	GetInputName() *string
}

type DeleteMediaConnectFlowInputRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0381f478-7d53-4076-9d5f-27680a6f73e7
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	InputName *string `json:"InputName,omitempty" xml:"InputName,omitempty"`
}

func (s DeleteMediaConnectFlowInputRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowInputRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowInputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeleteMediaConnectFlowInputRequest) GetInputName() *string {
	return s.InputName
}

func (s *DeleteMediaConnectFlowInputRequest) SetFlowId(v string) *DeleteMediaConnectFlowInputRequest {
	s.FlowId = &v
	return s
}

func (s *DeleteMediaConnectFlowInputRequest) SetInputName(v string) *DeleteMediaConnectFlowInputRequest {
	s.InputName = &v
	return s
}

func (s *DeleteMediaConnectFlowInputRequest) Validate() error {
	return dara.Validate(s)
}
