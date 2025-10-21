// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMemberHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateMemberHeaders
	GetWorkspace() *string
}

type CreateMemberHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca84d539167d4d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMemberHeaders) GoString() string {
	return s.String()
}

func (s *CreateMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMemberHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateMemberHeaders) SetCommonHeaders(v map[string]*string) *CreateMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMemberHeaders) SetWorkspace(v string) *CreateMemberHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateMemberHeaders) Validate() error {
	return dara.Validate(s)
}
