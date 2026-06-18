// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DeleteMem0ResponseBodyAccessDeniedDetail) *DeleteMem0ResponseBody
	GetAccessDeniedDetail() *DeleteMem0ResponseBodyAccessDeniedDetail
	SetData(v *DeleteMem0ResponseBodyData) *DeleteMem0ResponseBody
	GetData() *DeleteMem0ResponseBodyData
	SetRequestId(v string) *DeleteMem0ResponseBody
	GetRequestId() *string
}

type DeleteMem0ResponseBody struct {
	AccessDeniedDetail *DeleteMem0ResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The returned data.
	Data *DeleteMem0ResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMem0ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMem0ResponseBody) GetAccessDeniedDetail() *DeleteMem0ResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DeleteMem0ResponseBody) GetData() *DeleteMem0ResponseBodyData {
	return s.Data
}

func (s *DeleteMem0ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMem0ResponseBody) SetAccessDeniedDetail(v *DeleteMem0ResponseBodyAccessDeniedDetail) *DeleteMem0ResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteMem0ResponseBody) SetData(v *DeleteMem0ResponseBodyData) *DeleteMem0ResponseBody {
	s.Data = v
	return s
}

func (s *DeleteMem0ResponseBody) SetRequestId(v string) *DeleteMem0ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMem0ResponseBody) Validate() error {
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

type DeleteMem0ResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DeleteMem0ResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0ResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DeleteMem0ResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DeleteMem0ResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DeleteMem0ResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteMem0ResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0ResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteMem0ResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteMem0ResponseBodyData) SetTaskId(v int32) *DeleteMem0ResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteMem0ResponseBodyData) Validate() error {
	return dara.Validate(s)
}
