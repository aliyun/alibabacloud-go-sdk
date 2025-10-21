// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMemberHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetMemberHeaders
	GetWorkspace() *string
}

type GetMemberHeaders struct {
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

func (s GetMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMemberHeaders) GoString() string {
	return s.String()
}

func (s *GetMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMemberHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetMemberHeaders) SetCommonHeaders(v map[string]*string) *GetMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMemberHeaders) SetWorkspace(v string) *GetMemberHeaders {
	s.Workspace = &v
	return s
}

func (s *GetMemberHeaders) Validate() error {
	return dara.Validate(s)
}
