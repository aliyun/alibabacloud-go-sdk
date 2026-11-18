// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *DeleteDriveRequest
	GetDomainId() *string
	SetDriveId(v string) *DeleteDriveRequest
	GetDriveId() *string
	SetEndUserId(v string) *DeleteDriveRequest
	GetEndUserId() *string
	SetOfficeSiteId(v string) *DeleteDriveRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DeleteDriveRequest
	GetRegionId() *string
}

type DeleteDriveRequest struct {
	// The instance ID.
	//
	// example:
	//
	// dom-aaaa****
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the user-level storage resource.
	//
	// example:
	//
	// dri-aaaa****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
	// The ID of the end user. This parameter is used when locating the UPM drive by using DomainId or OfficeSiteId.
	//
	// example:
	//
	// abc****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDriveRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveRequest) GoString() string {
	return s.String()
}

func (s *DeleteDriveRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *DeleteDriveRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteDriveRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DeleteDriveRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DeleteDriveRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDriveRequest) SetDomainId(v string) *DeleteDriveRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDriveRequest) SetDriveId(v string) *DeleteDriveRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteDriveRequest) SetEndUserId(v string) *DeleteDriveRequest {
	s.EndUserId = &v
	return s
}

func (s *DeleteDriveRequest) SetOfficeSiteId(v string) *DeleteDriveRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DeleteDriveRequest) SetRegionId(v string) *DeleteDriveRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDriveRequest) Validate() error {
	return dara.Validate(s)
}
