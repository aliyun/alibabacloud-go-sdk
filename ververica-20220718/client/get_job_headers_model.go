// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetJobHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetJobHeaders
	GetWorkspace() *string
}

type GetJobHeaders struct {
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

func (s GetJobHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetJobHeaders) GoString() string {
	return s.String()
}

func (s *GetJobHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetJobHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetJobHeaders) SetCommonHeaders(v map[string]*string) *GetJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetJobHeaders) SetWorkspace(v string) *GetJobHeaders {
	s.Workspace = &v
	return s
}

func (s *GetJobHeaders) Validate() error {
	return dara.Validate(s)
}
