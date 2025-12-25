// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempPreviewStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *TempPreviewStatusResponseBodyAccessDeniedDetail) *TempPreviewStatusResponseBody
	GetAccessDeniedDetail() *TempPreviewStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *TempPreviewStatusResponseBody
	GetCode() *int64
	SetMessage(v string) *TempPreviewStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *TempPreviewStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *TempPreviewStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *TempPreviewStatusResponseBody
	GetSuccess() *bool
}

type TempPreviewStatusResponseBody struct {
	AccessDeniedDetail *TempPreviewStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TempPreviewStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponseBody) GetAccessDeniedDetail() *TempPreviewStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *TempPreviewStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *TempPreviewStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TempPreviewStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempPreviewStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *TempPreviewStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TempPreviewStatusResponseBody) SetAccessDeniedDetail(v *TempPreviewStatusResponseBodyAccessDeniedDetail) *TempPreviewStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *TempPreviewStatusResponseBody) SetCode(v int64) *TempPreviewStatusResponseBody {
	s.Code = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetMessage(v string) *TempPreviewStatusResponseBody {
	s.Message = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetRequestId(v string) *TempPreviewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetStatus(v string) *TempPreviewStatusResponseBody {
	s.Status = &v
	return s
}

func (s *TempPreviewStatusResponseBody) SetSuccess(v bool) *TempPreviewStatusResponseBody {
	s.Success = &v
	return s
}

func (s *TempPreviewStatusResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TempPreviewStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s TempPreviewStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s TempPreviewStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *TempPreviewStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *TempPreviewStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
