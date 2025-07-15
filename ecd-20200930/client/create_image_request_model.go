// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCleanUserdata(v bool) *CreateImageRequest
	GetAutoCleanUserdata() *bool
	SetDataSnapshotIds(v []*string) *CreateImageRequest
	GetDataSnapshotIds() []*string
	SetDescription(v string) *CreateImageRequest
	GetDescription() *string
	SetDesktopId(v string) *CreateImageRequest
	GetDesktopId() *string
	SetDiskType(v string) *CreateImageRequest
	GetDiskType() *string
	SetImageName(v string) *CreateImageRequest
	GetImageName() *string
	SetImageResourceType(v string) *CreateImageRequest
	GetImageResourceType() *string
	SetRegionId(v string) *CreateImageRequest
	GetRegionId() *string
	SetSnapshotId(v string) *CreateImageRequest
	GetSnapshotId() *string
	SetSnapshotIds(v []*string) *CreateImageRequest
	GetSnapshotIds() []*string
}

type CreateImageRequest struct {
	// Specifies whether to clear private data of users. If you set AutoCleanUserdata to `true`, the custom image clears the data directories, excluding the `Administrator` and `Public` directories, in the `C:\\Users` directory.
	//
	// example:
	//
	// false
	AutoCleanUserdata *bool     `json:"AutoCleanUserdata,omitempty" xml:"AutoCleanUserdata,omitempty"`
	DataSnapshotIds   []*string `json:"DataSnapshotIds,omitempty" xml:"DataSnapshotIds,omitempty" type:"Repeated"`
	// The description of the custom image. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The disk data that is contained in the custom image.
	//
	// Valid values:
	//
	// - SYSTEM: only contain data from system disks.
	//
	// - ALL: contain data from system disks and user disks. [default]
	//
	// example:
	//
	// ALL
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The name of the image. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testImageName
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// This parameter is not publicly available.
	//
	// example:
	//
	// To be hidden.
	ImageResourceType *string `json:"ImageResourceType,omitempty" xml:"ImageResourceType,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-2zefuwk8l6ytcgd3bf4o
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The IDs of the snapshots.
	SnapshotIds []*string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty" type:"Repeated"`
}

func (s CreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) GetAutoCleanUserdata() *bool {
	return s.AutoCleanUserdata
}

func (s *CreateImageRequest) GetDataSnapshotIds() []*string {
	return s.DataSnapshotIds
}

func (s *CreateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *CreateImageRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateImageRequest) GetImageResourceType() *string {
	return s.ImageResourceType
}

func (s *CreateImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateImageRequest) GetSnapshotIds() []*string {
	return s.SnapshotIds
}

func (s *CreateImageRequest) SetAutoCleanUserdata(v bool) *CreateImageRequest {
	s.AutoCleanUserdata = &v
	return s
}

func (s *CreateImageRequest) SetDataSnapshotIds(v []*string) *CreateImageRequest {
	s.DataSnapshotIds = v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetDesktopId(v string) *CreateImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateImageRequest) SetDiskType(v string) *CreateImageRequest {
	s.DiskType = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetImageResourceType(v string) *CreateImageRequest {
	s.ImageResourceType = &v
	return s
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotId(v string) *CreateImageRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateImageRequest) SetSnapshotIds(v []*string) *CreateImageRequest {
	s.SnapshotIds = v
	return s
}

func (s *CreateImageRequest) Validate() error {
	return dara.Validate(s)
}
