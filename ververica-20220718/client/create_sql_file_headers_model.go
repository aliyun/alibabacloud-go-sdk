// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlFileHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSqlFileHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateSqlFileHeaders
	GetWorkspace() *string
}

type CreateSqlFileHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateSqlFileHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlFileHeaders) GoString() string {
	return s.String()
}

func (s *CreateSqlFileHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSqlFileHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateSqlFileHeaders) SetCommonHeaders(v map[string]*string) *CreateSqlFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSqlFileHeaders) SetWorkspace(v string) *CreateSqlFileHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateSqlFileHeaders) Validate() error {
	return dara.Validate(s)
}
