// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteMcpServerRequest
	GetName() *string
}

type DeleteMcpServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServerRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *DeleteMcpServerRequest) SetName(v string) *DeleteMcpServerRequest {
	s.Name = &v
	return s
}

func (s *DeleteMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
