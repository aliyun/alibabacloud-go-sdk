// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanMaliciousFileByTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *ListScanMaliciousFileByTaskRequest
	GetDigest() *string
	SetInstanceId(v string) *ListScanMaliciousFileByTaskRequest
	GetInstanceId() *string
	SetLevel(v string) *ListScanMaliciousFileByTaskRequest
	GetLevel() *string
	SetPageNo(v int32) *ListScanMaliciousFileByTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanMaliciousFileByTaskRequest
	GetPageSize() *int32
	SetRepoId(v string) *ListScanMaliciousFileByTaskRequest
	GetRepoId() *string
	SetScanTaskId(v string) *ListScanMaliciousFileByTaskRequest
	GetScanTaskId() *string
	SetTag(v string) *ListScanMaliciousFileByTaskRequest
	GetTag() *string
}

type ListScanMaliciousFileByTaskRequest struct {
	// The image digest.
	//
	// example:
	//
	// sha256:aa4bffff6406785e930dab1f94c9a1297ba22xxxx71d71504a015764*********
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-gu94qynvpwk*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The severity of the malicious file.
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
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The image repository ID.
	//
	// example:
	//
	// crr-h1y4l279wb8*****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 79ee6bc2-3288-4c56-b967-**********
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The image tag.
	//
	// example:
	//
	// V6.11
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListScanMaliciousFileByTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScanMaliciousFileByTaskRequest) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskRequest) GetDigest() *string {
	return s.Digest
}

func (s *ListScanMaliciousFileByTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScanMaliciousFileByTaskRequest) GetLevel() *string {
	return s.Level
}

func (s *ListScanMaliciousFileByTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanMaliciousFileByTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanMaliciousFileByTaskRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *ListScanMaliciousFileByTaskRequest) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListScanMaliciousFileByTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *ListScanMaliciousFileByTaskRequest) SetDigest(v string) *ListScanMaliciousFileByTaskRequest {
	s.Digest = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetInstanceId(v string) *ListScanMaliciousFileByTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetLevel(v string) *ListScanMaliciousFileByTaskRequest {
	s.Level = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetPageNo(v int32) *ListScanMaliciousFileByTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetPageSize(v int32) *ListScanMaliciousFileByTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetRepoId(v string) *ListScanMaliciousFileByTaskRequest {
	s.RepoId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetScanTaskId(v string) *ListScanMaliciousFileByTaskRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetTag(v string) *ListScanMaliciousFileByTaskRequest {
	s.Tag = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) Validate() error {
	return dara.Validate(s)
}
