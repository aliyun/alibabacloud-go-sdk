// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateFolderHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateFolderHeaders
	GetWorkspace() *string
}

type CreateFolderHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateFolderHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderHeaders) GoString() string {
	return s.String()
}

func (s *CreateFolderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateFolderHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateFolderHeaders) SetCommonHeaders(v map[string]*string) *CreateFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateFolderHeaders) SetWorkspace(v string) *CreateFolderHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateFolderHeaders) Validate() error {
	return dara.Validate(s)
}
