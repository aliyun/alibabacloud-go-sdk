// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentsShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentsShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *ListAgentsShrinkRequest
	GetParamsShrink() *string
}

type ListAgentsShrinkRequest struct {
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

func (s ListAgentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentsShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentsShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *ListAgentsShrinkRequest) SetId(v string) *ListAgentsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListAgentsShrinkRequest) SetJsonrpc(v string) *ListAgentsShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentsShrinkRequest) SetParamsShrink(v string) *ListAgentsShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *ListAgentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
