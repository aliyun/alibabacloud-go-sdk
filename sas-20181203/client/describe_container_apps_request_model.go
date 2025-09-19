// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerAppsRequest
	GetClusterId() *string
	SetCurrentPage(v int32) *DescribeContainerAppsRequest
	GetCurrentPage() *int32
	SetFieldValue(v string) *DescribeContainerAppsRequest
	GetFieldValue() *string
	SetPageSize(v int32) *DescribeContainerAppsRequest
	GetPageSize() *int32
}

type DescribeContainerAppsRequest struct {
	// The ID of the container cluster.
	//
	// example:
	//
	// ca0a686115432429ca26cf780f5e9fff5
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The application value that you want to query. Fuzzy match is supported.
	//
	// example:
	//
	// cas-adad-qeqwe
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeContainerAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerAppsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeContainerAppsRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeContainerAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeContainerAppsRequest) SetClusterId(v string) *DescribeContainerAppsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerAppsRequest) SetCurrentPage(v int32) *DescribeContainerAppsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeContainerAppsRequest) SetFieldValue(v string) *DescribeContainerAppsRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeContainerAppsRequest) SetPageSize(v int32) *DescribeContainerAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeContainerAppsRequest) Validate() error {
	return dara.Validate(s)
}
