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
	// The content of the policy. Specifies whether the authorized user can access and use the resources and features defined in the policy.
	//
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
	// The description of the policy.
	//
	// example:
	//
	// test
	AbacPolicyDesc *string `json:"AbacPolicyDesc,omitempty" xml:"AbacPolicyDesc,omitempty"`
	// The name of the policy. The name must be unique for the tenant.
	//
	// This parameter is required.
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
