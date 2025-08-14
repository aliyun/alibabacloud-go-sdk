// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *DeleteMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *DeleteMediaConnectFlowOutputRequest
	GetOutputName() *string
}

type DeleteMediaConnectFlowOutputRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the output that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s DeleteMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeleteMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *DeleteMediaConnectFlowOutputRequest) SetFlowId(v string) *DeleteMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputRequest) SetOutputName(v string) *DeleteMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
