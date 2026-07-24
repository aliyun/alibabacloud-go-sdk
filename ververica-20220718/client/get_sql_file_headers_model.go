// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlFileHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSqlFileHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetSqlFileHeaders
	GetWorkspace() *string
}

type GetSqlFileHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetSqlFileHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSqlFileHeaders) GoString() string {
	return s.String()
}

func (s *GetSqlFileHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSqlFileHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetSqlFileHeaders) SetCommonHeaders(v map[string]*string) *GetSqlFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSqlFileHeaders) SetWorkspace(v string) *GetSqlFileHeaders {
	s.Workspace = &v
	return s
}

func (s *GetSqlFileHeaders) Validate() error {
	return dara.Validate(s)
}
