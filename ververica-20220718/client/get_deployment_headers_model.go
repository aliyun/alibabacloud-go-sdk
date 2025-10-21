// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeploymentHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeploymentHeaders
	GetWorkspace() *string
}

type GetDeploymentHeaders struct {
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

func (s GetDeploymentHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentHeaders) GoString() string {
	return s.String()
}

func (s *GetDeploymentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeploymentHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeploymentHeaders) SetCommonHeaders(v map[string]*string) *GetDeploymentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeploymentHeaders) SetWorkspace(v string) *GetDeploymentHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeploymentHeaders) Validate() error {
	return dara.Validate(s)
}
