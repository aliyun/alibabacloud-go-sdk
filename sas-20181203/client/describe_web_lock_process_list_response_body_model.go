// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockProcessListResponseBody
	GetCurrentPage() *int32
	SetList(v []*DescribeWebLockProcessListResponseBodyList) *DescribeWebLockProcessListResponseBody
	GetList() []*DescribeWebLockProcessListResponseBodyList
	SetPageSize(v int32) *DescribeWebLockProcessListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebLockProcessListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockProcessListResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockProcessListResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of details about the process.
	List []*DescribeWebLockProcessListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BF880
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of processes.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockProcessListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockProcessListResponseBody) GetList() []*DescribeWebLockProcessListResponseBodyList {
	return s.List
}

func (s *DescribeWebLockProcessListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockProcessListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockProcessListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockProcessListResponseBody) SetCurrentPage(v int32) *DescribeWebLockProcessListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBody) SetList(v []*DescribeWebLockProcessListResponseBodyList) *DescribeWebLockProcessListResponseBody {
	s.List = v
	return s
}

func (s *DescribeWebLockProcessListResponseBody) SetPageSize(v int32) *DescribeWebLockProcessListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBody) SetRequestId(v string) *DescribeWebLockProcessListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBody) SetTotalCount(v int32) *DescribeWebLockProcessListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebLockProcessListResponseBodyList struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 33
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// test_ecs
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// cron
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The path to the process.
	//
	// example:
	//
	// /usr/sbin/cron
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// Indicates whether the process is added to the process whitelist. Valid values:
	//
	// 	- **1**: The process is added to the process whitelist.
	//
	// 	- **0**: The process is not added to the process whitelist.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockProcessListResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessListResponseBodyList) GetCount() *string {
	return s.Count
}

func (s *DescribeWebLockProcessListResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeWebLockProcessListResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeWebLockProcessListResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeWebLockProcessListResponseBodyList) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWebLockProcessListResponseBodyList) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *DescribeWebLockProcessListResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWebLockProcessListResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockProcessListResponseBodyList) SetCount(v string) *DescribeWebLockProcessListResponseBodyList {
	s.Count = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetInstanceName(v string) *DescribeWebLockProcessListResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetInternetIp(v string) *DescribeWebLockProcessListResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetIntranetIp(v string) *DescribeWebLockProcessListResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetProcessName(v string) *DescribeWebLockProcessListResponseBodyList {
	s.ProcessName = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetProcessPath(v string) *DescribeWebLockProcessListResponseBodyList {
	s.ProcessPath = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetStatus(v int32) *DescribeWebLockProcessListResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) SetUuid(v string) *DescribeWebLockProcessListResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockProcessListResponseBodyList) Validate() error {
	return dara.Validate(s)
}
