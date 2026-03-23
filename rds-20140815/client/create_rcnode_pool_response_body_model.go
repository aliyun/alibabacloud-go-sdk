// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdSets(v []*string) *CreateRCNodePoolResponseBody
	GetInstanceIdSets() []*string
	SetNodePoolId(v string) *CreateRCNodePoolResponseBody
	GetNodePoolId() *string
	SetOrderId(v string) *CreateRCNodePoolResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateRCNodePoolResponseBody
	GetRequestId() *string
}

type CreateRCNodePoolResponseBody struct {
	InstanceIdSets []*string `json:"InstanceIdSets,omitempty" xml:"InstanceIdSets,omitempty" type:"Repeated"`
	NodePoolId     *string   `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	OrderId        *string   `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRCNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRCNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRCNodePoolResponseBody) GetInstanceIdSets() []*string {
	return s.InstanceIdSets
}

func (s *CreateRCNodePoolResponseBody) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *CreateRCNodePoolResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateRCNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRCNodePoolResponseBody) SetInstanceIdSets(v []*string) *CreateRCNodePoolResponseBody {
	s.InstanceIdSets = v
	return s
}

func (s *CreateRCNodePoolResponseBody) SetNodePoolId(v string) *CreateRCNodePoolResponseBody {
	s.NodePoolId = &v
	return s
}

func (s *CreateRCNodePoolResponseBody) SetOrderId(v string) *CreateRCNodePoolResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateRCNodePoolResponseBody) SetRequestId(v string) *CreateRCNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRCNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
