// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNapalStreamMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v *SendNapalStreamMessageResponseBodyMessage) *SendNapalStreamMessageResponseBody
	GetMessage() *SendNapalStreamMessageResponseBodyMessage
	SetRequestId(v string) *SendNapalStreamMessageResponseBody
	GetRequestId() *string
	SetTask(v *SendNapalStreamMessageResponseBodyTask) *SendNapalStreamMessageResponseBody
	GetTask() *SendNapalStreamMessageResponseBodyTask
	SetTaskArtifactUpdate(v *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) *SendNapalStreamMessageResponseBody
	GetTaskArtifactUpdate() *SendNapalStreamMessageResponseBodyTaskArtifactUpdate
	SetTaskStatusUpdate(v *SendNapalStreamMessageResponseBodyTaskStatusUpdate) *SendNapalStreamMessageResponseBody
	GetTaskStatusUpdate() *SendNapalStreamMessageResponseBodyTaskStatusUpdate
}

type SendNapalStreamMessageResponseBody struct {
	// This field is mutually exclusive with Task, TaskStatusUpdate, and TaskArtifactUpdate. When this field is returned, no task is created, and the stream closes after sending one Message. This API does not currently return this type. This field is reserved for protocol compatibility only.
	Message *SendNapalStreamMessageResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Task *SendNapalStreamMessageResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The task artifact update object.
	TaskArtifactUpdate *SendNapalStreamMessageResponseBodyTaskArtifactUpdate `json:"TaskArtifactUpdate,omitempty" xml:"TaskArtifactUpdate,omitempty" type:"Struct"`
	// The task status update object.
	TaskStatusUpdate *SendNapalStreamMessageResponseBodyTaskStatusUpdate `json:"TaskStatusUpdate,omitempty" xml:"TaskStatusUpdate,omitempty" type:"Struct"`
}

func (s SendNapalStreamMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBody) GetMessage() *SendNapalStreamMessageResponseBodyMessage {
	return s.Message
}

func (s *SendNapalStreamMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendNapalStreamMessageResponseBody) GetTask() *SendNapalStreamMessageResponseBodyTask {
	return s.Task
}

func (s *SendNapalStreamMessageResponseBody) GetTaskArtifactUpdate() *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	return s.TaskArtifactUpdate
}

func (s *SendNapalStreamMessageResponseBody) GetTaskStatusUpdate() *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	return s.TaskStatusUpdate
}

func (s *SendNapalStreamMessageResponseBody) SetMessage(v *SendNapalStreamMessageResponseBodyMessage) *SendNapalStreamMessageResponseBody {
	s.Message = v
	return s
}

func (s *SendNapalStreamMessageResponseBody) SetRequestId(v string) *SendNapalStreamMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBody) SetTask(v *SendNapalStreamMessageResponseBodyTask) *SendNapalStreamMessageResponseBody {
	s.Task = v
	return s
}

func (s *SendNapalStreamMessageResponseBody) SetTaskArtifactUpdate(v *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) *SendNapalStreamMessageResponseBody {
	s.TaskArtifactUpdate = v
	return s
}

func (s *SendNapalStreamMessageResponseBody) SetTaskStatusUpdate(v *SendNapalStreamMessageResponseBodyTaskStatusUpdate) *SendNapalStreamMessageResponseBody {
	s.TaskStatusUpdate = v
	return s
}

