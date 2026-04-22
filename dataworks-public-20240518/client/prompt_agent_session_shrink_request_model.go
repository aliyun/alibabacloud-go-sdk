// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PromptAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *PromptAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type PromptAgentSessionShrinkRequest struct {
	// example:
	//
	// 1021418411
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s PromptAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *PromptAgentSessionShrinkRequest) SetId(v string) *PromptAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) SetJsonrpc(v string) *PromptAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) SetParamsShrink(v string) *PromptAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
