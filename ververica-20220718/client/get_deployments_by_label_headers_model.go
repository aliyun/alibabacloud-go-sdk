// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByLabelHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentsByLabelHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentsByLabelHeaders
	GetWorkspace() *string
}

type GetDeploymentsByLabelHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeploymentsByLabelHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByLabelHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByLabelHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentsByLabelHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentsByLabelHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentsByLabelHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentsByLabelHeaders) SetWorkspace(v string) *GetDeploymentsByLabelHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentsByLabelHeaders) Validate() error {
	return dara.Validate(s)
}
