// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetMediaConnectFlowRequest
	GetFlowId() *string
}

type GetMediaConnectFlowRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34900dc6-90ec-4968-af3c-fcd87f231a5f
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetMediaConnectFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowRequest) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaConnectFlowRequest) SetFlowId(v string) *GetMediaConnectFlowRequest {
	s.FlowId = &v
	return s
}

func (s *GetMediaConnectFlowRequest) Validate() error {
	return dara.Validate(s)
}
