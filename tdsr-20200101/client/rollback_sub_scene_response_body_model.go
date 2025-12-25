// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackSubSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RollbackSubSceneResponseBodyAccessDeniedDetail) *RollbackSubSceneResponseBody
	GetAccessDeniedDetail() *RollbackSubSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *RollbackSubSceneResponseBody
	GetCode() *int64
	SetMessage(v string) *RollbackSubSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *RollbackSubSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackSubSceneResponseBody
	GetSuccess() *bool
}

type RollbackSubSceneResponseBody struct {
	AccessDeniedDetail *RollbackSubSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// sucess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3BCAD49D-2AC1-13EB-AC19-8C7A46CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RollbackSubSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackSubSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponseBody) GetAccessDeniedDetail() *RollbackSubSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RollbackSubSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RollbackSubSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RollbackSubSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackSubSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackSubSceneResponseBody) SetAccessDeniedDetail(v *RollbackSubSceneResponseBodyAccessDeniedDetail) *RollbackSubSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RollbackSubSceneResponseBody) SetCode(v int64) *RollbackSubSceneResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetMessage(v string) *RollbackSubSceneResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetRequestId(v string) *RollbackSubSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackSubSceneResponseBody) SetSuccess(v bool) *RollbackSubSceneResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackSubSceneResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RollbackSubSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s RollbackSubSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RollbackSubSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RollbackSubSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RollbackSubSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
