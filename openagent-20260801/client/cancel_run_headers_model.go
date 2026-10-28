// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRunHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CancelRunHeaders
	GetCommonHeaders() map[string]*string
	SetXOpenAgentIdempotencyKey(v string) *CancelRunHeaders
	GetXOpenAgentIdempotencyKey() *string
}

type CancelRunHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The idempotency key of the request. Use the same value and the same business parameters when retrying the same cancellation request.
	//
	// example:
	//
	// cancel-example-001
	XOpenAgentIdempotencyKey *string `json:"X-OpenAgent-Idempotency-Key,omitempty" xml:"X-OpenAgent-Idempotency-Key,omitempty"`
}

func (s CancelRunHeaders) String() string {
	return dara.Prettify(s)
}

func (s CancelRunHeaders) GoString() string {
	return s.String()
}

func (s *CancelRunHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CancelRunHeaders) GetXOpenAgentIdempotencyKey() *string {
	return s.XOpenAgentIdempotencyKey
}

func (s *CancelRunHeaders) SetCommonHeaders(v map[string]*string) *CancelRunHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelRunHeaders) SetXOpenAgentIdempotencyKey(v string) *CancelRunHeaders {
	s.XOpenAgentIdempotencyKey = &v
	return s
}

func (s *CancelRunHeaders) Validate() error {
	return dara.Validate(s)
}
