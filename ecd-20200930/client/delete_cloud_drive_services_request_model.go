// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v []*string) *DeleteCloudDriveServicesRequest
	GetCdsId() []*string
	SetRegionId(v string) *DeleteCloudDriveServicesRequest
	GetRegionId() *string
}

type DeleteCloudDriveServicesRequest struct {
	// The IDs of the Wuying cloud disks.
	//
	// This parameter is required.
	CdsId []*string `json:"CdsId,omitempty" xml:"CdsId,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCloudDriveServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveServicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveServicesRequest) GetCdsId() []*string {
	return s.CdsId
}

func (s *DeleteCloudDriveServicesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudDriveServicesRequest) SetCdsId(v []*string) *DeleteCloudDriveServicesRequest {
	s.CdsId = v
	return s
}

func (s *DeleteCloudDriveServicesRequest) SetRegionId(v string) *DeleteCloudDriveServicesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudDriveServicesRequest) Validate() error {
	return dara.Validate(s)
}
