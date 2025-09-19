// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterImageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetOpaClusterImageListRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *GetOpaClusterImageListRequest
	GetCurrentPage() *int32
	SetImageName(v string) *GetOpaClusterImageListRequest
	GetImageName() *string
	SetPageSize(v int32) *GetOpaClusterImageListRequest
	GetPageSize() *int32
}

type GetOpaClusterImageListRequest struct {
	// The ID of the cluster to which the container belongs.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// c4af4fdf38a98496a9b63c2be5dae****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// testImage
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetOpaClusterImageListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterImageListRequest) GoString() string {
	return s.String()
}

func (s *GetOpaClusterImageListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaClusterImageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOpaClusterImageListRequest) GetImageName() *string {
	return s.ImageName
}

func (s *GetOpaClusterImageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpaClusterImageListRequest) SetClusterId(v string) *GetOpaClusterImageListRequest {
	s.ClusterId = &v
	return s
}

func (s *GetOpaClusterImageListRequest) SetCurrentPage(v int32) *GetOpaClusterImageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOpaClusterImageListRequest) SetImageName(v string) *GetOpaClusterImageListRequest {
	s.ImageName = &v
	return s
}

func (s *GetOpaClusterImageListRequest) SetPageSize(v int32) *GetOpaClusterImageListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpaClusterImageListRequest) Validate() error {
	return dara.Validate(s)
}
