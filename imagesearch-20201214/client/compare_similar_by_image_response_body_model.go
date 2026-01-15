// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSimilarByImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *CompareSimilarByImageResponseBodyAccessDeniedDetail) *CompareSimilarByImageResponseBody
	GetAccessDeniedDetail() *CompareSimilarByImageResponseBodyAccessDeniedDetail
	SetCode(v int32) *CompareSimilarByImageResponseBody
	GetCode() *int32
	SetMsg(v string) *CompareSimilarByImageResponseBody
	GetMsg() *string
	SetRequestId(v string) *CompareSimilarByImageResponseBody
	GetRequestId() *string
	SetScore(v float64) *CompareSimilarByImageResponseBody
	GetScore() *float64
	SetSuccess(v bool) *CompareSimilarByImageResponseBody
	GetSuccess() *bool
}

type CompareSimilarByImageResponseBody struct {
	AccessDeniedDetail *CompareSimilarByImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0.85
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CompareSimilarByImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageResponseBody) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponseBody) GetAccessDeniedDetail() *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *CompareSimilarByImageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CompareSimilarByImageResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CompareSimilarByImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompareSimilarByImageResponseBody) GetScore() *float64 {
	return s.Score
}

func (s *CompareSimilarByImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CompareSimilarByImageResponseBody) SetAccessDeniedDetail(v *CompareSimilarByImageResponseBodyAccessDeniedDetail) *CompareSimilarByImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetCode(v int32) *CompareSimilarByImageResponseBody {
	s.Code = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetMsg(v string) *CompareSimilarByImageResponseBody {
	s.Msg = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetRequestId(v string) *CompareSimilarByImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetScore(v float64) *CompareSimilarByImageResponseBody {
	s.Score = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetSuccess(v bool) *CompareSimilarByImageResponseBody {
	s.Success = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CompareSimilarByImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CompareSimilarByImageResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}
