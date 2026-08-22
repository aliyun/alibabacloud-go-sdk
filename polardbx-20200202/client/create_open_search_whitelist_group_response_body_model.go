// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchWhitelistGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *CreateOpenSearchWhitelistGroupResponseBody
	GetAccessDeniedDetail() *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail
	SetData(v *CreateOpenSearchWhitelistGroupResponseBodyData) *CreateOpenSearchWhitelistGroupResponseBody
	GetData() *CreateOpenSearchWhitelistGroupResponseBodyData
	SetRequestId(v string) *CreateOpenSearchWhitelistGroupResponseBody
	GetRequestId() *string
}

type CreateOpenSearchWhitelistGroupResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned result.
	Data *CreateOpenSearchWhitelistGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6352AC16-76BF-5135-B1EA-ED49293526E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpenSearchWhitelistGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) GetAccessDeniedDetail() *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) GetData() *CreateOpenSearchWhitelistGroupResponseBodyData {
	return s.Data
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) SetAccessDeniedDetail(v *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *CreateOpenSearchWhitelistGroupResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) SetData(v *CreateOpenSearchWhitelistGroupResponseBodyData) *CreateOpenSearchWhitelistGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) SetRequestId(v string) *CreateOpenSearchWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBody) Validate() error {
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

type CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The type of the authentication principal.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The diagnostic information.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// NoPermissionType
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateOpenSearchWhitelistGroupResponseBodyData struct {
	// The group ID, which is a globally unique identifier generated by the system for the group.
	//
	// example:
	//
	// 237509538
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// special
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateOpenSearchWhitelistGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchWhitelistGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyData) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyData) SetGroupId(v string) *CreateOpenSearchWhitelistGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyData) SetGroupName(v string) *CreateOpenSearchWhitelistGroupResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *CreateOpenSearchWhitelistGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
