// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyRules(v *CreateEgressProxyRulesResponseBodyEgressProxyRules) *CreateEgressProxyRulesResponseBody
	GetEgressProxyRules() *CreateEgressProxyRulesResponseBodyEgressProxyRules
	SetRequestId(v string) *CreateEgressProxyRulesResponseBody
	GetRequestId() *string
}

type CreateEgressProxyRulesResponseBody struct {
	// The list of egress rules grouped by effective phase (PreDns, PreRequest, or PostResponse).
	EgressProxyRules *CreateEgressProxyRulesResponseBodyEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEgressProxyRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponseBody) GetEgressProxyRules() *CreateEgressProxyRulesResponseBodyEgressProxyRules {
	return s.EgressProxyRules
}

func (s *CreateEgressProxyRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEgressProxyRulesResponseBody) SetEgressProxyRules(v *CreateEgressProxyRulesResponseBodyEgressProxyRules) *CreateEgressProxyRulesResponseBody {
	s.EgressProxyRules = v
	return s
}

func (s *CreateEgressProxyRulesResponseBody) SetRequestId(v string) *CreateEgressProxyRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBody) Validate() error {
	if s.EgressProxyRules != nil {
		if err := s.EgressProxyRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEgressProxyRulesResponseBodyEgressProxyRules struct {
	// The list of rules in the PostResponse phase (after the response is returned).
	PostResponse []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse `json:"PostResponse,omitempty" xml:"PostResponse,omitempty" type:"Repeated"`
	// The list of rules for the pre-DNS stage (before DNS resolution).
	PreDns []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns `json:"PreDns,omitempty" xml:"PreDns,omitempty" type:"Repeated"`
	// The list of rules in the PreRequest stage (before request forwarding).
	PreRequest []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest `json:"PreRequest,omitempty" xml:"PreRequest,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRules) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) GetPostResponse() []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse {
	return s.PostResponse
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) GetPreDns() []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns {
	return s.PreDns
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) GetPreRequest() []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest {
	return s.PreRequest
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) SetPostResponse(v []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) *CreateEgressProxyRulesResponseBodyEgressProxyRules {
	s.PostResponse = v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) SetPreDns(v []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) *CreateEgressProxyRulesResponseBodyEgressProxyRules {
	s.PreDns = v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) SetPreRequest(v []*CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) *CreateEgressProxyRulesResponseBodyEgressProxyRules {
	s.PreRequest = v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRules) Validate() error {
	if s.PostResponse != nil {
		for _, item := range s.PostResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreDns != nil {
		for _, item := range s.PreDns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PreRequest != nil {
		for _, item := range s.PreRequest {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse struct {
	// The ID of the egress rule.
	//
	// example:
	//
	// epr-gw80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress rule within its stage. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) SetEgressProxyRuleId(v string) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) SetPriority(v int32) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPostResponse) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns struct {
	// The ID of the egress rule.
	//
	// example:
	//
	// epr-gw80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress rule within its stage. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) SetEgressProxyRuleId(v string) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) SetPriority(v int32) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreDns) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest struct {
	// The ID of the egress rule.
	//
	// example:
	//
	// epr-gw80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress rule within its stage. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 10
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) SetEgressProxyRuleId(v string) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) SetPriority(v int32) *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRulesResponseBodyEgressProxyRulesPreRequest) Validate() error {
	return dara.Validate(s)
}
