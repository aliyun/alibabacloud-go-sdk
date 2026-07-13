// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiChangeLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ListAtiChangeLogsResponseBodyAccessDeniedDetail) *ListAtiChangeLogsResponseBody
	GetAccessDeniedDetail() *ListAtiChangeLogsResponseBodyAccessDeniedDetail
	SetLogs(v *ListAtiChangeLogsResponseBodyLogs) *ListAtiChangeLogsResponseBody
	GetLogs() *ListAtiChangeLogsResponseBodyLogs
	SetPageNumber(v int32) *ListAtiChangeLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAtiChangeLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAtiChangeLogsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListAtiChangeLogsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListAtiChangeLogsResponseBody
	GetTotalPages() *int32
}

type ListAtiChangeLogsResponseBody struct {
	// The details of the access denial. This field is returned only when RAM authentication fails.
	AccessDeniedDetail *ListAtiChangeLogsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Logs               *ListAtiChangeLogsResponseBodyLogs               `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Struct"`
	// The current page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B57C121B-A45F-44D8-A9B2-13E5A5044195
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the address list.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 123
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListAtiChangeLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsResponseBody) GetAccessDeniedDetail() *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ListAtiChangeLogsResponseBody) GetLogs() *ListAtiChangeLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListAtiChangeLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAtiChangeLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAtiChangeLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAtiChangeLogsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListAtiChangeLogsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListAtiChangeLogsResponseBody) SetAccessDeniedDetail(v *ListAtiChangeLogsResponseBodyAccessDeniedDetail) *ListAtiChangeLogsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetLogs(v *ListAtiChangeLogsResponseBodyLogs) *ListAtiChangeLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetPageNumber(v int32) *ListAtiChangeLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetPageSize(v int32) *ListAtiChangeLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetRequestId(v string) *ListAtiChangeLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetTotalItems(v int32) *ListAtiChangeLogsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListAtiChangeLogsResponseBody) SetTotalPages(v int32) *ListAtiChangeLogsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListAtiChangeLogsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Logs != nil {
		if err := s.Logs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAtiChangeLogsResponseBodyAccessDeniedDetail struct {
	// The unauthorized operation that was attempted.
	//
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGateway
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
	// 10469733312XXX
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
	// The reason for the authentication failure. Valid values:
	//
	// - ExplicitDeny: explicit deny.
	//
	// - ImplicitDeny: implicit deny.
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

func (s ListAtiChangeLogsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ListAtiChangeLogsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ListAtiChangeLogsResponseBodyLogs struct {
	Log []*ListAtiChangeLogsResponseBodyLogsLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Repeated"`
}

func (s ListAtiChangeLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsResponseBodyLogs) GetLog() []*ListAtiChangeLogsResponseBodyLogsLog {
	return s.Log
}

func (s *ListAtiChangeLogsResponseBodyLogs) SetLog(v []*ListAtiChangeLogsResponseBodyLogsLog) *ListAtiChangeLogsResponseBodyLogs {
	s.Log = v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogs) Validate() error {
	if s.Log != nil {
		for _, item := range s.Log {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAtiChangeLogsResponseBodyLogsLog struct {
	AgentId         *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus     *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	CreateTimestamp *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	CreatorSubType  *string `json:"CreatorSubType,omitempty" xml:"CreatorSubType,omitempty"`
	CreatorType     *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	ErrorMessage    *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	LogId           *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	OperationName   *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	OperationType   *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OperatorAccount *string `json:"OperatorAccount,omitempty" xml:"OperatorAccount,omitempty"`
	Result          *string `json:"Result,omitempty" xml:"Result,omitempty"`
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Timestamp       *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UpdateTimestamp *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAtiChangeLogsResponseBodyLogsLog) String() string {
	return dara.Prettify(s)
}

func (s ListAtiChangeLogsResponseBodyLogsLog) GoString() string {
	return s.String()
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetCreatorSubType() *string {
	return s.CreatorSubType
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetLogId() *string {
	return s.LogId
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetOperationName() *string {
	return s.OperationName
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetOperationType() *string {
	return s.OperationType
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetOperatorAccount() *string {
	return s.OperatorAccount
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetResult() *string {
	return s.Result
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) GetVersion() *string {
	return s.Version
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetAgentId(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.AgentId = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetAgentStatus(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.AgentStatus = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetCreateTimestamp(v int64) *ListAtiChangeLogsResponseBodyLogsLog {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetCreatorSubType(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.CreatorSubType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetCreatorType(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.CreatorType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetErrorMessage(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.ErrorMessage = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetLogId(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.LogId = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetOperationName(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.OperationName = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetOperationType(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.OperationType = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetOperatorAccount(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.OperatorAccount = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetResult(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.Result = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetSourceIp(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.SourceIp = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetTimestamp(v int64) *ListAtiChangeLogsResponseBodyLogsLog {
	s.Timestamp = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetUpdateTimestamp(v int64) *ListAtiChangeLogsResponseBodyLogsLog {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) SetVersion(v string) *ListAtiChangeLogsResponseBodyLogsLog {
	s.Version = &v
	return s
}

func (s *ListAtiChangeLogsResponseBodyLogsLog) Validate() error {
	return dara.Validate(s)
}
