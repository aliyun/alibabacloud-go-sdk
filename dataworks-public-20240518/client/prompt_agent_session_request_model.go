// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PromptAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v *PromptAgentSessionRequestParams) *PromptAgentSessionRequest
	GetParams() *PromptAgentSessionRequestParams
}

type PromptAgentSessionRequest struct {
	// The ID passed by the requester. The value is returned as is.
	//
	// example:
	//
	// 1021418411
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Set the value to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameters.
	Params *PromptAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s PromptAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionRequest) GetParams() *PromptAgentSessionRequestParams {
	return s.Params
}

func (s *PromptAgentSessionRequest) SetId(v string) *PromptAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionRequest) SetJsonrpc(v string) *PromptAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionRequest) SetParams(v *PromptAgentSessionRequestParams) *PromptAgentSessionRequest {
	s.Params = v
	return s
}

func (s *PromptAgentSessionRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromptAgentSessionRequestParams struct {
	// The extended metadata.
	//
	// 	Notice: If the name of the agent bound to the specified session is dataworks_ai_assistant_agent (AI Assistant), you must provide the instance ID of the AI Assistant in the Context.agent.instanceId field of the extended metadata.</notice>
	Meta *PromptAgentSessionRequestParamsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// The array of user message content blocks. For more information, see https://agentclientprotocol.com/protocol/content.
	Prompt []*PromptAgentSessionRequestParamsPrompt `json:"Prompt,omitempty" xml:"Prompt,omitempty" type:"Repeated"`
	// The ID of the target session. If the session does not exist, an SSE error frame is returned.
	//
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s PromptAgentSessionRequestParams) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequestParams) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequestParams) GetMeta() *PromptAgentSessionRequestParamsMeta {
	return s.Meta
}

func (s *PromptAgentSessionRequestParams) GetPrompt() []*PromptAgentSessionRequestParamsPrompt {
	return s.Prompt
}

func (s *PromptAgentSessionRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *PromptAgentSessionRequestParams) SetMeta(v *PromptAgentSessionRequestParamsMeta) *PromptAgentSessionRequestParams {
	s.Meta = v
	return s
}

func (s *PromptAgentSessionRequestParams) SetPrompt(v []*PromptAgentSessionRequestParamsPrompt) *PromptAgentSessionRequestParams {
	s.Prompt = v
	return s
}

func (s *PromptAgentSessionRequestParams) SetSessionId(v string) *PromptAgentSessionRequestParams {
	s.SessionId = &v
	return s
}

func (s *PromptAgentSessionRequestParams) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	if s.Prompt != nil {
		for _, item := range s.Prompt {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PromptAgentSessionRequestParamsMeta struct {
	// The map-type context. In custom agent scenarios, you can use this value to replace specific placeholder parameters.
	//
	// example:
	//
	// {
	//
	//    "key1": "value1",
	//
	//    "key2": "value2"
	//
	// }
	Context interface{} `json:"Context,omitempty" xml:"Context,omitempty"`
}

func (s PromptAgentSessionRequestParamsMeta) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequestParamsMeta) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequestParamsMeta) GetContext() interface{} {
	return s.Context
}

func (s *PromptAgentSessionRequestParamsMeta) SetContext(v interface{}) *PromptAgentSessionRequestParamsMeta {
	s.Context = v
	return s
}

func (s *PromptAgentSessionRequestParamsMeta) Validate() error {
	return dara.Validate(s)
}

type PromptAgentSessionRequestParamsPrompt struct {
	// The file description.
	//
	// example:
	//
	// Valid when Type is set to resource_link. Example: Sales order details.csv
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DataWorks extended prompt metadata.
	Meta *PromptAgentSessionRequestParamsPromptMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// The MIME type of the file.
	//
	// example:
	//
	// Valid when Type is set to resource_link. Example: text/csv
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	// The file name.
	//
	// example:
	//
	// Valid when Type is set to resource_link. Example: xxx.csv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// Valid only when Type is set to resource_link.
	//
	// example:
	//
	// 1231231
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// **Text content**
	//
	// example:
	//
	// Valid when Type is set to text. Example: Sales amount in the last 7 days
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The file title.
	//
	// example:
	//
	// Valid when Type is set to resource_link. Example: Sales order details.csv
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// **Content block type**
	//
	// example:
	//
	// Valid values: text and resource_link
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URI of the file.
	//
	// example:
	//
	// Valid when Type is set to resource_link. Example: oss://${bucket}/${ossKey}
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s PromptAgentSessionRequestParamsPrompt) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequestParamsPrompt) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequestParamsPrompt) GetDescription() *string {
	return s.Description
}

func (s *PromptAgentSessionRequestParamsPrompt) GetMeta() *PromptAgentSessionRequestParamsPromptMeta {
	return s.Meta
}

func (s *PromptAgentSessionRequestParamsPrompt) GetMimeType() *string {
	return s.MimeType
}

func (s *PromptAgentSessionRequestParamsPrompt) GetName() *string {
	return s.Name
}

func (s *PromptAgentSessionRequestParamsPrompt) GetSize() *int64 {
	return s.Size
}

func (s *PromptAgentSessionRequestParamsPrompt) GetText() *string {
	return s.Text
}

func (s *PromptAgentSessionRequestParamsPrompt) GetTitle() *string {
	return s.Title
}

func (s *PromptAgentSessionRequestParamsPrompt) GetType() *string {
	return s.Type
}

func (s *PromptAgentSessionRequestParamsPrompt) GetUri() *string {
	return s.Uri
}

func (s *PromptAgentSessionRequestParamsPrompt) SetDescription(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Description = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetMeta(v *PromptAgentSessionRequestParamsPromptMeta) *PromptAgentSessionRequestParamsPrompt {
	s.Meta = v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetMimeType(v string) *PromptAgentSessionRequestParamsPrompt {
	s.MimeType = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetName(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Name = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetSize(v int64) *PromptAgentSessionRequestParamsPrompt {
	s.Size = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetText(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Text = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetTitle(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Title = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetType(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Type = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) SetUri(v string) *PromptAgentSessionRequestParamsPrompt {
	s.Uri = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPrompt) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromptAgentSessionRequestParamsPromptMeta struct {
	// Specifies whether to hide the prompt from the user. For example, if a user asks a question in the chat dialog box, such as "Sales amount in the last 7 days", and the calling system retrieves some business domain knowledge through RAG before calling the API operation, the knowledge needs to be added to the agent context but not displayed to the user. In this case, you can set this parameter to true.
	//
	// example:
	//
	// true
	Hide *bool `json:"Hide,omitempty" xml:"Hide,omitempty"`
}

func (s PromptAgentSessionRequestParamsPromptMeta) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequestParamsPromptMeta) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequestParamsPromptMeta) GetHide() *bool {
	return s.Hide
}

func (s *PromptAgentSessionRequestParamsPromptMeta) SetHide(v bool) *PromptAgentSessionRequestParamsPromptMeta {
	s.Hide = &v
	return s
}

func (s *PromptAgentSessionRequestParamsPromptMeta) Validate() error {
	return dara.Validate(s)
}
