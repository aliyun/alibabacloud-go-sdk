// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetMediaConnectFlowInputRequest
	GetFlowId() *string
	SetWithInternalVip(v string) *GetMediaConnectFlowInputRequest
	GetWithInternalVip() *string
}

type GetMediaConnectFlowInputRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId          *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	WithInternalVip *string `json:"WithInternalVip,omitempty" xml:"WithInternalVip,omitempty"`
}

func (s GetMediaConnectFlowInputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowInputRequest) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowInputRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaConnectFlowInputRequest) GetWithInternalVip() *string {
	return s.WithInternalVip
}

func (s *GetMediaConnectFlowInputRequest) SetFlowId(v string) *GetMediaConnectFlowInputRequest {
	s.FlowId = &v
	return s
}

func (s *GetMediaConnectFlowInputRequest) SetWithInternalVip(v string) *GetMediaConnectFlowInputRequest {
	s.WithInternalVip = &v
	return s
}

func (s *GetMediaConnectFlowInputRequest) Validate() error {
	return dara.Validate(s)
}
