// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DeleteRCDiskRequest
	GetDiskId() *string
	SetRegionId(v string) *DeleteRCDiskRequest
	GetRegionId() *string
}

type DeleteRCDiskRequest struct {
	// The ID of the cloud disk that you want to release.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcd-wz9c8isqly8637zw****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRCDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DeleteRCDiskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCDiskRequest) SetDiskId(v string) *DeleteRCDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DeleteRCDiskRequest) SetRegionId(v string) *DeleteRCDiskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCDiskRequest) Validate() error {
	return dara.Validate(s)
}
