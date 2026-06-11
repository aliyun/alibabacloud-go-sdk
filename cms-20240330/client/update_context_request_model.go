// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateContextRequest
	GetContent() *string
	SetExperience(v map[string]interface{}) *UpdateContextRequest
	GetExperience() map[string]interface{}
	SetMetadata(v map[string]interface{}) *UpdateContextRequest
	GetMetadata() map[string]interface{}
	SetPayload(v map[string]interface{}) *UpdateContextRequest
	GetPayload() map[string]interface{}
	SetTriggerCondition(v string) *UpdateContextRequest
	GetTriggerCondition() *string
}

type UpdateContextRequest struct {
	// The updated text for the long-term memory.
	//
	// example:
	//
	// Users prefer to first view the SLS error logs, index configuration, and the most recent Agent execution trace.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The experience object.
	//
	// example:
	//
	// {
	//
	// 	"taskType": "troubleshooting",
	//
	// 	"complexity": "medium",
	//
	// 	"confidence": 0.95
	//
	// }
	Experience map[string]interface{} `json:"experience,omitempty" xml:"experience,omitempty"`
	// A set of key-value pairs to attach to an object for storing custom information.
	//
	// example:
	//
	// {"sessionId":"test_session_001"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The payload to update.
	//
	// example:
	//
	// {
	//
	//     "userId": "u-10001",
	//
	//     "agentId": "sls-agent",
	//
	//     "appId": "console",
	//
	//     "categories": [
	//
	//       "preference"
	//
	//     ],
	//
	//     "source": "user_confirmed",
	//
	//     "topic": "debugging_preference",
	//
	//     "immutable": false,
	//
	//     "createdAt": 1776319200,
	//
	//     "updatedAt": 1776319200
	//
	//   }
	Payload map[string]interface{} `json:"payload,omitempty" xml:"payload,omitempty"`
	// The trigger condition.
	//
	// example:
	//
	// Identify and troubleshoot SLs issues
	TriggerCondition *string `json:"triggerCondition,omitempty" xml:"triggerCondition,omitempty"`
}

func (s UpdateContextRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextRequest) GoString() string {
	return s.String()
}

func (s *UpdateContextRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateContextRequest) GetExperience() map[string]interface{} {
	return s.Experience
}

func (s *UpdateContextRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateContextRequest) GetPayload() map[string]interface{} {
	return s.Payload
}

func (s *UpdateContextRequest) GetTriggerCondition() *string {
	return s.TriggerCondition
}

func (s *UpdateContextRequest) SetContent(v string) *UpdateContextRequest {
	s.Content = &v
	return s
}

func (s *UpdateContextRequest) SetExperience(v map[string]interface{}) *UpdateContextRequest {
	s.Experience = v
	return s
}

func (s *UpdateContextRequest) SetMetadata(v map[string]interface{}) *UpdateContextRequest {
	s.Metadata = v
	return s
}

func (s *UpdateContextRequest) SetPayload(v map[string]interface{}) *UpdateContextRequest {
	s.Payload = v
	return s
}

func (s *UpdateContextRequest) SetTriggerCondition(v string) *UpdateContextRequest {
	s.TriggerCondition = &v
	return s
}

func (s *UpdateContextRequest) Validate() error {
	return dara.Validate(s)
}
