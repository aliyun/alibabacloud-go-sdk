// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomConnectorsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCustomConnectorsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListCustomConnectorsHeaders
	GetWorkspace() *string
}

type ListCustomConnectorsHeaders struct {
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

func (s ListCustomConnectorsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCustomConnectorsHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomConnectorsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCustomConnectorsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListCustomConnectorsHeaders) SetCommonHeaders(v map[string]*string) *ListCustomConnectorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomConnectorsHeaders) SetWorkspace(v string) *ListCustomConnectorsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListCustomConnectorsHeaders) Validate() error {
	return dara.Validate(s)
}
