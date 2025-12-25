// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PublishStatusResponseBodyAccessDeniedDetail) *PublishStatusResponseBody
	GetAccessDeniedDetail() *PublishStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *PublishStatusResponseBody
	GetCode() *int64
	SetMessage(v string) *PublishStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *PublishStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *PublishStatusResponseBody
	GetSuccess() *bool
	SetSyncStatus(v string) *PublishStatusResponseBody
	GetSyncStatus() *string
}

type PublishStatusResponseBody struct {
	AccessDeniedDetail *PublishStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 344794c32937474a9c59eb130936****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// succeed
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
}

func (s PublishStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishStatusResponseBody) GoString() string {
	return s.String()
}

func (s *PublishStatusResponseBody) GetAccessDeniedDetail() *PublishStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PublishStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PublishStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *PublishStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishStatusResponseBody) GetSyncStatus() *string {
	return s.SyncStatus
}

func (s *PublishStatusResponseBody) SetAccessDeniedDetail(v *PublishStatusResponseBodyAccessDeniedDetail) *PublishStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PublishStatusResponseBody) SetCode(v int64) *PublishStatusResponseBody {
	s.Code = &v
	return s
}

func (s *PublishStatusResponseBody) SetMessage(v string) *PublishStatusResponseBody {
	s.Message = &v
	return s
}

func (s *PublishStatusResponseBody) SetRequestId(v string) *PublishStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishStatusResponseBody) SetStatus(v string) *PublishStatusResponseBody {
	s.Status = &v
	return s
}

func (s *PublishStatusResponseBody) SetSuccess(v bool) *PublishStatusResponseBody {
	s.Success = &v
	return s
}

func (s *PublishStatusResponseBody) SetSyncStatus(v string) *PublishStatusResponseBody {
	s.SyncStatus = &v
	return s
}

func (s *PublishStatusResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PublishStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PublishStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PublishStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PublishStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
