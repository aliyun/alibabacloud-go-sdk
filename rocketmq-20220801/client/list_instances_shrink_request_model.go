// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListInstancesShrinkRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesShrinkRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesShrinkRequest
	GetResourceGroupId() *string
	SetSeriesCodesShrink(v string) *ListInstancesShrinkRequest
	GetSeriesCodesShrink() *string
	SetStorageSecretKey(v string) *ListInstancesShrinkRequest
	GetStorageSecretKey() *string
	SetTags(v string) *ListInstancesShrinkRequest
	GetTags() *string
}

type ListInstancesShrinkRequest struct {
	// The filter condition that is used to query instances. If you do not configure this parameter, all instances are queried.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 100000000.
	//
	// If you set this parameter to a value smaller than 1, the system uses 1 as the value. If you set this parameter to a value greater than 100000000, the system uses 100000000 as the value.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// Value values: 10 to 200.
	//
	// If you set this parameter to a value smaller than 10, the system uses 10 as the value. If you set this parameter to a value greater than 200, the system uses 200 as the value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance.
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	SeriesCodesShrink *string `json:"seriesCodes,omitempty" xml:"seriesCodes,omitempty"`
	// The storage encryption key.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// The tags that are used to filter instances.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesShrinkRequest) GetSeriesCodesShrink() *string {
	return s.SeriesCodesShrink
}

func (s *ListInstancesShrinkRequest) GetStorageSecretKey() *string {
	return s.StorageSecretKey
}

func (s *ListInstancesShrinkRequest) GetTags() *string {
	return s.Tags
}

func (s *ListInstancesShrinkRequest) SetFilter(v string) *ListInstancesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetSeriesCodesShrink(v string) *ListInstancesShrinkRequest {
	s.SeriesCodesShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStorageSecretKey(v string) *ListInstancesShrinkRequest {
	s.StorageSecretKey = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTags(v string) *ListInstancesShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ListInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
