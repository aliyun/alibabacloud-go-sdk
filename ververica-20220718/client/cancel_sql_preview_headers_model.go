// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPreviewHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CancelSqlPreviewHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CancelSqlPreviewHeaders
	GetWorkspace() *string
}

type CancelSqlPreviewHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CancelSqlPreviewHeaders) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPreviewHeaders) GoString() string {
	return s.String()
}

func (s *CancelSqlPreviewHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CancelSqlPreviewHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CancelSqlPreviewHeaders) SetCommonHeaders(v map[string]*string) *CancelSqlPreviewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelSqlPreviewHeaders) SetWorkspace(v string) *CancelSqlPreviewHeaders {
	s.Workspace = &v
	return s
}

func (s *CancelSqlPreviewHeaders) Validate() error {
	return dara.Validate(s)
}
