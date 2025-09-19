// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCmdline(v string) *DescribePropertyProcDetailRequest
	GetCmdline() *string
	SetCurrentPage(v int32) *DescribePropertyProcDetailRequest
	GetCurrentPage() *int32
	SetExtend(v string) *DescribePropertyProcDetailRequest
	GetExtend() *string
	SetName(v string) *DescribePropertyProcDetailRequest
	GetName() *string
	SetPageSize(v int32) *DescribePropertyProcDetailRequest
	GetPageSize() *int32
	SetProcTimeEnd(v int64) *DescribePropertyProcDetailRequest
	GetProcTimeEnd() *int64
	SetProcTimeStart(v int64) *DescribePropertyProcDetailRequest
	GetProcTimeStart() *int64
	SetRemark(v string) *DescribePropertyProcDetailRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribePropertyProcDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetUser(v string) *DescribePropertyProcDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyProcDetailRequest
	GetUuid() *string
}

type DescribePropertyProcDetailRequest struct {
	// The startup parameter of the process.
	//
	// example:
	//
	// ./8888
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy search by process name is supported. If you want to use fuzzy search, set the parameter to 1. If you set the parameter to a different value or leave the parameter empty, fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// 8888
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The timestamp when the process ends. Unit: milliseconds.
	//
	// example:
	//
	// 1649587453000
	ProcTimeEnd *int64 `json:"ProcTimeEnd,omitempty" xml:"ProcTimeEnd,omitempty"`
	// The timestamp when the process starts. Unit: milliseconds.
	//
	// example:
	//
	// 1648809853000
	ProcTimeStart *int64 `json:"ProcTimeStart,omitempty" xml:"ProcTimeStart,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the account ID.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The user who runs the process.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyProcDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribePropertyProcDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyProcDetailRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribePropertyProcDetailRequest) GetName() *string {
	return s.Name
}

func (s *DescribePropertyProcDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyProcDetailRequest) GetProcTimeEnd() *int64 {
	return s.ProcTimeEnd
}

func (s *DescribePropertyProcDetailRequest) GetProcTimeStart() *int64 {
	return s.ProcTimeStart
}

func (s *DescribePropertyProcDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyProcDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribePropertyProcDetailRequest) GetUser() *string {
	return s.User
}

func (s *DescribePropertyProcDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyProcDetailRequest) SetCmdline(v string) *DescribePropertyProcDetailRequest {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetCurrentPage(v int32) *DescribePropertyProcDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetExtend(v string) *DescribePropertyProcDetailRequest {
	s.Extend = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetName(v string) *DescribePropertyProcDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetPageSize(v int32) *DescribePropertyProcDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetProcTimeEnd(v int64) *DescribePropertyProcDetailRequest {
	s.ProcTimeEnd = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetProcTimeStart(v int64) *DescribePropertyProcDetailRequest {
	s.ProcTimeStart = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetRemark(v string) *DescribePropertyProcDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribePropertyProcDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetUser(v string) *DescribePropertyProcDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) SetUuid(v string) *DescribePropertyProcDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyProcDetailRequest) Validate() error {
	return dara.Validate(s)
}
