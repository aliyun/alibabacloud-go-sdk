// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PublishHotspotResponseBodyAccessDeniedDetail) *PublishHotspotResponseBody
	GetAccessDeniedDetail() *PublishHotspotResponseBodyAccessDeniedDetail
	SetData(v map[string]interface{}) *PublishHotspotResponseBody
	GetData() map[string]interface{}
	SetErrMessage(v string) *PublishHotspotResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *PublishHotspotResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishHotspotResponseBody
	GetSuccess() *bool
}

type PublishHotspotResponseBody struct {
	AccessDeniedDetail *PublishHotspotResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 2345****
	Data       map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrMessage *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishHotspotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotResponseBody) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponseBody) GetAccessDeniedDetail() *PublishHotspotResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PublishHotspotResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *PublishHotspotResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PublishHotspotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishHotspotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishHotspotResponseBody) SetAccessDeniedDetail(v *PublishHotspotResponseBodyAccessDeniedDetail) *PublishHotspotResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PublishHotspotResponseBody) SetData(v map[string]interface{}) *PublishHotspotResponseBody {
	s.Data = v
	return s
}

func (s *PublishHotspotResponseBody) SetErrMessage(v string) *PublishHotspotResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PublishHotspotResponseBody) SetRequestId(v string) *PublishHotspotResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotResponseBody) SetSuccess(v bool) *PublishHotspotResponseBody {
	s.Success = &v
	return s
}

func (s *PublishHotspotResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishHotspotResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PublishHotspotResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PublishHotspotResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PublishHotspotResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
