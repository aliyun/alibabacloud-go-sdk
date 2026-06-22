// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccesskeyLeakListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeAccesskeyLeakListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeAccesskeyLeakListRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeAccesskeyLeakListRequest
	GetQuery() *string
	SetResourceDirectoryAccountId(v int64) *DescribeAccesskeyLeakListRequest
	GetResourceDirectoryAccountId() *int64
	SetStartTs(v int64) *DescribeAccesskeyLeakListRequest
	GetStartTs() *int64
	SetStatus(v string) *DescribeAccesskeyLeakListRequest
	GetStatus() *string
}

type DescribeAccesskeyLeakListRequest struct {
	// The page number of the current page in a paged query. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page for a paged query. The maximum value of PageSize is 100. The default number of entries per page is 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The AccessKey ID that you want to query. Only exact match is supported.
	//
	// example:
	//
	// yourAccessKeyID
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the member account in the resource directory (Alibaba Cloud account).
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The discovery time of the leaked information that you want to query. All AccessKey leak information discovered after this point in time is returned. This parameter is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614155361489
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The processing status of the leaked AccessKey information that you want to query. Valid values:
	//
	// - **pending**: unprocessed
	//
	// - **dealed**: processed.
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAccesskeyLeakListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccesskeyLeakListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAccesskeyLeakListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccesskeyLeakListRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeAccesskeyLeakListRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeAccesskeyLeakListRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeAccesskeyLeakListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccesskeyLeakListRequest) SetCurrentPage(v int32) *DescribeAccesskeyLeakListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetPageSize(v int32) *DescribeAccesskeyLeakListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetQuery(v string) *DescribeAccesskeyLeakListRequest {
	s.Query = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetResourceDirectoryAccountId(v int64) *DescribeAccesskeyLeakListRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetStartTs(v int64) *DescribeAccesskeyLeakListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) SetStatus(v string) *DescribeAccesskeyLeakListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAccesskeyLeakListRequest) Validate() error {
	return dara.Validate(s)
}
