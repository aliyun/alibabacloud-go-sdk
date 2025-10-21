// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavepointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSavepointHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateSavepointHeaders
	GetWorkspace() *string
}

type CreateSavepointHeaders struct {
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

func (s CreateSavepointHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSavepointHeaders) GoString() string {
	return s.String()
}

func (s *CreateSavepointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSavepointHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateSavepointHeaders) SetCommonHeaders(v map[string]*string) *CreateSavepointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSavepointHeaders) SetWorkspace(v string) *CreateSavepointHeaders {
	s.Workspace = &v
	return s
}

func (s *CreateSavepointHeaders) Validate() error {
	return dara.Validate(s)
}
