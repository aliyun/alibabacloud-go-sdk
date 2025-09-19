// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSecurityScanCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeImageSecurityScanCountRequest
	GetClusterId() *string
	SetDealed(v string) *DescribeImageSecurityScanCountRequest
	GetDealed() *string
	SetImageDigest(v string) *DescribeImageSecurityScanCountRequest
	GetImageDigest() *string
	SetImageTag(v string) *DescribeImageSecurityScanCountRequest
	GetImageTag() *string
	SetImageUuid(v string) *DescribeImageSecurityScanCountRequest
	GetImageUuid() *string
	SetRepoId(v string) *DescribeImageSecurityScanCountRequest
	GetRepoId() *string
	SetRepoInstanceId(v string) *DescribeImageSecurityScanCountRequest
	GetRepoInstanceId() *string
	SetRepoRegionId(v string) *DescribeImageSecurityScanCountRequest
	GetRepoRegionId() *string
	SetScanRange(v []*string) *DescribeImageSecurityScanCountRequest
	GetScanRange() []*string
	SetUuids(v []*string) *DescribeImageSecurityScanCountRequest
	GetUuids() []*string
}

type DescribeImageSecurityScanCountRequest struct {
	// The ID of the cluster that you want to scan.
	//
	// example:
	//
	// cdbbe7aa56cbf4b8f830f83718d26****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The handling status. Valid values:
	//
	// 	- **Y**: handled.
	//
	// 	- **N**: unhandled.
	//
	// 	- **A**: all.
	//
	// example:
	//
	// N
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The SHA-256 value of the image digest.
	//
	// example:
	//
	// a7978d51f5eddf7612ab15ae46bd4b4257bf59da77c2aafc9d9d8ab41bb3****
	ImageDigest *string `json:"ImageDigest,omitempty" xml:"ImageDigest,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// c958b80f-prd_default-9bb0****
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The UUID of the image.
	//
	// example:
	//
	// 325bfa067ae6c678e59e8a1b34cc****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
	// The ID of the Container Registry repository.
	//
	// example:
	//
	// 3df5b5a1f2339eb7ebc7d474b8d4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the Container Registry instance.
	//
	// >  You can call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to obtain the ID.
	//
	// example:
	//
	// cri-p2jahwuuwuk7****
	RepoInstanceId *string `json:"RepoInstanceId,omitempty" xml:"RepoInstanceId,omitempty"`
	// The region ID of the Container Registry repository.
	//
	// example:
	//
	// cn-beijing
	RepoRegionId *string `json:"RepoRegionId,omitempty" xml:"RepoRegionId,omitempty"`
	// The assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
	// The IDs of the instances that you want to scan.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s DescribeImageSecurityScanCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSecurityScanCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSecurityScanCountRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageSecurityScanCountRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeImageSecurityScanCountRequest) GetImageDigest() *string {
	return s.ImageDigest
}

func (s *DescribeImageSecurityScanCountRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeImageSecurityScanCountRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageSecurityScanCountRequest) GetRepoId() *string {
	return s.RepoId
}

func (s *DescribeImageSecurityScanCountRequest) GetRepoInstanceId() *string {
	return s.RepoInstanceId
}

func (s *DescribeImageSecurityScanCountRequest) GetRepoRegionId() *string {
	return s.RepoRegionId
}

func (s *DescribeImageSecurityScanCountRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageSecurityScanCountRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *DescribeImageSecurityScanCountRequest) SetClusterId(v string) *DescribeImageSecurityScanCountRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetDealed(v string) *DescribeImageSecurityScanCountRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetImageDigest(v string) *DescribeImageSecurityScanCountRequest {
	s.ImageDigest = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetImageTag(v string) *DescribeImageSecurityScanCountRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetImageUuid(v string) *DescribeImageSecurityScanCountRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetRepoId(v string) *DescribeImageSecurityScanCountRequest {
	s.RepoId = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetRepoInstanceId(v string) *DescribeImageSecurityScanCountRequest {
	s.RepoInstanceId = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetRepoRegionId(v string) *DescribeImageSecurityScanCountRequest {
	s.RepoRegionId = &v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetScanRange(v []*string) *DescribeImageSecurityScanCountRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) SetUuids(v []*string) *DescribeImageSecurityScanCountRequest {
	s.Uuids = v
	return s
}

func (s *DescribeImageSecurityScanCountRequest) Validate() error {
	return dara.Validate(s)
}
