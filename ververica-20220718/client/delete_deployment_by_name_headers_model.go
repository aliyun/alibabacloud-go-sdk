// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentByNameHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDeploymentByNameHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *DeleteDeploymentByNameHeaders
	GetWorkspace() *string
}

type DeleteDeploymentByNameHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8c34d
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteDeploymentByNameHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentByNameHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentByNameHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDeploymentByNameHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteDeploymentByNameHeaders) SetCommonHeaders(v map[string]*string) *DeleteDeploymentByNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDeploymentByNameHeaders) SetWorkspace(v string) *DeleteDeploymentByNameHeaders {
	s.Workspace = &v
	return s
}

func (s *DeleteDeploymentByNameHeaders) Validate() error {
	return dara.Validate(s)
}
