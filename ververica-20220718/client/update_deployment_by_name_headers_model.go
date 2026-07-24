// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentByNameHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateDeploymentByNameHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *UpdateDeploymentByNameHeaders
	GetWorkspace() *string
}

type UpdateDeploymentByNameHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdateDeploymentByNameHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentByNameHeaders) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentByNameHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateDeploymentByNameHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdateDeploymentByNameHeaders) SetCommonHeaders(v map[string]*string) *UpdateDeploymentByNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateDeploymentByNameHeaders) SetWorkspace(v string) *UpdateDeploymentByNameHeaders {
	s.Workspace = &v
	return s
}

func (s *UpdateDeploymentByNameHeaders) Validate() error {
	return dara.Validate(s)
}
