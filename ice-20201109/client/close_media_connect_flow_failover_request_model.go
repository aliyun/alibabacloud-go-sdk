// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMediaConnectFlowFailoverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *CloseMediaConnectFlowFailoverRequest
	GetFlowId() *string
}

type CloseMediaConnectFlowFailoverRequest struct {
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s CloseMediaConnectFlowFailoverRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseMediaConnectFlowFailoverRequest) GoString() string {
	return s.String()
}

func (s *CloseMediaConnectFlowFailoverRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *CloseMediaConnectFlowFailoverRequest) SetFlowId(v string) *CloseMediaConnectFlowFailoverRequest {
	s.FlowId = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverRequest) Validate() error {
	return dara.Validate(s)
}
