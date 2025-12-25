// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PredImageResponseBodyAccessDeniedDetail) *PredImageResponseBody
	GetAccessDeniedDetail() *PredImageResponseBodyAccessDeniedDetail
	SetCode(v int64) *PredImageResponseBody
	GetCode() *int64
	SetMessage(v string) *PredImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *PredImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PredImageResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *PredImageResponseBody
	GetTaskId() *string
}

type PredImageResponseBody struct {
	AccessDeniedDetail *PredImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 1234****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PredImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PredImageResponseBody) GoString() string {
	return s.String()
}

func (s *PredImageResponseBody) GetAccessDeniedDetail() *PredImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PredImageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PredImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PredImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PredImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PredImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PredImageResponseBody) SetAccessDeniedDetail(v *PredImageResponseBodyAccessDeniedDetail) *PredImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PredImageResponseBody) SetCode(v int64) *PredImageResponseBody {
	s.Code = &v
	return s
}

func (s *PredImageResponseBody) SetMessage(v string) *PredImageResponseBody {
	s.Message = &v
	return s
}

func (s *PredImageResponseBody) SetRequestId(v string) *PredImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredImageResponseBody) SetSuccess(v bool) *PredImageResponseBody {
	s.Success = &v
	return s
}

func (s *PredImageResponseBody) SetTaskId(v string) *PredImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *PredImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PredImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PredImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PredImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PredImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PredImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PredImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
