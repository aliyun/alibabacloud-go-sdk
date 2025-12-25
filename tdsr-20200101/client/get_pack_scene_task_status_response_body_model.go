// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackSceneTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) *GetPackSceneTaskStatusResponseBody
	GetAccessDeniedDetail() *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetPackSceneTaskStatusResponseBody
	GetCode() *int64
	SetData(v *GetPackSceneTaskStatusResponseBodyData) *GetPackSceneTaskStatusResponseBody
	GetData() *GetPackSceneTaskStatusResponseBodyData
	SetMessage(v string) *GetPackSceneTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPackSceneTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPackSceneTaskStatusResponseBody
	GetSuccess() *bool
}

type GetPackSceneTaskStatusResponseBody struct {
	AccessDeniedDetail *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPackSceneTaskStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A8CD0AD9-8A92-455A-A984-A7E4B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPackSceneTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPackSceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponseBody) GetAccessDeniedDetail() *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetPackSceneTaskStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetPackSceneTaskStatusResponseBody) GetData() *GetPackSceneTaskStatusResponseBodyData {
	return s.Data
}

func (s *GetPackSceneTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPackSceneTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPackSceneTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPackSceneTaskStatusResponseBody) SetAccessDeniedDetail(v *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) *GetPackSceneTaskStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetCode(v int64) *GetPackSceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetData(v *GetPackSceneTaskStatusResponseBodyData) *GetPackSceneTaskStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetMessage(v string) *GetPackSceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetRequestId(v string) *GetPackSceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) SetSuccess(v bool) *GetPackSceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBody) Validate() error {
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

type GetPackSceneTaskStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetPackSceneTaskStatusResponseBodyData struct {
	// example:
	//
	// 100
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPackSceneTaskStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPackSceneTaskStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusResponseBodyData) GetProgress() *int64 {
	return s.Progress
}

func (s *GetPackSceneTaskStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPackSceneTaskStatusResponseBodyData) SetProgress(v int64) *GetPackSceneTaskStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyData) SetStatus(v string) *GetPackSceneTaskStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPackSceneTaskStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
