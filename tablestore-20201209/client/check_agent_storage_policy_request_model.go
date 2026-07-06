// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAgentStoragePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentStorageName(v string) *CheckAgentStoragePolicyRequest
	GetAgentStorageName() *string
	SetPolicy(v string) *CheckAgentStoragePolicyRequest
	GetPolicy() *string
}

type CheckAgentStoragePolicyRequest struct {
	// The agent storage name.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-test
	AgentStorageName *string `json:"AgentStorageName,omitempty" xml:"AgentStorageName,omitempty"`
	// The agent storage access control policy in JSON format.
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
}

func (s CheckAgentStoragePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAgentStoragePolicyRequest) GoString() string {
	return s.String()
}

func (s *CheckAgentStoragePolicyRequest) GetAgentStorageName() *string {
	return s.AgentStorageName
}

func (s *CheckAgentStoragePolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CheckAgentStoragePolicyRequest) SetAgentStorageName(v string) *CheckAgentStoragePolicyRequest {
	s.AgentStorageName = &v
	return s
}

func (s *CheckAgentStoragePolicyRequest) SetPolicy(v string) *CheckAgentStoragePolicyRequest {
	s.Policy = &v
	return s
}

func (s *CheckAgentStoragePolicyRequest) Validate() error {
	return dara.Validate(s)
}
