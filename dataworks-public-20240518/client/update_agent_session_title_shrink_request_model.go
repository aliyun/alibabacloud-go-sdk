// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSessionTitleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UpdateAgentSessionTitleShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *UpdateAgentSessionTitleShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *UpdateAgentSessionTitleShrinkRequest
	GetParamsShrink() *string
}

type UpdateAgentSessionTitleShrinkRequest struct {
	// example:
	//
	// updateagentsessiontitle-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s UpdateAgentSessionTitleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateAgentSessionTitleShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UpdateAgentSessionTitleShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *UpdateAgentSessionTitleShrinkRequest) SetId(v string) *UpdateAgentSessionTitleShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateAgentSessionTitleShrinkRequest) SetJsonrpc(v string) *UpdateAgentSessionTitleShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *UpdateAgentSessionTitleShrinkRequest) SetParamsShrink(v string) *UpdateAgentSessionTitleShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *UpdateAgentSessionTitleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
