// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginLayoutDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetOriginLayoutDataResponseBodyAccessDeniedDetail) *GetOriginLayoutDataResponseBody
	GetAccessDeniedDetail() *GetOriginLayoutDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetOriginLayoutDataResponseBody
	GetCode() *int64
	SetData(v string) *GetOriginLayoutDataResponseBody
	GetData() *string
	SetMessage(v string) *GetOriginLayoutDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOriginLayoutDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOriginLayoutDataResponseBody
	GetSuccess() *bool
}

type GetOriginLayoutDataResponseBody struct {
	AccessDeniedDetail *GetOriginLayoutDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
}

func (s GetOriginLayoutDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponseBody) GetAccessDeniedDetail() *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetOriginLayoutDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetOriginLayoutDataResponseBody) GetData() *string {
	return s.Data
}

func (s *GetOriginLayoutDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOriginLayoutDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginLayoutDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOriginLayoutDataResponseBody) SetAccessDeniedDetail(v *GetOriginLayoutDataResponseBodyAccessDeniedDetail) *GetOriginLayoutDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetCode(v int64) *GetOriginLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetData(v string) *GetOriginLayoutDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetMessage(v string) *GetOriginLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetRequestId(v string) *GetOriginLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) SetSuccess(v bool) *GetOriginLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetOriginLayoutDataResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOriginLayoutDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetOriginLayoutDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOriginLayoutDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetOriginLayoutDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetOriginLayoutDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
