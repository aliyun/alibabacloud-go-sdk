// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *ConfigLayer4RemarkRequest
	GetListeners() *string
}

type ConfigLayer4RemarkRequest struct {
	// The port forwarding rule that you want to manage.
	//
	// This parameter is a string that consists of JSON arrays. Each element in a JSON array indicates a port forwarding rule. You can perform this operation only on one port forwarding rule at a time.
	//
	// > You can call the [DescribeNetworkRules](https://help.aliyun.com/document_detail/157484.html) to query existing port forwarding rules.
	//
	// Each port forwarding rule contains the following fields:
	//
	// 	- **InstanceId**: the ID of the instance. This field is required and must be of the STRING type.
	//
	// 	- **Protocol**: the forwarding protocol. This field is required and must be of the STRING type. Valid values: **tcp*	- and **udp**.
	//
	// 	- **FrontendPort**: the forwarding port. This field is required and must be of the INTEGER type.
	//
	// 	- **Remark**: the remarks of the port forwarding rule. This field is required and must be of the STRING type. The value can contain letters, digits, and some special characters, such as `, . + - 	- / _`. The value can be up to 200 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"InstanceId\\":\\"ddosDip-sg-4hr2b3l****\\",\\"FrontendPort\\":2020,\\"Protocol\\":\\"udp\\",\\"Remark\\":\\"test\\"}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s ConfigLayer4RemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RemarkRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RemarkRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RemarkRequest) SetListeners(v string) *ConfigLayer4RemarkRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RemarkRequest) Validate() error {
	return dara.Validate(s)
}
