// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHoneypotNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteHoneypotNodeRequest
	GetLang() *string
	SetNodeId(v string) *DeleteHoneypotNodeRequest
	GetNodeId() *string
}

type DeleteHoneypotNodeRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
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
	// 670baeee-86c4-46b9-8200-a2c38141a453
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DeleteHoneypotNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHoneypotNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoneypotNodeRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteHoneypotNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteHoneypotNodeRequest) SetLang(v string) *DeleteHoneypotNodeRequest {
	s.Lang = &v
	return s
}

func (s *DeleteHoneypotNodeRequest) SetNodeId(v string) *DeleteHoneypotNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteHoneypotNodeRequest) Validate() error {
	return dara.Validate(s)
}
