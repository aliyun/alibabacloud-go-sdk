// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSingleConnDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetSingleConnDataResponseBodyAccessDeniedDetail) *GetSingleConnDataResponseBody
	GetAccessDeniedDetail() *GetSingleConnDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetSingleConnDataResponseBody
	GetCode() *int64
	SetList(v []*GetSingleConnDataResponseBodyList) *GetSingleConnDataResponseBody
	GetList() []*GetSingleConnDataResponseBodyList
	SetMessage(v string) *GetSingleConnDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSingleConnDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSingleConnDataResponseBody
	GetSuccess() *bool
	SetVersion(v string) *GetSingleConnDataResponseBody
	GetVersion() *string
}

type GetSingleConnDataResponseBody struct {
	AccessDeniedDetail *GetSingleConnDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                               `json:"Code,omitempty" xml:"Code,omitempty"`
	List []*GetSingleConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// example:
	//
	// 2.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetSingleConnDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSingleConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBody) GetAccessDeniedDetail() *GetSingleConnDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetSingleConnDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSingleConnDataResponseBody) GetList() []*GetSingleConnDataResponseBodyList {
	return s.List
}

func (s *GetSingleConnDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSingleConnDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSingleConnDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSingleConnDataResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetSingleConnDataResponseBody) SetAccessDeniedDetail(v *GetSingleConnDataResponseBodyAccessDeniedDetail) *GetSingleConnDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetSingleConnDataResponseBody) SetCode(v int64) *GetSingleConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetList(v []*GetSingleConnDataResponseBodyList) *GetSingleConnDataResponseBody {
	s.List = v
	return s
}

func (s *GetSingleConnDataResponseBody) SetMessage(v string) *GetSingleConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetRequestId(v string) *GetSingleConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetSuccess(v bool) *GetSingleConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetSingleConnDataResponseBody) SetVersion(v string) *GetSingleConnDataResponseBody {
	s.Version = &v
	return s
}

func (s *GetSingleConnDataResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSingleConnDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetSingleConnDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSingleConnDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetSingleConnDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetSingleConnDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetSingleConnDataResponseBodyList struct {
	// ID
	//
	// example:
	//
	// 1#234@abc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1#567#def
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	// example:
	//
	// outer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSingleConnDataResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s GetSingleConnDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetSingleConnDataResponseBodyList) GetId() *string {
	return s.Id
}

func (s *GetSingleConnDataResponseBodyList) GetMapId() *string {
	return s.MapId
}

func (s *GetSingleConnDataResponseBodyList) GetType() *string {
	return s.Type
}

func (s *GetSingleConnDataResponseBodyList) SetId(v string) *GetSingleConnDataResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetSingleConnDataResponseBodyList) SetMapId(v string) *GetSingleConnDataResponseBodyList {
	s.MapId = &v
	return s
}

func (s *GetSingleConnDataResponseBodyList) SetType(v string) *GetSingleConnDataResponseBodyList {
	s.Type = &v
	return s
}

func (s *GetSingleConnDataResponseBodyList) Validate() error {
	return dara.Validate(s)
}
