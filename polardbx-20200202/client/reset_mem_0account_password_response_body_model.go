// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetMem0AccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) *ResetMem0AccountPasswordResponseBody
	GetAccessDeniedDetail() *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail
	SetData(v *ResetMem0AccountPasswordResponseBodyData) *ResetMem0AccountPasswordResponseBody
	GetData() *ResetMem0AccountPasswordResponseBodyData
	SetRequestId(v string) *ResetMem0AccountPasswordResponseBody
	GetRequestId() *string
}

type ResetMem0AccountPasswordResponseBody struct {
	AccessDeniedDetail *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The data struct.
	Data *ResetMem0AccountPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A501A191-BD70-5E50-98A9-C2A486A82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetMem0AccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetMem0AccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetMem0AccountPasswordResponseBody) GetAccessDeniedDetail() *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ResetMem0AccountPasswordResponseBody) GetData() *ResetMem0AccountPasswordResponseBodyData {
	return s.Data
}

func (s *ResetMem0AccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetMem0AccountPasswordResponseBody) SetAccessDeniedDetail(v *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) *ResetMem0AccountPasswordResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ResetMem0AccountPasswordResponseBody) SetData(v *ResetMem0AccountPasswordResponseBodyData) *ResetMem0AccountPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetMem0AccountPasswordResponseBody) SetRequestId(v string) *ResetMem0AccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBody) Validate() error {
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

type ResetMem0AccountPasswordResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ResetMem0AccountPasswordResponseBodyData struct {
	// API KEY
	//
	// example:
	//
	// aafdf2e7d0988ef***
	Mem0ApiKey *string `json:"Mem0ApiKey,omitempty" xml:"Mem0ApiKey,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetMem0AccountPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetMem0AccountPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetMem0AccountPasswordResponseBodyData) GetMem0ApiKey() *string {
	return s.Mem0ApiKey
}

func (s *ResetMem0AccountPasswordResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ResetMem0AccountPasswordResponseBodyData) SetMem0ApiKey(v string) *ResetMem0AccountPasswordResponseBodyData {
	s.Mem0ApiKey = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyData) SetTaskId(v int32) *ResetMem0AccountPasswordResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ResetMem0AccountPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
