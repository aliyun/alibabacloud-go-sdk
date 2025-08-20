// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoBuildRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *GetRepoBuildRecordStatusRequest
	GetBuildRecordId() *string
	SetInstanceId(v string) *GetRepoBuildRecordStatusRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepoBuildRecordStatusRequest
	GetRepoId() *string
}

type GetRepoBuildRecordStatusRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-jnzm47ihjmgc****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetRepoBuildRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoBuildRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusRequest) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *GetRepoBuildRecordStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoBuildRecordStatusRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoBuildRecordStatusRequest) SetBuildRecordId(v string) *GetRepoBuildRecordStatusRequest {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordStatusRequest) SetInstanceId(v string) *GetRepoBuildRecordStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoBuildRecordStatusRequest) SetRepoId(v string) *GetRepoBuildRecordStatusRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoBuildRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
