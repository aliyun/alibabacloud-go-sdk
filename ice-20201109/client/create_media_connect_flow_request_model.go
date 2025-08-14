// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConnectFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *CreateMediaConnectFlowRequest
	GetFlowName() *string
	SetFlowRegion(v string) *CreateMediaConnectFlowRequest
	GetFlowRegion() *string
}

type CreateMediaConnectFlowRequest struct {
	// The flow name.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliTestFlow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The region in which the flow resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	FlowRegion *string `json:"FlowRegion,omitempty" xml:"FlowRegion,omitempty"`
}

func (s CreateMediaConnectFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConnectFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaConnectFlowRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *CreateMediaConnectFlowRequest) GetFlowRegion() *string {
	return s.FlowRegion
}

func (s *CreateMediaConnectFlowRequest) SetFlowName(v string) *CreateMediaConnectFlowRequest {
	s.FlowName = &v
	return s
}

func (s *CreateMediaConnectFlowRequest) SetFlowRegion(v string) *CreateMediaConnectFlowRequest {
	s.FlowRegion = &v
	return s
}

func (s *CreateMediaConnectFlowRequest) Validate() error {
	return dara.Validate(s)
}
