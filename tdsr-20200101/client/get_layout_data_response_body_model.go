// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayoutDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetLayoutDataResponseBodyAccessDeniedDetail) *GetLayoutDataResponseBody
	GetAccessDeniedDetail() *GetLayoutDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetLayoutDataResponseBody
	GetCode() *int64
	SetData(v string) *GetLayoutDataResponseBody
	GetData() *string
	SetMessage(v string) *GetLayoutDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLayoutDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLayoutDataResponseBody
	GetSuccess() *bool
}

type GetLayoutDataResponseBody struct {
	AccessDeniedDetail *GetLayoutDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s GetLayoutDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLayoutDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponseBody) GetAccessDeniedDetail() *GetLayoutDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetLayoutDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetLayoutDataResponseBody) GetData() *string {
	return s.Data
}

func (s *GetLayoutDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLayoutDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLayoutDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLayoutDataResponseBody) SetAccessDeniedDetail(v *GetLayoutDataResponseBodyAccessDeniedDetail) *GetLayoutDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetLayoutDataResponseBody) SetCode(v int64) *GetLayoutDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetData(v string) *GetLayoutDataResponseBody {
	s.Data = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetMessage(v string) *GetLayoutDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetRequestId(v string) *GetLayoutDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayoutDataResponseBody) SetSuccess(v bool) *GetLayoutDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetLayoutDataResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLayoutDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetLayoutDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetLayoutDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetLayoutDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetLayoutDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
