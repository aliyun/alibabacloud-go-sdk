// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRepoTagRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepoTagRequest
	GetRepoId() *string
	SetTag(v string) *GetRepoTagRequest
	GetTag() *string
}

type GetRepoTagRequest struct {
	// The return value of status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Set the value to **GetRepoTag**.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The number of milliseconds that have elapsed since the image was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoTagRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoTagRequest) GetTag() *string {
	return s.Tag
}

func (s *GetRepoTagRequest) SetInstanceId(v string) *GetRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagRequest) SetRepoId(v string) *GetRepoTagRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagRequest) SetTag(v string) *GetRepoTagRequest {
	s.Tag = &v
	return s
}

func (s *GetRepoTagRequest) Validate() error {
	return dara.Validate(s)
}
