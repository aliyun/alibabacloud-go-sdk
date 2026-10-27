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
	// The instance access key. The key must belong to the Creative Agent instance being called. This key is different from the AccessKey Secret used for OpenAPI signing. Do not expose it publicly.
	//
	// This parameter is required.
	//
	// example:
	//
	// qis_xxx
	XQIAgentApiKey *string `json:"X-QI-Agent-Api-Key,omitempty" xml:"X-QI-Agent-Api-Key,omitempty"`
	// The published Creative Agent instance ID. Must correspond to the X-QI-Agent-Api-Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// AIGC_xxx
	XQIInstanceId *string `json:"X-QI-Instance-Id,omitempty" xml:"X-QI-Instance-Id,omitempty"`
	// The session identifier. Can be omitted for the first turn. Obtain it from the qi_session_id field in the SSE response and pass it back in subsequent turns. Use no more than 128 characters consisting of letters, digits, periods, underscores, colons, or hyphens.
	//
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
