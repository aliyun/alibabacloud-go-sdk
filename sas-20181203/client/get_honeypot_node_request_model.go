// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetHoneypotNodeRequest
	GetLang() *string
	SetNodeId(v string) *GetHoneypotNodeRequest
	GetNodeId() *string
}

type GetHoneypotNodeRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the management node.
	//
	// > You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to query the IDs of management nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// d3dd3864-4e02-4abd-8b6a-8f5f6fec4715
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s GetHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeRequest) GetLang() *string {
	return s.Lang
}

func (s *GetHoneypotNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHoneypotNodeRequest) SetLang(v string) *GetHoneypotNodeRequest {
	s.Lang = &v
	return s
}

func (s *GetHoneypotNodeRequest) SetNodeId(v string) *GetHoneypotNodeRequest {
	s.NodeId = &v
	return s
}

func (s *GetHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
