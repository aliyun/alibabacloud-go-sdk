// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *ResumeMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *ResumeMediaConnectFlowOutputRequest
	GetOutputName() *string
}

type ResumeMediaConnectFlowOutputRequest struct {
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s ResumeMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *ResumeMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *ResumeMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *ResumeMediaConnectFlowOutputRequest) SetFlowId(v string) *ResumeMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputRequest) SetOutputName(v string) *ResumeMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
