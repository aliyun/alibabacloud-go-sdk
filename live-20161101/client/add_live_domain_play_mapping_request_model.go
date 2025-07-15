// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainPlayMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddLiveDomainPlayMappingRequest
	GetOwnerId() *int64
	SetPlayDomain(v string) *AddLiveDomainPlayMappingRequest
	GetPlayDomain() *string
	SetPullDomain(v string) *AddLiveDomainPlayMappingRequest
	GetPullDomain() *string
	SetRegionId(v string) *AddLiveDomainPlayMappingRequest
	GetRegionId() *string
}

type AddLiveDomainPlayMappingRequest struct {
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

func (s AddLiveDomainPlayMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainPlayMappingRequest) GoString() string {
	return s.String()
}

func (s *AddLiveDomainPlayMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddLiveDomainPlayMappingRequest) GetPlayDomain() *string {
	return s.PlayDomain
}

func (s *AddLiveDomainPlayMappingRequest) GetPullDomain() *string {
	return s.PullDomain
}

func (s *AddLiveDomainPlayMappingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddLiveDomainPlayMappingRequest) SetOwnerId(v int64) *AddLiveDomainPlayMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *AddLiveDomainPlayMappingRequest) SetPlayDomain(v string) *AddLiveDomainPlayMappingRequest {
	s.PlayDomain = &v
	return s
}

func (s *AddLiveDomainPlayMappingRequest) SetPullDomain(v string) *AddLiveDomainPlayMappingRequest {
	s.PullDomain = &v
	return s
}

func (s *AddLiveDomainPlayMappingRequest) SetRegionId(v string) *AddLiveDomainPlayMappingRequest {
	s.RegionId = &v
	return s
}

func (s *AddLiveDomainPlayMappingRequest) Validate() error {
	return dara.Validate(s)
}
