// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicPreCheckImageScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigests(v string) *PublicPreCheckImageScanTaskRequest
	GetDigests() *string
	SetInstanceIds(v string) *PublicPreCheckImageScanTaskRequest
	GetInstanceIds() *string
	SetRegionIds(v string) *PublicPreCheckImageScanTaskRequest
	GetRegionIds() *string
	SetRegistryTypes(v string) *PublicPreCheckImageScanTaskRequest
	GetRegistryTypes() *string
	SetRepoIds(v string) *PublicPreCheckImageScanTaskRequest
	GetRepoIds() *string
	SetRepoNames(v string) *PublicPreCheckImageScanTaskRequest
	GetRepoNames() *string
	SetRepoNamespaces(v string) *PublicPreCheckImageScanTaskRequest
	GetRepoNamespaces() *string
	SetSourceIp(v string) *PublicPreCheckImageScanTaskRequest
	GetSourceIp() *string
	SetTags(v string) *PublicPreCheckImageScanTaskRequest
	GetTags() *string
}

type PublicPreCheckImageScanTaskRequest struct {
	// The SHA-256 value of the image digest. Separate multiple SHA-256 values with commas (,).
	//
	// example:
	//
	// 6a5e103187b31a94592a47a5858617f7****
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
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The tag that is added to the image. Separate multiple tags with commas (,).
	//
	// example:
	//
	// 0.2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s PublicPreCheckImageScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PublicPreCheckImageScanTaskRequest) GoString() string {
	return s.String()
}

func (s *PublicPreCheckImageScanTaskRequest) GetDigests() *string {
	return s.Digests
}

func (s *PublicPreCheckImageScanTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *PublicPreCheckImageScanTaskRequest) GetRegionIds() *string {
	return s.RegionIds
}

func (s *PublicPreCheckImageScanTaskRequest) GetRegistryTypes() *string {
	return s.RegistryTypes
}

func (s *PublicPreCheckImageScanTaskRequest) GetRepoIds() *string {
	return s.RepoIds
}

func (s *PublicPreCheckImageScanTaskRequest) GetRepoNames() *string {
	return s.RepoNames
}

func (s *PublicPreCheckImageScanTaskRequest) GetRepoNamespaces() *string {
	return s.RepoNamespaces
}

func (s *PublicPreCheckImageScanTaskRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *PublicPreCheckImageScanTaskRequest) GetTags() *string {
	return s.Tags
}

func (s *PublicPreCheckImageScanTaskRequest) SetDigests(v string) *PublicPreCheckImageScanTaskRequest {
	s.Digests = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetInstanceIds(v string) *PublicPreCheckImageScanTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetRegionIds(v string) *PublicPreCheckImageScanTaskRequest {
	s.RegionIds = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetRegistryTypes(v string) *PublicPreCheckImageScanTaskRequest {
	s.RegistryTypes = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetRepoIds(v string) *PublicPreCheckImageScanTaskRequest {
	s.RepoIds = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetRepoNames(v string) *PublicPreCheckImageScanTaskRequest {
	s.RepoNames = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetRepoNamespaces(v string) *PublicPreCheckImageScanTaskRequest {
	s.RepoNamespaces = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetSourceIp(v string) *PublicPreCheckImageScanTaskRequest {
	s.SourceIp = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) SetTags(v string) *PublicPreCheckImageScanTaskRequest {
	s.Tags = &v
	return s
}

func (s *PublicPreCheckImageScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
