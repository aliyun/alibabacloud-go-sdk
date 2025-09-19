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
	// The number of the page to return. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.\\
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
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
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the ID.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The beginning of the time range to query. You can query all AccessKey pair leaks that are detected later than this time point. The value of this parameter is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1614155361489
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// Specifies whether an AccessKey pair leak is handled. Valid values:
	//
	// 	- **pending**: unhandled
	//
	// 	- **dealed**: handled
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
