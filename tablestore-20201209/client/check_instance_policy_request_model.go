// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstancePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CheckInstancePolicyRequest
	GetInstanceName() *string
	SetPolicy(v string) *CheckInstancePolicyRequest
	GetPolicy() *string
}

type CheckInstancePolicyRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// workshop-bj-ots1
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
}

func (s CheckInstancePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstancePolicyRequest) GoString() string {
	return s.String()
}

func (s *CheckInstancePolicyRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CheckInstancePolicyRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CheckInstancePolicyRequest) SetInstanceName(v string) *CheckInstancePolicyRequest {
	s.InstanceName = &v
	return s
}

func (s *CheckInstancePolicyRequest) SetPolicy(v string) *CheckInstancePolicyRequest {
	s.Policy = &v
	return s
}

func (s *CheckInstancePolicyRequest) Validate() error {
	return dara.Validate(s)
}
