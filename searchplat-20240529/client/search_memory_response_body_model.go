// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *SearchMemoryResponseBody
	GetLatency() *int32
	SetRequestId(v string) *SearchMemoryResponseBody
	GetRequestId() *string
	SetResult(v *SearchMemoryResponseBodyResult) *SearchMemoryResponseBody
	GetResult() *SearchMemoryResponseBodyResult
	SetStatus(v string) *SearchMemoryResponseBody
	GetStatus() *string
}

type SearchMemoryResponseBody struct {
	Latency   *int32                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *SearchMemoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SearchMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *SearchMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMemoryResponseBody) GetResult() *SearchMemoryResponseBodyResult {
	return s.Result
}

func (s *SearchMemoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SearchMemoryResponseBody) SetLatency(v int32) *SearchMemoryResponseBody {
	s.Latency = &v
	return s
}

func (s *SearchMemoryResponseBody) SetRequestId(v string) *SearchMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMemoryResponseBody) SetResult(v *SearchMemoryResponseBodyResult) *SearchMemoryResponseBody {
	s.Result = v
	return s
}

func (s *SearchMemoryResponseBody) SetStatus(v string) *SearchMemoryResponseBody {
	s.Status = &v
	return s
}

func (s *SearchMemoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoryResponseBodyResult struct {
	Memory *SearchMemoryResponseBodyResultMemory `json:"memory,omitempty" xml:"memory,omitempty" type:"Struct"`
	Skill  *SearchMemoryResponseBodyResultSkill  `json:"skill,omitempty" xml:"skill,omitempty" type:"Struct"`
}

func (s SearchMemoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBodyResult) GetMemory() *SearchMemoryResponseBodyResultMemory {
	return s.Memory
}

func (s *SearchMemoryResponseBodyResult) GetSkill() *SearchMemoryResponseBodyResultSkill {
	return s.Skill
}

func (s *SearchMemoryResponseBodyResult) SetMemory(v *SearchMemoryResponseBodyResultMemory) *SearchMemoryResponseBodyResult {
	s.Memory = v
	return s
}

func (s *SearchMemoryResponseBodyResult) SetSkill(v *SearchMemoryResponseBodyResultSkill) *SearchMemoryResponseBodyResult {
	s.Skill = v
	return s
}

func (s *SearchMemoryResponseBodyResult) Validate() error {
	if s.Memory != nil {
		if err := s.Memory.Validate(); err != nil {
			return err
		}
	}
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoryResponseBodyResultMemory struct {
	Results []*SearchMemoryResponseBodyResultMemoryResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32                                         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchMemoryResponseBodyResultMemory) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBodyResultMemory) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBodyResultMemory) GetResults() []*SearchMemoryResponseBodyResultMemoryResults {
	return s.Results
}

func (s *SearchMemoryResponseBodyResultMemory) GetTotal() *int32 {
	return s.Total
}

func (s *SearchMemoryResponseBodyResultMemory) SetResults(v []*SearchMemoryResponseBodyResultMemoryResults) *SearchMemoryResponseBodyResultMemory {
	s.Results = v
	return s
}

func (s *SearchMemoryResponseBodyResultMemory) SetTotal(v int32) *SearchMemoryResponseBodyResultMemory {
	s.Total = &v
	return s
}

func (s *SearchMemoryResponseBodyResultMemory) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMemoryResponseBodyResultMemoryResults struct {
	AgentId  *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Memory   *string `json:"memory,omitempty" xml:"memory,omitempty"`
	MemoryId *string `json:"memory_id,omitempty" xml:"memory_id,omitempty"`
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s SearchMemoryResponseBodyResultMemoryResults) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBodyResultMemoryResults) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBodyResultMemoryResults) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoryResponseBodyResultMemoryResults) GetMemory() *string {
	return s.Memory
}

func (s *SearchMemoryResponseBodyResultMemoryResults) GetMemoryId() *string {
	return s.MemoryId
}

func (s *SearchMemoryResponseBodyResultMemoryResults) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoryResponseBodyResultMemoryResults) SetAgentId(v string) *SearchMemoryResponseBodyResultMemoryResults {
	s.AgentId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultMemoryResults) SetMemory(v string) *SearchMemoryResponseBodyResultMemoryResults {
	s.Memory = &v
	return s
}

func (s *SearchMemoryResponseBodyResultMemoryResults) SetMemoryId(v string) *SearchMemoryResponseBodyResultMemoryResults {
	s.MemoryId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultMemoryResults) SetUserId(v string) *SearchMemoryResponseBodyResultMemoryResults {
	s.UserId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultMemoryResults) Validate() error {
	return dara.Validate(s)
}

type SearchMemoryResponseBodyResultSkill struct {
	Results []*SearchMemoryResponseBodyResultSkillResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
	Total   *int32                                        `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchMemoryResponseBodyResultSkill) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBodyResultSkill) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBodyResultSkill) GetResults() []*SearchMemoryResponseBodyResultSkillResults {
	return s.Results
}

func (s *SearchMemoryResponseBodyResultSkill) GetTotal() *int32 {
	return s.Total
}

func (s *SearchMemoryResponseBodyResultSkill) SetResults(v []*SearchMemoryResponseBodyResultSkillResults) *SearchMemoryResponseBodyResultSkill {
	s.Results = v
	return s
}

func (s *SearchMemoryResponseBodyResultSkill) SetTotal(v int32) *SearchMemoryResponseBodyResultSkill {
	s.Total = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkill) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMemoryResponseBodyResultSkillResults struct {
	AgentId     *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	SkillId     *string `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SearchMemoryResponseBodyResultSkillResults) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoryResponseBodyResultSkillResults) GoString() string {
	return s.String()
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetDescription() *string {
	return s.Description
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetName() *string {
	return s.Name
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetSkillId() *string {
	return s.SkillId
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoryResponseBodyResultSkillResults) GetVersion() *string {
	return s.Version
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetAgentId(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.AgentId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetDescription(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.Description = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetName(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.Name = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetSkillId(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.SkillId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetUserId(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.UserId = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) SetVersion(v string) *SearchMemoryResponseBodyResultSkillResults {
	s.Version = &v
	return s
}

func (s *SearchMemoryResponseBodyResultSkillResults) Validate() error {
	return dara.Validate(s)
}
