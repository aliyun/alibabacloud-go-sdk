// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstalledSoftwaresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListInstalledSoftwaresRequest
	GetClusterId() *string
	SetPageNumber(v string) *ListInstalledSoftwaresRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListInstalledSoftwaresRequest
	GetPageSize() *string
}

type ListInstalledSoftwaresRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s ListInstalledSoftwaresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstalledSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwaresRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstalledSoftwaresRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListInstalledSoftwaresRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListInstalledSoftwaresRequest) SetClusterId(v string) *ListInstalledSoftwaresRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstalledSoftwaresRequest) SetPageNumber(v string) *ListInstalledSoftwaresRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstalledSoftwaresRequest) SetPageSize(v string) *ListInstalledSoftwaresRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstalledSoftwaresRequest) Validate() error {
	return dara.Validate(s)
}
