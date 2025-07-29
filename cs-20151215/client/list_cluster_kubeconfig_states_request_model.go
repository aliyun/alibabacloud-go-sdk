// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterKubeconfigStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListClusterKubeconfigStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClusterKubeconfigStatesRequest
	GetPageSize() *int32
}

type ListClusterKubeconfigStatesRequest struct {
	// The page number.
	//
	// 	- Valid values: ≥ 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: 10 to 50.
	//
	// 	- Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListClusterKubeconfigStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterKubeconfigStatesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterKubeconfigStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterKubeconfigStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterKubeconfigStatesRequest) SetPageNumber(v int32) *ListClusterKubeconfigStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterKubeconfigStatesRequest) SetPageSize(v int32) *ListClusterKubeconfigStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterKubeconfigStatesRequest) Validate() error {
	return dara.Validate(s)
}
