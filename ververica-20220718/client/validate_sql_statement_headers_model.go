// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSqlStatementHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ValidateSqlStatementHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ValidateSqlStatementHeaders
	GetWorkspace() *string
}

type ValidateSqlStatementHeaders struct {
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

func (s ValidateSqlStatementHeaders) String() string {
	return dara.Prettify(s)
}

func (s ValidateSqlStatementHeaders) GoString() string {
	return s.String()
}

func (s *ValidateSqlStatementHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ValidateSqlStatementHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ValidateSqlStatementHeaders) SetCommonHeaders(v map[string]*string) *ValidateSqlStatementHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ValidateSqlStatementHeaders) SetWorkspace(v string) *ValidateSqlStatementHeaders {
	s.Workspace = &v
	return s
}

func (s *ValidateSqlStatementHeaders) Validate() error {
	return dara.Validate(s)
}
