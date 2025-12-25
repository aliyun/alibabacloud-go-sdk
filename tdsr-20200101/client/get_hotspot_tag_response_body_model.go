// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetHotspotTagResponseBodyAccessDeniedDetail) *GetHotspotTagResponseBody
	GetAccessDeniedDetail() *GetHotspotTagResponseBodyAccessDeniedDetail
	SetData(v string) *GetHotspotTagResponseBody
	GetData() *string
	SetErrMessage(v string) *GetHotspotTagResponseBody
	GetErrMessage() *string
	SetObjectString(v string) *GetHotspotTagResponseBody
	GetObjectString() *string
	SetRequestId(v string) *GetHotspotTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotTagResponseBody
	GetSuccess() *bool
}

type GetHotspotTagResponseBody struct {
	AccessDeniedDetail *GetHotspotTagResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// {"watermarkImg":[],"enabledTitleTag":0}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// "{}"
	ObjectString *string `json:"ObjectString,omitempty" xml:"ObjectString,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponseBody) GetAccessDeniedDetail() *GetHotspotTagResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetHotspotTagResponseBody) GetData() *string {
	return s.Data
}

func (s *GetHotspotTagResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetHotspotTagResponseBody) GetObjectString() *string {
	return s.ObjectString
}

func (s *GetHotspotTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotTagResponseBody) SetAccessDeniedDetail(v *GetHotspotTagResponseBodyAccessDeniedDetail) *GetHotspotTagResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetHotspotTagResponseBody) SetData(v string) *GetHotspotTagResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetErrMessage(v string) *GetHotspotTagResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetObjectString(v string) *GetHotspotTagResponseBody {
	s.ObjectString = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetRequestId(v string) *GetHotspotTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotTagResponseBody) SetSuccess(v bool) *GetHotspotTagResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotTagResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotspotTagResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetHotspotTagResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTagResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetHotspotTagResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetHotspotTagResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
