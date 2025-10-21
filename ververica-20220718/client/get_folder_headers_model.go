// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFolderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFolderHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetFolderHeaders
	GetWorkspace() *string
}

type GetFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetFolderHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFolderHeaders) GoString() string {
	return s.String()
}

func (s *GetFolderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFolderHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetFolderHeaders) SetCommonHeaders(v map[string]*string) *GetFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFolderHeaders) SetWorkspace(v string) *GetFolderHeaders {
	s.Workspace = &v
	return s
}

func (s *GetFolderHeaders) Validate() error {
	return dara.Validate(s)
}
