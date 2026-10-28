// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateSessionHeaders
	GetCommonHeaders() map[string]*string
	SetXOpenAgentIdempotencyKey(v string) *UpdateSessionHeaders
	GetXOpenAgentIdempotencyKey() *string
}

type UpdateSessionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The idempotency key for the request. Use the same value when you retry the same update request.
	//
	// example:
	//
	// update-session-example-001
	XOpenAgentIdempotencyKey *string `json:"X-OpenAgent-Idempotency-Key,omitempty" xml:"X-OpenAgent-Idempotency-Key,omitempty"`
}

func (s UpdateSessionHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSessionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateSessionHeaders) GetXOpenAgentIdempotencyKey() *string {
	return s.XOpenAgentIdempotencyKey
}

func (s *UpdateSessionHeaders) SetCommonHeaders(v map[string]*string) *UpdateSessionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSessionHeaders) SetXOpenAgentIdempotencyKey(v string) *UpdateSessionHeaders {
	s.XOpenAgentIdempotencyKey = &v
	return s
}

func (s *UpdateSessionHeaders) Validate() error {
	return dara.Validate(s)
}
