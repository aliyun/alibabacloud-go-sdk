// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeImagesRequest
	GetEnsRegionId() *string
	SetImageId(v string) *DescribeImagesRequest
	GetImageId() *string
	SetImageName(v string) *DescribeImagesRequest
	GetImageName() *string
	SetPageNumber(v string) *DescribeImagesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeImagesRequest
	GetPageSize() *string
	SetSnapshotId(v string) *DescribeImagesRequest
	GetSnapshotId() *string
	SetStatus(v string) *DescribeImagesRequest
	GetStatus() *string
}

type DescribeImagesRequest struct {
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-dalian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the image. You can specify only one image ID.
	//
	// Custom images and public images are supported.
	//
	// example:
	//
	// m-5qm2r6xo7njrpdkx04q1o****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length The name must start with a letter and cannot start with `acs:` or `aliyun`. The name cannot contain `http://` or `https://`. The name can contain letters, digits, periods (.), colons (:), underscores (_), and hyphens (-).
	//
	// By default, this parameter is left empty, which indicates that the original name is retained.
	//
	// example:
	//
	// centos_6_08_64_20G_a****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The page number. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 50
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// mock-clone_snapshot_id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// This parameter is not currently in use.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeImagesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImagesRequest) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeImagesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeImagesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeImagesRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeImagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeImagesRequest) SetEnsRegionId(v string) *DescribeImagesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeImagesRequest) SetImageId(v string) *DescribeImagesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagesRequest) SetImageName(v string) *DescribeImagesRequest {
	s.ImageName = &v
	return s
}

func (s *DescribeImagesRequest) SetPageNumber(v string) *DescribeImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeImagesRequest) SetPageSize(v string) *DescribeImagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImagesRequest) SetSnapshotId(v string) *DescribeImagesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImagesRequest) SetStatus(v string) *DescribeImagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeImagesRequest) Validate() error {
	return dara.Validate(s)
}
