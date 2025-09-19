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
	// Sets the page number from which to start displaying the query results. The default value is 1, indicating that the display starts from the first page.
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
	// Specifies the maximum number of data entries to display per page in a paginated query. The default number of data entries per page is 20, and if the PageSize parameter is empty, it will default to returning 20 data entries.
	//
	// > It is recommended that the PageSize value is not empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Probe status. Values:
	//
	// - **0**: Enabled
	//
	// - **1**: Disabled
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
