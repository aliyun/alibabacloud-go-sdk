// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMem0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CreateMem0ResponseBodyAccessDeniedDetail) *CreateMem0ResponseBody
	GetAccessDeniedDetail() *CreateMem0ResponseBodyAccessDeniedDetail
	SetData(v *CreateMem0ResponseBodyData) *CreateMem0ResponseBody
	GetData() *CreateMem0ResponseBodyData
	SetRequestId(v string) *CreateMem0ResponseBody
	GetRequestId() *string
}

type CreateMem0ResponseBody struct {
	AccessDeniedDetail *CreateMem0ResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data structure.
	Data *CreateMem0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMem0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMem0ResponseBody) GetAccessDeniedDetail() *CreateMem0ResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CreateMem0ResponseBody) GetData() *CreateMem0ResponseBodyData {
	return s.Data
}

func (s *CreateMem0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMem0ResponseBody) SetAccessDeniedDetail(v *CreateMem0ResponseBodyAccessDeniedDetail) *CreateMem0ResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CreateMem0ResponseBody) SetData(v *CreateMem0ResponseBodyData) *CreateMem0ResponseBody {
	s.Data = v
	return s
}

func (s *CreateMem0ResponseBody) SetRequestId(v string) *CreateMem0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMem0ResponseBody) Validate() error {
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

type CreateMem0ResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CreateMem0ResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0ResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CreateMem0ResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CreateMem0ResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type CreateMem0ResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMem0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMem0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMem0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *CreateMem0ResponseBodyData) SetTaskId(v int32) *CreateMem0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateMem0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
