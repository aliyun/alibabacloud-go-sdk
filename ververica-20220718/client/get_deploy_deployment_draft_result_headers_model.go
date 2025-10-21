// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeployDeploymentDraftResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDeployDeploymentDraftResultHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetDeployDeploymentDraftResultHeaders
	GetWorkspace() *string
}

type GetDeployDeploymentDraftResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 710d6a64d8****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetDeployDeploymentDraftResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDeployDeploymentDraftResultHeaders) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDeployDeploymentDraftResultHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetDeployDeploymentDraftResultHeaders) SetCommonHeaders(v map[string]*string) *GetDeployDeploymentDraftResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDeployDeploymentDraftResultHeaders) SetWorkspace(v string) *GetDeployDeploymentDraftResultHeaders {
	s.Workspace = &v
	return s
}

func (s *GetDeployDeploymentDraftResultHeaders) Validate() error {
	return dara.Validate(s)
}
