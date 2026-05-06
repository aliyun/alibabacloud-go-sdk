// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeylessServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListKeylessServersRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListKeylessServersRequest
	GetPageSize() *int64
	SetSiteId(v int64) *ListKeylessServersRequest
	GetSiteId() *int64
}

type ListKeylessServersRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListKeylessServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKeylessServersRequest) GoString() string {
	return s.String()
}

func (s *ListKeylessServersRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListKeylessServersRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListKeylessServersRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListKeylessServersRequest) SetPageNumber(v int64) *ListKeylessServersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeylessServersRequest) SetPageSize(v int64) *ListKeylessServersRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeylessServersRequest) SetSiteId(v int64) *ListKeylessServersRequest {
	s.SiteId = &v
	return s
}

func (s *ListKeylessServersRequest) Validate() error {
	return dara.Validate(s)
}
