// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizKey(v string) *GetModelDeploymentSpecRequest
	GetBizKey() *string
	SetModelVersion(v string) *GetModelDeploymentSpecRequest
	GetModelVersion() *string
	SetProfileId(v string) *GetModelDeploymentSpecRequest
	GetProfileId() *string
	SetResourceSelections(v string) *GetModelDeploymentSpecRequest
	GetResourceSelections() *string
	SetWorkspaceId(v string) *GetModelDeploymentSpecRequest
	GetWorkspaceId() *string
}

type GetModelDeploymentSpecRequest struct {
	// example:
	//
	// cmu-biz
	BizKey *string `json:"BizKey,omitempty" xml:"BizKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prf_5cd37a1c7eaa***c6829dbb02
	ProfileId *string `json:"ProfileId,omitempty" xml:"ProfileId,omitempty"`
	// example:
	//
	// %7B%22MemberType%22%3A%22Default%22%7D
	ResourceSelections *string `json:"ResourceSelections,omitempty" xml:"ResourceSelections,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 295949
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetModelDeploymentSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentSpecRequest) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentSpecRequest) GetBizKey() *string {
	return s.BizKey
}

func (s *GetModelDeploymentSpecRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *GetModelDeploymentSpecRequest) GetProfileId() *string {
	return s.ProfileId
}

func (s *GetModelDeploymentSpecRequest) GetResourceSelections() *string {
	return s.ResourceSelections
}

func (s *GetModelDeploymentSpecRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetModelDeploymentSpecRequest) SetBizKey(v string) *GetModelDeploymentSpecRequest {
	s.BizKey = &v
	return s
}

func (s *GetModelDeploymentSpecRequest) SetModelVersion(v string) *GetModelDeploymentSpecRequest {
	s.ModelVersion = &v
	return s
}

func (s *GetModelDeploymentSpecRequest) SetProfileId(v string) *GetModelDeploymentSpecRequest {
	s.ProfileId = &v
	return s
}

func (s *GetModelDeploymentSpecRequest) SetResourceSelections(v string) *GetModelDeploymentSpecRequest {
	s.ResourceSelections = &v
	return s
}

func (s *GetModelDeploymentSpecRequest) SetWorkspaceId(v string) *GetModelDeploymentSpecRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetModelDeploymentSpecRequest) Validate() error {
	return dara.Validate(s)
}
