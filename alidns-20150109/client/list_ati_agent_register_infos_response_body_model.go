// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiAgentRegisterInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) *ListAtiAgentRegisterInfosResponseBody
	GetAccessDeniedDetail() *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail
	SetAgentRegisterInfos(v *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) *ListAtiAgentRegisterInfosResponseBody
	GetAgentRegisterInfos() *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos
	SetMaxResults(v int32) *ListAtiAgentRegisterInfosResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAtiAgentRegisterInfosResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListAtiAgentRegisterInfosResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiAgentRegisterInfosResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAtiAgentRegisterInfosResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListAtiAgentRegisterInfosResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListAtiAgentRegisterInfosResponseBody
	GetTotalPages() *int32
}

type ListAtiAgentRegisterInfosResponseBody struct {
	// The details about the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	AgentRegisterInfos *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos `json:"AgentRegisterInfos,omitempty" xml:"AgentRegisterInfos,omitempty" type:"Struct"`
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page set for the paged query. This is the paging size. Maximum value: **100**. Default value: **20**. Settings determine how many rows are displayed per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 11
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListAtiAgentRegisterInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetAccessDeniedDetail() *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetAgentRegisterInfos() *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos {
	return s.AgentRegisterInfos
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListAtiAgentRegisterInfosResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetAccessDeniedDetail(v *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) *ListAtiAgentRegisterInfosResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetAgentRegisterInfos(v *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) *ListAtiAgentRegisterInfosResponseBody {
	s.AgentRegisterInfos = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetMaxResults(v int32) *ListAtiAgentRegisterInfosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetNextToken(v string) *ListAtiAgentRegisterInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetPageNumber(v int32) *ListAtiAgentRegisterInfosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetPageSize(v int32) *ListAtiAgentRegisterInfosResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetRequestId(v string) *ListAtiAgentRegisterInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetTotalItems(v int32) *ListAtiAgentRegisterInfosResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) SetTotalPages(v int32) *ListAtiAgentRegisterInfosResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.AgentRegisterInfos != nil {
		if err := s.AgentRegisterInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// CreateUser
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authorization principal.
	//
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authorization principal.
	//
	// example:
	//
	// 1046973331XXXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The identity type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encrypted diagnostic message.
	//
	// example:
	//
	// AQFohtp4aIbaeEXXXXQxNjFDLUIzMzgtNTXXXX05NkFCLUI2RkY5XXXXzAzQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The cause of the authentication failure. Valid values:
	//
	// - ExplicitDeny: Explicit denial.
	//
	// - ImplicitDeny: Implicit denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos struct {
	AgentRegisterInfo []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo `json:"AgentRegisterInfo,omitempty" xml:"AgentRegisterInfo,omitempty" type:"Repeated"`
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) GetAgentRegisterInfo() []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	return s.AgentRegisterInfo
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) SetAgentRegisterInfo(v []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos {
	s.AgentRegisterInfo = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfos) Validate() error {
	if s.AgentRegisterInfo != nil {
		for _, item := range s.AgentRegisterInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo struct {
	AgentDisplayName    *string                                                                            `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	AgentHost           *string                                                                            `json:"AgentHost,omitempty" xml:"AgentHost,omitempty"`
	AgentId             *string                                                                            `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentRegisterInfoId *string                                                                            `json:"AgentRegisterInfoId,omitempty" xml:"AgentRegisterInfoId,omitempty"`
	AgentVersion        *string                                                                            `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	AtiName             *string                                                                            `json:"AtiName,omitempty" xml:"AtiName,omitempty"`
	CreateTimestamp     *string                                                                            `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Endpoints           *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	Status              *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTimestamp     *string                                                                            `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAgentHost() *string {
	return s.AgentHost
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAgentRegisterInfoId() *string {
	return s.AgentRegisterInfoId
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetAtiName() *string {
	return s.AtiName
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetCreateTimestamp() *string {
	return s.CreateTimestamp
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetEndpoints() *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints {
	return s.Endpoints
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetStatus() *string {
	return s.Status
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) GetUpdateTimestamp() *string {
	return s.UpdateTimestamp
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAgentDisplayName(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AgentDisplayName = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAgentHost(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AgentHost = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAgentId(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AgentId = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAgentRegisterInfoId(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AgentRegisterInfoId = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAgentVersion(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AgentVersion = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetAtiName(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.AtiName = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetCreateTimestamp(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetEndpoints(v *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.Endpoints = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetStatus(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.Status = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) SetUpdateTimestamp(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfo) Validate() error {
	if s.Endpoints != nil {
		if err := s.Endpoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints struct {
	Endpoint []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) GetEndpoint() []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint {
	return s.Endpoint
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) SetEndpoint(v []*ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints {
	s.Endpoint = v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpoints) Validate() error {
	if s.Endpoint != nil {
		for _, item := range s.Endpoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint struct {
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) GetProtocol() *string {
	return s.Protocol
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) SetProtocol(v string) *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint {
	s.Protocol = &v
	return s
}

func (s *ListAtiAgentRegisterInfosResponseBodyAgentRegisterInfosAgentRegisterInfoEndpointsEndpoint) Validate() error {
	return dara.Validate(s)
}
