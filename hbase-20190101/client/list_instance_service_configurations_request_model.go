// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceServiceConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstanceServiceConfigurationsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *ListInstanceServiceConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceServiceConfigurationsRequest
	GetPageSize() *int32
}

type ListInstanceServiceConfigurationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-t4naqsay5gn****
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

func (s ListInstanceServiceConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceServiceConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceServiceConfigurationsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstanceServiceConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceServiceConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceServiceConfigurationsRequest) SetClusterId(v string) *ListInstanceServiceConfigurationsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageNumber(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) SetPageSize(v int32) *ListInstanceServiceConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceServiceConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
