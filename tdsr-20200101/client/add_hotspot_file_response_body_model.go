// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHotspotFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddHotspotFileResponseBodyAccessDeniedDetail) *AddHotspotFileResponseBody
	GetAccessDeniedDetail() *AddHotspotFileResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddHotspotFileResponseBody
	GetCode() *int64
	SetData(v map[string]interface{}) *AddHotspotFileResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *AddHotspotFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddHotspotFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddHotspotFileResponseBody
	GetSuccess() *bool
}

type AddHotspotFileResponseBody struct {
	AccessDeniedDetail *AddHotspotFileResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Code               *int64                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{}                        `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddHotspotFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddHotspotFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddHotspotFileResponseBody) GetAccessDeniedDetail() *AddHotspotFileResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddHotspotFileResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddHotspotFileResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddHotspotFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddHotspotFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddHotspotFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddHotspotFileResponseBody) SetAccessDeniedDetail(v *AddHotspotFileResponseBodyAccessDeniedDetail) *AddHotspotFileResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddHotspotFileResponseBody) SetCode(v int64) *AddHotspotFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetData(v map[string]interface{}) *AddHotspotFileResponseBody {
	s.Data = v
	return s
}

func (s *AddHotspotFileResponseBody) SetMessage(v string) *AddHotspotFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetRequestId(v string) *AddHotspotFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHotspotFileResponseBody) SetSuccess(v bool) *AddHotspotFileResponseBody {
	s.Success = &v
	return s
}

func (s *AddHotspotFileResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddHotspotFileResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddHotspotFileResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddHotspotFileResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddHotspotFileResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddHotspotFileResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
