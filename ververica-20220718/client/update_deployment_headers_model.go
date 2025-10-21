// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDeploymentHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateDeploymentHeaders
	GetWorkspace() *string
}

type UpdateDeploymentHeaders struct {
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

func (s UpdateDeploymentHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDeploymentHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDeploymentHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentHeaders) SetWorkspace(v string) *UpdateDeploymentHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateDeploymentHeaders) Validate() error {
	return dara.Validate(s)
}
