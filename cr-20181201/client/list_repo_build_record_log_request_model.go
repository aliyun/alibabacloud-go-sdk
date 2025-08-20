// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRecordLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildRecordId(v string) *ListRepoBuildRecordLogRequest
	GetBuildRecordId() *string
	SetInstanceId(v string) *ListRepoBuildRecordLogRequest
	GetInstanceId() *string
	SetOffset(v int32) *ListRepoBuildRecordLogRequest
	GetOffset() *int32
	SetRepoId(v string) *ListRepoBuildRecordLogRequest
	GetRepoId() *string
}

type ListRepoBuildRecordLogRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// C5B4D5D7-A1C6-4E9B-ABD2-401361C4****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-nmbv37dlv5d3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The offset of log lines.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-z4dvahhku9wv4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRecordLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRecordLogRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogRequest) GetBuildRecordId() *string {
	return s.BuildRecordId
}

func (s *ListRepoBuildRecordLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRepoBuildRecordLogRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListRepoBuildRecordLogRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListRepoBuildRecordLogRequest) SetBuildRecordId(v string) *ListRepoBuildRecordLogRequest {
	s.BuildRecordId = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetInstanceId(v string) *ListRepoBuildRecordLogRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetOffset(v int32) *ListRepoBuildRecordLogRequest {
	s.Offset = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetRepoId(v string) *ListRepoBuildRecordLogRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) Validate() error {
	return dara.Validate(s)
}
