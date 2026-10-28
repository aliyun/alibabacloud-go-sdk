// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcProbeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeIdcProbeListRequest
	GetCurrentPage() *int32
	SetIdcName(v string) *DescribeIdcProbeListRequest
	GetIdcName() *string
	SetPageSize(v int32) *DescribeIdcProbeListRequest
	GetPageSize() *int32
	SetStatus(v int32) *DescribeIdcProbeListRequest
	GetStatus() *int32
}

type DescribeIdcProbeListRequest struct {
	// Specifies the page number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the IDC.
	//
	// example:
	//
	// 3K IDC
	IdcName *string `json:"IdcName,omitempty" xml:"IdcName,omitempty"`
	// Specifies the maximum number of entries to return on each page for a paging query. The default number of entries per page is 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Specify a value for the PageSize parameter.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the probe. Valid values:
	//
	// - **0**: Enabled.
	//
	// - **1**: Disabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIdcProbeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcProbeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdcProbeListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeIdcProbeListRequest) GetIdcName() *string {
	return s.IdcName
}

func (s *DescribeIdcProbeListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIdcProbeListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeIdcProbeListRequest) SetCurrentPage(v int32) *DescribeIdcProbeListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeIdcProbeListRequest) SetIdcName(v string) *DescribeIdcProbeListRequest {
	s.IdcName = &v
	return s
}

func (s *DescribeIdcProbeListRequest) SetPageSize(v int32) *DescribeIdcProbeListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIdcProbeListRequest) SetStatus(v int32) *DescribeIdcProbeListRequest {
	s.Status = &v
	return s
}

func (s *DescribeIdcProbeListRequest) Validate() error {
	return dara.Validate(s)
}
