// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRoomPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *AddRoomPlanResponseBodyAccessDeniedDetail) *AddRoomPlanResponseBody
	GetAccessDeniedDetail() *AddRoomPlanResponseBodyAccessDeniedDetail
	SetCode(v int64) *AddRoomPlanResponseBody
	GetCode() *int64
	SetData(v *AddRoomPlanResponseBodyData) *AddRoomPlanResponseBody
	GetData() *AddRoomPlanResponseBodyData
	SetMessage(v string) *AddRoomPlanResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddRoomPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRoomPlanResponseBody
	GetSuccess() *bool
}

type AddRoomPlanResponseBody struct {
	AccessDeniedDetail *AddRoomPlanResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddRoomPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRoomPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRoomPlanResponseBody) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponseBody) GetAccessDeniedDetail() *AddRoomPlanResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *AddRoomPlanResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AddRoomPlanResponseBody) GetData() *AddRoomPlanResponseBodyData {
	return s.Data
}

func (s *AddRoomPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRoomPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRoomPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRoomPlanResponseBody) SetAccessDeniedDetail(v *AddRoomPlanResponseBodyAccessDeniedDetail) *AddRoomPlanResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *AddRoomPlanResponseBody) SetCode(v int64) *AddRoomPlanResponseBody {
	s.Code = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetData(v *AddRoomPlanResponseBodyData) *AddRoomPlanResponseBody {
	s.Data = v
	return s
}

func (s *AddRoomPlanResponseBody) SetMessage(v string) *AddRoomPlanResponseBody {
	s.Message = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetRequestId(v string) *AddRoomPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRoomPlanResponseBody) SetSuccess(v bool) *AddRoomPlanResponseBody {
	s.Success = &v
	return s
}

func (s *AddRoomPlanResponseBody) Validate() error {
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

type AddRoomPlanResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s AddRoomPlanResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s AddRoomPlanResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetAuthAction(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) SetPolicyType(v string) *AddRoomPlanResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *AddRoomPlanResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type AddRoomPlanResponseBodyData struct {
	// example:
	//
	// LTAI5t9kjkiudsnsu****
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// ""
	Callback *string `json:"Callback,omitempty" xml:"Callback,omitempty"`
	// example:
	//
	// 123/
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 1640315897
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// oss.aliyun.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAy****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// ngEWGzttc3v2gJWCxEEt****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s AddRoomPlanResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddRoomPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddRoomPlanResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *AddRoomPlanResponseBodyData) GetCallback() *string {
	return s.Callback
}

func (s *AddRoomPlanResponseBodyData) GetDir() *string {
	return s.Dir
}

func (s *AddRoomPlanResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *AddRoomPlanResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *AddRoomPlanResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *AddRoomPlanResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *AddRoomPlanResponseBodyData) SetAccessId(v string) *AddRoomPlanResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetCallback(v string) *AddRoomPlanResponseBodyData {
	s.Callback = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetDir(v string) *AddRoomPlanResponseBodyData {
	s.Dir = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetExpire(v string) *AddRoomPlanResponseBodyData {
	s.Expire = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetHost(v string) *AddRoomPlanResponseBodyData {
	s.Host = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetPolicy(v string) *AddRoomPlanResponseBodyData {
	s.Policy = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) SetSignature(v string) *AddRoomPlanResponseBodyData {
	s.Signature = &v
	return s
}

func (s *AddRoomPlanResponseBodyData) Validate() error {
	return dara.Validate(s)
}
