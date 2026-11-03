// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeJobHistoryRequest
	GetAppId() *string
	SetCurrentPage(v int64) *DescribeJobHistoryRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *DescribeJobHistoryRequest
	GetPageSize() *int64
	SetState(v string) *DescribeJobHistoryRequest
	GetState() *string
}

type DescribeJobHistoryRequest struct {
	// The ID of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1a7a07-abcb-4652-a1d3-2d57f415****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page in the paging Settings for a paged query. Valid values: 0 to 10000.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The running status of the job. Valid values:
	//
	// - **0**: Not run.
	//
	// - **1**: Succeeded.
	//
	// - **2**: Failed.
	//
	// - **3**: Running.
	//
	// example:
	//
	// 1
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeJobHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobHistoryRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeJobHistoryRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeJobHistoryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeJobHistoryRequest) GetState() *string {
	return s.State
}

func (s *DescribeJobHistoryRequest) SetAppId(v string) *DescribeJobHistoryRequest {
	s.AppId = &v
	return s
}

func (s *DescribeJobHistoryRequest) SetCurrentPage(v int64) *DescribeJobHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeJobHistoryRequest) SetPageSize(v int64) *DescribeJobHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeJobHistoryRequest) SetState(v string) *DescribeJobHistoryRequest {
	s.State = &v
	return s
}

func (s *DescribeJobHistoryRequest) Validate() error {
	return dara.Validate(s)
}
