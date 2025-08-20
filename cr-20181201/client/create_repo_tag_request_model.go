// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromTag(v string) *CreateRepoTagRequest
	GetFromTag() *string
	SetInstanceId(v string) *CreateRepoTagRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *CreateRepoTagRequest
	GetNamespaceName() *string
	SetRepoName(v string) *CreateRepoTagRequest
	GetRepoName() *string
	SetToTag(v string) *CreateRepoTagRequest
	GetToTag() *string
}

type CreateRepoTagRequest struct {
	// The source image tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	FromTag *string `json:"FromTag,omitempty" xml:"FromTag,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The image tag that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2
	ToTag *string `json:"ToTag,omitempty" xml:"ToTag,omitempty"`
}

func (s CreateRepoTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTagRequest) GetFromTag() *string {
	return s.FromTag
}

func (s *CreateRepoTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoTagRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateRepoTagRequest) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateRepoTagRequest) GetToTag() *string {
	return s.ToTag
}

func (s *CreateRepoTagRequest) SetFromTag(v string) *CreateRepoTagRequest {
	s.FromTag = &v
	return s
}

func (s *CreateRepoTagRequest) SetInstanceId(v string) *CreateRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTagRequest) SetNamespaceName(v string) *CreateRepoTagRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRepoTagRequest) SetRepoName(v string) *CreateRepoTagRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepoTagRequest) SetToTag(v string) *CreateRepoTagRequest {
	s.ToTag = &v
	return s
}

func (s *CreateRepoTagRequest) Validate() error {
	return dara.Validate(s)
}
