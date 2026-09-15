// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAigcChatCompletionStreamHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AigcChatCompletionStreamHeaders
	GetCommonHeaders() map[string]*string
	SetXQIAgentApiKey(v string) *AigcChatCompletionStreamHeaders
	GetXQIAgentApiKey() *string
	SetXQIInstanceId(v string) *AigcChatCompletionStreamHeaders
	GetXQIInstanceId() *string
	SetXQISessionId(v string) *AigcChatCompletionStreamHeaders
	GetXQISessionId() *string
}

type AigcChatCompletionStreamHeaders struct {
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
	// AIGC_xxx
	XQIInstanceId *string `json:"X-QI-Instance-Id,omitempty" xml:"X-QI-Instance-Id,omitempty"`
	// example:
	//
	// session-xxx
	XQISessionId *string `json:"X-QI-Session-Id,omitempty" xml:"X-QI-Session-Id,omitempty"`
}

func (s AigcChatCompletionStreamHeaders) String() string {
	return dara.Prettify(s)
}

func (s AigcChatCompletionStreamHeaders) GoString() string {
	return s.String()
}

func (s *AigcChatCompletionStreamHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AigcChatCompletionStreamHeaders) GetXQIAgentApiKey() *string {
	return s.XQIAgentApiKey
}

func (s *AigcChatCompletionStreamHeaders) GetXQIInstanceId() *string {
	return s.XQIInstanceId
}

func (s *AigcChatCompletionStreamHeaders) GetXQISessionId() *string {
	return s.XQISessionId
}

func (s *AigcChatCompletionStreamHeaders) SetCommonHeaders(v map[string]*string) *AigcChatCompletionStreamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AigcChatCompletionStreamHeaders) SetXQIAgentApiKey(v string) *AigcChatCompletionStreamHeaders {
	s.XQIAgentApiKey = &v
	return s
}

func (s *AigcChatCompletionStreamHeaders) SetXQIInstanceId(v string) *AigcChatCompletionStreamHeaders {
	s.XQIInstanceId = &v
	return s
}

func (s *AigcChatCompletionStreamHeaders) SetXQISessionId(v string) *AigcChatCompletionStreamHeaders {
	s.XQISessionId = &v
	return s
}

func (s *AigcChatCompletionStreamHeaders) Validate() error {
	return dara.Validate(s)
}
