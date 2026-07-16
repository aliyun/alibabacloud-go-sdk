// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolardbxSupabaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) *CreatePolardbxSupabaseInstanceResponseBody
	GetAccessDeniedDetail() *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail
	SetData(v *CreatePolardbxSupabaseInstanceResponseBodyData) *CreatePolardbxSupabaseInstanceResponseBody
	GetData() *CreatePolardbxSupabaseInstanceResponseBodyData
	SetRequestId(v string) *CreatePolardbxSupabaseInstanceResponseBody
	GetRequestId() *string
}

type CreatePolardbxSupabaseInstanceResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The creation result.
	Data *CreatePolardbxSupabaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolardbxSupabaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePolardbxSupabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) GetAccessDeniedDetail() *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) GetData() *CreatePolardbxSupabaseInstanceResponseBodyData {
	return s.Data
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) SetAccessDeniedDetail(v *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) *CreatePolardbxSupabaseInstanceResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) SetData(v *CreatePolardbxSupabaseInstanceResponseBodyData) *CreatePolardbxSupabaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) SetRequestId(v string) *CreatePolardbxSupabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBody) Validate() error {
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

type CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail struct {
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
	// The authentication principal type.
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

func (s CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePolardbxSupabaseInstanceResponseBodyData struct {
	// The endpoint.
	//
	// example:
	//
	// test2.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxsp-xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 265325896860727
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The port.
	//
	// example:
	//
	// 3300
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreatePolardbxSupabaseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePolardbxSupabaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) SetConnectionString(v string) *CreatePolardbxSupabaseInstanceResponseBodyData {
	s.ConnectionString = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) SetDBInstanceId(v string) *CreatePolardbxSupabaseInstanceResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) SetOrderId(v int64) *CreatePolardbxSupabaseInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) SetPort(v int32) *CreatePolardbxSupabaseInstanceResponseBodyData {
	s.Port = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
