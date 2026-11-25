// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRepoBuildRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *CancelRepoBuildRecordRequest
	GetBuildRecordId() *string
	SetInstanceId(v string) *CancelRepoBuildRecordRequest
	GetInstanceId() *string
	SetRepoId(v string) *CancelRepoBuildRecordRequest
	GetRepoId() *string
}

type CancelRepoBuildRecordRequest struct {
	// The build record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74FDBA62-30C0-4F22-BE7B-F1D36FD1****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The image repository ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CancelRepoBuildRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordRequest) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *CancelRepoBuildRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelRepoBuildRecordRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CancelRepoBuildRecordRequest) SetBuildRecordId(v string) *CancelRepoBuildRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *CancelRepoBuildRecordRequest) SetInstanceId(v string) *CancelRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelRepoBuildRecordRequest) SetRepoId(v string) *CancelRepoBuildRecordRequest {
	s.RepoId = &v
	return s
}

func (s *CancelRepoBuildRecordRequest) Validate() error {
	return dara.Validate(s)
}
