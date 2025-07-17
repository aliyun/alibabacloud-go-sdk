// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbacPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListAbacPoliciesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAbacPoliciesResponseBody
	GetErrorMessage() *string
	SetPolicyList(v []*ListAbacPoliciesResponseBodyPolicyList) *ListAbacPoliciesResponseBody
	GetPolicyList() []*ListAbacPoliciesResponseBodyPolicyList
	SetRequestId(v string) *ListAbacPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAbacPoliciesResponseBody
	GetSuccess() *bool
	SetTid(v int64) *ListAbacPoliciesResponseBody
	GetTid() *int64
	SetTotalCount(v int64) *ListAbacPoliciesResponseBody
	GetTotalCount() *int64
}

type ListAbacPoliciesResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PolicyList   []*ListAbacPoliciesResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAbacPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAbacPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAbacPoliciesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAbacPoliciesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAbacPoliciesResponseBody) GetPolicyList() []*ListAbacPoliciesResponseBodyPolicyList {
	return s.PolicyList
}

func (s *ListAbacPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAbacPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAbacPoliciesResponseBody) GetTid() *int64 {
	return s.Tid
}

func (s *ListAbacPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAbacPoliciesResponseBody) SetErrorCode(v string) *ListAbacPoliciesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetErrorMessage(v string) *ListAbacPoliciesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetPolicyList(v []*ListAbacPoliciesResponseBodyPolicyList) *ListAbacPoliciesResponseBody {
	s.PolicyList = v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetRequestId(v string) *ListAbacPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetSuccess(v bool) *ListAbacPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetTid(v int64) *ListAbacPoliciesResponseBody {
	s.Tid = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) SetTotalCount(v int64) *ListAbacPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAbacPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAbacPoliciesResponseBodyPolicyList struct {
	// example:
	//
	// {
	//
	//   "Statement": [
	//
	//     {
	//
	//       "Action": "*",
	//
	//       "Effect": "Allow",
	//
	//       "Resource": "*",
	//
	//       "Condition": {
	//
	//         "StringEqualsIgnoreCase": {
	//
	//           "dms:DbType": [
	//
	//             "redis"
	//
	//           ]
	//
	//         }
	//
	//       }
	//
	//     }
	//
	//   ],
	//
	//   "Version": "1"
	//
	// }
	AbacPolicyContent *string `json:"AbacPolicyContent,omitempty" xml:"AbacPolicyContent,omitempty"`
	// example:
	//
	// test
	AbacPolicyDesc *string `json:"AbacPolicyDesc,omitempty" xml:"AbacPolicyDesc,omitempty"`
	// example:
	//
	// 12****
	AbacPolicyId *int64 `json:"AbacPolicyId,omitempty" xml:"AbacPolicyId,omitempty"`
	// example:
	//
	// policy_test
	AbacPolicyName *string `json:"AbacPolicyName,omitempty" xml:"AbacPolicyName,omitempty"`
	// example:
	//
	// USER_DEFINE
	AbacPolicySource *string `json:"AbacPolicySource,omitempty" xml:"AbacPolicySource,omitempty"`
	// example:
	//
	// 51****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
}

func (s ListAbacPoliciesResponseBodyPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAbacPoliciesResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetAbacPolicyContent() *string {
	return s.AbacPolicyContent
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetAbacPolicyDesc() *string {
	return s.AbacPolicyDesc
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetAbacPolicyId() *int64 {
	return s.AbacPolicyId
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetAbacPolicyName() *string {
	return s.AbacPolicyName
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetAbacPolicySource() *string {
	return s.AbacPolicySource
}

func (s *ListAbacPoliciesResponseBodyPolicyList) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetAbacPolicyContent(v string) *ListAbacPoliciesResponseBodyPolicyList {
	s.AbacPolicyContent = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetAbacPolicyDesc(v string) *ListAbacPoliciesResponseBodyPolicyList {
	s.AbacPolicyDesc = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetAbacPolicyId(v int64) *ListAbacPoliciesResponseBodyPolicyList {
	s.AbacPolicyId = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetAbacPolicyName(v string) *ListAbacPoliciesResponseBodyPolicyList {
	s.AbacPolicyName = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetAbacPolicySource(v string) *ListAbacPoliciesResponseBodyPolicyList {
	s.AbacPolicySource = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) SetCreatorId(v int64) *ListAbacPoliciesResponseBodyPolicyList {
	s.CreatorId = &v
	return s
}

func (s *ListAbacPoliciesResponseBodyPolicyList) Validate() error {
	return dara.Validate(s)
}
