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
	// example:
	//
	// 1021418411
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                          `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *PromptAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
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
	Meta   *PromptAgentSessionRequestParamsMeta     `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	Prompt []*PromptAgentSessionRequestParamsPrompt `json:"Prompt,omitempty" xml:"Prompt,omitempty" type:"Repeated"`
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
	Description *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Meta        *PromptAgentSessionRequestParamsPromptMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	MimeType    *string                                    `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	Name        *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	Size        *int64                                     `json:"Size,omitempty" xml:"Size,omitempty"`
	Text        *string                                    `json:"Text,omitempty" xml:"Text,omitempty"`
	Title       *string                                    `json:"Title,omitempty" xml:"Title,omitempty"`
	Type        *string                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	Uri         *string                                    `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	// example:
	//
	// true or false
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
