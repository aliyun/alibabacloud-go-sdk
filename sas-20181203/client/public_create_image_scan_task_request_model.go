// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicCreateImageScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigests(v string) *PublicCreateImageScanTaskRequest
	GetDigests() *string
	SetInstanceIds(v string) *PublicCreateImageScanTaskRequest
	GetInstanceIds() *string
	SetRegionIds(v string) *PublicCreateImageScanTaskRequest
	GetRegionIds() *string
	SetRegistryTypes(v string) *PublicCreateImageScanTaskRequest
	GetRegistryTypes() *string
	SetRepoIds(v string) *PublicCreateImageScanTaskRequest
	GetRepoIds() *string
	SetRepoNames(v string) *PublicCreateImageScanTaskRequest
	GetRepoNames() *string
	SetRepoNamespaces(v string) *PublicCreateImageScanTaskRequest
	GetRepoNamespaces() *string
	SetSourceIp(v string) *PublicCreateImageScanTaskRequest
	GetSourceIp() *string
	SetTags(v string) *PublicCreateImageScanTaskRequest
	GetTags() *string
}

type PublicCreateImageScanTaskRequest struct {
	// The SHA-256 value of the image digest. Separate multiple SHA-256 values with commas (,).
	//
	// example:
	//
	// 6a5e103187b31a94592a47a5858617f7a6c
	Digests *string `json:"Digests,omitempty" xml:"Digests,omitempty"`
	// The ID of the Container Registry instance in which the image repository is created. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// i-uf6j8vq9l4r5ntht****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the image. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIds *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The type of the image repository. Separate multiple types with commas (,). Valid values:
	//
	// 	- **acr**
	//
	// 	- **harbor**
	//
	// 	- **quay**
	//
	// example:
	//
	// acr
	RegistryTypes *string `json:"RegistryTypes,omitempty" xml:"RegistryTypes,omitempty"`
	// The ID of the image repository. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// crr-vridcl4****
	RepoIds *string `json:"RepoIds,omitempty" xml:"RepoIds,omitempty"`
	// The name of the image repository. Separate multiple names with commas (,).
	//
	// example:
	//
	// centos
	RepoNames *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
	// The namespace to which the image repository belongs. Separate multiple namespaces with commas (,).
	//
	// example:
	//
	// hanghai-namespace
	RepoNamespaces *string `json:"RepoNamespaces,omitempty" xml:"RepoNamespaces,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168..XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The tag that is added to the image. Separate multiple tags with commas (,).
	//
	// example:
	//
	// 0.2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s PublicCreateImageScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PublicCreateImageScanTaskRequest) GoString() string {
	return s.String()
}

func (s *PublicCreateImageScanTaskRequest) GetDigests() *string {
	return s.Digests
}

func (s *PublicCreateImageScanTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *PublicCreateImageScanTaskRequest) GetRegionIds() *string {
	return s.RegionIds
}

func (s *PublicCreateImageScanTaskRequest) GetRegistryTypes() *string {
	return s.RegistryTypes
}

func (s *PublicCreateImageScanTaskRequest) GetRepoIds() *string {
	return s.RepoIds
}

func (s *PublicCreateImageScanTaskRequest) GetRepoNames() *string {
	return s.RepoNames
}

func (s *PublicCreateImageScanTaskRequest) GetRepoNamespaces() *string {
	return s.RepoNamespaces
}

func (s *PublicCreateImageScanTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PublicCreateImageScanTaskRequest) GetTags() *string {
	return s.Tags
}

func (s *PublicCreateImageScanTaskRequest) SetDigests(v string) *PublicCreateImageScanTaskRequest {
	s.Digests = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetInstanceIds(v string) *PublicCreateImageScanTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetRegionIds(v string) *PublicCreateImageScanTaskRequest {
	s.RegionIds = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetRegistryTypes(v string) *PublicCreateImageScanTaskRequest {
	s.RegistryTypes = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetRepoIds(v string) *PublicCreateImageScanTaskRequest {
	s.RepoIds = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetRepoNames(v string) *PublicCreateImageScanTaskRequest {
	s.RepoNames = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetRepoNamespaces(v string) *PublicCreateImageScanTaskRequest {
	s.RepoNamespaces = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetSourceIp(v string) *PublicCreateImageScanTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) SetTags(v string) *PublicCreateImageScanTaskRequest {
	s.Tags = &v
	return s
}

func (s *PublicCreateImageScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
