// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListInstancesRequest
	GetFilter() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetSeriesCodes(v []*string) *ListInstancesRequest
	GetSeriesCodes() []*string
	SetStorageSecretKey(v string) *ListInstancesRequest
	GetStorageSecretKey() *string
	SetTags(v string) *ListInstancesRequest
	GetTags() *string
}

type ListInstancesRequest struct {
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
	SeriesCodes []*string `json:"seriesCodes,omitempty" xml:"seriesCodes,omitempty" type:"Repeated"`
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

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetSeriesCodes() []*string {
	return s.SeriesCodes
}

func (s *ListInstancesRequest) GetStorageSecretKey() *string {
	return s.StorageSecretKey
}

func (s *ListInstancesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListInstancesRequest) SetFilter(v string) *ListInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetSeriesCodes(v []*string) *ListInstancesRequest {
	s.SeriesCodes = v
	return s
}

func (s *ListInstancesRequest) SetStorageSecretKey(v string) *ListInstancesRequest {
	s.StorageSecretKey = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v string) *ListInstancesRequest {
	s.Tags = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
