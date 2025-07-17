// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAbacPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbacPolicyContent(v string) *UpdateAbacPolicyRequest
	GetAbacPolicyContent() *string
	SetAbacPolicyDesc(v string) *UpdateAbacPolicyRequest
	GetAbacPolicyDesc() *string
	SetAbacPolicyId(v int64) *UpdateAbacPolicyRequest
	GetAbacPolicyId() *int64
	SetAbacPolicyName(v string) *UpdateAbacPolicyRequest
	GetAbacPolicyName() *string
	SetTid(v int64) *UpdateAbacPolicyRequest
	GetTid() *int64
}

type UpdateAbacPolicyRequest struct {
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
	// This parameter is required.
	//
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
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateAbacPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAbacPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAbacPolicyRequest) GetAbacPolicyContent() *string {
	return s.AbacPolicyContent
}

func (s *UpdateAbacPolicyRequest) GetAbacPolicyDesc() *string {
	return s.AbacPolicyDesc
}

func (s *UpdateAbacPolicyRequest) GetAbacPolicyId() *int64 {
	return s.AbacPolicyId
}

func (s *UpdateAbacPolicyRequest) GetAbacPolicyName() *string {
	return s.AbacPolicyName
}

func (s *UpdateAbacPolicyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateAbacPolicyRequest) SetAbacPolicyContent(v string) *UpdateAbacPolicyRequest {
	s.AbacPolicyContent = &v
	return s
}

func (s *UpdateAbacPolicyRequest) SetAbacPolicyDesc(v string) *UpdateAbacPolicyRequest {
	s.AbacPolicyDesc = &v
	return s
}

func (s *UpdateAbacPolicyRequest) SetAbacPolicyId(v int64) *UpdateAbacPolicyRequest {
	s.AbacPolicyId = &v
	return s
}

func (s *UpdateAbacPolicyRequest) SetAbacPolicyName(v string) *UpdateAbacPolicyRequest {
	s.AbacPolicyName = &v
	return s
}

func (s *UpdateAbacPolicyRequest) SetTid(v int64) *UpdateAbacPolicyRequest {
	s.Tid = &v
	return s
}

func (s *UpdateAbacPolicyRequest) Validate() error {
	return dara.Validate(s)
}
