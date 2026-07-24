// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlExecutionResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FetchSqlExecutionResultHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *FetchSqlExecutionResultHeaders
	GetWorkspace() *string
}

type FetchSqlExecutionResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s FetchSqlExecutionResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlExecutionResultHeaders) GoString() string {
	return s.String()
}

func (s *FetchSqlExecutionResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FetchSqlExecutionResultHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *FetchSqlExecutionResultHeaders) SetCommonHeaders(v map[string]*string) *FetchSqlExecutionResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FetchSqlExecutionResultHeaders) SetWorkspace(v string) *FetchSqlExecutionResultHeaders {
	s.Workspace = &v
	return s
}

func (s *FetchSqlExecutionResultHeaders) Validate() error {
	return dara.Validate(s)
}
