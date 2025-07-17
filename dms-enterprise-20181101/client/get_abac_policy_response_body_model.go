// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAbacPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetAbacPolicyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAbacPolicyResponseBody
	GetErrorMessage() *string
	SetPolicy(v *GetAbacPolicyResponseBodyPolicy) *GetAbacPolicyResponseBody
	GetPolicy() *GetAbacPolicyResponseBodyPolicy
	SetRequestId(v string) *GetAbacPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAbacPolicyResponseBody
	GetSuccess() *bool
}

type GetAbacPolicyResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string                          `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Policy       *GetAbacPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// example:
	//
	// 2B7844DE-A0C3-50ED-A796-8F07D377144C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAbacPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAbacPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAbacPolicyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAbacPolicyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAbacPolicyResponseBody) GetPolicy() *GetAbacPolicyResponseBodyPolicy {
	return s.Policy
}

func (s *GetAbacPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAbacPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAbacPolicyResponseBody) SetErrorCode(v string) *GetAbacPolicyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAbacPolicyResponseBody) SetErrorMessage(v string) *GetAbacPolicyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAbacPolicyResponseBody) SetPolicy(v *GetAbacPolicyResponseBodyPolicy) *GetAbacPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetAbacPolicyResponseBody) SetRequestId(v string) *GetAbacPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAbacPolicyResponseBody) SetSuccess(v bool) *GetAbacPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *GetAbacPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAbacPolicyResponseBodyPolicy struct {
	// example:
	//
	// 3
	AuthorizedQuantity *string `json:"AuthorizedQuantity,omitempty" xml:"AuthorizedQuantity,omitempty"`
	// example:
	//
	// 51****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
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
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// example:
	//
	// test
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// example:
	//
	// 12****
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// policy_test
	PolicyName   *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicySource *string `json:"PolicySource,omitempty" xml:"PolicySource,omitempty"`
}

func (s GetAbacPolicyResponseBodyPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetAbacPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetAbacPolicyResponseBodyPolicy) GetAuthorizedQuantity() *string {
	return s.AuthorizedQuantity
}

func (s *GetAbacPolicyResponseBodyPolicy) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *GetAbacPolicyResponseBodyPolicy) GetPolicyContent() *string {
	return s.PolicyContent
}

func (s *GetAbacPolicyResponseBodyPolicy) GetPolicyDesc() *string {
	return s.PolicyDesc
}

func (s *GetAbacPolicyResponseBodyPolicy) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *GetAbacPolicyResponseBodyPolicy) GetPolicyName() *string {
	return s.PolicyName
}

func (s *GetAbacPolicyResponseBodyPolicy) GetPolicySource() *string {
	return s.PolicySource
}

func (s *GetAbacPolicyResponseBodyPolicy) SetAuthorizedQuantity(v string) *GetAbacPolicyResponseBodyPolicy {
	s.AuthorizedQuantity = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetCreatorId(v int64) *GetAbacPolicyResponseBodyPolicy {
	s.CreatorId = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetPolicyContent(v string) *GetAbacPolicyResponseBodyPolicy {
	s.PolicyContent = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetPolicyDesc(v string) *GetAbacPolicyResponseBodyPolicy {
	s.PolicyDesc = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetPolicyId(v int64) *GetAbacPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetPolicyName(v string) *GetAbacPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) SetPolicySource(v string) *GetAbacPolicyResponseBodyPolicy {
	s.PolicySource = &v
	return s
}

func (s *GetAbacPolicyResponseBodyPolicy) Validate() error {
	return dara.Validate(s)
}
