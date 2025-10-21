// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSqlPreviewHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubmitSqlPreviewHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *SubmitSqlPreviewHeaders
	GetWorkspace() *string
}

type SubmitSqlPreviewHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s SubmitSqlPreviewHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubmitSqlPreviewHeaders) GoString() string {
	return s.String()
}

func (s *SubmitSqlPreviewHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubmitSqlPreviewHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *SubmitSqlPreviewHeaders) SetCommonHeaders(v map[string]*string) *SubmitSqlPreviewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubmitSqlPreviewHeaders) SetWorkspace(v string) *SubmitSqlPreviewHeaders {
	s.Workspace = &v
	return s
}

func (s *SubmitSqlPreviewHeaders) Validate() error {
	return dara.Validate(s)
}
