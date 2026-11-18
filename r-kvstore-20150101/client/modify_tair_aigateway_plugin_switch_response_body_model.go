// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTairAIGatewayPluginSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyTairAIGatewayPluginSwitchResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ModifyTairAIGatewayPluginSwitchResponseBody
	GetRequestId() *string
}

type ModifyTairAIGatewayPluginSwitchResponseBody struct {
	// The ID of the Tair AI gateway instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTairAIGatewayPluginSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTairAIGatewayPluginSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTairAIGatewayPluginSwitchResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyTairAIGatewayPluginSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTairAIGatewayPluginSwitchResponseBody) SetInstanceId(v string) *ModifyTairAIGatewayPluginSwitchResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchResponseBody) SetRequestId(v string) *ModifyTairAIGatewayPluginSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTairAIGatewayPluginSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
