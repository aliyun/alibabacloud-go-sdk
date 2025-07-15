// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpMetadataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetSpMetadataRequest
	GetDirectoryId() *string
	SetOfficeSiteId(v string) *GetSpMetadataRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *GetSpMetadataRequest
	GetRegionId() *string
}

type GetSpMetadataRequest struct {
	// The workspace ID. This parameter is the same as `OfficeSiteId`. We recommend that you use `OfficeSiteId` to replace `DirectoryId`. You can specify only `DirectoryId` or `OfficeSiteId`.
	//
	// example:
	//
	// cn-hangzhou+dir-400695****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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

func (s GetSpMetadataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSpMetadataRequest) GoString() string {
	return s.String()
}

func (s *GetSpMetadataRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetSpMetadataRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *GetSpMetadataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSpMetadataRequest) SetDirectoryId(v string) *GetSpMetadataRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetSpMetadataRequest) SetOfficeSiteId(v string) *GetSpMetadataRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetSpMetadataRequest) SetRegionId(v string) *GetSpMetadataRequest {
	s.RegionId = &v
	return s
}

func (s *GetSpMetadataRequest) Validate() error {
	return dara.Validate(s)
}
