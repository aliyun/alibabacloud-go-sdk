// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenMediaConnectFlowFailoverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *OpenMediaConnectFlowFailoverRequest
	GetFlowId() *string
}

type OpenMediaConnectFlowFailoverRequest struct {
	// The ID of the MediaConnect flow.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s OpenMediaConnectFlowFailoverRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenMediaConnectFlowFailoverRequest) GoString() string {
	return s.String()
}

func (s *OpenMediaConnectFlowFailoverRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *OpenMediaConnectFlowFailoverRequest) SetFlowId(v string) *OpenMediaConnectFlowFailoverRequest {
	s.FlowId = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverRequest) Validate() error {
	return dara.Validate(s)
}
