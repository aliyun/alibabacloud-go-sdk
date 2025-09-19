// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeVulListPageRequest
	GetCurrentPage() *int32
	SetCveId(v string) *DescribeVulListPageRequest
	GetCveId() *string
	SetPageSize(v int32) *DescribeVulListPageRequest
	GetPageSize() *int32
	SetRaspDefend(v int32) *DescribeVulListPageRequest
	GetRaspDefend() *int32
	SetVulNameLike(v string) *DescribeVulListPageRequest
	GetVulNameLike() *string
	SetVulType(v string) *DescribeVulListPageRequest
	GetVulType() *string
}

type DescribeVulListPageRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the vulnerability.
	//
	// example:
	//
	// CVE-2022-44702
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the application protection feature is supported. Valid values:
	//
	// - **0**: no.
	//
	// - **1**: yes.
	//
	// example:
	//
	// 0
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// RCE vulnerability
	VulNameLike *string `json:"VulNameLike,omitempty" xml:"VulNameLike,omitempty"`
	// The type of the vulnerabilities. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **app**: Application vulnerability that is detected by using web scanner.
	//
	// example:
	//
	// cve
	VulType *string `json:"VulType,omitempty" xml:"VulType,omitempty"`
}

func (s DescribeVulListPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulListPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVulListPageRequest) GetCveId() *string {
	return s.CveId
}

func (s *DescribeVulListPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVulListPageRequest) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeVulListPageRequest) GetVulNameLike() *string {
	return s.VulNameLike
}

func (s *DescribeVulListPageRequest) GetVulType() *string {
	return s.VulType
}

func (s *DescribeVulListPageRequest) SetCurrentPage(v int32) *DescribeVulListPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulListPageRequest) SetCveId(v string) *DescribeVulListPageRequest {
	s.CveId = &v
	return s
}

func (s *DescribeVulListPageRequest) SetPageSize(v int32) *DescribeVulListPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulListPageRequest) SetRaspDefend(v int32) *DescribeVulListPageRequest {
	s.RaspDefend = &v
	return s
}

func (s *DescribeVulListPageRequest) SetVulNameLike(v string) *DescribeVulListPageRequest {
	s.VulNameLike = &v
	return s
}

func (s *DescribeVulListPageRequest) SetVulType(v string) *DescribeVulListPageRequest {
	s.VulType = &v
	return s
}

func (s *DescribeVulListPageRequest) Validate() error {
	return dara.Validate(s)
}
