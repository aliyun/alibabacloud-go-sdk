// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteJobHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteJobHeaders
	GetWorkspace() *string
}

type DeleteJobHeaders struct {
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

func (s DeleteJobHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobHeaders) GoString() string {
	return s.String()
}

func (s *DeleteJobHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteJobHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteJobHeaders) SetCommonHeaders(v map[string]*string) *DeleteJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteJobHeaders) SetWorkspace(v string) *DeleteJobHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteJobHeaders) Validate() error {
	return dara.Validate(s)
}
