// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *GetRepoTagScanStatusRequest
	GetDigest() *string
	SetInstanceId(v string) *GetRepoTagScanStatusRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepoTagScanStatusRequest
	GetRepoId() *string
	SetScanTaskId(v string) *GetRepoTagScanStatusRequest
	GetScanTaskId() *string
	SetScanType(v string) *GetRepoTagScanStatusRequest
	GetScanType() *string
	SetTag(v string) *GetRepoTagScanStatusRequest
	GetTag() *string
}

type GetRepoTagScanStatusRequest struct {
	// The image digest.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1d529
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-uf082u9dg8do****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 838152F9-F725-5A52-A344-8972D65AC045
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	ScanType   *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The image tag.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagScanStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusRequest) GetDigest() *string {
	return s.Digest
}

func (s *GetRepoTagScanStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoTagScanStatusRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoTagScanStatusRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *GetRepoTagScanStatusRequest) GetScanType() *string {
	return s.ScanType
}

func (s *GetRepoTagScanStatusRequest) GetTag() *string {
	return s.Tag
}

func (s *GetRepoTagScanStatusRequest) SetDigest(v string) *GetRepoTagScanStatusRequest {
	s.Digest = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetInstanceId(v string) *GetRepoTagScanStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetRepoId(v string) *GetRepoTagScanStatusRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetScanTaskId(v string) *GetRepoTagScanStatusRequest {
	s.ScanTaskId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetScanType(v string) *GetRepoTagScanStatusRequest {
	s.ScanType = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetTag(v string) *GetRepoTagScanStatusRequest {
	s.Tag = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) Validate() error {
	return dara.Validate(s)
}
