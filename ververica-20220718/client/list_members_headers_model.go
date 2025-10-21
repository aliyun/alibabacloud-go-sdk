// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListMembersHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListMembersHeaders
	GetWorkspace() *string
}

type ListMembersHeaders struct {
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

func (s ListMembersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListMembersHeaders) GoString() string {
	return s.String()
}

func (s *ListMembersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListMembersHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListMembersHeaders) SetCommonHeaders(v map[string]*string) *ListMembersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMembersHeaders) SetWorkspace(v string) *ListMembersHeaders {
	s.Workspace = &v
	return s
}

func (s *ListMembersHeaders) Validate() error {
	return dara.Validate(s)
}
