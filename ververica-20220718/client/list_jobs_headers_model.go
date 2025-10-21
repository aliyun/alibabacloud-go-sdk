// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListJobsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListJobsHeaders
	GetWorkspace() *string
}

type ListJobsHeaders struct {
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

func (s ListJobsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListJobsHeaders) GoString() string {
	return s.String()
}

func (s *ListJobsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListJobsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListJobsHeaders) SetCommonHeaders(v map[string]*string) *ListJobsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListJobsHeaders) SetWorkspace(v string) *ListJobsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListJobsHeaders) Validate() error {
	return dara.Validate(s)
}
