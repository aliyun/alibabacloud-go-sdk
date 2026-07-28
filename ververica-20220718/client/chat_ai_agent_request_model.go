// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatAiAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHitlDecisions(v []*ChatAiAgentRequestHitlDecisions) *ChatAiAgentRequest
	GetHitlDecisions() []*ChatAiAgentRequestHitlDecisions
	SetRefs(v *ChatAiAgentRequestRefs) *ChatAiAgentRequest
	GetRefs() *ChatAiAgentRequestRefs
	SetSessionId(v string) *ChatAiAgentRequest
	GetSessionId() *string
	SetUserMessage(v string) *ChatAiAgentRequest
	GetUserMessage() *string
}

type ChatAiAgentRequest struct {
	// The list of Human-in-the-Loop (HITL) approval decisions, used to resume a session interrupted by a hitlPending event.
	HitlDecisions []*ChatAiAgentRequestHitlDecisions `json:"hitlDecisions,omitempty" xml:"hitlDecisions,omitempty" type:"Repeated"`
	// The resource references, including jobs and skill lists.
	Refs *ChatAiAgentRequestRefs `json:"refs,omitempty" xml:"refs,omitempty" type:"Struct"`
	// The session ID. If not specified, the server generates one. For multi-turn conversations, pass the same value across requests.
	//
	// example:
	//
	// 019F8CC7-EAD3-5E06-B0BF-3A2A0638B3DD-deliverData-20260723102220-VM8X0A5VZQ
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user natural language input. The value must be 1 to 64,000 characters in length.
	//
	// example:
	//
	// “”
	UserMessage *string `json:"userMessage,omitempty" xml:"userMessage,omitempty"`
}

func (s ChatAiAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentRequest) GoString() string {
	return s.String()
}

func (s *ChatAiAgentRequest) GetHitlDecisions() []*ChatAiAgentRequestHitlDecisions {
	return s.HitlDecisions
}

func (s *ChatAiAgentRequest) GetRefs() *ChatAiAgentRequestRefs {
	return s.Refs
}

func (s *ChatAiAgentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatAiAgentRequest) GetUserMessage() *string {
	return s.UserMessage
}

func (s *ChatAiAgentRequest) SetHitlDecisions(v []*ChatAiAgentRequestHitlDecisions) *ChatAiAgentRequest {
	s.HitlDecisions = v
	return s
}

func (s *ChatAiAgentRequest) SetRefs(v *ChatAiAgentRequestRefs) *ChatAiAgentRequest {
	s.Refs = v
	return s
}

func (s *ChatAiAgentRequest) SetSessionId(v string) *ChatAiAgentRequest {
	s.SessionId = &v
	return s
}

func (s *ChatAiAgentRequest) SetUserMessage(v string) *ChatAiAgentRequest {
	s.UserMessage = &v
	return s
}

func (s *ChatAiAgentRequest) Validate() error {
	if s.HitlDecisions != nil {
		for _, item := range s.HitlDecisions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Refs != nil {
		if err := s.Refs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatAiAgentRequestHitlDecisions struct {
	// The approval decision. Valid values: approve and deny.
	//
	// example:
	//
	// approve
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The approval item ID corresponding to the hitlPending event.
	//
	// example:
	//
	// -
	HitlId *string `json:"hitlId,omitempty" xml:"hitlId,omitempty"`
}

func (s ChatAiAgentRequestHitlDecisions) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentRequestHitlDecisions) GoString() string {
	return s.String()
}

func (s *ChatAiAgentRequestHitlDecisions) GetDecision() *string {
	return s.Decision
}

func (s *ChatAiAgentRequestHitlDecisions) GetHitlId() *string {
	return s.HitlId
}

func (s *ChatAiAgentRequestHitlDecisions) SetDecision(v string) *ChatAiAgentRequestHitlDecisions {
	s.Decision = &v
	return s
}

func (s *ChatAiAgentRequestHitlDecisions) SetHitlId(v string) *ChatAiAgentRequestHitlDecisions {
	s.HitlId = &v
	return s
}

func (s *ChatAiAgentRequestHitlDecisions) Validate() error {
	return dara.Validate(s)
}

type ChatAiAgentRequestRefs struct {
	// The list of job references.
	Jobs []*ChatAiAgentRequestRefsJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// The list of skills to inject.
	Skills []*string `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
}

func (s ChatAiAgentRequestRefs) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentRequestRefs) GoString() string {
	return s.String()
}

func (s *ChatAiAgentRequestRefs) GetJobs() []*ChatAiAgentRequestRefsJobs {
	return s.Jobs
}

func (s *ChatAiAgentRequestRefs) GetSkills() []*string {
	return s.Skills
}

func (s *ChatAiAgentRequestRefs) SetJobs(v []*ChatAiAgentRequestRefsJobs) *ChatAiAgentRequestRefs {
	s.Jobs = v
	return s
}

func (s *ChatAiAgentRequestRefs) SetSkills(v []*string) *ChatAiAgentRequestRefs {
	s.Skills = v
	return s
}

func (s *ChatAiAgentRequestRefs) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatAiAgentRequestRefsJobs struct {
	// Deployment ID
	//
	// example:
	//
	// 2a63abb7-7ae7-4902-9970-fe5cff4bd7c1
	DeploymentId *string `json:"deploymentId,omitempty" xml:"deploymentId,omitempty"`
	// Job ID
	//
	// example:
	//
	// ccb853c3-1d5a-438d-bf98-346815ad875a
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
}

func (s ChatAiAgentRequestRefsJobs) String() string {
	return dara.Prettify(s)
}

func (s ChatAiAgentRequestRefsJobs) GoString() string {
	return s.String()
}

func (s *ChatAiAgentRequestRefsJobs) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *ChatAiAgentRequestRefsJobs) GetJobId() *string {
	return s.JobId
}

func (s *ChatAiAgentRequestRefsJobs) SetDeploymentId(v string) *ChatAiAgentRequestRefsJobs {
	s.DeploymentId = &v
	return s
}

func (s *ChatAiAgentRequestRefsJobs) SetJobId(v string) *ChatAiAgentRequestRefsJobs {
	s.JobId = &v
	return s
}

func (s *ChatAiAgentRequestRefsJobs) Validate() error {
	return dara.Validate(s)
}
