// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchSqlPreviewResultsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FetchSqlPreviewResultsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *FetchSqlPreviewResultsHeaders
	GetWorkspace() *string
}

type FetchSqlPreviewResultsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s FetchSqlPreviewResultsHeaders) String() string {
	return dara.Prettify(s)
}

func (s FetchSqlPreviewResultsHeaders) GoString() string {
	return s.String()
}

func (s *FetchSqlPreviewResultsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FetchSqlPreviewResultsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *FetchSqlPreviewResultsHeaders) SetCommonHeaders(v map[string]*string) *FetchSqlPreviewResultsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FetchSqlPreviewResultsHeaders) SetWorkspace(v string) *FetchSqlPreviewResultsHeaders {
	s.Workspace = &v
	return s
}

func (s *FetchSqlPreviewResultsHeaders) Validate() error {
	return dara.Validate(s)
}
