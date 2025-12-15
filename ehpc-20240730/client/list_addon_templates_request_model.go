// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonNames(v []*string) *ListAddonTemplatesRequest
	GetAddonNames() []*string
	SetClusterCategory(v string) *ListAddonTemplatesRequest
	GetClusterCategory() *string
	SetPageNumber(v int64) *ListAddonTemplatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAddonTemplatesRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListAddonTemplatesRequest
	GetRegionId() *string
}

type ListAddonTemplatesRequest struct {
	// The addon names.
	AddonNames []*string `json:"AddonNames,omitempty" xml:"AddonNames,omitempty" type:"Repeated"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The page number of the page to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAddonTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesRequest) GetAddonNames() []*string {
	return s.AddonNames
}

func (s *ListAddonTemplatesRequest) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *ListAddonTemplatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAddonTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAddonTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAddonTemplatesRequest) SetAddonNames(v []*string) *ListAddonTemplatesRequest {
	s.AddonNames = v
	return s
}

func (s *ListAddonTemplatesRequest) SetClusterCategory(v string) *ListAddonTemplatesRequest {
	s.ClusterCategory = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetPageNumber(v int64) *ListAddonTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetPageSize(v int64) *ListAddonTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAddonTemplatesRequest) SetRegionId(v string) *ListAddonTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAddonTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
