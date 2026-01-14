// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DeleteForwardingRulesRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *DeleteForwardingRulesRequest
	GetClientToken() *string
	SetForwardingRuleIds(v []*string) *DeleteForwardingRulesRequest
	GetForwardingRuleIds() []*string
	SetListenerId(v string) *DeleteForwardingRulesRequest
	GetListenerId() *string
	SetRegionId(v string) *DeleteForwardingRulesRequest
	GetRegionId() *string
}

type DeleteForwardingRulesRequest struct {
	// The GA instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4q****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The forwarding rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// frule-bp19a3t3yzr21q3****
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteForwardingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DeleteForwardingRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteForwardingRulesRequest) GetForwardingRuleIds() []*string {
	return s.ForwardingRuleIds
}

func (s *DeleteForwardingRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *DeleteForwardingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteForwardingRulesRequest) SetAcceleratorId(v string) *DeleteForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetClientToken(v string) *DeleteForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetForwardingRuleIds(v []*string) *DeleteForwardingRulesRequest {
	s.ForwardingRuleIds = v
	return s
}

func (s *DeleteForwardingRulesRequest) SetListenerId(v string) *DeleteForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) SetRegionId(v string) *DeleteForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteForwardingRulesRequest) Validate() error {
	return dara.Validate(s)
}
