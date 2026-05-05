// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackageAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgents(v []*DescribeCreditPackageAgentsResponseBodyAgents) *DescribeCreditPackageAgentsResponseBody
	GetAgents() []*DescribeCreditPackageAgentsResponseBodyAgents
	SetMaxResults(v int32) *DescribeCreditPackageAgentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCreditPackageAgentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCreditPackageAgentsResponseBody
	GetRequestId() *string
}

type DescribeCreditPackageAgentsResponseBody struct {
	Agents []*DescribeCreditPackageAgentsResponseBodyAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCreditPackageAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageAgentsResponseBody) GetAgents() []*DescribeCreditPackageAgentsResponseBodyAgents {
	return s.Agents
}

func (s *DescribeCreditPackageAgentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCreditPackageAgentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCreditPackageAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreditPackageAgentsResponseBody) SetAgents(v []*DescribeCreditPackageAgentsResponseBodyAgents) *DescribeCreditPackageAgentsResponseBody {
	s.Agents = v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBody) SetMaxResults(v int32) *DescribeCreditPackageAgentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBody) SetNextToken(v string) *DescribeCreditPackageAgentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBody) SetRequestId(v string) *DescribeCreditPackageAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBody) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCreditPackageAgentsResponseBodyAgents struct {
	// Agent ID
	//
	// example:
	//
	// agent-abc
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 2026-04-01 10:00:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cp-inst-001
	CreditPackageId *string `json:"CreditPackageId,omitempty" xml:"CreditPackageId,omitempty"`
	// example:
	//
	// 2026-10-01 10:00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// jvs-copilot.standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 2000
	TotalCredit *int64 `json:"TotalCredit,omitempty" xml:"TotalCredit,omitempty"`
	// example:
	//
	// 100
	UsedCredit *int64 `json:"UsedCredit,omitempty" xml:"UsedCredit,omitempty"`
	// example:
	//
	// 80
	WarnPercent *int32 `json:"WarnPercent,omitempty" xml:"WarnPercent,omitempty"`
}

func (s DescribeCreditPackageAgentsResponseBodyAgents) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackageAgentsResponseBodyAgents) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetCreditPackageId() *string {
	return s.CreditPackageId
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetTotalCredit() *int64 {
	return s.TotalCredit
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetUsedCredit() *int64 {
	return s.UsedCredit
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) GetWarnPercent() *int32 {
	return s.WarnPercent
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetAgentId(v string) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.AgentId = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetCreatedTime(v string) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetCreditPackageId(v string) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.CreditPackageId = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetExpiredTime(v string) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetInstanceType(v string) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.InstanceType = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetTotalCredit(v int64) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.TotalCredit = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetUsedCredit(v int64) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.UsedCredit = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) SetWarnPercent(v int32) *DescribeCreditPackageAgentsResponseBodyAgents {
	s.WarnPercent = &v
	return s
}

func (s *DescribeCreditPackageAgentsResponseBodyAgents) Validate() error {
	return dara.Validate(s)
}
