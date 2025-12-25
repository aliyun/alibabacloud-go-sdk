// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelBuildResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *LabelBuildResponseBodyAccessDeniedDetail) *LabelBuildResponseBody
	GetAccessDeniedDetail() *LabelBuildResponseBodyAccessDeniedDetail
	SetCode(v int64) *LabelBuildResponseBody
	GetCode() *int64
	SetMessage(v string) *LabelBuildResponseBody
	GetMessage() *string
	SetRequestId(v string) *LabelBuildResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *LabelBuildResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *LabelBuildResponseBody
	GetTaskId() *string
}

type LabelBuildResponseBody struct {
	AccessDeniedDetail *LabelBuildResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// 234****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s LabelBuildResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LabelBuildResponseBody) GoString() string {
	return s.String()
}

func (s *LabelBuildResponseBody) GetAccessDeniedDetail() *LabelBuildResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *LabelBuildResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *LabelBuildResponseBody) GetMessage() *string {
	return s.Message
}

func (s *LabelBuildResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LabelBuildResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *LabelBuildResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *LabelBuildResponseBody) SetAccessDeniedDetail(v *LabelBuildResponseBodyAccessDeniedDetail) *LabelBuildResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *LabelBuildResponseBody) SetCode(v int64) *LabelBuildResponseBody {
	s.Code = &v
	return s
}

func (s *LabelBuildResponseBody) SetMessage(v string) *LabelBuildResponseBody {
	s.Message = &v
	return s
}

func (s *LabelBuildResponseBody) SetRequestId(v string) *LabelBuildResponseBody {
	s.RequestId = &v
	return s
}

func (s *LabelBuildResponseBody) SetSuccess(v bool) *LabelBuildResponseBody {
	s.Success = &v
	return s
}

func (s *LabelBuildResponseBody) SetTaskId(v string) *LabelBuildResponseBody {
	s.TaskId = &v
	return s
}

func (s *LabelBuildResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LabelBuildResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s LabelBuildResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s LabelBuildResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetAuthAction(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) SetPolicyType(v string) *LabelBuildResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *LabelBuildResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
