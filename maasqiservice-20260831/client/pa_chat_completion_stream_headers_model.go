// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaChatCompletionStreamHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PaChatCompletionStreamHeaders
	GetCommonHeaders() map[string]*string
	SetXQIAgentApiKey(v string) *PaChatCompletionStreamHeaders
	GetXQIAgentApiKey() *string
	SetXQIInstanceId(v string) *PaChatCompletionStreamHeaders
	GetXQIInstanceId() *string
	SetXQISessionId(v string) *PaChatCompletionStreamHeaders
	GetXQISessionId() *string
}

type PaChatCompletionStreamHeaders struct {
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
	// PA_xxx
	XQIInstanceId *string `json:"X-QI-Instance-Id,omitempty" xml:"X-QI-Instance-Id,omitempty"`
	// example:
	//
	// session-xxx
	XQISessionId *string `json:"X-QI-Session-Id,omitempty" xml:"X-QI-Session-Id,omitempty"`
}

func (s PaChatCompletionStreamHeaders) String() string {
	return dara.Prettify(s)
}

func (s PaChatCompletionStreamHeaders) GoString() string {
	return s.String()
}

func (s *PaChatCompletionStreamHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PaChatCompletionStreamHeaders) GetXQIAgentApiKey() *string {
	return s.XQIAgentApiKey
}

func (s *PaChatCompletionStreamHeaders) GetXQIInstanceId() *string {
	return s.XQIInstanceId
}

func (s *PaChatCompletionStreamHeaders) GetXQISessionId() *string {
	return s.XQISessionId
}

func (s *PaChatCompletionStreamHeaders) SetCommonHeaders(v map[string]*string) *PaChatCompletionStreamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PaChatCompletionStreamHeaders) SetXQIAgentApiKey(v string) *PaChatCompletionStreamHeaders {
	s.XQIAgentApiKey = &v
	return s
}

func (s *PaChatCompletionStreamHeaders) SetXQIInstanceId(v string) *PaChatCompletionStreamHeaders {
	s.XQIInstanceId = &v
	return s
}

func (s *PaChatCompletionStreamHeaders) SetXQISessionId(v string) *PaChatCompletionStreamHeaders {
	s.XQISessionId = &v
	return s
}

func (s *PaChatCompletionStreamHeaders) Validate() error {
	return dara.Validate(s)
}
