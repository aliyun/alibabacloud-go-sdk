// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetAccessDeniedDetail() *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail
	SetAgentRegisterInfoId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetAgentRegisterInfoId() *string
	SetErrorCode(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetErrorMessage() *string
	SetMessage(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetMessage() *string
	SetPrecheckStatus(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetPrecheckStatus() *string
	SetRequestId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetSuccess() *bool
	SetVerifyTimestamp(v int64) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody
	GetVerifyTimestamp() *int64
}

type VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody struct {
	AccessDeniedDetail *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 2074753647748672512
	AgentRegisterInfoId *string `json:"AgentRegisterInfoId,omitempty" xml:"AgentRegisterInfoId,omitempty"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// PrecheckFailedOnTooManyVmSnapshot
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// ACME DNS-01 域名验证通过
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// PASSED
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1782572400000
	VerifyTimestamp *int64 `json:"VerifyTimestamp,omitempty" xml:"VerifyTimestamp,omitempty"`
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetAccessDeniedDetail() *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetAgentRegisterInfoId() *string {
	return s.AgentRegisterInfoId
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetPrecheckStatus() *string {
	return s.PrecheckStatus
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) GetVerifyTimestamp() *int64 {
	return s.VerifyTimestamp
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetAccessDeniedDetail(v *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetAgentRegisterInfoId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.AgentRegisterInfoId = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetErrorCode(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetErrorMessage(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetMessage(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetPrecheckStatus(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.PrecheckStatus = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetRequestId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetSuccess(v bool) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) SetVerifyTimestamp(v int64) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody {
	s.VerifyTimestamp = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// RemoveRspDomainServerHoldStatusForGateway
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 2015555733387XXXX
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 10469733312XXX
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaNIARXXXXUQwNjE0LUQzN0XXXXVEQy1BQzExLTMzXXXXNTkxRjk1Ng==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// DlpSend
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetAuthAction(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) SetPolicyType(v string) *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *VerifyAtiAgentRegisterInfoAcmeChallengeRecordResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
