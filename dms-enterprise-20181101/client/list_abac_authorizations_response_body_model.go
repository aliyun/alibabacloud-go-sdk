// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacAuthorizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationList(v []*ListAbacAuthorizationsResponseBodyAuthorizationList) *ListAbacAuthorizationsResponseBody
	GetAuthorizationList() []*ListAbacAuthorizationsResponseBodyAuthorizationList
	SetErrorCode(v string) *ListAbacAuthorizationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAbacAuthorizationsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAbacAuthorizationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAbacAuthorizationsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAbacAuthorizationsResponseBody
	GetTotalCount() *int64
}

type ListAbacAuthorizationsResponseBody struct {
	AuthorizationList []*ListAbacAuthorizationsResponseBodyAuthorizationList `json:"AuthorizationList,omitempty" xml:"AuthorizationList,omitempty" type:"Repeated"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAbacAuthorizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAbacAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAbacAuthorizationsResponseBody) GetAuthorizationList() []*ListAbacAuthorizationsResponseBodyAuthorizationList {
	return s.AuthorizationList
}

func (s *ListAbacAuthorizationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAbacAuthorizationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAbacAuthorizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAbacAuthorizationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAbacAuthorizationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAbacAuthorizationsResponseBody) SetAuthorizationList(v []*ListAbacAuthorizationsResponseBodyAuthorizationList) *ListAbacAuthorizationsResponseBody {
	s.AuthorizationList = v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) SetErrorCode(v string) *ListAbacAuthorizationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) SetErrorMessage(v string) *ListAbacAuthorizationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) SetRequestId(v string) *ListAbacAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) SetSuccess(v bool) *ListAbacAuthorizationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) SetTotalCount(v int64) *ListAbacAuthorizationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAbacAuthorizationsResponseBodyAuthorizationList struct {
	// example:
	//
	// 32****
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// example:
	//
	// 51****
	IdentityId *int64 `json:"IdentityId,omitempty" xml:"IdentityId,omitempty"`
	// example:
	//
	// test_user
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// example:
	//
	// 12****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// policy_test
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// USER_DEFINE
	PolicySource *string `json:"PolicySource,omitempty" xml:"PolicySource,omitempty"`
}

func (s ListAbacAuthorizationsResponseBodyAuthorizationList) String() string {
	return dara.Prettify(s)
}

func (s ListAbacAuthorizationsResponseBodyAuthorizationList) GoString() string {
	return s.String()
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetIdentityId() *int64 {
	return s.IdentityId
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetIdentityName() *string {
	return s.IdentityName
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetIdentityType() *string {
	return s.IdentityType
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) GetPolicySource() *string {
	return s.PolicySource
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetAuthorizationId(v int64) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.AuthorizationId = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetIdentityId(v int64) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.IdentityId = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetIdentityName(v string) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.IdentityName = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetIdentityType(v string) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.IdentityType = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetPolicyId(v int64) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.PolicyId = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetPolicyName(v string) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.PolicyName = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) SetPolicySource(v string) *ListAbacAuthorizationsResponseBodyAuthorizationList {
	s.PolicySource = &v
	return s
}

func (s *ListAbacAuthorizationsResponseBodyAuthorizationList) Validate() error {
	return dara.Validate(s)
}
