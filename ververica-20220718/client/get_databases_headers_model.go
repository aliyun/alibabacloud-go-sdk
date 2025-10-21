// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabasesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDatabasesHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDatabasesHeaders
	GetWorkspace() *string
}

type GetDatabasesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDatabasesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDatabasesHeaders) GoString() string {
	return s.String()
}

func (s *GetDatabasesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDatabasesHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDatabasesHeaders) SetCommonHeaders(v map[string]*string) *GetDatabasesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDatabasesHeaders) SetWorkspace(v string) *GetDatabasesHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDatabasesHeaders) Validate() error {
	return dara.Validate(s)
}
