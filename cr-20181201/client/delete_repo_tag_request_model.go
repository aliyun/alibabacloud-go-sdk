// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteRepoTagRequest
	GetInstanceId() *string
	SetRepoId(v string) *DeleteRepoTagRequest
	GetRepoId() *string
	SetTag(v string) *DeleteRepoTagRequest
	GetTag() *string
}

type DeleteRepoTagRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The tag of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DeleteRepoTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteRepoTagRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DeleteRepoTagRequest) GetTag() *string {
	return s.Tag
}

func (s *DeleteRepoTagRequest) SetInstanceId(v string) *DeleteRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoTagRequest) SetRepoId(v string) *DeleteRepoTagRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepoTagRequest) SetTag(v string) *DeleteRepoTagRequest {
	s.Tag = &v
	return s
}

func (s *DeleteRepoTagRequest) Validate() error {
	return dara.Validate(s)
}
