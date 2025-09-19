// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListOssScanConfigRequest
	GetCurrentPage() *int32
	SetName(v string) *ListOssScanConfigRequest
	GetName() *string
	SetPageSize(v int32) *ListOssScanConfigRequest
	GetPageSize() *int32
}

type ListOssScanConfigRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The policy name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOssScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOssScanConfigRequest) GoString() string {
	return s.String()
}

func (s *ListOssScanConfigRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOssScanConfigRequest) GetName() *string {
	return s.Name
}

func (s *ListOssScanConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOssScanConfigRequest) SetCurrentPage(v int32) *ListOssScanConfigRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOssScanConfigRequest) SetName(v string) *ListOssScanConfigRequest {
	s.Name = &v
	return s
}

func (s *ListOssScanConfigRequest) SetPageSize(v int32) *ListOssScanConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
