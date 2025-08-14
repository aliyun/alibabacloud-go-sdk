// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *ForbidMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *ForbidMediaConnectFlowOutputRequest
	GetOutputName() *string
}

type ForbidMediaConnectFlowOutputRequest struct {
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s ForbidMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ForbidMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *ForbidMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ForbidMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *ForbidMediaConnectFlowOutputRequest) SetFlowId(v string) *ForbidMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputRequest) SetOutputName(v string) *ForbidMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *ForbidMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
