// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchWhitelistGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *ModifyOpenSearchWhitelistGroupResponseBody
	GetAccessDeniedDetail() *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail
	SetData(v *ModifyOpenSearchWhitelistGroupResponseBodyData) *ModifyOpenSearchWhitelistGroupResponseBody
	GetData() *ModifyOpenSearchWhitelistGroupResponseBodyData
	SetRequestId(v string) *ModifyOpenSearchWhitelistGroupResponseBody
	GetRequestId() *string
}

type ModifyOpenSearchWhitelistGroupResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The creation result.
	Data *ModifyOpenSearchWhitelistGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOpenSearchWhitelistGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchWhitelistGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) GetAccessDeniedDetail() *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) GetData() *ModifyOpenSearchWhitelistGroupResponseBodyData {
	return s.Data
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) SetAccessDeniedDetail(v *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) *ModifyOpenSearchWhitelistGroupResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) SetData(v *ModifyOpenSearchWhitelistGroupResponseBodyData) *ModifyOpenSearchWhitelistGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) SetRequestId(v string) *ModifyOpenSearchWhitelistGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBody) Validate() error {
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

type ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail struct {
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
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the permission denial.
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

func (s ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ModifyOpenSearchWhitelistGroupResponseBodyData struct {
	// The group ID, which is a globally unique identifier generated by the system for the group.
	//
	// example:
	//
	// g-00mzurifez86htk8fn90
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the task was last updated, in timestamp format.
	//
	// example:
	//
	// 2025-09-17T02:27:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ModifyOpenSearchWhitelistGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchWhitelistGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyData) SetGroupId(v string) *ModifyOpenSearchWhitelistGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyData) SetUpdateTime(v string) *ModifyOpenSearchWhitelistGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
