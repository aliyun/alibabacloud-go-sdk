// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNapalStreamMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v *SendNapalStreamMessageRequestConfiguration) *SendNapalStreamMessageRequest
	GetConfiguration() *SendNapalStreamMessageRequestConfiguration
	SetMessage(v *SendNapalStreamMessageRequestMessage) *SendNapalStreamMessageRequest
	GetMessage() *SendNapalStreamMessageRequestMessage
	SetMetadata(v map[string]*string) *SendNapalStreamMessageRequest
	GetMetadata() map[string]*string
}

type SendNapalStreamMessageRequest struct {
	// The request configuration object.
	Configuration *SendNapalStreamMessageRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The message object that contains user input and session context information.
	Message *SendNapalStreamMessageRequestMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
	// The additional request information.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s SendNapalStreamMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageRequest) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageRequest) GetConfiguration() *SendNapalStreamMessageRequestConfiguration {
	return s.Configuration
}

func (s *SendNapalStreamMessageRequest) GetMessage() *SendNapalStreamMessageRequestMessage {
	return s.Message
}

func (s *SendNapalStreamMessageRequest) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *SendNapalStreamMessageRequest) SetConfiguration(v *SendNapalStreamMessageRequestConfiguration) *SendNapalStreamMessageRequest {
	s.Configuration = v
	return s
}

func (s *SendNapalStreamMessageRequest) SetMessage(v *SendNapalStreamMessageRequestMessage) *SendNapalStreamMessageRequest {
	s.Message = v
	return s
}

func (s *SendNapalStreamMessageRequest) SetMetadata(v map[string]*string) *SendNapalStreamMessageRequest {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageRequest) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.Message != nil {
		if err := s.Message.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendNapalStreamMessageRequestConfiguration struct {
	// The accepted output modes. Default value: ["text/event-stream"], which indicates that SSE streaming responses are accepted.
	AcceptedOutputModes []*string `json:"AcceptedOutputModes,omitempty" xml:"AcceptedOutputModes,omitempty" type:"Repeated"`
	// The history message length. Controls the number of historical messages carried in multi-turn conversations. Default value: 20.
	//
	// example:
	//
	// 20
	HistoryLength *int32 `json:"HistoryLength,omitempty" xml:"HistoryLength,omitempty"`
	// Specifies whether to return immediately. Valid values:
	//
	// - false (default): Returns responses in streaming mode.
	//
	// - true: Returns the task ID immediately and processes the request asynchronously.
	//
	// example:
	//
	// false
	ReturnImmediately *bool `json:"ReturnImmediately,omitempty" xml:"ReturnImmediately,omitempty"`
}

func (s SendNapalStreamMessageRequestConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageRequestConfiguration) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageRequestConfiguration) GetAcceptedOutputModes() []*string {
	return s.AcceptedOutputModes
}

func (s *SendNapalStreamMessageRequestConfiguration) GetHistoryLength() *int32 {
	return s.HistoryLength
}

func (s *SendNapalStreamMessageRequestConfiguration) GetReturnImmediately() *bool {
	return s.ReturnImmediately
}

func (s *SendNapalStreamMessageRequestConfiguration) SetAcceptedOutputModes(v []*string) *SendNapalStreamMessageRequestConfiguration {
	s.AcceptedOutputModes = v
	return s
}

func (s *SendNapalStreamMessageRequestConfiguration) SetHistoryLength(v int32) *SendNapalStreamMessageRequestConfiguration {
	s.HistoryLength = &v
	return s
}

func (s *SendNapalStreamMessageRequestConfiguration) SetReturnImmediately(v bool) *SendNapalStreamMessageRequestConfiguration {
	s.ReturnImmediately = &v
	return s
}

func (s *SendNapalStreamMessageRequestConfiguration) Validate() error {
	return dara.Validate(s)
}

type SendNapalStreamMessageRequestMessage struct {
	// The session context ID. Do not specify this parameter for the first conversation. The server creates a new session. For multi-turn conversations, pass the contextId from the previous response to maintain context continuity.
	//
	// example:
	//
	// context-xxx
	ContextId *string `json:"ContextId,omitempty" xml:"ContextId,omitempty"`
	// The list of extension information.
	Extensions []*string `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The message ID. If not specified, the server automatically generates one.
	//
	// example:
	//
	// m_msijl2sv_pcfge8r7l
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The extended metadata, used to pass additional context information.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The list of message content parts. Multiple parts are supported.
	Parts []*SendNapalStreamMessageRequestMessageParts `json:"Parts,omitempty" xml:"Parts,omitempty" type:"Repeated"`
	// The list of referenced historical task IDs, used for context association.
	ReferenceTaskIds []*string `json:"ReferenceTaskIds,omitempty" xml:"ReferenceTaskIds,omitempty" type:"Repeated"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The task ID. Pass the ID of the previous task in follow-up conversation scenarios.
	//
	// example:
	//
	// task-xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SendNapalStreamMessageRequestMessage) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageRequestMessage) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageRequestMessage) GetContextId() *string {
	return s.ContextId
}

