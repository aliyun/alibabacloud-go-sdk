// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotUpdateJobResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotUpdateJobResultHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetHotUpdateJobResultHeaders
	GetWorkspace() *string
}

type GetHotUpdateJobResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetHotUpdateJobResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotUpdateJobResultHeaders) GoString() string {
	return s.String()
}

func (s *GetHotUpdateJobResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotUpdateJobResultHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetHotUpdateJobResultHeaders) SetCommonHeaders(v map[string]*string) *GetHotUpdateJobResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotUpdateJobResultHeaders) SetWorkspace(v string) *GetHotUpdateJobResultHeaders {
	s.Workspace = &v
	return s
}

func (s *GetHotUpdateJobResultHeaders) Validate() error {
	return dara.Validate(s)
}
