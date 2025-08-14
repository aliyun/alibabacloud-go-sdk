// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *DeleteMediaConnectFlowRequest
	GetFlowId() *string
}

type DeleteMediaConnectFlowRequest struct {
	// The flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0381f478-7d53-4076-9d5f-27680a6f73e7
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DeleteMediaConnectFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *DeleteMediaConnectFlowRequest) SetFlowId(v string) *DeleteMediaConnectFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DeleteMediaConnectFlowRequest) Validate() error {
	return dara.Validate(s)
}
