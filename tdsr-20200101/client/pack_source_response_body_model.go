// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PackSourceResponseBodyAccessDeniedDetail) *PackSourceResponseBody
	GetAccessDeniedDetail() *PackSourceResponseBodyAccessDeniedDetail
	SetCode(v int64) *PackSourceResponseBody
	GetCode() *int64
	SetData(v *PackSourceResponseBodyData) *PackSourceResponseBody
	GetData() *PackSourceResponseBodyData
	SetMessage(v string) *PackSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *PackSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PackSourceResponseBody
	GetSuccess() *bool
}

type PackSourceResponseBody struct {
	AccessDeniedDetail *PackSourceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PackSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PackSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PackSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PackSourceResponseBody) GetAccessDeniedDetail() *PackSourceResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PackSourceResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PackSourceResponseBody) GetData() *PackSourceResponseBodyData {
	return s.Data
}

func (s *PackSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PackSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PackSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PackSourceResponseBody) SetAccessDeniedDetail(v *PackSourceResponseBodyAccessDeniedDetail) *PackSourceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PackSourceResponseBody) SetCode(v int64) *PackSourceResponseBody {
	s.Code = &v
	return s
}

func (s *PackSourceResponseBody) SetData(v *PackSourceResponseBodyData) *PackSourceResponseBody {
	s.Data = v
	return s
}

func (s *PackSourceResponseBody) SetMessage(v string) *PackSourceResponseBody {
	s.Message = &v
	return s
}

func (s *PackSourceResponseBody) SetRequestId(v string) *PackSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PackSourceResponseBody) SetSuccess(v bool) *PackSourceResponseBody {
	s.Success = &v
	return s
}

func (s *PackSourceResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PackSourceResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PackSourceResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PackSourceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PackSourceResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PackSourceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PackSourceResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type PackSourceResponseBodyData struct {
	// example:
	//
	// hjsyuyiuwe7wehg****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PackSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PackSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *PackSourceResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PackSourceResponseBodyData) SetTaskId(v string) *PackSourceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PackSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
