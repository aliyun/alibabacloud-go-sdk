// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByIpHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentsByIpHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentsByIpHeaders
	GetWorkspace() *string
}

type GetDeploymentsByIpHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentsByIpHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByIpHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByIpHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentsByIpHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentsByIpHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentsByIpHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentsByIpHeaders) SetWorkspace(v string) *GetDeploymentsByIpHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentsByIpHeaders) Validate() error {
	return dara.Validate(s)
}
