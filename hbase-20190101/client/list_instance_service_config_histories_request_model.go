// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstanceServiceConfigHistoriesRequest
	GetClusterId() *string
	SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceServiceConfigHistoriesRequest
	GetPageSize() *int32
}

type ListInstanceServiceConfigHistoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstanceServiceConfigHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigHistoriesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigHistoriesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstanceServiceConfigHistoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceServiceConfigHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetClusterId(v string) *ListInstanceServiceConfigHistoriesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageNumber(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) SetPageSize(v int32) *ListInstanceServiceConfigHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceServiceConfigHistoriesRequest) Validate() error {
	return dara.Validate(s)
}
