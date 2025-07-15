// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainMappingRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainMappingRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainMappingRequest
	GetRegionId() *string
}

type DescribeLiveDomainMappingRequest struct {
	// The domain name for which you want to query the mappings. The following types of domain names are supported:
	//
	// 	- Ingest domain
	//
	// 	- Main streaming domain
	//
	// 	- Sub-streaming domain
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveDomainMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMappingRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMappingRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainMappingRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainMappingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainMappingRequest) SetDomainName(v string) *DescribeLiveDomainMappingRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMappingRequest) SetOwnerId(v int64) *DescribeLiveDomainMappingRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainMappingRequest) SetRegionId(v string) *DescribeLiveDomainMappingRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainMappingRequest) Validate() error {
	return dara.Validate(s)
}
