// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookReleasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePlaybookReleasesRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribePlaybookReleasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePlaybookReleasesRequest
	GetPageSize() *int32
	SetPlaybookUuid(v string) *DescribePlaybookReleasesRequest
	GetPlaybookUuid() *string
}

type DescribePlaybookReleasesRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. If you do not specify the PageSize parameter, 10 entries are returned by default.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookReleasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookReleasesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookReleasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePlaybookReleasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePlaybookReleasesRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookReleasesRequest) SetLang(v string) *DescribePlaybookReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageNumber(v int32) *DescribePlaybookReleasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPageSize(v int32) *DescribePlaybookReleasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) SetPlaybookUuid(v string) *DescribePlaybookReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookReleasesRequest) Validate() error {
	return dara.Validate(s)
}
