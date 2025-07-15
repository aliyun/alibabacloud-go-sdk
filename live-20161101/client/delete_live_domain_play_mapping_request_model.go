// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDomainPlayMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLiveDomainPlayMappingRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *DeleteLiveDomainPlayMappingRequest
	GetPlayDomain() *string
	SetPullDomain(v string) *DeleteLiveDomainPlayMappingRequest
	GetPullDomain() *string
	SetRegionId(v string) *DeleteLiveDomainPlayMappingRequest
	GetRegionId() *string
}

type DeleteLiveDomainPlayMappingRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The sub-streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	PullDomain *string `json:"PullDomain,omitempty" xml:"PullDomain,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLiveDomainPlayMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDomainPlayMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveDomainPlayMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveDomainPlayMappingRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *DeleteLiveDomainPlayMappingRequest) GetPullDomain() *string {
	return s.PullDomain
}

func (s *DeleteLiveDomainPlayMappingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveDomainPlayMappingRequest) SetOwnerId(v int64) *DeleteLiveDomainPlayMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingRequest) SetPlayDomain(v string) *DeleteLiveDomainPlayMappingRequest {
	s.PlayDomain = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingRequest) SetPullDomain(v string) *DeleteLiveDomainPlayMappingRequest {
	s.PullDomain = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingRequest) SetRegionId(v string) *DeleteLiveDomainPlayMappingRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveDomainPlayMappingRequest) Validate() error {
	return dara.Validate(s)
}
