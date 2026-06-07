// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetMcpServerRequest
	GetName() *string
}

type GetMcpServerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerRequest) GoString() string {
	return s.String()
}

func (s *GetMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *GetMcpServerRequest) SetName(v string) *GetMcpServerRequest {
	s.Name = &v
	return s
}

func (s *GetMcpServerRequest) Validate() error {
	return dara.Validate(s)
}