func (s *SendNapalStreamMessageResponseBody) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	if s.TaskArtifactUpdate != nil {
		if err := s.TaskArtifactUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.TaskStatusUpdate != nil {
		if err := s.TaskStatusUpdate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyMessage struct {
	// The session context ID.
	//
	// example:
	//
	// context-07b0**bcc2
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	Extensions []*string `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The message ID.
	//
	// example:
	//
	// message-fd6e**9949
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The list of message content parts.
	Parts []*SendNapalStreamMessageResponseBodyMessageParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
	// Reserved field. This parameter is not returned by the current operation.
	ReferenceTaskIds []*string `json:"ReferenceTaskIds,omitempty" xml:"ReferenceTaskIds,omitempty" type:"Repeated"`
	// The message role.
	//
	// example:
	//
	// agent
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Reserved field. A directly returned Message does not create a task, so this field is empty. This API does not currently return a top-level Message.
	//
	// example:
	//
	// task-reserved
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyMessage) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetExtensions() []*string {
	return s.Extensions
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetMessageId() *string {
	return s.MessageId
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetParts() []*SendNapalStreamMessageResponseBodyMessageParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetReferenceTaskIds() []*string {
	return s.ReferenceTaskIds
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetRole() *string {
	return s.Role
}

func (s *SendNapalStreamMessageResponseBodyMessage) GetTaskId() *string {
	return s.TaskId
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetContextId(v string) *SendNapalStreamMessageResponseBodyMessage {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetExtensions(v []*string) *SendNapalStreamMessageResponseBodyMessage {
	s.Extensions = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetMessageId(v string) *SendNapalStreamMessageResponseBodyMessage {
	s.MessageId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageResponseBodyMessage {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetParts(v []*SendNapalStreamMessageResponseBodyMessageParts) *SendNapalStreamMessageResponseBodyMessage {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetReferenceTaskIds(v []*string) *SendNapalStreamMessageResponseBodyMessage {
	s.ReferenceTaskIds = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetRole(v string) *SendNapalStreamMessageResponseBodyMessage {
	s.Role = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) SetTaskId(v string) *SendNapalStreamMessageResponseBodyMessage {
	s.TaskId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessage) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyMessageParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The text content.
	//
	// example:
	//
	// The current instance is running normally
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyMessageParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyMessageParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyMessageParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyMessageParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyMessageParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyMessageParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetText(v string) *SendNapalStreamMessageResponseBodyMessageParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyMessageParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyMessageParts) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTask struct {
	// The list of task artifacts.
	Artifacts []*SendNapalStreamMessageResponseBodyTaskArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The session context ID. Used to maintain context continuity in multi-turn conversations.
	//
	// example:
	//
	// context-07b0**bcc2
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// The list of historical messages.
	History []*SendNapalStreamMessageResponseBodyTaskHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// task-38cZ**MAVKu
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The extended metadata, used to pass additional context information.
	//
	// example:
	//
	// {"usage":"{totalTokens=327672}"}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The task status object.
	Status *SendNapalStreamMessageResponseBodyTaskStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s SendNapalStreamMessageResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTask) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTask) GetArtifacts() []*SendNapalStreamMessageResponseBodyTaskArtifacts {
	return s.Artifacts
}

func (s *SendNapalStreamMessageResponseBodyTask) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageResponseBodyTask) GetHistory() []*SendNapalStreamMessageResponseBodyTaskHistory {
	return s.History
}

func (s *SendNapalStreamMessageResponseBodyTask) GetId() *string {
	return s.Id
}

func (s *SendNapalStreamMessageResponseBodyTask) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyTask) GetStatus() *SendNapalStreamMessageResponseBodyTaskStatus {
	return s.Status
}

func (s *SendNapalStreamMessageResponseBodyTask) SetArtifacts(v []*SendNapalStreamMessageResponseBodyTaskArtifacts) *SendNapalStreamMessageResponseBodyTask {
	s.Artifacts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) SetContextId(v string) *SendNapalStreamMessageResponseBodyTask {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) SetHistory(v []*SendNapalStreamMessageResponseBodyTaskHistory) *SendNapalStreamMessageResponseBodyTask {
	s.History = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) SetId(v string) *SendNapalStreamMessageResponseBodyTask {
	s.Id = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageResponseBodyTask {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) SetStatus(v *SendNapalStreamMessageResponseBodyTaskStatus) *SendNapalStreamMessageResponseBodyTask {
	s.Status = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTask) Validate() error {
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskArtifacts struct {
	// The unique identifier of the artifact.
	//
	// example:
	//
	// output
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Instance health inspection results
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	Extensions []*string `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the artifact.
	//
	// example:
	//
	// Inspection report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of content parts.
	Parts []*SendNapalStreamMessageResponseBodyTaskArtifactsParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
}

func (s SendNapalStreamMessageResponseBodyTaskArtifacts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskArtifacts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetDescription() *string {
	return s.Description
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetExtensions() []*string {
	return s.Extensions
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetName() *string {
	return s.Name
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) GetParts() []*SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetArtifactId(v string) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetDescription(v string) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.Description = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetExtensions(v []*string) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.Extensions = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetName(v string) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.Name = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) SetParts(v []*SendNapalStreamMessageResponseBodyTaskArtifactsParts) *SendNapalStreamMessageResponseBodyTaskArtifacts {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifacts) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskArtifactsParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The report text fragment.
	//
	// example:
	//
	// Diagnostic results
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactsParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactsParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetText(v string) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyTaskArtifactsParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactsParts) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTaskHistory struct {
	// The session context ID.
	//
	// example:
	//
	// context-07b0**bcc2
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	Extensions []*string `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The message ID.
	//
	// example:
	//
	// message-fd6e**9949
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The list of message content parts.
	Parts []*SendNapalStreamMessageResponseBodyTaskHistoryParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
	// Reserved field. This parameter is not returned by the current operation.
	ReferenceTaskIds []*string `json:"ReferenceTaskIds,omitempty" xml:"ReferenceTaskIds,omitempty" type:"Repeated"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-38cZ**MAVKu
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskHistory) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskHistory) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetExtensions() []*string {
	return s.Extensions
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetMessageId() *string {
	return s.MessageId
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetParts() []*SendNapalStreamMessageResponseBodyTaskHistoryParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetReferenceTaskIds() []*string {
	return s.ReferenceTaskIds
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetRole() *string {
	return s.Role
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) GetTaskId() *string {
	return s.TaskId
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetContextId(v string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetExtensions(v []*string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.Extensions = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetMessageId(v string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.MessageId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetParts(v []*SendNapalStreamMessageResponseBodyTaskHistoryParts) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetReferenceTaskIds(v []*string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.ReferenceTaskIds = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetRole(v string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.Role = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) SetTaskId(v string) *SendNapalStreamMessageResponseBodyTaskHistory {
	s.TaskId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistory) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskHistoryParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The text content.
	//
	// example:
	//
	// Diagnose this instance ngw-xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskHistoryParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskHistoryParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetText(v string) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyTaskHistoryParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskHistoryParts) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTaskStatus struct {
	// The message object defined by the A2A protocol. It contains the sender role, one or more content parts, and optional session and task context. When returned as a top-level field, it indicates a direct reply from the agent and is mutually exclusive with Task, TaskStatusUpdate, and TaskArtifactUpdate. The stream closes immediately after this message is returned. When returned as Status.Message, it represents a descriptive message associated with the task status. This API does not currently return a top-level Message. This field is reserved for protocol compatibility only.
	Message *SendNapalStreamMessageResponseBodyTaskStatusMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The task state. Valid values:
	//
	// - TASK_STATE_SUBMITTED: The task has been submitted.
	//
	// - TASK_STATE_WORKING: The task is being executed.
	//
	// - TASK_STATE_COMPLETED: The task has been completed.
	//
	// - TASK_STATE_FAILED: The task has failed.
	//
	// example:
	//
	// TASK_STATE_SUBMITTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status timestamp in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-07T06:08:10Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatus) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatus) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) GetMessage() *SendNapalStreamMessageResponseBodyTaskStatusMessage {
	return s.Message
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) GetState() *string {
	return s.State
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) GetTimestamp() *string {
	return s.Timestamp
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) SetMessage(v *SendNapalStreamMessageResponseBodyTaskStatusMessage) *SendNapalStreamMessageResponseBodyTaskStatus {
	s.Message = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) SetState(v string) *SendNapalStreamMessageResponseBodyTaskStatus {
	s.State = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) SetTimestamp(v string) *SendNapalStreamMessageResponseBodyTaskStatus {
	s.Timestamp = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatus) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusMessage struct {
	// The message ID.
	//
	// example:
	//
	// message-fd6e**9949
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The list of message content parts.
	Parts []*SendNapalStreamMessageResponseBodyTaskStatusMessageParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
	// The message role.
	//
	// example:
	//
	// agent
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusMessage) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusMessage) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) GetMessageId() *string {
	return s.MessageId
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) GetParts() []*SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) GetRole() *string {
	return s.Role
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) SetMessageId(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessage {
	s.MessageId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) SetParts(v []*SendNapalStreamMessageResponseBodyTaskStatusMessageParts) *SendNapalStreamMessageResponseBodyTaskStatusMessage {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) SetRole(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessage {
	s.Role = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessage) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusMessageParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The text content. The natural language instruction entered by the user, such as a diagnostic request or question consultation.
	//
	// example:
	//
	// The current instance is running normally
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusMessageParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetText(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyTaskStatusMessageParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusMessageParts) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTaskArtifactUpdate struct {
	// Indicates whether the content is appended. A value of `true` indicates that the current Text is appended to the end of the existing report content. A value of `false` indicates that the existing content is overwritten.
	//
	// example:
	//
	// true
	Append *bool `json:"Append,omitempty" xml:"Append,omitempty"`
	// The artifact object.
	Artifact *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact `json:"Artifact,omitempty" xml:"Artifact,omitempty" type:"Struct"`
	// The session context ID.
	//
	// example:
	//
	// context-07b0**bcc2
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Indicates whether this is the last chunk. A value of `true` indicates that the report content has been fully pushed and no more events will follow.
	//
	// example:
	//
	// false
	LastChunk *bool `json:"LastChunk,omitempty" xml:"LastChunk,omitempty"`
	// The task ID.
	//
	// example:
	//
	// task-38cZ**MAVKu
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdate) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GetAppend() *bool {
	return s.Append
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GetArtifact() *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	return s.Artifact
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GetLastChunk() *bool {
	return s.LastChunk
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) GetTaskId() *string {
	return s.TaskId
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) SetAppend(v bool) *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	s.Append = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) SetArtifact(v *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	s.Artifact = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) SetContextId(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) SetLastChunk(v bool) *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	s.LastChunk = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) SetTaskId(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdate {
	s.TaskId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdate) Validate() error {
	if s.Artifact != nil {
		if err := s.Artifact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact struct {
	// The unique identifier of the artifact.
	//
	// example:
	//
	// output
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Detailed inspection report
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	Extensions []*string `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the artifact.
	//
	// example:
	//
	// Inspection report
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of content parts.
	Parts []*SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetDescription() *string {
	return s.Description
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetExtensions() []*string {
	return s.Extensions
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetName() *string {
	return s.Name
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) GetParts() []*SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetArtifactId(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.ArtifactId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetDescription(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.Description = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetExtensions(v []*string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.Extensions = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetName(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.Name = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) SetParts(v []*SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifact) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The report text fragment.
	//
	// example:
	//
	// Instance status is normal
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetText(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskArtifactUpdateArtifactParts) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdate struct {
	// The session context ID.
	//
	// example:
	//
	// context-07b0**bcc2
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// Indicates whether this is a final event. A value of true indicates that the task has ended (completed or failed) and no more events will be pushed after this.
	//
	// example:
	//
	// false
	Final *bool `json:"Final,omitempty" xml:"Final,omitempty"`
	// The metadata object that contains step execution information.
	Metadata *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// The task status object.
	Status *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// The task ID.
	//
	// example:
	//
	// task-38cZ**MAVKu
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdate) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdate) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) GetFinal() *bool {
	return s.Final
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) GetMetadata() *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata {
	return s.Metadata
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) GetStatus() *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus {
	return s.Status
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) GetTaskId() *string {
	return s.TaskId
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) SetContextId(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) SetFinal(v bool) *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	s.Final = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) SetMetadata(v *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) SetStatus(v *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	s.Status = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) SetTaskId(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdate {
	s.TaskId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdate) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata struct {
	// The step execution information.
	Step *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep `json:"Step,omitempty" xml:"Step,omitempty" type:"Struct"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) GetStep() *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	return s.Step
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) SetStep(v *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata {
	s.Step = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadata) Validate() error {
	if s.Step != nil {
		if err := s.Step.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep struct {
	// The execution duration of the step. Unit: milliseconds.
	//
	// example:
	//
	// 203
	CostTime *int64 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The unique identifier of the step.
	//
	// example:
	//
	// 30688
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the step encountered an error during execution.
	//
	// example:
	//
	// false
	IsError *bool `json:"IsError,omitempty" xml:"IsError,omitempty"`
	// The step name.
	//
	// example:
	//
	// load_skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The input parameters of the step.
	//
	// example:
	//
	// {"skill_id":"239"}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// The execution result of the step.
	//
	// example:
	//
	// success
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// The number of retries.
	//
	// example:
	//
	// 1
	RetryCount *int64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	// The content displayed on the frontend.
	//
	// example:
	//
	// load_skill
	UiContent *string `json:"UiContent,omitempty" xml:"UiContent,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetCostTime() *int64 {
	return s.CostTime
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetId() *string {
	return s.Id
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetIsError() *bool {
	return s.IsError
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetName() *string {
	return s.Name
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetParams() interface{} {
	return s.Params
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetResult() interface{} {
	return s.Result
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetRetryCount() *int64 {
	return s.RetryCount
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) GetUiContent() *string {
	return s.UiContent
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetCostTime(v int64) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.CostTime = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetId(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.Id = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetIsError(v bool) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.IsError = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetName(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.Name = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetParams(v interface{}) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.Params = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetResult(v interface{}) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.Result = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetRetryCount(v int64) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.RetryCount = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) SetUiContent(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep {
	s.UiContent = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateMetadataStep) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus struct {
	// The message body object.
	Message *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - TASK_STATE_WORKING: The task is running.
	//
	// - TASK_STATE_COMPLETED: The task is completed.
	//
	// - TASK_STATE_FAILED: The task has failed.
	//
	// - TASK_STATE_CANCELED: The task is canceled.
	//
	// example:
	//
	// TASK_STATE_WORKING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status timestamp in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-07T06:08:30Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) GetMessage() *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage {
	return s.Message
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) GetState() *string {
	return s.State
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) GetTimestamp() *string {
	return s.Timestamp
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) SetMessage(v *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus {
	s.Message = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) SetState(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus {
	s.State = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) SetTimestamp(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus {
	s.Timestamp = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatus) Validate() error {
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage struct {
	// The message ID.
	//
	// example:
	//
	// message-fd6e**9949
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The list of message content parts.
	Parts []*SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) GetMessageId() *string {
	return s.MessageId
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) GetParts() []*SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	return s.Parts
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) GetRole() *string {
	return s.Role
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) SetMessageId(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage {
	s.MessageId = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) SetParts(v []*SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) SetRole(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage {
	s.Role = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessage) Validate() error {
	if s.Parts != nil {
		for _, item := range s.Parts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts struct {
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// reserved.bin
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// application/octet-stream
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The text content.
	//
	// example:
	//
	// Query traffic
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Reserved field. This parameter is not returned by the current operation.
	//
	// example:
	//
	// https://example.com/reserved.bin
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetData(v interface{}) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetFilename(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetMediaType(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetRaw(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetText(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) SetUrl(v string) *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageResponseBodyTaskStatusUpdateStatusMessageParts) Validate() error {
	return dara.Validate(s)
}
