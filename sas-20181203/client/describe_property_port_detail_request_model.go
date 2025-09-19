// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindIp(v string) *DescribePropertyPortDetailRequest
	GetBindIp() *string
	SetCurrentPage(v int32) *DescribePropertyPortDetailRequest
	GetCurrentPage() *int32
	SetExtend(v string) *DescribePropertyPortDetailRequest
	GetExtend() *string
	SetPageSize(v int32) *DescribePropertyPortDetailRequest
	GetPageSize() *int32
	SetPort(v string) *DescribePropertyPortDetailRequest
	GetPort() *string
	SetProcName(v string) *DescribePropertyPortDetailRequest
	GetProcName() *string
	SetRemark(v string) *DescribePropertyPortDetailRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribePropertyPortDetailRequest
	GetResourceDirectoryAccountId() *int64
	SetUuid(v string) *DescribePropertyPortDetailRequest
	GetUuid() *string
}

type DescribePropertyPortDetailRequest struct {
	// The IP address bound to the port.
	//
	// example:
	//
	// 0.0.XX.XX
	BindIp *string `json:"BindIp,omitempty" xml:"BindIp,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy search by port number is supported. If you want to use fuzzy search, set the parameter to **1**. If you set the parameter to a different value or leave the parameter empty, fuzzy search is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The listening port of the server.
	//
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the server process.
	//
	// example:
	//
	// sshd
	ProcName *string `json:"ProcName,omitempty" xml:"ProcName,omitempty"`
	// The name or IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 50d213b4-3a35-427a-b8a5-04b0c7e1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyPortDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailRequest) GetBindIp() *string {
	return s.BindIp
}

func (s *DescribePropertyPortDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyPortDetailRequest) GetExtend() *string {
	return s.Extend
}

func (s *DescribePropertyPortDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyPortDetailRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyPortDetailRequest) GetProcName() *string {
	return s.ProcName
}

func (s *DescribePropertyPortDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyPortDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribePropertyPortDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyPortDetailRequest) SetBindIp(v string) *DescribePropertyPortDetailRequest {
	s.BindIp = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetCurrentPage(v int32) *DescribePropertyPortDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetExtend(v string) *DescribePropertyPortDetailRequest {
	s.Extend = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetPageSize(v int32) *DescribePropertyPortDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetPort(v string) *DescribePropertyPortDetailRequest {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetProcName(v string) *DescribePropertyPortDetailRequest {
	s.ProcName = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetRemark(v string) *DescribePropertyPortDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribePropertyPortDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) SetUuid(v string) *DescribePropertyPortDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyPortDetailRequest) Validate() error {
	return dara.Validate(s)
}
