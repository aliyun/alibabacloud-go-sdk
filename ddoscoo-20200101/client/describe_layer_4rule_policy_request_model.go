// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *DescribeLayer4RulePolicyRequest
	GetListeners() *string
}

type DescribeLayer4RulePolicyRequest struct {
	// The port forwarding rule that you want to query.
	//
	// This parameter is a string that consists of JSON arrays. Each element in a JSON array indicates a port forwarding rule. You can query only one port forwarding rule at a time.
	//
	// > You can call the [DescribeNetworkRules](https://help.aliyun.com/document_detail/157484.html) to query existing port forwarding rules.
	//
	// Each port forwarding rule contains the following fields:
	//
	// 	- **InstanceId**: the ID of the instance. This field is required and must be of the string type.
	//
	// 	- **Protocol**: the forwarding protocol. This field is required and must be of the string type. Valid values: **tcp*	- and **udp**.
	//
	// 	- **FrontendPort**: the forwarding port. This field is required and must be of the integer type.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"InstanceId\\":\\"ddosDip-sg-4hr2b3l****\\",\\"FrontendPort\\":2020,\\"Protocol\\":\\"udp\\"}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s DescribeLayer4RulePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulePolicyRequest) GetListeners() *string {
	return s.Listeners
}

func (s *DescribeLayer4RulePolicyRequest) SetListeners(v string) *DescribeLayer4RulePolicyRequest {
	s.Listeners = &v
	return s
}

func (s *DescribeLayer4RulePolicyRequest) Validate() error {
	return dara.Validate(s)
}
