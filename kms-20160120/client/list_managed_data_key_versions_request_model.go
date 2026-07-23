// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedDataKeyVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataKeyName(v string) *ListManagedDataKeyVersionsRequest
	GetDataKeyName() *string
	SetPageNumber(v int32) *ListManagedDataKeyVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListManagedDataKeyVersionsRequest
	GetPageSize() *int32
}

type ListManagedDataKeyVersionsRequest struct {
	// The name of the managed data key (DK) to query. This parameter is required.
	//
	// example:
	//
	// example-data-key
	DataKeyName *string `json:"DataKeyName,omitempty" xml:"DataKeyName,omitempty"`
	// The page number. The value must be an integer greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListManagedDataKeyVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManagedDataKeyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListManagedDataKeyVersionsRequest) GetDataKeyName() *string {
	return s.DataKeyName
}

func (s *ListManagedDataKeyVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListManagedDataKeyVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListManagedDataKeyVersionsRequest) SetDataKeyName(v string) *ListManagedDataKeyVersionsRequest {
	s.DataKeyName = &v
	return s
}

func (s *ListManagedDataKeyVersionsRequest) SetPageNumber(v int32) *ListManagedDataKeyVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListManagedDataKeyVersionsRequest) SetPageSize(v int32) *ListManagedDataKeyVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListManagedDataKeyVersionsRequest) Validate() error {
	return dara.Validate(s)
}
