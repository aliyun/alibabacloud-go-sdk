// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeRightAngleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *OptimizeRightAngleResponseBodyAccessDeniedDetail) *OptimizeRightAngleResponseBody
	GetAccessDeniedDetail() *OptimizeRightAngleResponseBodyAccessDeniedDetail
	SetCode(v int64) *OptimizeRightAngleResponseBody
	GetCode() *int64
	SetMessage(v string) *OptimizeRightAngleResponseBody
	GetMessage() *string
	SetRequestId(v string) *OptimizeRightAngleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OptimizeRightAngleResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *OptimizeRightAngleResponseBody
	GetTaskId() *string
}

type OptimizeRightAngleResponseBody struct {
	AccessDeniedDetail *OptimizeRightAngleResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 2345****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OptimizeRightAngleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OptimizeRightAngleResponseBody) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponseBody) GetAccessDeniedDetail() *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *OptimizeRightAngleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *OptimizeRightAngleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OptimizeRightAngleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OptimizeRightAngleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OptimizeRightAngleResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *OptimizeRightAngleResponseBody) SetAccessDeniedDetail(v *OptimizeRightAngleResponseBodyAccessDeniedDetail) *OptimizeRightAngleResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetCode(v int64) *OptimizeRightAngleResponseBody {
	s.Code = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetMessage(v string) *OptimizeRightAngleResponseBody {
	s.Message = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetRequestId(v string) *OptimizeRightAngleResponseBody {
	s.RequestId = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetSuccess(v bool) *OptimizeRightAngleResponseBody {
	s.Success = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) SetTaskId(v string) *OptimizeRightAngleResponseBody {
	s.TaskId = &v
	return s
}

func (s *OptimizeRightAngleResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OptimizeRightAngleResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s OptimizeRightAngleResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s OptimizeRightAngleResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetAuthAction(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) SetPolicyType(v string) *OptimizeRightAngleResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *OptimizeRightAngleResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
