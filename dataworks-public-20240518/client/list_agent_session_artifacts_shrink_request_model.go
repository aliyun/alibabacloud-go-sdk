// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionArtifactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentSessionArtifactsShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentSessionArtifactsShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *ListAgentSessionArtifactsShrinkRequest
	GetParamsShrink() *string
}

type ListAgentSessionArtifactsShrinkRequest struct {
	// example:
	//
	// 10001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ListAgentSessionArtifactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionArtifactsShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionArtifactsShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *ListAgentSessionArtifactsShrinkRequest) SetId(v string) *ListAgentSessionArtifactsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListAgentSessionArtifactsShrinkRequest) SetJsonrpc(v string) *ListAgentSessionArtifactsShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionArtifactsShrinkRequest) SetParamsShrink(v string) *ListAgentSessionArtifactsShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *ListAgentSessionArtifactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
