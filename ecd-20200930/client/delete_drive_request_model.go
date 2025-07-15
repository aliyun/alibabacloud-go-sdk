// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DeleteDriveRequest
	GetDriveId() *string
	SetRegionId(v string) *DeleteDriveRequest
	GetRegionId() *string
}

type DeleteDriveRequest struct {
	// example:
	//
	// dri-aaaa****
	DriveId *string `json:"DriveId,omitempty" xml:"DriveId,omitempty"`
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

func (s *DeleteDriveRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DeleteDriveRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDriveRequest) SetDriveId(v string) *DeleteDriveRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteDriveRequest) SetRegionId(v string) *DeleteDriveRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDriveRequest) Validate() error {
	return dara.Validate(s)
}
