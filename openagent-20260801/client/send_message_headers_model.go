// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendMessageHeaders
	GetCommonHeaders() map[string]*string
	SetXOpenAgentIdempotencyKey(v string) *SendMessageHeaders
	GetXOpenAgentIdempotencyKey() *string
}

type SendMessageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The idempotency key for the request. Use the same key and the same business parameters when retrying the same request.
	//
	// example:
	//
	// send-message-example-001
	XOpenAgentIdempotencyKey *string `json:"X-OpenAgent-Idempotency-Key,omitempty" xml:"X-OpenAgent-Idempotency-Key,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendMessageHeaders) GetXOpenAgentIdempotencyKey() *string {
	return s.XOpenAgentIdempotencyKey
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXOpenAgentIdempotencyKey(v string) *SendMessageHeaders {
	s.XOpenAgentIdempotencyKey = &v
	return s
}

func (s *SendMessageHeaders) Validate() error {
	return dara.Validate(s)
}
