// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeLiveDomainResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ChangeLiveDomainResourceGroupRequest
	GetDomainName() *string
	SetNewResourceGroupId(v string) *ChangeLiveDomainResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerId(v int64) *ChangeLiveDomainResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ChangeLiveDomainResourceGroupRequest
	GetRegionId() *string
}

type ChangeLiveDomainResourceGroupRequest struct {
	// The ingest domain or streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the resource group to which the domain name is moved.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-****ke6uuxw****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChangeLiveDomainResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeLiveDomainResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeLiveDomainResourceGroupRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ChangeLiveDomainResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeLiveDomainResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeLiveDomainResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChangeLiveDomainResourceGroupRequest) SetDomainName(v string) *ChangeLiveDomainResourceGroupRequest {
	s.DomainName = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeLiveDomainResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupRequest) SetOwnerId(v int64) *ChangeLiveDomainResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupRequest) SetRegionId(v string) *ChangeLiveDomainResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeLiveDomainResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
