// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterLabelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetOpaClusterLabelListRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *GetOpaClusterLabelListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetOpaClusterLabelListRequest
	GetPageSize() *int32
	SetTagName(v string) *GetOpaClusterLabelListRequest
	GetTagName() *string
}

type GetOpaClusterLabelListRequest struct {
	// The ID of the cluster to which the container belongs.
	//
	// >  You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of clusters.
	//
	// example:
	//
	// c556c8133b5ad4378b7fc533ddbda****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 200. If you leave the PageSize parameter empty, 200 entries are returned by default. Maximum value: 200.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the tag.
	//
	// example:
	//
	// test
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetOpaClusterLabelListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterLabelListRequest) GoString() string {
	return s.String()
}

func (s *GetOpaClusterLabelListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetOpaClusterLabelListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetOpaClusterLabelListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpaClusterLabelListRequest) GetTagName() *string {
	return s.TagName
}

func (s *GetOpaClusterLabelListRequest) SetClusterId(v string) *GetOpaClusterLabelListRequest {
	s.ClusterId = &v
	return s
}

func (s *GetOpaClusterLabelListRequest) SetCurrentPage(v int32) *GetOpaClusterLabelListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetOpaClusterLabelListRequest) SetPageSize(v int32) *GetOpaClusterLabelListRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpaClusterLabelListRequest) SetTagName(v string) *GetOpaClusterLabelListRequest {
	s.TagName = &v
	return s
}

func (s *GetOpaClusterLabelListRequest) Validate() error {
	return dara.Validate(s)
}
