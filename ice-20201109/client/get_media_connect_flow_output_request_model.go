// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetMediaConnectFlowOutputRequest
	GetFlowId() *string
	SetOutputName(v string) *GetMediaConnectFlowOutputRequest
	GetOutputName() *string
}

type GetMediaConnectFlowOutputRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0381f478-7d53-4076-9d5f-27680a6f73e7
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The name of the output that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestOutput
	OutputName *string `json:"OutputName,omitempty" xml:"OutputName,omitempty"`
}

func (s GetMediaConnectFlowOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowOutputRequest) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowOutputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaConnectFlowOutputRequest) GetOutputName() *string {
	return s.OutputName
}

func (s *GetMediaConnectFlowOutputRequest) SetFlowId(v string) *GetMediaConnectFlowOutputRequest {
	s.FlowId = &v
	return s
}

func (s *GetMediaConnectFlowOutputRequest) SetOutputName(v string) *GetMediaConnectFlowOutputRequest {
	s.OutputName = &v
	return s
}

func (s *GetMediaConnectFlowOutputRequest) Validate() error {
	return dara.Validate(s)
}
