// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourcePackStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetSourcePackStatusResponseBodyAccessDeniedDetail) *GetSourcePackStatusResponseBody
	GetAccessDeniedDetail() *GetSourcePackStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetSourcePackStatusResponseBody
	GetCode() *int64
	SetData(v *GetSourcePackStatusResponseBodyData) *GetSourcePackStatusResponseBody
	GetData() *GetSourcePackStatusResponseBodyData
	SetMessage(v string) *GetSourcePackStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSourcePackStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSourcePackStatusResponseBody
	GetSuccess() *bool
	SetUrl(v string) *GetSourcePackStatusResponseBody
	GetUrl() *string
}

type GetSourcePackStatusResponseBody struct {
	AccessDeniedDetail *GetSourcePackStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSourcePackStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetSourcePackStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourcePackStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponseBody) GetAccessDeniedDetail() *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetSourcePackStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSourcePackStatusResponseBody) GetData() *GetSourcePackStatusResponseBodyData {
	return s.Data
}

func (s *GetSourcePackStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSourcePackStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourcePackStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSourcePackStatusResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetSourcePackStatusResponseBody) SetAccessDeniedDetail(v *GetSourcePackStatusResponseBodyAccessDeniedDetail) *GetSourcePackStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetCode(v int64) *GetSourcePackStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetData(v *GetSourcePackStatusResponseBodyData) *GetSourcePackStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetMessage(v string) *GetSourcePackStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetRequestId(v string) *GetSourcePackStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetSuccess(v bool) *GetSourcePackStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) SetUrl(v string) *GetSourcePackStatusResponseBody {
	s.Url = &v
	return s
}

func (s *GetSourcePackStatusResponseBody) Validate() error {
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

type GetSourcePackStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetSourcePackStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSourcePackStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetSourcePackStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetSourcePackStatusResponseBodyData struct {
	// example:
	//
	// 100
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSourcePackStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSourcePackStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusResponseBodyData) GetProgress() *int64 {
	return s.Progress
}

func (s *GetSourcePackStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSourcePackStatusResponseBodyData) SetProgress(v int64) *GetSourcePackStatusResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyData) SetStatus(v string) *GetSourcePackStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSourcePackStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
