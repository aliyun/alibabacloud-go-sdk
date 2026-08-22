// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0ConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeContext0ConfigResponseBodyAccessDeniedDetail) *DescribeContext0ConfigResponseBody
	GetAccessDeniedDetail() *DescribeContext0ConfigResponseBodyAccessDeniedDetail
	SetData(v *DescribeContext0ConfigResponseBodyData) *DescribeContext0ConfigResponseBody
	GetData() *DescribeContext0ConfigResponseBodyData
	SetRequestId(v string) *DescribeContext0ConfigResponseBody
	GetRequestId() *string
}

type DescribeContext0ConfigResponseBody struct {
	AccessDeniedDetail *DescribeContext0ConfigResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	Data               *DescribeContext0ConfigResponseBodyData               `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContext0ConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0ConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContext0ConfigResponseBody) GetAccessDeniedDetail() *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeContext0ConfigResponseBody) GetData() *DescribeContext0ConfigResponseBodyData {
	return s.Data
}

func (s *DescribeContext0ConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContext0ConfigResponseBody) SetAccessDeniedDetail(v *DescribeContext0ConfigResponseBodyAccessDeniedDetail) *DescribeContext0ConfigResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeContext0ConfigResponseBody) SetData(v *DescribeContext0ConfigResponseBodyData) *DescribeContext0ConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContext0ConfigResponseBody) SetRequestId(v string) *DescribeContext0ConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContext0ConfigResponseBody) Validate() error {
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

type DescribeContext0ConfigResponseBodyAccessDeniedDetail struct {
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

func (s DescribeContext0ConfigResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0ConfigResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeContext0ConfigResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeContext0ConfigResponseBodyData struct {
	// example:
	//
	// admin-key-example
	Context0AdminKey *string `json:"Context0AdminKey,omitempty" xml:"Context0AdminKey,omitempty"`
	// example:
	//
	// context0-example
	Context0InstanceName *string `json:"Context0InstanceName,omitempty" xml:"Context0InstanceName,omitempty"`
	// example:
	//
	// owner-key-example
	Context0OwnerKey *string `json:"Context0OwnerKey,omitempty" xml:"Context0OwnerKey,omitempty"`
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// pxs-*********
	OpenSearchInstanceName *string `json:"OpenSearchInstanceName,omitempty" xml:"OpenSearchInstanceName,omitempty"`
}

func (s DescribeContext0ConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0ConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContext0ConfigResponseBodyData) GetContext0AdminKey() *string {
	return s.Context0AdminKey
}

func (s *DescribeContext0ConfigResponseBodyData) GetContext0InstanceName() *string {
	return s.Context0InstanceName
}

func (s *DescribeContext0ConfigResponseBodyData) GetContext0OwnerKey() *string {
	return s.Context0OwnerKey
}

func (s *DescribeContext0ConfigResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0ConfigResponseBodyData) GetOpenSearchInstanceName() *string {
	return s.OpenSearchInstanceName
}

func (s *DescribeContext0ConfigResponseBodyData) SetContext0AdminKey(v string) *DescribeContext0ConfigResponseBodyData {
	s.Context0AdminKey = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyData) SetContext0InstanceName(v string) *DescribeContext0ConfigResponseBodyData {
	s.Context0InstanceName = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyData) SetContext0OwnerKey(v string) *DescribeContext0ConfigResponseBodyData {
	s.Context0OwnerKey = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyData) SetDBInstanceName(v string) *DescribeContext0ConfigResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyData) SetOpenSearchInstanceName(v string) *DescribeContext0ConfigResponseBodyData {
	s.OpenSearchInstanceName = &v
	return s
}

func (s *DescribeContext0ConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
