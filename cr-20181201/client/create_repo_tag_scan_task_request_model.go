// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTagScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigest(v string) *CreateRepoTagScanTaskRequest
	GetDigest() *string
	SetInstanceId(v string) *CreateRepoTagScanTaskRequest
	GetInstanceId() *string
	SetRepoId(v string) *CreateRepoTagScanTaskRequest
	GetRepoId() *string
	SetScanService(v string) *CreateRepoTagScanTaskRequest
	GetScanService() *string
	SetScanType(v string) *CreateRepoTagScanTaskRequest
	GetScanType() *string
	SetTag(v string) *CreateRepoTagScanTaskRequest
	GetTag() *string
}

type CreateRepoTagScanTaskRequest struct {
	// The image digest.
	//
	// example:
	//
	// sha256:815386ebbe9a3490f38785ab11bda34ec8dacf4634af77b8912832d4f85dca04
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The image repository ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The scan engine type:
	//
	// - `SAS_SCAN_SERVICE`: The cloud security scan engine. This is a paid service.
	//
	// - `ACR_SCAN_SERVICE`: The ACR scan engine.
	//
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	ScanType    *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The image tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRepoTagScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTagScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskRequest) GetDigest() *string {
	return s.Digest
}

func (s *CreateRepoTagScanTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateRepoTagScanTaskRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateRepoTagScanTaskRequest) GetScanService() *string {
	return s.ScanService
}

func (s *CreateRepoTagScanTaskRequest) GetScanType() *string {
	return s.ScanType
}

func (s *CreateRepoTagScanTaskRequest) GetTag() *string {
	return s.Tag
}

func (s *CreateRepoTagScanTaskRequest) SetDigest(v string) *CreateRepoTagScanTaskRequest {
	s.Digest = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetInstanceId(v string) *CreateRepoTagScanTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetRepoId(v string) *CreateRepoTagScanTaskRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetScanService(v string) *CreateRepoTagScanTaskRequest {
	s.ScanService = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetScanType(v string) *CreateRepoTagScanTaskRequest {
	s.ScanType = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetTag(v string) *CreateRepoTagScanTaskRequest {
	s.Tag = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
