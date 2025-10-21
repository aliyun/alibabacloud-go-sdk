// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavepointsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSavepointsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListSavepointsHeaders
	GetWorkspace() *string
}

type ListSavepointsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListSavepointsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSavepointsHeaders) GoString() string {
	return s.String()
}

func (s *ListSavepointsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSavepointsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListSavepointsHeaders) SetCommonHeaders(v map[string]*string) *ListSavepointsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSavepointsHeaders) SetWorkspace(v string) *ListSavepointsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListSavepointsHeaders) Validate() error {
	return dara.Validate(s)
}
