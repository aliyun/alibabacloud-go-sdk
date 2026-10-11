// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRuleGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyRuleGroupId(v string) *CreateEgressProxyRuleGroupResponseBody
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRules(v *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) *CreateEgressProxyRuleGroupResponseBody
	GetEgressProxyRules() *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules
	SetRequestId(v string) *CreateEgressProxyRuleGroupResponseBody
	GetRequestId() *string
}

type CreateEgressProxyRuleGroupResponseBody struct {
	// The ID of the egress proxy rule group.
	//
	// example:
	//
	// eprg-uf60966a9633667c51d7
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// The list of egress rules grouped by phase (PreDns, PreRequest, and PostResponse).
	EgressProxyRules *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules `json:"EgressProxyRules,omitempty" xml:"EgressProxyRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 019F3A67-2D1B-537B-B08D-B2DEB69E8E8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEgressProxyRuleGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponseBody) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *CreateEgressProxyRuleGroupResponseBody) GetEgressProxyRules() *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules {
	return s.EgressProxyRules
}

func (s *CreateEgressProxyRuleGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEgressProxyRuleGroupResponseBody) SetEgressProxyRuleGroupId(v string) *CreateEgressProxyRuleGroupResponseBody {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBody) SetEgressProxyRules(v *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) *CreateEgressProxyRuleGroupResponseBody {
	s.EgressProxyRules = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBody) SetRequestId(v string) *CreateEgressProxyRuleGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBody) Validate() error {
	if s.EgressProxyRules != nil {
		if err := s.EgressProxyRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEgressProxyRuleGroupResponseBodyEgressProxyRules struct {
	// The list of rules in the PostResponse stage, which is executed after the response is returned.
	PostResponse []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse `json:"PostResponse,omitempty" xml:"PostResponse,omitempty" type:"Repeated"`
	// The list of rules in the PreDns stage, which is executed before DNS resolution.
	PreDns []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns `json:"PreDns,omitempty" xml:"PreDns,omitempty" type:"Repeated"`
	// The list of rules in the PreRequest stage, which is executed before the request is forwarded.
	PreRequest []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest `json:"PreRequest,omitempty" xml:"PreRequest,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) GetPostResponse() []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse {
	return s.PostResponse
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) GetPreDns() []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns {
	return s.PreDns
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) GetPreRequest() []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest {
	return s.PreRequest
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) SetPostResponse(v []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules {
	s.PostResponse = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) SetPreDns(v []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules {
	s.PreDns = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) SetPreRequest(v []*CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules {
	s.PreRequest = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRules) Validate() error {
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

type CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse struct {
	// The ID of the egress proxy rule.
	//
	// example:
	//
	// epr-g2w80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress proxy rule in the stage to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) SetEgressProxyRuleId(v string) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) SetPriority(v int32) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPostResponse) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns struct {
	// The ID of the egress proxy rule.
	//
	// example:
	//
	// epr-gw80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress proxy rule in the stage to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) SetEgressProxyRuleId(v string) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) SetPriority(v int32) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreDns) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest struct {
	// The ID of the egress proxy rule.
	//
	// example:
	//
	// epr-g1w80bb6a91dbb775****
	EgressProxyRuleId *string `json:"EgressProxyRuleId,omitempty" xml:"EgressProxyRuleId,omitempty"`
	// The priority of the egress proxy rule in the stage to which it belongs. Valid values: 1 to 1000. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) GetEgressProxyRuleId() *string {
	return s.EgressProxyRuleId
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) SetEgressProxyRuleId(v string) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest {
	s.EgressProxyRuleId = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) SetPriority(v int32) *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest {
	s.Priority = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponseBodyEgressProxyRulesPreRequest) Validate() error {
	return dara.Validate(s)
}
