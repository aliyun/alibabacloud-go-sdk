// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardingRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListForwardingRulesRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *ListForwardingRulesRequest
	GetClientToken() *string
	SetForwardingRuleId(v string) *ListForwardingRulesRequest
	GetForwardingRuleId() *string
	SetListenerId(v string) *ListForwardingRulesRequest
	GetListenerId() *string
	SetMaxResults(v int32) *ListForwardingRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListForwardingRulesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListForwardingRulesRequest
	GetRegionId() *string
}

type ListForwardingRulesRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qzk****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// frule-bp19a3t3yzr21q3****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1s0vzbi5bxlx5pw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If this is your first query or no subsequent query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListForwardingRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListForwardingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListForwardingRulesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListForwardingRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListForwardingRulesRequest) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *ListForwardingRulesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListForwardingRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListForwardingRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListForwardingRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListForwardingRulesRequest) SetAcceleratorId(v string) *ListForwardingRulesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetClientToken(v string) *ListForwardingRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetForwardingRuleId(v string) *ListForwardingRulesRequest {
	s.ForwardingRuleId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetListenerId(v string) *ListForwardingRulesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListForwardingRulesRequest) SetMaxResults(v int32) *ListForwardingRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListForwardingRulesRequest) SetNextToken(v string) *ListForwardingRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListForwardingRulesRequest) SetRegionId(v string) *ListForwardingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListForwardingRulesRequest) Validate() error {
	return dara.Validate(s)
}
