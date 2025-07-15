// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveFotaUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *ApproveFotaUpdateRequest
	GetAppVersion() *string
	SetDesktopId(v string) *ApproveFotaUpdateRequest
	GetDesktopId() *string
	SetRegionId(v string) *ApproveFotaUpdateRequest
	GetRegionId() *string
}

type ApproveFotaUpdateRequest struct {
	// Mirror version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1-D-20220513.143129
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-138dsptkrt00u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApproveFotaUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveFotaUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ApproveFotaUpdateRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ApproveFotaUpdateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApproveFotaUpdateRequest) SetAppVersion(v string) *ApproveFotaUpdateRequest {
	s.AppVersion = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetDesktopId(v string) *ApproveFotaUpdateRequest {
	s.DesktopId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetRegionId(v string) *ApproveFotaUpdateRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) Validate() error {
	return dara.Validate(s)
}
