// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonIdsShrink(v string) *ListAddonsShrinkRequest
	GetAddonIdsShrink() *string
	SetClusterId(v string) *ListAddonsShrinkRequest
	GetClusterId() *string
	SetPageNumber(v int32) *ListAddonsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAddonsShrinkRequest
	GetPageSize() *int32
}

type ListAddonsShrinkRequest struct {
	// The addon IDs.
	AddonIdsShrink *string `json:"AddonIds,omitempty" xml:"AddonIds,omitempty"`
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

func (s ListAddonsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsShrinkRequest) GetAddonIdsShrink() *string {
	return s.AddonIdsShrink
}

func (s *ListAddonsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAddonsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAddonsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAddonsShrinkRequest) SetAddonIdsShrink(v string) *ListAddonsShrinkRequest {
	s.AddonIdsShrink = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetClusterId(v string) *ListAddonsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetPageNumber(v int32) *ListAddonsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAddonsShrinkRequest) SetPageSize(v int32) *ListAddonsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAddonsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
