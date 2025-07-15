// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStagingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainStagingConfigRequest
	GetDomainName() *string
	SetFunctionNames(v string) *DescribeLiveDomainStagingConfigRequest
	GetFunctionNames() *string
	SetOwnerId(v int64) *DescribeLiveDomainStagingConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainStagingConfigRequest
	GetRegionId() *string
}

type DescribeLiveDomainStagingConfigRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// developer.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The feature that you want to query. Separate multiple features with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliauth
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveDomainStagingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStagingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStagingConfigRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainStagingConfigRequest) GetFunctionNames() *string {
	return s.FunctionNames
}

func (s *DescribeLiveDomainStagingConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainStagingConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainStagingConfigRequest) SetDomainName(v string) *DescribeLiveDomainStagingConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigRequest) SetFunctionNames(v string) *DescribeLiveDomainStagingConfigRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigRequest) SetOwnerId(v int64) *DescribeLiveDomainStagingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigRequest) SetRegionId(v string) *DescribeLiveDomainStagingConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainStagingConfigRequest) Validate() error {
	return dara.Validate(s)
}
