// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopySceneTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) *GetCopySceneTaskStatusResponseBody
	GetAccessDeniedDetail() *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetCopySceneTaskStatusResponseBody
	GetCode() *int64
	SetData(v *GetCopySceneTaskStatusResponseBodyData) *GetCopySceneTaskStatusResponseBody
	GetData() *GetCopySceneTaskStatusResponseBodyData
	SetMessage(v string) *GetCopySceneTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCopySceneTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCopySceneTaskStatusResponseBody
	GetSuccess() *bool
}

type GetCopySceneTaskStatusResponseBody struct {
	AccessDeniedDetail *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetCopySceneTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCopySceneTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCopySceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponseBody) GetAccessDeniedDetail() *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetCopySceneTaskStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetCopySceneTaskStatusResponseBody) GetData() *GetCopySceneTaskStatusResponseBodyData {
	return s.Data
}

func (s *GetCopySceneTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCopySceneTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCopySceneTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCopySceneTaskStatusResponseBody) SetAccessDeniedDetail(v *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) *GetCopySceneTaskStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetCode(v int64) *GetCopySceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetData(v *GetCopySceneTaskStatusResponseBodyData) *GetCopySceneTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetMessage(v string) *GetCopySceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetRequestId(v string) *GetCopySceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) SetSuccess(v bool) *GetCopySceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBody) Validate() error {
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

type GetCopySceneTaskStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetCopySceneTaskStatusResponseBodyData struct {
	// example:
	//
	// 100
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCopySceneTaskStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCopySceneTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusResponseBodyData) GetProgress() *int64 {
	return s.Progress
}

func (s *GetCopySceneTaskStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCopySceneTaskStatusResponseBodyData) SetProgress(v int64) *GetCopySceneTaskStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyData) SetStatus(v string) *GetCopySceneTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCopySceneTaskStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
