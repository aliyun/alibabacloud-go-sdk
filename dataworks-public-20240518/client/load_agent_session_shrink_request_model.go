// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *LoadAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *LoadAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type LoadAgentSessionShrinkRequest struct {
	// example:
	//
	// 4as3dasf654a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s LoadAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *LoadAgentSessionShrinkRequest) SetId(v string) *LoadAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) SetJsonrpc(v string) *LoadAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) SetParamsShrink(v string) *LoadAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
