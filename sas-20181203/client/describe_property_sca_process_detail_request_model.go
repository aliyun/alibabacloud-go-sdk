// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaProcessDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribePropertyScaProcessDetailRequest
	GetBizType() *string
	SetCmdline(v string) *DescribePropertyScaProcessDetailRequest
	GetCmdline() *string
	SetCurrentPage(v int32) *DescribePropertyScaProcessDetailRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribePropertyScaProcessDetailRequest
	GetPageSize() *int32
	SetPid(v string) *DescribePropertyScaProcessDetailRequest
	GetPid() *string
	SetRemark(v string) *DescribePropertyScaProcessDetailRequest
	GetRemark() *string
	SetUuid(v string) *DescribePropertyScaProcessDetailRequest
	GetUuid() *string
}

type DescribePropertyScaProcessDetailRequest struct {
	// The type of the application process. Default value: **java**. Valid values:
	//
	// 	- **java**: Java process.
	//
	// 	- **php**: PHP process.
	//
	// example:
	//
	// java
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The startup parameter.
	//
	// >  This parameter supports only prefix queries. Fuzzy match is not supported.
	//
	// example:
	//
	// java -jar
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 10. If you leave this parameter empty, 10 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process ID.
	//
	// >  Only exact match is supported.
	//
	// example:
	//
	// 756
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The information about the server that you want to query. The value can be the public IP address, private IP address, or name of the server. Fuzzy match is supported.
	//
	// example:
	//
	// 10.167.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The UUID of the server.
	//
	// >
	//
	// 	- You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// 	- Only exact match is supported.
	//
	// example:
	//
	// D0D6E6E4-CB8C-4897-B852-46AEFDA0****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyScaProcessDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaProcessDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaProcessDetailRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribePropertyScaProcessDetailRequest) GetCmdline() *string {
	return s.Cmdline
}

func (s *DescribePropertyScaProcessDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaProcessDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaProcessDetailRequest) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyScaProcessDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyScaProcessDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyScaProcessDetailRequest) SetBizType(v string) *DescribePropertyScaProcessDetailRequest {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetCmdline(v string) *DescribePropertyScaProcessDetailRequest {
	s.Cmdline = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetCurrentPage(v int32) *DescribePropertyScaProcessDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetPageSize(v int32) *DescribePropertyScaProcessDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetPid(v string) *DescribePropertyScaProcessDetailRequest {
	s.Pid = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetRemark(v string) *DescribePropertyScaProcessDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) SetUuid(v string) *DescribePropertyScaProcessDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaProcessDetailRequest) Validate() error {
	return dara.Validate(s)
}
