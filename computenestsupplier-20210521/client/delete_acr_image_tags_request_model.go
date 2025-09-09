// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *DeleteAcrImageTagsRequest
	GetArtifactType() *string
	SetClientToken(v string) *DeleteAcrImageTagsRequest
	GetClientToken() *string
	SetRegionId(v string) *DeleteAcrImageTagsRequest
	GetRegionId() *string
	SetRepoId(v string) *DeleteAcrImageTagsRequest
	GetRepoId() *string
	SetTag(v string) *DeleteAcrImageTagsRequest
	GetTag() *string
}

type DeleteAcrImageTagsRequest struct {
	// example:
	//
	// AcrImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// 788E7CP0EN9D51P
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
	// crr-3gqhkza0wbxxxxxx
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DeleteAcrImageTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageTagsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageTagsRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *DeleteAcrImageTagsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAcrImageTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAcrImageTagsRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteAcrImageTagsRequest) GetTag() *string {
	return s.Tag
}

func (s *DeleteAcrImageTagsRequest) SetArtifactType(v string) *DeleteAcrImageTagsRequest {
	s.ArtifactType = &v
	return s
}

func (s *DeleteAcrImageTagsRequest) SetClientToken(v string) *DeleteAcrImageTagsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAcrImageTagsRequest) SetRegionId(v string) *DeleteAcrImageTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAcrImageTagsRequest) SetRepoId(v string) *DeleteAcrImageTagsRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteAcrImageTagsRequest) SetTag(v string) *DeleteAcrImageTagsRequest {
	s.Tag = &v
	return s
}

func (s *DeleteAcrImageTagsRequest) Validate() error {
	return dara.Validate(s)
}
