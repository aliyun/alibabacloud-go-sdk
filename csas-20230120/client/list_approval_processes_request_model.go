// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApprovalProcessesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListApprovalProcessesRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *ListApprovalProcessesRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListApprovalProcessesRequest
	GetPolicyId() *string
	SetPolicyType(v string) *ListApprovalProcessesRequest
	GetPolicyType() *string
	SetProcessIds(v []*string) *ListApprovalProcessesRequest
	GetProcessIds() []*string
	SetProcessName(v string) *ListApprovalProcessesRequest
	GetProcessName() *string
	SetSaseUserId(v string) *ListApprovalProcessesRequest
	GetSaseUserId() *string
	SetUsername(v string) *ListApprovalProcessesRequest
	GetUsername() *string
}

type ListApprovalProcessesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pid-6d6ad77d5b52****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string   `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	ProcessIds []*string `json:"ProcessIds,omitempty" xml:"ProcessIds,omitempty" type:"Repeated"`
	// example:
	//
	// test
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListApprovalProcessesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApprovalProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListApprovalProcessesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListApprovalProcessesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApprovalProcessesRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListApprovalProcessesRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListApprovalProcessesRequest) GetProcessIds() []*string {
	return s.ProcessIds
}

func (s *ListApprovalProcessesRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *ListApprovalProcessesRequest) GetSaseUserId() *string {
	return s.SaseUserId
}

func (s *ListApprovalProcessesRequest) GetUsername() *string {
	return s.Username
}

func (s *ListApprovalProcessesRequest) SetCurrentPage(v int64) *ListApprovalProcessesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetPageSize(v int64) *ListApprovalProcessesRequest {
	s.PageSize = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetPolicyId(v string) *ListApprovalProcessesRequest {
	s.PolicyId = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetPolicyType(v string) *ListApprovalProcessesRequest {
	s.PolicyType = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetProcessIds(v []*string) *ListApprovalProcessesRequest {
	s.ProcessIds = v
	return s
}

func (s *ListApprovalProcessesRequest) SetProcessName(v string) *ListApprovalProcessesRequest {
	s.ProcessName = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetSaseUserId(v string) *ListApprovalProcessesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListApprovalProcessesRequest) SetUsername(v string) *ListApprovalProcessesRequest {
	s.Username = &v
	return s
}

func (s *ListApprovalProcessesRequest) Validate() error {
	return dara.Validate(s)
}
