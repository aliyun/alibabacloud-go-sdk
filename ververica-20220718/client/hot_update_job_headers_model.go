// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotUpdateJobHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotUpdateJobHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *HotUpdateJobHeaders
	GetWorkspace() *string
}

type HotUpdateJobHeaders struct {
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

func (s HotUpdateJobHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotUpdateJobHeaders) GoString() string {
	return s.String()
}

func (s *HotUpdateJobHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotUpdateJobHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *HotUpdateJobHeaders) SetCommonHeaders(v map[string]*string) *HotUpdateJobHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotUpdateJobHeaders) SetWorkspace(v string) *HotUpdateJobHeaders {
	s.Workspace = &v
	return s
}

func (s *HotUpdateJobHeaders) Validate() error {
	return dara.Validate(s)
}
