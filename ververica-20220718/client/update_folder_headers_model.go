// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFolderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateFolderHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateFolderHeaders
	GetWorkspace() *string
}

type UpdateFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f89a0c1ca8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateFolderHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateFolderHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFolderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateFolderHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateFolderHeaders) SetCommonHeaders(v map[string]*string) *UpdateFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFolderHeaders) SetWorkspace(v string) *UpdateFolderHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateFolderHeaders) Validate() error {
	return dara.Validate(s)
}
