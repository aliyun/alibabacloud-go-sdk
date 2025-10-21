// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSavepointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSavepointHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteSavepointHeaders
	GetWorkspace() *string
}

type DeleteSavepointHeaders struct {
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

func (s DeleteSavepointHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSavepointHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSavepointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSavepointHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteSavepointHeaders) SetCommonHeaders(v map[string]*string) *DeleteSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSavepointHeaders) SetWorkspace(v string) *DeleteSavepointHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteSavepointHeaders) Validate() error {
	return dara.Validate(s)
}
