// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGuiChatCompletionStreamHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GuiChatCompletionStreamHeaders
	GetCommonHeaders() map[string]*string
	SetXQIAgentApiKey(v string) *GuiChatCompletionStreamHeaders
	GetXQIAgentApiKey() *string
	SetXQIInstanceId(v string) *GuiChatCompletionStreamHeaders
	GetXQIInstanceId() *string
	SetXQISessionId(v string) *GuiChatCompletionStreamHeaders
	GetXQISessionId() *string
}

type GuiChatCompletionStreamHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qis_xxx
	XQIAgentApiKey *string `json:"X-QI-Agent-Api-Key,omitempty" xml:"X-QI-Agent-Api-Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// GUI_xxx
	XQIInstanceId *string `json:"X-QI-Instance-Id,omitempty" xml:"X-QI-Instance-Id,omitempty"`
	// example:
	//
	// session-xxx
	XQISessionId *string `json:"X-QI-Session-Id,omitempty" xml:"X-QI-Session-Id,omitempty"`
}

func (s GuiChatCompletionStreamHeaders) String() string {
	return dara.Prettify(s)
}

func (s GuiChatCompletionStreamHeaders) GoString() string {
	return s.String()
}

func (s *GuiChatCompletionStreamHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GuiChatCompletionStreamHeaders) GetXQIAgentApiKey() *string {
	return s.XQIAgentApiKey
}

func (s *GuiChatCompletionStreamHeaders) GetXQIInstanceId() *string {
	return s.XQIInstanceId
}

func (s *GuiChatCompletionStreamHeaders) GetXQISessionId() *string {
	return s.XQISessionId
}

func (s *GuiChatCompletionStreamHeaders) SetCommonHeaders(v map[string]*string) *GuiChatCompletionStreamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GuiChatCompletionStreamHeaders) SetXQIAgentApiKey(v string) *GuiChatCompletionStreamHeaders {
	s.XQIAgentApiKey = &v
	return s
}

func (s *GuiChatCompletionStreamHeaders) SetXQIInstanceId(v string) *GuiChatCompletionStreamHeaders {
	s.XQIInstanceId = &v
	return s
}

func (s *GuiChatCompletionStreamHeaders) SetXQISessionId(v string) *GuiChatCompletionStreamHeaders {
	s.XQISessionId = &v
	return s
}

func (s *GuiChatCompletionStreamHeaders) Validate() error {
	return dara.Validate(s)
}
