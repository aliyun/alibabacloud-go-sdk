// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwaresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListSoftwaresRequest
	GetCategory() *string
	SetClusterId(v string) *ListSoftwaresRequest
	GetClusterId() *string
	SetName(v string) *ListSoftwaresRequest
	GetName() *string
	SetOsInfos(v []*ListSoftwaresRequestOsInfos) *ListSoftwaresRequest
	GetOsInfos() []*ListSoftwaresRequestOsInfos
	SetPageNumber(v string) *ListSoftwaresRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListSoftwaresRequest
	GetPageSize() *string
}

type ListSoftwaresRequest struct {
	// The application category.
	//
	// example:
	//
	// NWP
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system (OS) information.
	OsInfos []*ListSoftwaresRequestOsInfos `json:"OsInfos,omitempty" xml:"OsInfos,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) GetCategory() *string {
	return s.Category
}

func (s *ListSoftwaresRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSoftwaresRequest) GetName() *string {
	return s.Name
}

func (s *ListSoftwaresRequest) GetOsInfos() []*ListSoftwaresRequestOsInfos {
	return s.OsInfos
}

func (s *ListSoftwaresRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListSoftwaresRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSoftwaresRequest) SetCategory(v string) *ListSoftwaresRequest {
	s.Category = &v
	return s
}

func (s *ListSoftwaresRequest) SetClusterId(v string) *ListSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *ListSoftwaresRequest) SetName(v string) *ListSoftwaresRequest {
	s.Name = &v
	return s
}

func (s *ListSoftwaresRequest) SetOsInfos(v []*ListSoftwaresRequestOsInfos) *ListSoftwaresRequest {
	s.OsInfos = v
	return s
}

func (s *ListSoftwaresRequest) SetPageNumber(v string) *ListSoftwaresRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSoftwaresRequest) SetPageSize(v string) *ListSoftwaresRequest {
	s.PageSize = &v
	return s
}

func (s *ListSoftwaresRequest) Validate() error {
	if s.OsInfos != nil {
		for _, item := range s.OsInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwaresRequestOsInfos struct {
	// The OS architecture. Valid values:
	//
	// 	- x86_64
	//
	// 	- arm64
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image tag.
	//
	// example:
	//
	// CentOS_7.9_64
	OsTag *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresRequestOsInfos) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwaresRequestOsInfos) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequestOsInfos) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListSoftwaresRequestOsInfos) GetOsTag() *string {
	return s.OsTag
}

func (s *ListSoftwaresRequestOsInfos) SetArchitecture(v string) *ListSoftwaresRequestOsInfos {
	s.Architecture = &v
	return s
}

func (s *ListSoftwaresRequestOsInfos) SetOsTag(v string) *ListSoftwaresRequestOsInfos {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresRequestOsInfos) Validate() error {
	return dara.Validate(s)
}
