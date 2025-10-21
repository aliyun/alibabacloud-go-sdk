// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFolderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteFolderHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteFolderHeaders
	GetWorkspace() *string
}

type DeleteFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c96306e2b****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteFolderHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteFolderHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFolderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteFolderHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteFolderHeaders) SetCommonHeaders(v map[string]*string) *DeleteFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFolderHeaders) SetWorkspace(v string) *DeleteFolderHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteFolderHeaders) Validate() error {
	return dara.Validate(s)
}
