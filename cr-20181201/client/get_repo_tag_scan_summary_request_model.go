// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepoTagScanSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *GetRepoTagScanSummaryRequest
	GetDigest() *string
	SetInstanceId(v string) *GetRepoTagScanSummaryRequest
	GetInstanceId() *string
	SetRepoId(v string) *GetRepoTagScanSummaryRequest
	GetRepoId() *string
	SetScanTaskId(v string) *GetRepoTagScanSummaryRequest
	GetScanTaskId() *string
	SetTag(v string) *GetRepoTagScanSummaryRequest
	GetTag() *string
}

type GetRepoTagScanSummaryRequest struct {
	// The number of unknown-severity vulnerabilities.
	//
	// example:
	//
	// sha256:c9f370a4eb1c00d0b0d7212a0a9fa4a7697756c90f0f680afaf9737a25725f4c
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image tag.
	//
	// example:
	//
	// crr-c2i5yk6h6pu9d5o8
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 47A3E5A3-6AD4-5F02-93B8-59F778AE25D4
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The ID of the security scan task.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagScanSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepoTagScanSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryRequest) GetDigest() *string {
	return s.Digest
}

func (s *GetRepoTagScanSummaryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRepoTagScanSummaryRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *GetRepoTagScanSummaryRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *GetRepoTagScanSummaryRequest) GetTag() *string {
	return s.Tag
}

func (s *GetRepoTagScanSummaryRequest) SetDigest(v string) *GetRepoTagScanSummaryRequest {
	s.Digest = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetInstanceId(v string) *GetRepoTagScanSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetRepoId(v string) *GetRepoTagScanSummaryRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetScanTaskId(v string) *GetRepoTagScanSummaryRequest {
	s.ScanTaskId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetTag(v string) *GetRepoTagScanSummaryRequest {
	s.Tag = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) Validate() error {
	return dara.Validate(s)
}
