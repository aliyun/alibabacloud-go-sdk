// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetConnDataResponseBodyAccessDeniedDetail) *GetConnDataResponseBody
	GetAccessDeniedDetail() *GetConnDataResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetConnDataResponseBody
	GetCode() *int64
	SetExtend(v string) *GetConnDataResponseBody
	GetExtend() *string
	SetList(v []*GetConnDataResponseBodyList) *GetConnDataResponseBody
	GetList() []*GetConnDataResponseBodyList
	SetMessage(v string) *GetConnDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetConnDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetConnDataResponseBody
	GetSuccess() *bool
	SetVersion(v string) *GetConnDataResponseBody
	GetVersion() *string
}

type GetConnDataResponseBody struct {
	AccessDeniedDetail *GetConnDataResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	Extend *string                        `json:"Extend,omitempty" xml:"Extend,omitempty"`
	List   []*GetConnDataResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
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

func (s GetConnDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBody) GetAccessDeniedDetail() *GetConnDataResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetConnDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetConnDataResponseBody) GetExtend() *string {
	return s.Extend
}

func (s *GetConnDataResponseBody) GetList() []*GetConnDataResponseBodyList {
	return s.List
}

func (s *GetConnDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetConnDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetConnDataResponseBody) GetVersion() *string {
	return s.Version
}

func (s *GetConnDataResponseBody) SetAccessDeniedDetail(v *GetConnDataResponseBodyAccessDeniedDetail) *GetConnDataResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetConnDataResponseBody) SetCode(v int64) *GetConnDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnDataResponseBody) SetExtend(v string) *GetConnDataResponseBody {
	s.Extend = &v
	return s
}

func (s *GetConnDataResponseBody) SetList(v []*GetConnDataResponseBodyList) *GetConnDataResponseBody {
	s.List = v
	return s
}

func (s *GetConnDataResponseBody) SetMessage(v string) *GetConnDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnDataResponseBody) SetRequestId(v string) *GetConnDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnDataResponseBody) SetSuccess(v bool) *GetConnDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetConnDataResponseBody) SetVersion(v string) *GetConnDataResponseBody {
	s.Version = &v
	return s
}

func (s *GetConnDataResponseBody) Validate() error {
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

type GetConnDataResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetConnDataResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetConnDataResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetConnDataResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetConnDataResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetConnDataResponseBodyList struct {
	// ID
	//
	// example:
	//
	// 1#234#abc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1#456#def
	MapId *string `json:"MapId,omitempty" xml:"MapId,omitempty"`
	// example:
	//
	// outer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetConnDataResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s GetConnDataResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetConnDataResponseBodyList) GetId() *string {
	return s.Id
}

func (s *GetConnDataResponseBodyList) GetMapId() *string {
	return s.MapId
}

func (s *GetConnDataResponseBodyList) GetType() *string {
	return s.Type
}

func (s *GetConnDataResponseBodyList) SetId(v string) *GetConnDataResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetConnDataResponseBodyList) SetMapId(v string) *GetConnDataResponseBodyList {
	s.MapId = &v
	return s
}

func (s *GetConnDataResponseBodyList) SetType(v string) *GetConnDataResponseBodyList {
	s.Type = &v
	return s
}

func (s *GetConnDataResponseBodyList) Validate() error {
	return dara.Validate(s)
}
