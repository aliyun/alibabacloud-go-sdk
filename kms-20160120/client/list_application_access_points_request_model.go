// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListApplicationAccessPointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListApplicationAccessPointsRequest
	GetPageSize() *int32
}

type ListApplicationAccessPointsRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationAccessPointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListApplicationAccessPointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApplicationAccessPointsRequest) SetPageNumber(v int32) *ListApplicationAccessPointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationAccessPointsRequest) SetPageSize(v int32) *ListApplicationAccessPointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
