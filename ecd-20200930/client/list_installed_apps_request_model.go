// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *ListInstalledAppsRequest
	GetDesktopId() *string
	SetPageNumber(v int32) *ListInstalledAppsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstalledAppsRequest
	GetPageSize() *int32
}

type ListInstalledAppsRequest struct {
	// The cloud computer ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-39clsqyxr****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The page number.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInstalledAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledAppsRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledAppsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *ListInstalledAppsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstalledAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstalledAppsRequest) SetDesktopId(v string) *ListInstalledAppsRequest {
	s.DesktopId = &v
	return s
}

func (s *ListInstalledAppsRequest) SetPageNumber(v int32) *ListInstalledAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstalledAppsRequest) SetPageSize(v int32) *ListInstalledAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstalledAppsRequest) Validate() error {
	return dara.Validate(s)
}
