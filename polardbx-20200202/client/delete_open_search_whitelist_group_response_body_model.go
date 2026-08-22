// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSearchWhitelistGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *DeleteOpenSearchWhitelistGroupResponseBody
	GetAccessDeniedDetail() *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail
	SetData(v *DeleteOpenSearchWhitelistGroupResponseBodyData) *DeleteOpenSearchWhitelistGroupResponseBody
	GetData() *DeleteOpenSearchWhitelistGroupResponseBodyData
	SetRequestId(v string) *DeleteOpenSearchWhitelistGroupResponseBody
	GetRequestId() *string
}

type DeleteOpenSearchWhitelistGroupResponseBody struct {
	AccessDeniedDetail *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *DeleteOpenSearchWhitelistGroupResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// B87E2AB3-B7C9-4394-9160-7F639F732031
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOpenSearchWhitelistGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) GetAccessDeniedDetail() *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) GetData() *DeleteOpenSearchWhitelistGroupResponseBodyData {
	return s.Data
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) SetAccessDeniedDetail(v *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *DeleteOpenSearchWhitelistGroupResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) SetData(v *DeleteOpenSearchWhitelistGroupResponseBodyData) *DeleteOpenSearchWhitelistGroupResponseBody {
	s.Data = v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) SetRequestId(v string) *DeleteOpenSearchWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteOpenSearchWhitelistGroupResponseBodyData struct {
	// example:
	//
	// get app list success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteOpenSearchWhitelistGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSearchWhitelistGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyData) SetMessage(v string) *DeleteOpenSearchWhitelistGroupResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteOpenSearchWhitelistGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
