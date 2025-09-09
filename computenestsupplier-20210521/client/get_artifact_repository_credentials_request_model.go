// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactRepositoryCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *GetArtifactRepositoryCredentialsRequest
	GetArtifactType() *string
	SetDeployRegionId(v string) *GetArtifactRepositoryCredentialsRequest
	GetDeployRegionId() *string
}

type GetArtifactRepositoryCredentialsRequest struct {
	// The type of the deployment package. Valid values:
	//
	// 	- File: Object Storage Service (OSS) object.
	//
	// 	- AcrImage: container image.
	//
	// This parameter is required.
	//
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactRepositoryCredentialsRequest) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *GetArtifactRepositoryCredentialsRequest) SetArtifactType(v string) *GetArtifactRepositoryCredentialsRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsRequest) SetDeployRegionId(v string) *GetArtifactRepositoryCredentialsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
