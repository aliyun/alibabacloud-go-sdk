// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMemberHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteMemberHeaders
	GetWorkspace() *string
}

type DeleteMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemberHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMemberHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteMemberHeaders) SetCommonHeaders(v map[string]*string) *DeleteMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMemberHeaders) SetWorkspace(v string) *DeleteMemberHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteMemberHeaders) Validate() error {
	return dara.Validate(s)
}
