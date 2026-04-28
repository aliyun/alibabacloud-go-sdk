// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemorySkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMemorySkillResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMemorySkillResponseBody
	GetRequestId() *string
	SetResult(v *GetMemorySkillResponseBodyResult) *GetMemorySkillResponseBody
	GetResult() *GetMemorySkillResponseBodyResult
	SetStatus(v string) *GetMemorySkillResponseBody
	GetStatus() *string
}

type GetMemorySkillResponseBody struct {
	Latency   *int32                            `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                           `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMemorySkillResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                           `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMemorySkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemorySkillResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMemorySkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemorySkillResponseBody) GetResult() *GetMemorySkillResponseBodyResult {
	return s.Result
}

func (s *GetMemorySkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetMemorySkillResponseBody) SetLatency(v int32) *GetMemorySkillResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMemorySkillResponseBody) SetRequestId(v string) *GetMemorySkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemorySkillResponseBody) SetResult(v *GetMemorySkillResponseBodyResult) *GetMemorySkillResponseBody {
	s.Result = v
	return s
}

func (s *GetMemorySkillResponseBody) SetStatus(v string) *GetMemorySkillResponseBody {
	s.Status = &v
	return s
}

func (s *GetMemorySkillResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemorySkillResponseBodyResult struct {
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Files   *string `json:"files,omitempty" xml:"files,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	SkillId *string `json:"skill_id,omitempty" xml:"skill_id,omitempty"`
	UserId  *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetMemorySkillResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMemorySkillResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMemorySkillResponseBodyResult) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemorySkillResponseBodyResult) GetFiles() *string {
	return s.Files
}

func (s *GetMemorySkillResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetMemorySkillResponseBodyResult) GetSkillId() *string {
	return s.SkillId
}

func (s *GetMemorySkillResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *GetMemorySkillResponseBodyResult) SetAgentId(v string) *GetMemorySkillResponseBodyResult {
	s.AgentId = &v
	return s
}

func (s *GetMemorySkillResponseBodyResult) SetFiles(v string) *GetMemorySkillResponseBodyResult {
	s.Files = &v
	return s
}

func (s *GetMemorySkillResponseBodyResult) SetName(v string) *GetMemorySkillResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetMemorySkillResponseBodyResult) SetSkillId(v string) *GetMemorySkillResponseBodyResult {
	s.SkillId = &v
	return s
}

func (s *GetMemorySkillResponseBodyResult) SetUserId(v string) *GetMemorySkillResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetMemorySkillResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
