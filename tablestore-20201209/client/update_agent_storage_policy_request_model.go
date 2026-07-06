// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStoragePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *UpdateAgentStoragePolicyRequest
	GetAgentStorageName() *string
	SetPolicy(v string) *UpdateAgentStoragePolicyRequest
	GetPolicy() *string
	SetPolicyVersion(v int64) *UpdateAgentStoragePolicyRequest
	GetPolicyVersion() *int64
}

type UpdateAgentStoragePolicyRequest struct {
	// The name of the agent storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The access control policy of the agent storage in JSON format. For more information, see https://www.alibabacloud.com/help/en/ram/user-guide/policy-structure-and-syntax.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "Version": "1",
	//
	//     "Statement": [
	//
	//         {
	//
	//             "Action": [
	//
	//                 "ots:*"
	//
	//             ],
	//
	//             "Resource": [
	//
	//                 "acs:ots:*:13791xxxxxxxxxxx:agentstorage/myagentstorage*"
	//
	//             ],
	//
	//             "Principal": [
	//
	//                 "*"
	//
	//             ],
	//
	//             "Effect": "Allow",
	//
	//             "Condition": {
	//
	//                 "StringEquals": {
	//
	//                     "ots:TLSVersion": [
	//
	//                         "1.2"
	//
	//                     ]
	//
	//                 },
	//
	//                 "IpAddress": {
	//
	//                     "acs:SourceIp": [
	//
	//                         "192.168.0.1",
	//
	//                         "172.16.0.1"
	//
	//                     ]
	//
	//                 }
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the agent storage access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s UpdateAgentStoragePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStoragePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentStoragePolicyRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *UpdateAgentStoragePolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateAgentStoragePolicyRequest) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *UpdateAgentStoragePolicyRequest) SetAgentStorageName(v string) *UpdateAgentStoragePolicyRequest {
	s.AgentStorageName = &v
	return s
}

func (s *UpdateAgentStoragePolicyRequest) SetPolicy(v string) *UpdateAgentStoragePolicyRequest {
	s.Policy = &v
	return s
}

func (s *UpdateAgentStoragePolicyRequest) SetPolicyVersion(v int64) *UpdateAgentStoragePolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *UpdateAgentStoragePolicyRequest) Validate() error {
	return dara.Validate(s)
}
