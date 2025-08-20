// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanBaselineByTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *ListScanBaselineByTaskRequest
	GetDigest() *string
	SetInstanceId(v string) *ListScanBaselineByTaskRequest
	GetInstanceId() *string
	SetLevel(v string) *ListScanBaselineByTaskRequest
	GetLevel() *string
	SetPageNo(v int32) *ListScanBaselineByTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanBaselineByTaskRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListScanBaselineByTaskRequest
	GetRepoId() *string
	SetScanTaskId(v string) *ListScanBaselineByTaskRequest
	GetScanTaskId() *string
	SetTag(v string) *ListScanBaselineByTaskRequest
	GetTag() *string
}

type ListScanBaselineByTaskRequest struct {
	// The digest value of the image.
	//
	// example:
	//
	// sha256:1c89806cfaf66d2990e2cf1131ebd56ff24b133745a33abf1228*************
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The level of the baseline risk.
	//
	// example:
	//
	// High
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-**************
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 3e526d7e-4b45-4703-b046-***********
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The image version.
	//
	// example:
	//
	// 1.1.36
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListScanBaselineByTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScanBaselineByTaskRequest) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskRequest) GetDigest() *string {
	return s.Digest
}

func (s *ListScanBaselineByTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScanBaselineByTaskRequest) GetLevel() *string {
	return s.Level
}

func (s *ListScanBaselineByTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanBaselineByTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanBaselineByTaskRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListScanBaselineByTaskRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListScanBaselineByTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *ListScanBaselineByTaskRequest) SetDigest(v string) *ListScanBaselineByTaskRequest {
	s.Digest = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetInstanceId(v string) *ListScanBaselineByTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetLevel(v string) *ListScanBaselineByTaskRequest {
	s.Level = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetPageNo(v int32) *ListScanBaselineByTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetPageSize(v int32) *ListScanBaselineByTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetRepoId(v string) *ListScanBaselineByTaskRequest {
	s.RepoId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetScanTaskId(v string) *ListScanBaselineByTaskRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetTag(v string) *ListScanBaselineByTaskRequest {
	s.Tag = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) Validate() error {
	return dara.Validate(s)
}
