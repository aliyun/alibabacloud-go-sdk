// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeContainerTagsRequest
	GetAppName() *string
	SetClusterId(v string) *DescribeContainerTagsRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *DescribeContainerTagsRequest
	GetCurrentPage() *int32
	SetFieldName(v string) *DescribeContainerTagsRequest
	GetFieldName() *string
	SetFieldValue(v string) *DescribeContainerTagsRequest
	GetFieldValue() *string
	SetNamespace(v string) *DescribeContainerTagsRequest
	GetNamespace() *string
	SetPageSize(v int32) *DescribeContainerTagsRequest
	GetPageSize() *int32
}

type DescribeContainerTagsRequest struct {
	// The name of the application.
	//
	// example:
	//
	// node-exporter
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the cluster to which the container belongs.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// c22143730ab6e40b09ec7c1c51d4d****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the attribute that is used for the query. Valid values:
	//
	// 	- **namespace**: the namespace
	//
	// 	- **appName**: the application name
	//
	// 	- **image**: the image
	//
	// 	- **tag**: the tag
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the attribute that is used for the query.
	//
	// example:
	//
	// demo4
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-name-01
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of entries to return on each page. Default value: 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeContainerTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerTagsRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeContainerTagsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerTagsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeContainerTagsRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeContainerTagsRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeContainerTagsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeContainerTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerTagsRequest) SetAppName(v string) *DescribeContainerTagsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetClusterId(v string) *DescribeContainerTagsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetCurrentPage(v int32) *DescribeContainerTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetFieldName(v string) *DescribeContainerTagsRequest {
	s.FieldName = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetFieldValue(v string) *DescribeContainerTagsRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetNamespace(v string) *DescribeContainerTagsRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerTagsRequest) SetPageSize(v int32) *DescribeContainerTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerTagsRequest) Validate() error {
	return dara.Validate(s)
}
