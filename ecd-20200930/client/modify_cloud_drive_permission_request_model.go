// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudDrivePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *ModifyCloudDrivePermissionRequest
	GetCdsId() *string
	SetDownloadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest
	GetDownloadEndUserIds() []*string
	SetDownloadUploadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest
	GetDownloadUploadEndUserIds() []*string
	SetNoDownloadNoUploadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest
	GetNoDownloadNoUploadEndUserIds() []*string
	SetRegionId(v string) *ModifyCloudDrivePermissionRequest
	GetRegionId() *string
}

type ModifyCloudDrivePermissionRequest struct {
	// The ID of the cloud disk in Cloud Drive Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-60911*****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The IDs of the users who have the download permissions.
	DownloadEndUserIds []*string `json:"DownloadEndUserIds,omitempty" xml:"DownloadEndUserIds,omitempty" type:"Repeated"`
	// The IDs of the users who have the upload and download permissions.
	DownloadUploadEndUserIds     []*string `json:"DownloadUploadEndUserIds,omitempty" xml:"DownloadUploadEndUserIds,omitempty" type:"Repeated"`
	NoDownloadNoUploadEndUserIds []*string `json:"NoDownloadNoUploadEndUserIds,omitempty" xml:"NoDownloadNoUploadEndUserIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyCloudDrivePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudDrivePermissionRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudDrivePermissionRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *ModifyCloudDrivePermissionRequest) GetDownloadEndUserIds() []*string {
	return s.DownloadEndUserIds
}

func (s *ModifyCloudDrivePermissionRequest) GetDownloadUploadEndUserIds() []*string {
	return s.DownloadUploadEndUserIds
}

func (s *ModifyCloudDrivePermissionRequest) GetNoDownloadNoUploadEndUserIds() []*string {
	return s.NoDownloadNoUploadEndUserIds
}

func (s *ModifyCloudDrivePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudDrivePermissionRequest) SetCdsId(v string) *ModifyCloudDrivePermissionRequest {
	s.CdsId = &v
	return s
}

func (s *ModifyCloudDrivePermissionRequest) SetDownloadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest {
	s.DownloadEndUserIds = v
	return s
}

func (s *ModifyCloudDrivePermissionRequest) SetDownloadUploadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest {
	s.DownloadUploadEndUserIds = v
	return s
}

func (s *ModifyCloudDrivePermissionRequest) SetNoDownloadNoUploadEndUserIds(v []*string) *ModifyCloudDrivePermissionRequest {
	s.NoDownloadNoUploadEndUserIds = v
	return s
}

func (s *ModifyCloudDrivePermissionRequest) SetRegionId(v string) *ModifyCloudDrivePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudDrivePermissionRequest) Validate() error {
	return dara.Validate(s)
}
