// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpWhitelist(v string) *ListInstanceIpWhitelistRequest
	GetIpWhitelist() *string
	SetPageNumber(v int32) *ListInstanceIpWhitelistRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceIpWhitelistRequest
	GetPageSize() *int32
}

type ListInstanceIpWhitelistRequest struct {
	// IP whitelist.
	//
	// example:
	//
	// 0.0.0.0/0
	IpWhitelist *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ListInstanceIpWhitelistRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceIpWhitelistRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceIpWhitelistRequest) SetIpWhitelist(v string) *ListInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ListInstanceIpWhitelistRequest) SetPageNumber(v int32) *ListInstanceIpWhitelistRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceIpWhitelistRequest) SetPageSize(v int32) *ListInstanceIpWhitelistRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
