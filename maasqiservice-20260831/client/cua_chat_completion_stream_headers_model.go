// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCuaChatCompletionStreamHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CuaChatCompletionStreamHeaders
	GetCommonHeaders() map[string]*string
	SetXQIAgentApiKey(v string) *CuaChatCompletionStreamHeaders
	GetXQIAgentApiKey() *string
	SetXQIInstanceId(v string) *CuaChatCompletionStreamHeaders
	GetXQIInstanceId() *string
	SetXQISessionId(v string) *CuaChatCompletionStreamHeaders
	GetXQISessionId() *string
}

type CuaChatCompletionStreamHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qis_<instance_api_key>
	XQIAgentApiKey *string `json:"X-QI-Agent-Api-Key,omitempty" xml:"X-QI-Agent-Api-Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUA_<instance_id>
	XQIInstanceId *string `json:"X-QI-Instance-Id,omitempty" xml:"X-QI-Instance-Id,omitempty"`
	// example:
	//
	// session-cua-1
	XQISessionId *string `json:"X-QI-Session-Id,omitempty" xml:"X-QI-Session-Id,omitempty"`
}

func (s CuaChatCompletionStreamHeaders) String() string {
	return dara.Prettify(s)
}

func (s CuaChatCompletionStreamHeaders) GoString() string {
	return s.String()
}

func (s *CuaChatCompletionStreamHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CuaChatCompletionStreamHeaders) GetXQIAgentApiKey() *string {
	return s.XQIAgentApiKey
}

func (s *CuaChatCompletionStreamHeaders) GetXQIInstanceId() *string {
	return s.XQIInstanceId
}

func (s *CuaChatCompletionStreamHeaders) GetXQISessionId() *string {
	return s.XQISessionId
}

func (s *CuaChatCompletionStreamHeaders) SetCommonHeaders(v map[string]*string) *CuaChatCompletionStreamHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CuaChatCompletionStreamHeaders) SetXQIAgentApiKey(v string) *CuaChatCompletionStreamHeaders {
	s.XQIAgentApiKey = &v
	return s
}

func (s *CuaChatCompletionStreamHeaders) SetXQIInstanceId(v string) *CuaChatCompletionStreamHeaders {
	s.XQIInstanceId = &v
	return s
}

func (s *CuaChatCompletionStreamHeaders) SetXQISessionId(v string) *CuaChatCompletionStreamHeaders {
	s.XQISessionId = &v
	return s
}

func (s *CuaChatCompletionStreamHeaders) Validate() error {
	return dara.Validate(s)
}
