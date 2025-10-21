// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentDraftsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeploymentDraftsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListDeploymentDraftsHeaders
	GetWorkspace() *string
}

type ListDeploymentDraftsHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListDeploymentDraftsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentDraftsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentDraftsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeploymentDraftsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListDeploymentDraftsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentDraftsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentDraftsHeaders) SetWorkspace(v string) *ListDeploymentDraftsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListDeploymentDraftsHeaders) Validate() error {
	return dara.Validate(s)
}
