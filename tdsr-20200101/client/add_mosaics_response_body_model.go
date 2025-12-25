// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMosaicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddMosaicsResponseBodyAccessDeniedDetail) *AddMosaicsResponseBody
	GetAccessDeniedDetail() *AddMosaicsResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddMosaicsResponseBody
	GetCode() *int64
	SetMessage(v string) *AddMosaicsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddMosaicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddMosaicsResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *AddMosaicsResponseBody
	GetTaskId() *string
}

type AddMosaicsResponseBody struct {
	AccessDeniedDetail *AddMosaicsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3BCAD49D-2AC1-13EB-AC19-8C7A46C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// OvFuuwhfoAX8uIpxC/GJ****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AddMosaicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMosaicsResponseBody) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponseBody) GetAccessDeniedDetail() *AddMosaicsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddMosaicsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddMosaicsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddMosaicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMosaicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddMosaicsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AddMosaicsResponseBody) SetAccessDeniedDetail(v *AddMosaicsResponseBodyAccessDeniedDetail) *AddMosaicsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddMosaicsResponseBody) SetCode(v int64) *AddMosaicsResponseBody {
	s.Code = &v
	return s
}

func (s *AddMosaicsResponseBody) SetMessage(v string) *AddMosaicsResponseBody {
	s.Message = &v
	return s
}

func (s *AddMosaicsResponseBody) SetRequestId(v string) *AddMosaicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMosaicsResponseBody) SetSuccess(v bool) *AddMosaicsResponseBody {
	s.Success = &v
	return s
}

func (s *AddMosaicsResponseBody) SetTaskId(v string) *AddMosaicsResponseBody {
	s.TaskId = &v
	return s
}

func (s *AddMosaicsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMosaicsResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddMosaicsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddMosaicsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddMosaicsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddMosaicsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
