// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContextDBConfigResponseBodyAccessDeniedDetail) *DescribeContextDBConfigResponseBody
	GetAccessDeniedDetail() *DescribeContextDBConfigResponseBodyAccessDeniedDetail
	SetData(v *DescribeContextDBConfigResponseBodyData) *DescribeContextDBConfigResponseBody
	GetData() *DescribeContextDBConfigResponseBodyData
	SetRequestId(v string) *DescribeContextDBConfigResponseBody
	GetRequestId() *string
}

type DescribeContextDBConfigResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeContextDBConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *DescribeContextDBConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContextDBConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContextDBConfigResponseBody) GetAccessDeniedDetail() *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContextDBConfigResponseBody) GetData() *DescribeContextDBConfigResponseBodyData {
	return s.Data
}

func (s *DescribeContextDBConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContextDBConfigResponseBody) SetAccessDeniedDetail(v *DescribeContextDBConfigResponseBodyAccessDeniedDetail) *DescribeContextDBConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContextDBConfigResponseBody) SetData(v *DescribeContextDBConfigResponseBodyData) *DescribeContextDBConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContextDBConfigResponseBody) SetRequestId(v string) *DescribeContextDBConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContextDBConfigResponseBody) Validate() error {
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

type DescribeContextDBConfigResponseBodyAccessDeniedDetail struct {
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

func (s DescribeContextDBConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContextDBConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContextDBConfigResponseBodyData struct {
	// The admin key of the context service instance.
	//
	// example:
	//
	// ctx-admin-***
	ContextDBAdminKey *string `json:"ContextDBAdminKey,omitempty" xml:"ContextDBAdminKey,omitempty"`
	// The name of the context service instance.
	//
	// example:
	//
	// pxt-*********
	ContextDBInstanceName *string `json:"ContextDBInstanceName,omitempty" xml:"ContextDBInstanceName,omitempty"`
	// The owner key of the context service instance.
	//
	// example:
	//
	// ctx-***
	ContextDBOwnerKey *string `json:"ContextDBOwnerKey,omitempty" xml:"ContextDBOwnerKey,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the PolarDB-X Search instance.
	//
	// example:
	//
	// pxs-*********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
}

func (s DescribeContextDBConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContextDBConfigResponseBodyData) GetContextDBAdminKey() *string {
	return s.ContextDBAdminKey
}

func (s *DescribeContextDBConfigResponseBodyData) GetContextDBInstanceName() *string {
	return s.ContextDBInstanceName
}

func (s *DescribeContextDBConfigResponseBodyData) GetContextDBOwnerKey() *string {
	return s.ContextDBOwnerKey
}

func (s *DescribeContextDBConfigResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBConfigResponseBodyData) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *DescribeContextDBConfigResponseBodyData) SetContextDBAdminKey(v string) *DescribeContextDBConfigResponseBodyData {
	s.ContextDBAdminKey = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyData) SetContextDBInstanceName(v string) *DescribeContextDBConfigResponseBodyData {
	s.ContextDBInstanceName = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyData) SetContextDBOwnerKey(v string) *DescribeContextDBConfigResponseBodyData {
	s.ContextDBOwnerKey = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyData) SetDBInstanceName(v string) *DescribeContextDBConfigResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyData) SetOpenSearchInstanceName(v string) *DescribeContextDBConfigResponseBodyData {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *DescribeContextDBConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
