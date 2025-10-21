// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMemberHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateMemberHeaders
	GetWorkspace() *string
}

type UpdateMemberHeaders struct {
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

func (s UpdateMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemberHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMemberHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateMemberHeaders) SetCommonHeaders(v map[string]*string) *UpdateMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMemberHeaders) SetWorkspace(v string) *UpdateMemberHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateMemberHeaders) Validate() error {
	return dara.Validate(s)
}
