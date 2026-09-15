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
	SetNextToken(v string) *DescribePropertyProcDetailRequest
	GetNextToken() *string
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
	SetUseNextToken(v bool) *DescribePropertyProcDetailRequest
	GetUseNextToken() *bool
	SetUser(v string) *DescribePropertyProcDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyProcDetailRequest
	GetUuid() *string
}

type DescribePropertyProcDetailRequest struct {
	// The startup parameters of the process.
	//
	// example:
	//
	// ./8888
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// The page number of the page to return. Default value: **1**, which indicates that the first page is returned.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether fuzzy match is supported for the process name. Set this parameter to 1 to enable fuzzy match. Other values or an empty value indicate that fuzzy match is not supported.
	//
	// example:
	//
	// 1
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The process name.
	//
	// example:
	//
	// 8888
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that marks the current position from which to start reading. Leave this parameter empty to start reading from the beginning.
	//
	// > You do not need to specify this parameter for the first call. The NextToken value for the second call is included in the response of the first call. Each subsequent response contains the NextToken value for the next call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k+AtdhNE3kgQEK36GujZ5on+tWdc+4WoaoMP/kUNxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page in a paged query. Default value: **10**, which indicates that 10 entries of process Asset Fingerprints information are displayed per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The end timestamp of the process startup time range. Unit: milliseconds.
	//
	// example:
	//
	// 1649587453000
	ProcTimeEnd *int64 `json:"ProcTimeEnd,omitempty" xml:"ProcTimeEnd,omitempty"`
	// The start timestamp of the process startup time range. Unit: milliseconds.
	//
	// example:
	//
	// 1648809853000
	ProcTimeStart *int64 `json:"ProcTimeStart,omitempty" xml:"ProcTimeStart,omitempty"`
	// The name or IP address of the server that you want to query.
	//
	// example:
	//
	// 192.168.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member accounts in the resource folder.
	//
	// >You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to use the NextToken method to retrieve the vulnerability list data. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// false
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The information about the user that runs the process.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server that you want to query.
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

func (s *DescribePropertyProcDetailRequest) GetNextToken() *string {
	return s.NextToken
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

func (s *DescribePropertyProcDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
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

func (s *DescribePropertyProcDetailRequest) SetNextToken(v string) *DescribePropertyProcDetailRequest {
	s.NextToken = &v
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

func (s *DescribePropertyProcDetailRequest) SetUseNextToken(v bool) *DescribePropertyProcDetailRequest {
	s.UseNextToken = &v
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