func (s *SendNapalStreamMessageRequestMessage) GetExtensions() []*string {
	return s.Extensions
}

func (s *SendNapalStreamMessageRequestMessage) GetMessageId() *string {
	return s.MessageId
}

func (s *SendNapalStreamMessageRequestMessage) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *SendNapalStreamMessageRequestMessage) GetParts() []*SendNapalStreamMessageRequestMessageParts {
	return s.Parts
}

func (s *SendNapalStreamMessageRequestMessage) GetReferenceTaskIds() []*string {
	return s.ReferenceTaskIds
}

func (s *SendNapalStreamMessageRequestMessage) GetRole() *string {
	return s.Role
}

func (s *SendNapalStreamMessageRequestMessage) GetTaskId() *string {
	return s.TaskId
}

func (s *SendNapalStreamMessageRequestMessage) SetContextId(v string) *SendNapalStreamMessageRequestMessage {
	s.ContextId = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetExtensions(v []*string) *SendNapalStreamMessageRequestMessage {
	s.Extensions = v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetMessageId(v string) *SendNapalStreamMessageRequestMessage {
	s.MessageId = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetMetadata(v map[string]interface{}) *SendNapalStreamMessageRequestMessage {
	s.Metadata = v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetParts(v []*SendNapalStreamMessageRequestMessageParts) *SendNapalStreamMessageRequestMessage {
	s.Parts = v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetReferenceTaskIds(v []*string) *SendNapalStreamMessageRequestMessage {
	s.ReferenceTaskIds = v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetRole(v string) *SendNapalStreamMessageRequestMessage {
	s.Role = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) SetTaskId(v string) *SendNapalStreamMessageRequestMessage {
	s.TaskId = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessage) Validate() error {
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

type SendNapalStreamMessageRequestMessageParts struct {
	// The structured data, used to pass JSON-formatted structured content.
	//
	// example:
	//
	// {"key":"value"}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The file name.
	//
	// example:
	//
	// report.txt
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// The media type.
	//
	// example:
	//
	// application/json
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The raw content, used to pass non-text data.
	//
	// example:
	//
	// cmVzZXJ2ZWQ=
	Raw *string `json:"Raw,omitempty" xml:"Raw,omitempty"`
	// The text content. The natural language instruction entered by the user, such as a diagnostic request or question consultation.
	//
	// example:
	//
	// Diagnose this instance ngw-xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The file URL, used to pass file-type content.
	//
	// example:
	//
	// https://example.com/file.txt
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SendNapalStreamMessageRequestMessageParts) String() string {
	return dara.Prettify(s)
}

func (s SendNapalStreamMessageRequestMessageParts) GoString() string {
	return s.String()
}

func (s *SendNapalStreamMessageRequestMessageParts) GetData() interface{} {
	return s.Data
}

func (s *SendNapalStreamMessageRequestMessageParts) GetFilename() *string {
	return s.Filename
}

func (s *SendNapalStreamMessageRequestMessageParts) GetMediaType() *string {
	return s.MediaType
}

func (s *SendNapalStreamMessageRequestMessageParts) GetRaw() *string {
	return s.Raw
}

func (s *SendNapalStreamMessageRequestMessageParts) GetText() *string {
	return s.Text
}

func (s *SendNapalStreamMessageRequestMessageParts) GetUrl() *string {
	return s.Url
}

func (s *SendNapalStreamMessageRequestMessageParts) SetData(v interface{}) *SendNapalStreamMessageRequestMessageParts {
	s.Data = v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) SetFilename(v string) *SendNapalStreamMessageRequestMessageParts {
	s.Filename = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) SetMediaType(v string) *SendNapalStreamMessageRequestMessageParts {
	s.MediaType = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) SetRaw(v string) *SendNapalStreamMessageRequestMessageParts {
	s.Raw = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) SetText(v string) *SendNapalStreamMessageRequestMessageParts {
	s.Text = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) SetUrl(v string) *SendNapalStreamMessageRequestMessageParts {
	s.Url = &v
	return s
}

func (s *SendNapalStreamMessageRequestMessageParts) Validate() error {
	return dara.Validate(s)
}
