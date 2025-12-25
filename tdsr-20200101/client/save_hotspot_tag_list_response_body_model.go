// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveHotspotTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *SaveHotspotTagListResponseBodyAccessDeniedDetail) *SaveHotspotTagListResponseBody
	GetAccessDeniedDetail() *SaveHotspotTagListResponseBodyAccessDeniedDetail
	SetCode(v int64) *SaveHotspotTagListResponseBody
	GetCode() *int64
	SetMessage(v string) *SaveHotspotTagListResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveHotspotTagListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveHotspotTagListResponseBody
	GetSuccess() *bool
}

type SaveHotspotTagListResponseBody struct {
	AccessDeniedDetail *SaveHotspotTagListResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
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

func (s SaveHotspotTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagListResponseBody) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListResponseBody) GetAccessDeniedDetail() *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *SaveHotspotTagListResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SaveHotspotTagListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveHotspotTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveHotspotTagListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveHotspotTagListResponseBody) SetAccessDeniedDetail(v *SaveHotspotTagListResponseBodyAccessDeniedDetail) *SaveHotspotTagListResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetCode(v int64) *SaveHotspotTagListResponseBody {
	s.Code = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetMessage(v string) *SaveHotspotTagListResponseBody {
	s.Message = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetRequestId(v string) *SaveHotspotTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) SetSuccess(v bool) *SaveHotspotTagListResponseBody {
	s.Success = &v
	return s
}

func (s *SaveHotspotTagListResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SaveHotspotTagListResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s SaveHotspotTagListResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s SaveHotspotTagListResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetAuthAction(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) SetPolicyType(v string) *SaveHotspotTagListResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *SaveHotspotTagListResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
