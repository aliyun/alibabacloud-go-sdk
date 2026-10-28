// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSessionHeaders
	GetCommonHeaders() map[string]*string
	SetXOpenAgentIdempotencyKey(v string) *CreateSessionHeaders
	GetXOpenAgentIdempotencyKey() *string
}

type CreateSessionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The request idempotency key. Use the same value and the same business parameters when retrying the same request.
	//
	// example:
	//
	// create-session-example-001
	XOpenAgentIdempotencyKey *string `json:"X-OpenAgent-Idempotency-Key,omitempty" xml:"X-OpenAgent-Idempotency-Key,omitempty"`
}

func (s CreateSessionHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionHeaders) GoString() string {
	return s.String()
}

func (s *CreateSessionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSessionHeaders) GetXOpenAgentIdempotencyKey() *string {
	return s.XOpenAgentIdempotencyKey
}

func (s *CreateSessionHeaders) SetCommonHeaders(v map[string]*string) *CreateSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSessionHeaders) SetXOpenAgentIdempotencyKey(v string) *CreateSessionHeaders {
	s.XOpenAgentIdempotencyKey = &v
	return s
}

func (s *CreateSessionHeaders) Validate() error {
	return dara.Validate(s)
}
