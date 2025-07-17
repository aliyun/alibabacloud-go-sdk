// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAbacPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbacPolicyContent(v string) *CreateAbacPolicyRequest
	GetAbacPolicyContent() *string
	SetAbacPolicyDesc(v string) *CreateAbacPolicyRequest
	GetAbacPolicyDesc() *string
	SetAbacPolicyName(v string) *CreateAbacPolicyRequest
	GetAbacPolicyName() *string
	SetTid(v int64) *CreateAbacPolicyRequest
	GetTid() *int64
}

type CreateAbacPolicyRequest struct {
	// This parameter is required.
	//
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
	// policy_test
	AbacPolicyName *string `json:"AbacPolicyName,omitempty" xml:"AbacPolicyName,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateAbacPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAbacPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAbacPolicyRequest) GetAbacPolicyContent() *string {
	return s.AbacPolicyContent
}

func (s *CreateAbacPolicyRequest) GetAbacPolicyDesc() *string {
	return s.AbacPolicyDesc
}

func (s *CreateAbacPolicyRequest) GetAbacPolicyName() *string {
	return s.AbacPolicyName
}

func (s *CreateAbacPolicyRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateAbacPolicyRequest) SetAbacPolicyContent(v string) *CreateAbacPolicyRequest {
	s.AbacPolicyContent = &v
	return s
}

func (s *CreateAbacPolicyRequest) SetAbacPolicyDesc(v string) *CreateAbacPolicyRequest {
	s.AbacPolicyDesc = &v
	return s
}

func (s *CreateAbacPolicyRequest) SetAbacPolicyName(v string) *CreateAbacPolicyRequest {
	s.AbacPolicyName = &v
	return s
}

func (s *CreateAbacPolicyRequest) SetTid(v int64) *CreateAbacPolicyRequest {
	s.Tid = &v
	return s
}

func (s *CreateAbacPolicyRequest) Validate() error {
	return dara.Validate(s)
}
