// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectVerticalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RectVerticalResponseBodyAccessDeniedDetail) *RectVerticalResponseBody
	GetAccessDeniedDetail() *RectVerticalResponseBodyAccessDeniedDetail
	SetCode(v int64) *RectVerticalResponseBody
	GetCode() *int64
	SetMessage(v string) *RectVerticalResponseBody
	GetMessage() *string
	SetRequestId(v string) *RectVerticalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RectVerticalResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *RectVerticalResponseBody
	GetTaskId() *string
}

type RectVerticalResponseBody struct {
	AccessDeniedDetail *RectVerticalResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// syuwoiewyieiyy****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RectVerticalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RectVerticalResponseBody) GoString() string {
	return s.String()
}

func (s *RectVerticalResponseBody) GetAccessDeniedDetail() *RectVerticalResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RectVerticalResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RectVerticalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RectVerticalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RectVerticalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RectVerticalResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RectVerticalResponseBody) SetAccessDeniedDetail(v *RectVerticalResponseBodyAccessDeniedDetail) *RectVerticalResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RectVerticalResponseBody) SetCode(v int64) *RectVerticalResponseBody {
	s.Code = &v
	return s
}

func (s *RectVerticalResponseBody) SetMessage(v string) *RectVerticalResponseBody {
	s.Message = &v
	return s
}

func (s *RectVerticalResponseBody) SetRequestId(v string) *RectVerticalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectVerticalResponseBody) SetSuccess(v bool) *RectVerticalResponseBody {
	s.Success = &v
	return s
}

func (s *RectVerticalResponseBody) SetTaskId(v string) *RectVerticalResponseBody {
	s.TaskId = &v
	return s
}

func (s *RectVerticalResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RectVerticalResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s RectVerticalResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RectVerticalResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RectVerticalResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RectVerticalResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
