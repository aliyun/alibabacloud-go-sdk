// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SaveHotspotConfigResponseBodyAccessDeniedDetail) *SaveHotspotConfigResponseBody
	GetAccessDeniedDetail() *SaveHotspotConfigResponseBodyAccessDeniedDetail
	SetErrMessage(v string) *SaveHotspotConfigResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SaveHotspotConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveHotspotConfigResponseBody
	GetSuccess() *bool
}

type SaveHotspotConfigResponseBody struct {
	AccessDeniedDetail *SaveHotspotConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	ErrMessage         *string                                          `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveHotspotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponseBody) GetAccessDeniedDetail() *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SaveHotspotConfigResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SaveHotspotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveHotspotConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveHotspotConfigResponseBody) SetAccessDeniedDetail(v *SaveHotspotConfigResponseBodyAccessDeniedDetail) *SaveHotspotConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetErrMessage(v string) *SaveHotspotConfigResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetRequestId(v string) *SaveHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) SetSuccess(v bool) *SaveHotspotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SaveHotspotConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveHotspotConfigResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SaveHotspotConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SaveHotspotConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SaveHotspotConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
