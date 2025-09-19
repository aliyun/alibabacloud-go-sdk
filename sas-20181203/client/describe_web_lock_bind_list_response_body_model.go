// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockBindListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBindList(v []*DescribeWebLockBindListResponseBodyBindList) *DescribeWebLockBindListResponseBody
	GetBindList() []*DescribeWebLockBindListResponseBodyBindList
	SetCurrentPage(v int32) *DescribeWebLockBindListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebLockBindListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebLockBindListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockBindListResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockBindListResponseBody struct {
	// The information about the servers that have web tamper proofing enabled.
	BindList []*DescribeWebLockBindListResponseBodyBindList `json:"BindList,omitempty" xml:"BindList,omitempty" type:"Repeated"`
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of servers that have web tamper proofing enabled.
	//
	// example:
	//
	// 11409
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockBindListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockBindListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponseBody) GetBindList() []*DescribeWebLockBindListResponseBodyBindList {
	return s.BindList
}

func (s *DescribeWebLockBindListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockBindListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockBindListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockBindListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockBindListResponseBody) SetBindList(v []*DescribeWebLockBindListResponseBodyBindList) *DescribeWebLockBindListResponseBody {
	s.BindList = v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetCurrentPage(v int32) *DescribeWebLockBindListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetPageSize(v int32) *DescribeWebLockBindListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetRequestId(v string) *DescribeWebLockBindListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) SetTotalCount(v int32) *DescribeWebLockBindListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebLockBindListResponseBodyBindList struct {
	// The number of alerts.
	//
	// example:
	//
	// 2
	AuditCount *string `json:"AuditCount,omitempty" xml:"AuditCount,omitempty"`
	// The number of blocked tampering events.
	//
	// example:
	//
	// 10
	BlockCount *string `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// The number of protected directories.
	//
	// example:
	//
	// 5
	DirCount *string `json:"DirCount,omitempty" xml:"DirCount,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// testName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 54.169.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The operating system that the server runs.
	//
	// example:
	//
	// Linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The percentage of the starting progress of web tamper proofing. Valid values: 0 to 100.
	//
	// example:
	//
	// 99
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The error code for web tamper proofing. Valid values:
	//
	// 	- **2001**: The Security Center agent is offline.
	//
	// 	- **9999**: The connection timed out.
	//
	// example:
	//
	// 2001
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The exception details of web tamper proofing. Valid values:
	//
	// 	- **client offline**: The Security Center agent is offline.
	//
	// 	- **timeout**: The connection timed out.
	//
	// example:
	//
	// client offline
	ServiceDetail *string `json:"ServiceDetail,omitempty" xml:"ServiceDetail,omitempty"`
	// The status of web tamper proofing on the server. Valid values:
	//
	// 	- **stop**: Web tamper proofing is disabled.
	//
	// 	- **initializing**: Web tamper proofing is being enabled.
	//
	// 	- **exception**: Web tamper proofing is not running as expected.
	//
	// 	- **running**: Web tamper proofing is running.
	//
	// 	- **closing**: Web tamper proofing is being disabled.
	//
	// example:
	//
	// stop
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The protection status of the server. Valid values:
	//
	// 	- **on**: The server is protected.
	//
	// 	- **off**: The server is not protected.
	//
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// inet-12345****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockBindListResponseBodyBindList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockBindListResponseBodyBindList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetAuditCount() *string {
	return s.AuditCount
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetBlockCount() *string {
	return s.BlockCount
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetDirCount() *string {
	return s.DirCount
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetOs() *string {
	return s.Os
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetPercent() *int32 {
	return s.Percent
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetServiceDetail() *string {
	return s.ServiceDetail
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetStatus() *string {
	return s.Status
}

func (s *DescribeWebLockBindListResponseBodyBindList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetAuditCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.AuditCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetBlockCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.BlockCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetDirCount(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.DirCount = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetInstanceName(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.InstanceName = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetInternetIp(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.InternetIp = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetIntranetIp(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetOs(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Os = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetPercent(v int32) *DescribeWebLockBindListResponseBodyBindList {
	s.Percent = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceCode(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceCode = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceDetail(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceDetail = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetServiceStatus(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetStatus(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Status = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) SetUuid(v string) *DescribeWebLockBindListResponseBodyBindList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockBindListResponseBodyBindList) Validate() error {
	return dara.Validate(s)
}
