// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionArtifactMetaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionArtifactMetaShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionArtifactMetaShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *GetAgentSessionArtifactMetaShrinkRequest
	GetParamsShrink() *string
}

type GetAgentSessionArtifactMetaShrinkRequest struct {
	// example:
	//
	// 900335678024
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc      *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetAgentSessionArtifactMetaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) SetId(v string) *GetAgentSessionArtifactMetaShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) SetJsonrpc(v string) *GetAgentSessionArtifactMetaShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) SetParamsShrink(v string) *GetAgentSessionArtifactMetaShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetAgentSessionArtifactMetaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
