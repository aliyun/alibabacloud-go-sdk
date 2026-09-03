// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetIdpMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetIdpMetadataRequest
	GetDirectoryId() *string
	SetIdpMetadata(v string) *SetIdpMetadataRequest
	GetIdpMetadata() *string
	SetOfficeSiteId(v string) *SetIdpMetadataRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *SetIdpMetadataRequest
	GetRegionId() *string
}

type SetIdpMetadataRequest struct {
	// The office network ID, which has the same meaning as `OfficeSiteId`. We recommend that you stop using `DirectoryId` and use `OfficeSiteId` instead. You can specify only one of `DirectoryId` and `OfficeSiteId`, not both.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata of the identity provider (IdP).
	//
	// This parameter is required.
	//
	// example:
	//
	// &lt;EntityDescriptor ID********Descriptor&gt;
	IdpMetadata *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetIdpMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s SetIdpMetadataRequest) GoString() string {
	return s.String()
}

func (s *SetIdpMetadataRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetIdpMetadataRequest) GetIdpMetadata() *string {
	return s.IdpMetadata
}

func (s *SetIdpMetadataRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *SetIdpMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetIdpMetadataRequest) SetDirectoryId(v string) *SetIdpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetIdpMetadata(v string) *SetIdpMetadataRequest {
	s.IdpMetadata = &v
	return s
}

func (s *SetIdpMetadataRequest) SetOfficeSiteId(v string) *SetIdpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SetIdpMetadataRequest) SetRegionId(v string) *SetIdpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *SetIdpMetadataRequest) Validate() error {
	return dara.Validate(s)
}
