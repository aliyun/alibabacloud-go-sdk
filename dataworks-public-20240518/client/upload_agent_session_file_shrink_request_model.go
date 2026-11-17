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
	// The JSON-RPC correlation identifier for the current request. The Id in the response returns this value as is to correlate the request and response.
	//
	// example:
	//
	// attachment-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. The example uses 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The parameter object for the upload stage. The prepare step issues the upload URL, and the complete step confirms that the file has been written to the target session instance.
	//
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
