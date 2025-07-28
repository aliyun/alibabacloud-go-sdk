// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *UpdateInstancePolicyRequest
	GetInstanceName() *string
	SetPolicy(v string) *UpdateInstancePolicyRequest
	GetPolicy() *string
	SetPolicyVersion(v int64) *UpdateInstancePolicyRequest
	GetPolicyVersion() *int64
}

type UpdateInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-test-12345
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance policy in the JSON format.
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
	//                 "acs:ots:*:13791xxxxxxxxxxx:instance/myinstance*"
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
	// The version of the instance policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyVersion *int64 `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
}

func (s UpdateInstancePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstancePolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateInstancePolicyRequest) GetPolicyVersion() *int64 {
	return s.PolicyVersion
}

func (s *UpdateInstancePolicyRequest) SetInstanceName(v string) *UpdateInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstancePolicyRequest) SetPolicy(v string) *UpdateInstancePolicyRequest {
	s.Policy = &v
	return s
}

func (s *UpdateInstancePolicyRequest) SetPolicyVersion(v int64) *UpdateInstancePolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *UpdateInstancePolicyRequest) Validate() error {
	return dara.Validate(s)
}
