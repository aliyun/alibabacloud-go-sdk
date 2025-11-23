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
	// The content of the policy.
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
	// The description of the policy.
	//
	// example:
	//
	// test
	AbacPolicyDesc *string `json:"AbacPolicyDesc,omitempty" xml:"AbacPolicyDesc,omitempty"`
	// The ID of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	AbacPolicyId *int64 `json:"AbacPolicyId,omitempty" xml:"AbacPolicyId,omitempty"`
	// The name of the permission policy.
	//
	// example:
	//
	// policy_test
	AbacPolicyName *string `json:"AbacPolicyName,omitempty" xml:"AbacPolicyName,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the DMS console and move the pointer over the profile picture in the upper-right corner. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the "Manage DMS tenants" topic.
	//
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
