// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITaskEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAITaskEventsResponseBody
	GetCode() *string
	SetEvents(v []*ListAITaskEventsResponseBodyEvents) *ListAITaskEventsResponseBody
	GetEvents() []*ListAITaskEventsResponseBodyEvents
	SetHttpStatusCode(v int32) *ListAITaskEventsResponseBody
	GetHttpStatusCode() *int32
	SetNextToken(v string) *ListAITaskEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAITaskEventsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListAITaskEventsResponseBody
	GetSuccess() *string
	SetTaskId(v string) *ListAITaskEventsResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *ListAITaskEventsResponseBody
	GetTaskStatus() *string
	SetTaskType(v string) *ListAITaskEventsResponseBody
	GetTaskType() *string
}

type ListAITaskEventsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The events.
	Events []*ListAITaskEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAAAdDWBF2****w==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the AI task.
	//
	// 	- PENDING
	//
	// 	- WAITING
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAILURE
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the AI task.
	//
	// 	- GenerateTemplate: The AI task is used to generate a template.
	//
	// 	- FixTemplate: The AI task is used to fix a template.
	//
	// example:
	//
	// GenerateTemplate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListAITaskEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAITaskEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAITaskEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAITaskEventsResponseBody) GetEvents() []*ListAITaskEventsResponseBodyEvents {
	return s.Events
}

func (s *ListAITaskEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAITaskEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAITaskEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAITaskEventsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListAITaskEventsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAITaskEventsResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListAITaskEventsResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAITaskEventsResponseBody) SetCode(v string) *ListAITaskEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetEvents(v []*ListAITaskEventsResponseBodyEvents) *ListAITaskEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListAITaskEventsResponseBody) SetHttpStatusCode(v int32) *ListAITaskEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetNextToken(v string) *ListAITaskEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetRequestId(v string) *ListAITaskEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetSuccess(v string) *ListAITaskEventsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetTaskId(v string) *ListAITaskEventsResponseBody {
	s.TaskId = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetTaskStatus(v string) *ListAITaskEventsResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *ListAITaskEventsResponseBody) SetTaskType(v string) *ListAITaskEventsResponseBody {
	s.TaskType = &v
	return s
}

func (s *ListAITaskEventsResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAITaskEventsResponseBodyEvents struct {
	// The type of the agent that is used to execute the AI task.
	//
	// Valid values:
	//
	// 	- GenerateTemplateAgent
	//
	// 	- FixUserTemplateAgent
	//
	// example:
	//
	// GenerateTemplateAgent
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The time when the event was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-01T04:07:39
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The estimated execution time of the handler. Unit: seconds.
	//
	// example:
	//
	// 60
	EstimatedProcessingTime *string `json:"EstimatedProcessingTime,omitempty" xml:"EstimatedProcessingTime,omitempty"`
	// The details of the event.
	//
	// example:
	//
	// Document template generator started.
	EventData *string `json:"EventData,omitempty" xml:"EventData,omitempty"`
	// The execution state of the handler that process the AI task.
	//
	// Valid values:
	//
	// 	- SUCCESS
	//
	// 	- RUNNING
	//
	// 	- FAILURE
	//
	// example:
	//
	// RUNNING
	HandlerProcessStatus *string `json:"HandlerProcessStatus,omitempty" xml:"HandlerProcessStatus,omitempty"`
	// The type of the handler that is used to execute the task.
	//
	// Valid values:
	//
	// 	- TerraformTemplateGenerator
	//
	// 	- TemplateGenerator
	//
	// 	- ROSTemplateModifier
	//
	// 	- TerraformTemplateStaticFixer
	//
	// 	- TerraformTemplateDynamicFixer
	//
	// 	- DocumentTemplateGenerator
	//
	// 	- TerraformTemplateModifier
	//
	// 	- TemplateModifier
	//
	// 	- FixTemplateInputPreprocessor
	//
	// 	- TemplateStaticFixer
	//
	// 	- GenerateTemplateInputPreprocessor
	//
	// 	- ROSTemplateGenerator
	//
	// 	- TemplateDynamicFixer
	//
	// 	- BaseDynamicFixer
	//
	// 	- ROSTemplateStaticFixer
	//
	// 	- ROSTemplateDynamicFixer
	//
	// example:
	//
	// ROSTemplateGenerator
	HandlerType *string `json:"HandlerType,omitempty" xml:"HandlerType,omitempty"`
}

func (s ListAITaskEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAITaskEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListAITaskEventsResponseBodyEvents) GetAgentType() *string {
	return s.AgentType
}

func (s *ListAITaskEventsResponseBodyEvents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAITaskEventsResponseBodyEvents) GetEstimatedProcessingTime() *string {
	return s.EstimatedProcessingTime
}

func (s *ListAITaskEventsResponseBodyEvents) GetEventData() *string {
	return s.EventData
}

func (s *ListAITaskEventsResponseBodyEvents) GetHandlerProcessStatus() *string {
	return s.HandlerProcessStatus
}

func (s *ListAITaskEventsResponseBodyEvents) GetHandlerType() *string {
	return s.HandlerType
}

func (s *ListAITaskEventsResponseBodyEvents) SetAgentType(v string) *ListAITaskEventsResponseBodyEvents {
	s.AgentType = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) SetCreateTime(v string) *ListAITaskEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) SetEstimatedProcessingTime(v string) *ListAITaskEventsResponseBodyEvents {
	s.EstimatedProcessingTime = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) SetEventData(v string) *ListAITaskEventsResponseBodyEvents {
	s.EventData = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) SetHandlerProcessStatus(v string) *ListAITaskEventsResponseBodyEvents {
	s.HandlerProcessStatus = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) SetHandlerType(v string) *ListAITaskEventsResponseBodyEvents {
	s.HandlerType = &v
	return s
}

func (s *ListAITaskEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
