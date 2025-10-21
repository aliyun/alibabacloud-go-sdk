// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavepointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSavepointHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetSavepointHeaders
	GetWorkspace() *string
}

type GetSavepointHeaders struct {
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

func (s GetSavepointHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSavepointHeaders) GoString() string {
	return s.String()
}

func (s *GetSavepointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSavepointHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetSavepointHeaders) SetCommonHeaders(v map[string]*string) *GetSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSavepointHeaders) SetWorkspace(v string) *GetSavepointHeaders {
	s.Workspace = &v
	return s
}

func (s *GetSavepointHeaders) Validate() error {
	return dara.Validate(s)
}
