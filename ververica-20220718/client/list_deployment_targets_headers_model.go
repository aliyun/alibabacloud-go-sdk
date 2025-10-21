// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentTargetsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDeploymentTargetsHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *ListDeploymentTargetsHeaders
	GetWorkspace() *string
}

type ListDeploymentTargetsHeaders struct {
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

func (s ListDeploymentTargetsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentTargetsHeaders) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDeploymentTargetsHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListDeploymentTargetsHeaders) SetCommonHeaders(v map[string]*string) *ListDeploymentTargetsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDeploymentTargetsHeaders) SetWorkspace(v string) *ListDeploymentTargetsHeaders {
	s.Workspace = &v
	return s
}

func (s *ListDeploymentTargetsHeaders) Validate() error {
	return dara.Validate(s)
}
