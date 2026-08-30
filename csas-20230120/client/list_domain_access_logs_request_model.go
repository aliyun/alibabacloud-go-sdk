// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainAccessLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlockAction(v string) *ListDomainAccessLogsRequest
	GetBlockAction() *string
	SetCurrentPage(v int32) *ListDomainAccessLogsRequest
	GetCurrentPage() *int32
	SetDepartment(v string) *ListDomainAccessLogsRequest
	GetDepartment() *string
	SetEndTime(v int64) *ListDomainAccessLogsRequest
	GetEndTime() *int64
	SetPageSize(v int32) *ListDomainAccessLogsRequest
	GetPageSize() *int32
	SetPolicyType(v string) *ListDomainAccessLogsRequest
	GetPolicyType() *string
	SetRemoteHost(v string) *ListDomainAccessLogsRequest
	GetRemoteHost() *string
	SetStartTime(v int64) *ListDomainAccessLogsRequest
	GetStartTime() *int64
	SetUserName(v string) *ListDomainAccessLogsRequest
	GetUserName() *string
}

type ListDomainAccessLogsRequest struct {
	// The action taken upon a rule hit. Exact match is used. Valid values:
	//
	// - Audit: Audit.
	//
	// - Observe: Observe only.
	//
	// - WhiteList: Allowed by whitelist.
	//
	// - Block: Blocked.
	//
	// - Redirect: Redirected to a prompt page.
	//
	// example:
	//
	// Block
	BlockAction *string `json:"BlockAction,omitempty" xml:"BlockAction,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The department. Exact match is used.
	//
	// example:
	//
	// IT department
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// The end time of the query. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1754956800
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page in paging. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The policy type used to filter results.
	//
	// example:
	//
	// la_domain_white
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The destination domain name accessed. Exact match is used.
	//
	// example:
	//
	// www.example.com
	RemoteHost *string `json:"RemoteHost,omitempty" xml:"RemoteHost,omitempty"`
	// The start time of the query. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1754870400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The username. Exact match is used.
	//
	// example:
	//
	// zhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListDomainAccessLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainAccessLogsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainAccessLogsRequest) GetBlockAction() *string {
	return s.BlockAction
}

func (s *ListDomainAccessLogsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListDomainAccessLogsRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListDomainAccessLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDomainAccessLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainAccessLogsRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListDomainAccessLogsRequest) GetRemoteHost() *string {
	return s.RemoteHost
}

func (s *ListDomainAccessLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDomainAccessLogsRequest) GetUserName() *string {
	return s.UserName
}

func (s *ListDomainAccessLogsRequest) SetBlockAction(v string) *ListDomainAccessLogsRequest {
	s.BlockAction = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetCurrentPage(v int32) *ListDomainAccessLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetDepartment(v string) *ListDomainAccessLogsRequest {
	s.Department = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetEndTime(v int64) *ListDomainAccessLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetPageSize(v int32) *ListDomainAccessLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetPolicyType(v string) *ListDomainAccessLogsRequest {
	s.PolicyType = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetRemoteHost(v string) *ListDomainAccessLogsRequest {
	s.RemoteHost = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetStartTime(v int64) *ListDomainAccessLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDomainAccessLogsRequest) SetUserName(v string) *ListDomainAccessLogsRequest {
	s.UserName = &v
	return s
}

func (s *ListDomainAccessLogsRequest) Validate() error {
	return dara.Validate(s)
}
