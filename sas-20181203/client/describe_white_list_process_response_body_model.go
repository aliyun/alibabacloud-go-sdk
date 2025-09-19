// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeWhiteListProcessResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeWhiteListProcessResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWhiteListProcessResponseBody
	GetPageSize() *int32
	SetProcesses(v []*DescribeWhiteListProcessResponseBodyProcesses) *DescribeWhiteListProcessResponseBody
	GetProcesses() []*DescribeWhiteListProcessResponseBodyProcesses
	SetRequestId(v string) *DescribeWhiteListProcessResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWhiteListProcessResponseBody
	GetTotalCount() *int32
}

type DescribeWhiteListProcessResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 200
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the processes.
	Processes []*DescribeWhiteListProcessResponseBodyProcesses `json:"Processes,omitempty" xml:"Processes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D81DD78E-E006-5C65-A171-C8CB09XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 44
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWhiteListProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListProcessResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWhiteListProcessResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListProcessResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListProcessResponseBody) GetProcesses() []*DescribeWhiteListProcessResponseBodyProcesses {
	return s.Processes
}

func (s *DescribeWhiteListProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhiteListProcessResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWhiteListProcessResponseBody) SetCount(v int32) *DescribeWhiteListProcessResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) SetCurrentPage(v int32) *DescribeWhiteListProcessResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) SetPageSize(v int32) *DescribeWhiteListProcessResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) SetProcesses(v []*DescribeWhiteListProcessResponseBodyProcesses) *DescribeWhiteListProcessResponseBody {
	s.Processes = v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) SetRequestId(v string) *DescribeWhiteListProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) SetTotalCount(v int32) *DescribeWhiteListProcessResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWhiteListProcessResponseBodyProcesses struct {
	// The path to the process startup file.
	//
	// example:
	//
	// /root/bash1
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The primary key of the process information.
	//
	// example:
	//
	// 2100019
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The trust score of the process.
	//
	// example:
	//
	// 99
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The MD5 hash value of the process startup file.
	//
	// example:
	//
	// a28e8eba54ece1f3748d80e57dc89400
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 53090
	ProcessId *int32 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// vim
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The type of the process. Valid values:
	//
	// 	- **1**: trusted
	//
	// 	- **2**: suspicious
	//
	// 	- **3**: malicious
	//
	// example:
	//
	// 1
	ProcessType *int32 `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// Indicates whether the process is trusted. Valid values:
	//
	// 	- **1**: no
	//
	// 	- **2**: yes
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeWhiteListProcessResponseBodyProcesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListProcessResponseBodyProcesses) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetId() *int64 {
	return s.Id
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetLevel() *int32 {
	return s.Level
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetMd5() *string {
	return s.Md5
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetProcessId() *int32 {
	return s.ProcessId
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetProcessType() *int32 {
	return s.ProcessType
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetFilePath(v string) *DescribeWhiteListProcessResponseBodyProcesses {
	s.FilePath = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetId(v int64) *DescribeWhiteListProcessResponseBodyProcesses {
	s.Id = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetLevel(v int32) *DescribeWhiteListProcessResponseBodyProcesses {
	s.Level = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetMd5(v string) *DescribeWhiteListProcessResponseBodyProcesses {
	s.Md5 = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetProcessId(v int32) *DescribeWhiteListProcessResponseBodyProcesses {
	s.ProcessId = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetProcessName(v string) *DescribeWhiteListProcessResponseBodyProcesses {
	s.ProcessName = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetProcessType(v int32) *DescribeWhiteListProcessResponseBodyProcesses {
	s.ProcessType = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) SetStatus(v int32) *DescribeWhiteListProcessResponseBodyProcesses {
	s.Status = &v
	return s
}

func (s *DescribeWhiteListProcessResponseBodyProcesses) Validate() error {
	return dara.Validate(s)
}
