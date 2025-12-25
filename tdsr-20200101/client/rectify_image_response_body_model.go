// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectifyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *RectifyImageResponseBodyAccessDeniedDetail) *RectifyImageResponseBody
	GetAccessDeniedDetail() *RectifyImageResponseBodyAccessDeniedDetail
	SetCode(v int64) *RectifyImageResponseBody
	GetCode() *int64
	SetMessage(v string) *RectifyImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *RectifyImageResponseBody
	GetRequestId() *string
	SetSubSceneId(v string) *RectifyImageResponseBody
	GetSubSceneId() *string
	SetSuccess(v bool) *RectifyImageResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *RectifyImageResponseBody
	GetTaskId() *string
}

type RectifyImageResponseBody struct {
	AccessDeniedDetail *RectifyImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 2345****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1234****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RectifyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RectifyImageResponseBody) GoString() string {
	return s.String()
}

func (s *RectifyImageResponseBody) GetAccessDeniedDetail() *RectifyImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *RectifyImageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *RectifyImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RectifyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RectifyImageResponseBody) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *RectifyImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RectifyImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RectifyImageResponseBody) SetAccessDeniedDetail(v *RectifyImageResponseBodyAccessDeniedDetail) *RectifyImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *RectifyImageResponseBody) SetCode(v int64) *RectifyImageResponseBody {
	s.Code = &v
	return s
}

func (s *RectifyImageResponseBody) SetMessage(v string) *RectifyImageResponseBody {
	s.Message = &v
	return s
}

func (s *RectifyImageResponseBody) SetRequestId(v string) *RectifyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RectifyImageResponseBody) SetSubSceneId(v string) *RectifyImageResponseBody {
	s.SubSceneId = &v
	return s
}

func (s *RectifyImageResponseBody) SetSuccess(v bool) *RectifyImageResponseBody {
	s.Success = &v
	return s
}

func (s *RectifyImageResponseBody) SetTaskId(v string) *RectifyImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *RectifyImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RectifyImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s RectifyImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s RectifyImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *RectifyImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *RectifyImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
