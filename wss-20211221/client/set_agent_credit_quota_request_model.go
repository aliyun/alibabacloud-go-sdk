// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentCreditQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v []*string) *SetAgentCreditQuotaRequest
	GetAgentIds() []*string
	SetAgentType(v string) *SetAgentCreditQuotaRequest
	GetAgentType() *string
	SetBizType(v string) *SetAgentCreditQuotaRequest
	GetBizType() *string
	SetCreditQuota(v int32) *SetAgentCreditQuotaRequest
	GetCreditQuota() *int32
}

type SetAgentCreditQuotaRequest struct {
	// A list of Agent IDs.
	AgentIds []*string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// The Agent type.
	//
	// example:
	//
	// JVSClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The business type.
	//
	// example:
	//
	// BUSINESS
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The credit quota.
	//
	// example:
	//
	// 200
	CreditQuota *int32 `json:"CreditQuota,omitempty" xml:"CreditQuota,omitempty"`
}

func (s SetAgentCreditQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAgentCreditQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetAgentCreditQuotaRequest) GetAgentIds() []*string {
	return s.AgentIds
}

func (s *SetAgentCreditQuotaRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *SetAgentCreditQuotaRequest) GetBizType() *string {
	return s.BizType
}

func (s *SetAgentCreditQuotaRequest) GetCreditQuota() *int32 {
	return s.CreditQuota
}

func (s *SetAgentCreditQuotaRequest) SetAgentIds(v []*string) *SetAgentCreditQuotaRequest {
	s.AgentIds = v
	return s
}

func (s *SetAgentCreditQuotaRequest) SetAgentType(v string) *SetAgentCreditQuotaRequest {
	s.AgentType = &v
	return s
}

func (s *SetAgentCreditQuotaRequest) SetBizType(v string) *SetAgentCreditQuotaRequest {
	s.BizType = &v
	return s
}

func (s *SetAgentCreditQuotaRequest) SetCreditQuota(v int32) *SetAgentCreditQuotaRequest {
	s.CreditQuota = &v
	return s
}

func (s *SetAgentCreditQuotaRequest) Validate() error {
	return dara.Validate(s)
}
