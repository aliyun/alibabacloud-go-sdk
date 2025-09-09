// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *DeleteAcrImageRepositoriesRequest
	GetArtifactType() *string
	SetClientToken(v string) *DeleteAcrImageRepositoriesRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteAcrImageRepositoriesRequest
	GetRegionId() *string
	SetRepoId(v string) *DeleteAcrImageRepositoriesRequest
	GetRepoId() *string
}

type DeleteAcrImageRepositoriesRequest struct {
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crr-7x9rf32mkqoqulrn
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s DeleteAcrImageRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageRepositoriesRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *DeleteAcrImageRepositoriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAcrImageRepositoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAcrImageRepositoriesRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteAcrImageRepositoriesRequest) SetArtifactType(v string) *DeleteAcrImageRepositoriesRequest {
	s.ArtifactType = &v
	return s
}

func (s *DeleteAcrImageRepositoriesRequest) SetClientToken(v string) *DeleteAcrImageRepositoriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAcrImageRepositoriesRequest) SetRegionId(v string) *DeleteAcrImageRepositoriesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAcrImageRepositoriesRequest) SetRepoId(v string) *DeleteAcrImageRepositoriesRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteAcrImageRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
