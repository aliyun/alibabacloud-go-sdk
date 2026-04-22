// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CancelAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *CancelAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *CancelAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type CancelAgentSessionShrinkRequest struct {
	// example:
	//
	// 676303114031776
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s CancelAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *CancelAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CancelAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *CancelAgentSessionShrinkRequest) SetId(v string) *CancelAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *CancelAgentSessionShrinkRequest) SetJsonrpc(v string) *CancelAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *CancelAgentSessionShrinkRequest) SetParamsShrink(v string) *CancelAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *CancelAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
