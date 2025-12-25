// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredictionWallLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *PredictionWallLineResponseBodyAccessDeniedDetail) *PredictionWallLineResponseBody
	GetAccessDeniedDetail() *PredictionWallLineResponseBodyAccessDeniedDetail
	SetCode(v int64) *PredictionWallLineResponseBody
	GetCode() *int64
	SetMessage(v string) *PredictionWallLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *PredictionWallLineResponseBody
	GetRequestId() *string
	SetSubSceneId(v string) *PredictionWallLineResponseBody
	GetSubSceneId() *string
	SetSuccess(v bool) *PredictionWallLineResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *PredictionWallLineResponseBody
	GetTaskId() *string
}

type PredictionWallLineResponseBody struct {
	AccessDeniedDetail *PredictionWallLineResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 2345****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1234****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PredictionWallLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PredictionWallLineResponseBody) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponseBody) GetAccessDeniedDetail() *PredictionWallLineResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *PredictionWallLineResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PredictionWallLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PredictionWallLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PredictionWallLineResponseBody) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *PredictionWallLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PredictionWallLineResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PredictionWallLineResponseBody) SetAccessDeniedDetail(v *PredictionWallLineResponseBodyAccessDeniedDetail) *PredictionWallLineResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *PredictionWallLineResponseBody) SetCode(v int64) *PredictionWallLineResponseBody {
	s.Code = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetMessage(v string) *PredictionWallLineResponseBody {
	s.Message = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetRequestId(v string) *PredictionWallLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSubSceneId(v string) *PredictionWallLineResponseBody {
	s.SubSceneId = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetSuccess(v bool) *PredictionWallLineResponseBody {
	s.Success = &v
	return s
}

func (s *PredictionWallLineResponseBody) SetTaskId(v string) *PredictionWallLineResponseBody {
	s.TaskId = &v
	return s
}

func (s *PredictionWallLineResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PredictionWallLineResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s PredictionWallLineResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s PredictionWallLineResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetAuthAction(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) SetPolicyType(v string) *PredictionWallLineResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *PredictionWallLineResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
