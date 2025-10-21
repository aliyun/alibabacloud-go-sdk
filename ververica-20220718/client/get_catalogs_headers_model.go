// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCatalogsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetCatalogsHeaders
	GetWorkspace() *string
}

type GetCatalogsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetCatalogsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogsHeaders) GoString() string {
	return s.String()
}

func (s *GetCatalogsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCatalogsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetCatalogsHeaders) SetCommonHeaders(v map[string]*string) *GetCatalogsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCatalogsHeaders) SetWorkspace(v string) *GetCatalogsHeaders {
	s.Workspace = &v
	return s
}

func (s *GetCatalogsHeaders) Validate() error {
	return dara.Validate(s)
}
