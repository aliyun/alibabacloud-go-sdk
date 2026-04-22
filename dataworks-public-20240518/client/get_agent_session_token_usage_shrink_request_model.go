// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionTokenUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionTokenUsageShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionTokenUsageShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *GetAgentSessionTokenUsageShrinkRequest
	GetParamsShrink() *string
}

type GetAgentSessionTokenUsageShrinkRequest struct {
	// example:
	//
	// 1033814166
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetAgentSessionTokenUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageShrinkRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionTokenUsageShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionTokenUsageShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *GetAgentSessionTokenUsageShrinkRequest) SetId(v string) *GetAgentSessionTokenUsageShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionTokenUsageShrinkRequest) SetJsonrpc(v string) *GetAgentSessionTokenUsageShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionTokenUsageShrinkRequest) SetParamsShrink(v string) *GetAgentSessionTokenUsageShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetAgentSessionTokenUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
