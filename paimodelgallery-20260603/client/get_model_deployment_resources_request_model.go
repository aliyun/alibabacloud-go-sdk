// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizKey(v string) *GetModelDeploymentResourcesRequest
	GetBizKey() *string
	SetModelVersion(v string) *GetModelDeploymentResourcesRequest
	GetModelVersion() *string
	SetProfileId(v string) *GetModelDeploymentResourcesRequest
	GetProfileId() *string
	SetWorkspaceId(v string) *GetModelDeploymentResourcesRequest
	GetWorkspaceId() *string
}

type GetModelDeploymentResourcesRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 295949
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetModelDeploymentResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesRequest) GetBizKey() *string {
	return s.BizKey
}

func (s *GetModelDeploymentResourcesRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *GetModelDeploymentResourcesRequest) GetProfileId() *string {
	return s.ProfileId
}

func (s *GetModelDeploymentResourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetModelDeploymentResourcesRequest) SetBizKey(v string) *GetModelDeploymentResourcesRequest {
	s.BizKey = &v
	return s
}

func (s *GetModelDeploymentResourcesRequest) SetModelVersion(v string) *GetModelDeploymentResourcesRequest {
	s.ModelVersion = &v
	return s
}

func (s *GetModelDeploymentResourcesRequest) SetProfileId(v string) *GetModelDeploymentResourcesRequest {
	s.ProfileId = &v
	return s
}

func (s *GetModelDeploymentResourcesRequest) SetWorkspaceId(v string) *GetModelDeploymentResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetModelDeploymentResourcesRequest) Validate() error {
	return dara.Validate(s)
}
