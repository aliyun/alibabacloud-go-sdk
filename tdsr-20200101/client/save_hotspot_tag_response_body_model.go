// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SaveHotspotTagResponseBodyAccessDeniedDetail) *SaveHotspotTagResponseBody
	GetAccessDeniedDetail() *SaveHotspotTagResponseBodyAccessDeniedDetail
	SetErrMessage(v string) *SaveHotspotTagResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SaveHotspotTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveHotspotTagResponseBody
	GetSuccess() *bool
}

type SaveHotspotTagResponseBody struct {
	AccessDeniedDetail *SaveHotspotTagResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	ErrMessage         *string                                       `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveHotspotTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponseBody) GetAccessDeniedDetail() *SaveHotspotTagResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SaveHotspotTagResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SaveHotspotTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveHotspotTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveHotspotTagResponseBody) SetAccessDeniedDetail(v *SaveHotspotTagResponseBodyAccessDeniedDetail) *SaveHotspotTagResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SaveHotspotTagResponseBody) SetErrMessage(v string) *SaveHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetRequestId(v string) *SaveHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagResponseBody) SetSuccess(v bool) *SaveHotspotTagResponseBody {
	s.Success = &v
	return s
}

func (s *SaveHotspotTagResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveHotspotTagResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SaveHotspotTagResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SaveHotspotTagResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SaveHotspotTagResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
