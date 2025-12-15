// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonIds(v []*string) *ListAddonsRequest
	GetAddonIds() []*string
	SetClusterId(v string) *ListAddonsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *ListAddonsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAddonsRequest
	GetPageSize() *int32
}

type ListAddonsRequest struct {
	// The addon IDs.
	AddonIds []*string `json:"AddonIds,omitempty" xml:"AddonIds,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number of the page to return. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsRequest) GetAddonIds() []*string {
	return s.AddonIds
}

func (s *ListAddonsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAddonsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAddonsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAddonsRequest) SetAddonIds(v []*string) *ListAddonsRequest {
	s.AddonIds = v
	return s
}

func (s *ListAddonsRequest) SetClusterId(v string) *ListAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAddonsRequest) SetPageNumber(v int32) *ListAddonsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsRequest) SetPageSize(v int32) *ListAddonsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAddonsRequest) Validate() error {
	return dara.Validate(s)
}
