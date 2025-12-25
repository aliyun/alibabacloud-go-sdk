// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DetailProjectResponseBodyAccessDeniedDetail) *DetailProjectResponseBody
	GetAccessDeniedDetail() *DetailProjectResponseBodyAccessDeniedDetail
	SetBusinessId(v int64) *DetailProjectResponseBody
	GetBusinessId() *int64
	SetBusinessName(v string) *DetailProjectResponseBody
	GetBusinessName() *string
	SetCode(v int64) *DetailProjectResponseBody
	GetCode() *int64
	SetGmtCreate(v int64) *DetailProjectResponseBody
	GetGmtCreate() *int64
	SetGmtModified(v int64) *DetailProjectResponseBody
	GetGmtModified() *int64
	SetId(v string) *DetailProjectResponseBody
	GetId() *string
	SetMessage(v string) *DetailProjectResponseBody
	GetMessage() *string
	SetName(v string) *DetailProjectResponseBody
	GetName() *string
	SetRequestId(v string) *DetailProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetailProjectResponseBody
	GetSuccess() *bool
	SetToken(v string) *DetailProjectResponseBody
	GetToken() *string
}

type DetailProjectResponseBody struct {
	AccessDeniedDetail *DetailProjectResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 5244****
	BusinessId   *int64  `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1621236933677
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1621236933677
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1234****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Token
	//
	// example:
	//
	// d989623696ab4f87a80b8d5b0b00****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DetailProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetailProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetailProjectResponseBody) GetAccessDeniedDetail() *DetailProjectResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DetailProjectResponseBody) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *DetailProjectResponseBody) GetBusinessName() *string {
	return s.BusinessName
}

func (s *DetailProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DetailProjectResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DetailProjectResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DetailProjectResponseBody) GetId() *string {
	return s.Id
}

func (s *DetailProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetailProjectResponseBody) GetName() *string {
	return s.Name
}

func (s *DetailProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetailProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetailProjectResponseBody) GetToken() *string {
	return s.Token
}

func (s *DetailProjectResponseBody) SetAccessDeniedDetail(v *DetailProjectResponseBodyAccessDeniedDetail) *DetailProjectResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DetailProjectResponseBody) SetBusinessId(v int64) *DetailProjectResponseBody {
	s.BusinessId = &v
	return s
}

func (s *DetailProjectResponseBody) SetBusinessName(v string) *DetailProjectResponseBody {
	s.BusinessName = &v
	return s
}

func (s *DetailProjectResponseBody) SetCode(v int64) *DetailProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DetailProjectResponseBody) SetGmtCreate(v int64) *DetailProjectResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DetailProjectResponseBody) SetGmtModified(v int64) *DetailProjectResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DetailProjectResponseBody) SetId(v string) *DetailProjectResponseBody {
	s.Id = &v
	return s
}

func (s *DetailProjectResponseBody) SetMessage(v string) *DetailProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DetailProjectResponseBody) SetName(v string) *DetailProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DetailProjectResponseBody) SetRequestId(v string) *DetailProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailProjectResponseBody) SetSuccess(v bool) *DetailProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DetailProjectResponseBody) SetToken(v string) *DetailProjectResponseBody {
	s.Token = &v
	return s
}

func (s *DetailProjectResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetailProjectResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DetailProjectResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DetailProjectResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DetailProjectResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DetailProjectResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
