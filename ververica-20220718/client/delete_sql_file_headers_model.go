// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlFileHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSqlFileHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteSqlFileHeaders
	GetWorkspace() *string
}

type DeleteSqlFileHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteSqlFileHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlFileHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSqlFileHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSqlFileHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteSqlFileHeaders) SetCommonHeaders(v map[string]*string) *DeleteSqlFileHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSqlFileHeaders) SetWorkspace(v string) *DeleteSqlFileHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteSqlFileHeaders) Validate() error {
	return dara.Validate(s)
}
