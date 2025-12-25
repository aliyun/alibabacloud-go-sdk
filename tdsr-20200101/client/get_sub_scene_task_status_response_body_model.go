// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubSceneTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) *GetSubSceneTaskStatusResponseBody
	GetAccessDeniedDetail() *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail
	SetCode(v int64) *GetSubSceneTaskStatusResponseBody
	GetCode() *int64
	SetList(v []*GetSubSceneTaskStatusResponseBodyList) *GetSubSceneTaskStatusResponseBody
	GetList() []*GetSubSceneTaskStatusResponseBodyList
	SetMessage(v string) *GetSubSceneTaskStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSubSceneTaskStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSubSceneTaskStatusResponseBody
	GetSuccess() *bool
}

type GetSubSceneTaskStatusResponseBody struct {
	AccessDeniedDetail *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	List []*GetSubSceneTaskStatusResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s GetSubSceneTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBody) GetAccessDeniedDetail() *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetSubSceneTaskStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetSubSceneTaskStatusResponseBody) GetList() []*GetSubSceneTaskStatusResponseBodyList {
	return s.List
}

func (s *GetSubSceneTaskStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubSceneTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSubSceneTaskStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubSceneTaskStatusResponseBody) SetAccessDeniedDetail(v *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) *GetSubSceneTaskStatusResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetCode(v int64) *GetSubSceneTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetList(v []*GetSubSceneTaskStatusResponseBodyList) *GetSubSceneTaskStatusResponseBody {
	s.List = v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetMessage(v string) *GetSubSceneTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetRequestId(v string) *GetSubSceneTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) SetSuccess(v bool) *GetSubSceneTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBody) Validate() error {
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

type GetSubSceneTaskStatusResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthAction(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) SetPolicyType(v string) *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetSubSceneTaskStatusResponseBodyList struct {
	// example:
	//
	// 2001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 4638****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2345****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// init
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 3456***
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// example:
	//
	// cutimage
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSubSceneTaskStatusResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s GetSubSceneTaskStatusResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetId() *string {
	return s.Id
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetSceneId() *string {
	return s.SceneId
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *GetSubSceneTaskStatusResponseBodyList) GetType() *string {
	return s.Type
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorCode(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorCode = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetErrorMsg(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.ErrorMsg = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Id = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetStatus(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Status = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetSubSceneId(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.SubSceneId = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) SetType(v string) *GetSubSceneTaskStatusResponseBodyList {
	s.Type = &v
	return s
}

func (s *GetSubSceneTaskStatusResponseBodyList) Validate() error {
	return dara.Validate(s)
}
