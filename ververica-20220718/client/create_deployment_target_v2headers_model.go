// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetV2Headers interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDeploymentTargetV2Headers
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *CreateDeploymentTargetV2Headers
	GetWorkspace() *string
}

type CreateDeploymentTargetV2Headers struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bd5d90a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateDeploymentTargetV2Headers) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetV2Headers) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetV2Headers) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDeploymentTargetV2Headers) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateDeploymentTargetV2Headers) SetCommonHeaders(v map[string]*string) *CreateDeploymentTargetV2Headers {
	s.CommonHeaders = v
	return s
}

func (s *CreateDeploymentTargetV2Headers) SetWorkspace(v string) *CreateDeploymentTargetV2Headers {
	s.Workspace = &v
	return s
}

func (s *CreateDeploymentTargetV2Headers) Validate() error {
	return dara.Validate(s)
}
