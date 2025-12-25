// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PublishHotspotConfigResponseBodyAccessDeniedDetail) *PublishHotspotConfigResponseBody
	GetAccessDeniedDetail() *PublishHotspotConfigResponseBodyAccessDeniedDetail
	SetCode(v int64) *PublishHotspotConfigResponseBody
	GetCode() *int64
	SetMessage(v string) *PublishHotspotConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishHotspotConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishHotspotConfigResponseBody
	GetSuccess() *bool
}

type PublishHotspotConfigResponseBody struct {
	AccessDeniedDetail *PublishHotspotConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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
	// B28A2ECB-AB29-1E01****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishHotspotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigResponseBody) GetAccessDeniedDetail() *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PublishHotspotConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PublishHotspotConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishHotspotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishHotspotConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishHotspotConfigResponseBody) SetAccessDeniedDetail(v *PublishHotspotConfigResponseBodyAccessDeniedDetail) *PublishHotspotConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetCode(v int64) *PublishHotspotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetMessage(v string) *PublishHotspotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetRequestId(v string) *PublishHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) SetSuccess(v bool) *PublishHotspotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *PublishHotspotConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishHotspotConfigResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PublishHotspotConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PublishHotspotConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PublishHotspotConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
