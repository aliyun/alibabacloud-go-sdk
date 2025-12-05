// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByNameHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentsByNameHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentsByNameHeaders
	GetWorkspace() *string
}

type GetDeploymentsByNameHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentsByNameHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByNameHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByNameHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentsByNameHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentsByNameHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentsByNameHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentsByNameHeaders) SetWorkspace(v string) *GetDeploymentsByNameHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentsByNameHeaders) Validate() error {
	return dara.Validate(s)
}
