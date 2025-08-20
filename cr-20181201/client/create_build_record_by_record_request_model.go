// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBuildRecordByRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *CreateBuildRecordByRecordRequest
	GetBuildRecordId() *string
	SetInstanceId(v string) *CreateBuildRecordByRecordRequest
	GetInstanceId() *string
	SetRepoId(v string) *CreateBuildRecordByRecordRequest
	GetRepoId() *string
}

type CreateBuildRecordByRecordRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0A311FC5-B8C6-4332-80E4-539EB73****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-hnoq7j93or3k****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateBuildRecordByRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBuildRecordByRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordRequest) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *CreateBuildRecordByRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBuildRecordByRecordRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateBuildRecordByRecordRequest) SetBuildRecordId(v string) *CreateBuildRecordByRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRecordRequest) SetInstanceId(v string) *CreateBuildRecordByRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBuildRecordByRecordRequest) SetRepoId(v string) *CreateBuildRecordByRecordRequest {
	s.RepoId = &v
	return s
}

func (s *CreateBuildRecordByRecordRequest) Validate() error {
	return dara.Validate(s)
}
