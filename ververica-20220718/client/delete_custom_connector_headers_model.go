// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomConnectorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteCustomConnectorHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteCustomConnectorHeaders
	GetWorkspace() *string
}

type DeleteCustomConnectorHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteCustomConnectorHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomConnectorHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomConnectorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteCustomConnectorHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteCustomConnectorHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomConnectorHeaders) SetWorkspace(v string) *DeleteCustomConnectorHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteCustomConnectorHeaders) Validate() error {
	return dara.Validate(s)
}
