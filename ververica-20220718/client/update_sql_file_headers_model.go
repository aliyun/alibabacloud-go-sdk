// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSqlFileHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateSqlFileHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateSqlFileHeaders
	GetWorkspace() *string
}

type UpdateSqlFileHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateSqlFileHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateSqlFileHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSqlFileHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateSqlFileHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateSqlFileHeaders) SetCommonHeaders(v map[string]*string) *UpdateSqlFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSqlFileHeaders) SetWorkspace(v string) *UpdateSqlFileHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateSqlFileHeaders) Validate() error {
	return dara.Validate(s)
}
