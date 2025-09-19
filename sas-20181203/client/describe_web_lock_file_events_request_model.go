// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockFileEventsRequest
	GetCurrentPage() *int32
	SetDealed(v string) *DescribeWebLockFileEventsRequest
	GetDealed() *string
	SetPageSize(v int32) *DescribeWebLockFileEventsRequest
	GetPageSize() *int32
	SetProcessName(v string) *DescribeWebLockFileEventsRequest
	GetProcessName() *string
	SetRemark(v string) *DescribeWebLockFileEventsRequest
	GetRemark() *string
	SetTsBegin(v int64) *DescribeWebLockFileEventsRequest
	GetTsBegin() *int64
	SetTsEnd(v int64) *DescribeWebLockFileEventsRequest
	GetTsEnd() *int64
}

type DescribeWebLockFileEventsRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether the event on web tamper proofing is handled. Valid values:
	//
	// 	- **n**: The event on web tamper proofing is handled.
	//
	// 	- **y**: The event on web tamper proofing is not handled.
	//
	// example:
	//
	// n
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// sys_create
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The name of the asset.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the names of assets.
	//
	// example:
	//
	// test-ecs
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1660649981419
	TsBegin *int64 `json:"TsBegin,omitempty" xml:"TsBegin,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1660649981419
	TsEnd *int64 `json:"TsEnd,omitempty" xml:"TsEnd,omitempty"`
}

func (s DescribeWebLockFileEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockFileEventsRequest) GetDealed() *string {
	return s.Dealed
}

func (s *DescribeWebLockFileEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockFileEventsRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWebLockFileEventsRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeWebLockFileEventsRequest) GetTsBegin() *int64 {
	return s.TsBegin
}

func (s *DescribeWebLockFileEventsRequest) GetTsEnd() *int64 {
	return s.TsEnd
}

func (s *DescribeWebLockFileEventsRequest) SetCurrentPage(v int32) *DescribeWebLockFileEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetDealed(v string) *DescribeWebLockFileEventsRequest {
	s.Dealed = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetPageSize(v int32) *DescribeWebLockFileEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetProcessName(v string) *DescribeWebLockFileEventsRequest {
	s.ProcessName = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetRemark(v string) *DescribeWebLockFileEventsRequest {
	s.Remark = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetTsBegin(v int64) *DescribeWebLockFileEventsRequest {
	s.TsBegin = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) SetTsEnd(v int64) *DescribeWebLockFileEventsRequest {
	s.TsEnd = &v
	return s
}

func (s *DescribeWebLockFileEventsRequest) Validate() error {
	return dara.Validate(s)
}
