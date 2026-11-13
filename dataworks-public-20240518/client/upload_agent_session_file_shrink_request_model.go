// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSessionFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *UploadAgentSessionFileShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *UploadAgentSessionFileShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *UploadAgentSessionFileShrinkRequest
	GetParamsShrink() *string
}

type UploadAgentSessionFileShrinkRequest struct {
	// example:
	//
	// attachment-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s UploadAgentSessionFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UploadAgentSessionFileShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UploadAgentSessionFileShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *UploadAgentSessionFileShrinkRequest) SetId(v string) *UploadAgentSessionFileShrinkRequest {
	s.Id = &v
	return s
}

func (s *UploadAgentSessionFileShrinkRequest) SetJsonrpc(v string) *UploadAgentSessionFileShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *UploadAgentSessionFileShrinkRequest) SetParamsShrink(v string) *UploadAgentSessionFileShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *UploadAgentSessionFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
