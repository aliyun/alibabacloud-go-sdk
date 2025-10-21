// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomConnectorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RegisterCustomConnectorHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *RegisterCustomConnectorHeaders
	GetWorkspace() *string
}

type RegisterCustomConnectorHeaders struct {
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

func (s RegisterCustomConnectorHeaders) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomConnectorHeaders) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RegisterCustomConnectorHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *RegisterCustomConnectorHeaders) SetCommonHeaders(v map[string]*string) *RegisterCustomConnectorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RegisterCustomConnectorHeaders) SetWorkspace(v string) *RegisterCustomConnectorHeaders {
	s.Workspace = &v
	return s
}

func (s *RegisterCustomConnectorHeaders) Validate() error {
	return dara.Validate(s)
}
