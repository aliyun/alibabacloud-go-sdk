// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateDeploymentDraftResultHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetValidateDeploymentDraftResultHeaders
	GetCommonHeaders() map[string]*string
	SetWorkspace(v string) *GetValidateDeploymentDraftResultHeaders
	GetWorkspace() *string
}

type GetValidateDeploymentDraftResultHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetValidateDeploymentDraftResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetValidateDeploymentDraftResultHeaders) GoString() string {
	return s.String()
}

func (s *GetValidateDeploymentDraftResultHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetValidateDeploymentDraftResultHeaders) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetValidateDeploymentDraftResultHeaders) SetCommonHeaders(v map[string]*string) *GetValidateDeploymentDraftResultHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetValidateDeploymentDraftResultHeaders) SetWorkspace(v string) *GetValidateDeploymentDraftResultHeaders {
	s.Workspace = &v
	return s
}

func (s *GetValidateDeploymentDraftResultHeaders) Validate() error {
	return dara.Validate(s)
}
