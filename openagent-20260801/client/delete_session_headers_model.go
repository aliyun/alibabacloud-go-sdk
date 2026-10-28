// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSessionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSessionHeaders
	GetCommonHeaders() map[string]*string
	SetXOpenAgentIdempotencyKey(v string) *DeleteSessionHeaders
	GetXOpenAgentIdempotencyKey() *string
}

type DeleteSessionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The idempotency key of the request. Subsequent calls after deletion, including those using the same key, return 404 Session.NotFound.
	//
	// example:
	//
	// delete-session-example-001
	XOpenAgentIdempotencyKey *string `json:"X-OpenAgent-Idempotency-Key,omitempty" xml:"X-OpenAgent-Idempotency-Key,omitempty"`
}

func (s DeleteSessionHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSessionHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSessionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSessionHeaders) GetXOpenAgentIdempotencyKey() *string {
	return s.XOpenAgentIdempotencyKey
}

func (s *DeleteSessionHeaders) SetCommonHeaders(v map[string]*string) *DeleteSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSessionHeaders) SetXOpenAgentIdempotencyKey(v string) *DeleteSessionHeaders {
	s.XOpenAgentIdempotencyKey = &v
	return s
}

func (s *DeleteSessionHeaders) Validate() error {
	return dara.Validate(s)
}
