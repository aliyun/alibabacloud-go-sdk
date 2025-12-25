// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PackSceneResponseBodyAccessDeniedDetail) *PackSceneResponseBody
	GetAccessDeniedDetail() *PackSceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *PackSceneResponseBody
	GetCode() *int64
	SetData(v *PackSceneResponseBodyData) *PackSceneResponseBody
	GetData() *PackSceneResponseBodyData
	SetMessage(v string) *PackSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *PackSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PackSceneResponseBody
	GetSuccess() *bool
}

type PackSceneResponseBody struct {
	AccessDeniedDetail *PackSceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PackSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B76****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PackSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PackSceneResponseBody) GoString() string {
	return s.String()
}

func (s *PackSceneResponseBody) GetAccessDeniedDetail() *PackSceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PackSceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PackSceneResponseBody) GetData() *PackSceneResponseBodyData {
	return s.Data
}

func (s *PackSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PackSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PackSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PackSceneResponseBody) SetAccessDeniedDetail(v *PackSceneResponseBodyAccessDeniedDetail) *PackSceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PackSceneResponseBody) SetCode(v int64) *PackSceneResponseBody {
	s.Code = &v
	return s
}

func (s *PackSceneResponseBody) SetData(v *PackSceneResponseBodyData) *PackSceneResponseBody {
	s.Data = v
	return s
}

func (s *PackSceneResponseBody) SetMessage(v string) *PackSceneResponseBody {
	s.Message = &v
	return s
}

func (s *PackSceneResponseBody) SetRequestId(v string) *PackSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *PackSceneResponseBody) SetSuccess(v bool) *PackSceneResponseBody {
	s.Success = &v
	return s
}

func (s *PackSceneResponseBody) Validate() error {
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

type PackSceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PackSceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PackSceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PackSceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PackSceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PackSceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type PackSceneResponseBodyData struct {
	// example:
	//
	// hjsyuyiuwe7wehg****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PackSceneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PackSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *PackSceneResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PackSceneResponseBodyData) SetTaskId(v string) *PackSceneResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PackSceneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
