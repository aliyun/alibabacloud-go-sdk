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
	// The workspace ID. This parameter is the same as `OfficeSiteId`. We recommend that you use `OfficeSiteId` to replace `DirectoryId`. You can specify only `DirectoryId` or `OfficeSiteId`.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata of the IdP.
	//
	// This parameter is required.
	//
	// example:
	//
	// &lt;EntityDescriptor ID********Descriptor&gt;
	IdpMetadata *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
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
