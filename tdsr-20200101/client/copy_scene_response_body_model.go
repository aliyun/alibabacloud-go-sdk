// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopySceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CopySceneResponseBodyAccessDeniedDetail) *CopySceneResponseBody
	GetAccessDeniedDetail() *CopySceneResponseBodyAccessDeniedDetail
	SetCode(v int64) *CopySceneResponseBody
	GetCode() *int64
	SetData(v *CopySceneResponseBodyData) *CopySceneResponseBody
	GetData() *CopySceneResponseBodyData
	SetMessage(v string) *CopySceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopySceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CopySceneResponseBody
	GetSuccess() *bool
}

type CopySceneResponseBody struct {
	AccessDeniedDetail *CopySceneResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CopySceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F882EA7-3A1D-0113-94E4-70162C4B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopySceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopySceneResponseBody) GoString() string {
	return s.String()
}

func (s *CopySceneResponseBody) GetAccessDeniedDetail() *CopySceneResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CopySceneResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CopySceneResponseBody) GetData() *CopySceneResponseBodyData {
	return s.Data
}

func (s *CopySceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopySceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopySceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CopySceneResponseBody) SetAccessDeniedDetail(v *CopySceneResponseBodyAccessDeniedDetail) *CopySceneResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CopySceneResponseBody) SetCode(v int64) *CopySceneResponseBody {
	s.Code = &v
	return s
}

func (s *CopySceneResponseBody) SetData(v *CopySceneResponseBodyData) *CopySceneResponseBody {
	s.Data = v
	return s
}

func (s *CopySceneResponseBody) SetMessage(v string) *CopySceneResponseBody {
	s.Message = &v
	return s
}

func (s *CopySceneResponseBody) SetRequestId(v string) *CopySceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopySceneResponseBody) SetSuccess(v bool) *CopySceneResponseBody {
	s.Success = &v
	return s
}

func (s *CopySceneResponseBody) Validate() error {
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

type CopySceneResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CopySceneResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CopySceneResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CopySceneResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CopySceneResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CopySceneResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CopySceneResponseBodyData struct {
	// example:
	//
	// yuywey****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CopySceneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CopySceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *CopySceneResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CopySceneResponseBodyData) SetTaskId(v string) *CopySceneResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CopySceneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
