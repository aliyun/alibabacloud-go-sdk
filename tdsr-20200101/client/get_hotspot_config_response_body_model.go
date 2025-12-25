// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetHotspotConfigResponseBodyAccessDeniedDetail) *GetHotspotConfigResponseBody
	GetAccessDeniedDetail() *GetHotspotConfigResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetHotspotConfigResponseBody
	GetCode() *int64
	SetData(v string) *GetHotspotConfigResponseBody
	GetData() *string
	SetMessage(v string) *GetHotspotConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotConfigResponseBody
	GetSuccess() *bool
}

type GetHotspotConfigResponseBody struct {
	AccessDeniedDetail *GetHotspotConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// config
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotspotConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponseBody) GetAccessDeniedDetail() *GetHotspotConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetHotspotConfigResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetHotspotConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *GetHotspotConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotConfigResponseBody) SetAccessDeniedDetail(v *GetHotspotConfigResponseBodyAccessDeniedDetail) *GetHotspotConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetHotspotConfigResponseBody) SetCode(v int64) *GetHotspotConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetData(v string) *GetHotspotConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetMessage(v string) *GetHotspotConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetRequestId(v string) *GetHotspotConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotConfigResponseBody) SetSuccess(v bool) *GetHotspotConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotConfigResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotspotConfigResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetHotspotConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetHotspotConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetHotspotConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
