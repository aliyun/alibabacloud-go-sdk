// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowAllOutputNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowId(v string) *GetMediaConnectFlowAllOutputNameRequest
	GetFlowId() *string
}

type GetMediaConnectFlowAllOutputNameRequest struct {
	// The ID of the MediaConnect flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0381f478-7d53-4076-9d5f-27680a6f73e7
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetMediaConnectFlowAllOutputNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowAllOutputNameRequest) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowAllOutputNameRequest) GetFlowId() *string {
	return s.FlowId
}

func (s *GetMediaConnectFlowAllOutputNameRequest) SetFlowId(v string) *GetMediaConnectFlowAllOutputNameRequest {
	s.FlowId = &v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameRequest) Validate() error {
	return dara.Validate(s)
}
