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
	SetSource(v string) *ListOssScanConfigRequest
	GetSource() *string
}

type ListOssScanConfigRequest struct {
	// The page number of the current page in a paged query.
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
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The business source. Valid values:
	//
	// - **OSS**: OSS
	//
	// - **NAS**: NAS
	//
	// example:
	//
	// OSS
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

func (s *ListOssScanConfigRequest) GetSource() *string {
	return s.Source
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

func (s *ListOssScanConfigRequest) SetSource(v string) *ListOssScanConfigRequest {
	s.Source = &v
	return s
}

func (s *ListOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
