// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTablesHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetTablesHeaders
	GetWorkspace() *string
}

type GetTablesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetTablesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTablesHeaders) GoString() string {
	return s.String()
}

func (s *GetTablesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTablesHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetTablesHeaders) SetCommonHeaders(v map[string]*string) *GetTablesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTablesHeaders) SetWorkspace(v string) *GetTablesHeaders {
	s.Workspace = &v
	return s
}

func (s *GetTablesHeaders) Validate() error {
	return dara.Validate(s)
}
