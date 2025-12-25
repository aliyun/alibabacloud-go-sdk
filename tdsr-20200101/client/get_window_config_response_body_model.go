// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWindowConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetWindowConfigResponseBodyAccessDeniedDetail) *GetWindowConfigResponseBody
	GetAccessDeniedDetail() *GetWindowConfigResponseBodyAccessDeniedDetail
	SetData(v map[string]interface{}) *GetWindowConfigResponseBody
	GetData() map[string]interface{}
	SetErrMessage(v string) *GetWindowConfigResponseBody
	GetErrMessage() *string
	SetObjectString(v string) *GetWindowConfigResponseBody
	GetObjectString() *string
	SetRequestId(v string) *GetWindowConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWindowConfigResponseBody
	GetSuccess() *bool
}

type GetWindowConfigResponseBody struct {
	AccessDeniedDetail *GetWindowConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 2345****
	Data       map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// "{}"
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWindowConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWindowConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponseBody) GetAccessDeniedDetail() *GetWindowConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetWindowConfigResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetWindowConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetWindowConfigResponseBody) GetObjectString() *string {
	return s.ObjectString
}

func (s *GetWindowConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWindowConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWindowConfigResponseBody) SetAccessDeniedDetail(v *GetWindowConfigResponseBodyAccessDeniedDetail) *GetWindowConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetWindowConfigResponseBody) SetData(v map[string]interface{}) *GetWindowConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetWindowConfigResponseBody) SetErrMessage(v string) *GetWindowConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetObjectString(v string) *GetWindowConfigResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetRequestId(v string) *GetWindowConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWindowConfigResponseBody) SetSuccess(v bool) *GetWindowConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetWindowConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWindowConfigResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetWindowConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetWindowConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetWindowConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetWindowConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
