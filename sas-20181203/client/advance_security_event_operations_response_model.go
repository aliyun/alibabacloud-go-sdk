// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdvanceSecurityEventOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AdvanceSecurityEventOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AdvanceSecurityEventOperationsResponse
	GetStatusCode() *int32
	SetBody(v *AdvanceSecurityEventOperationsResponseBody) *AdvanceSecurityEventOperationsResponse
	GetBody() *AdvanceSecurityEventOperationsResponseBody
}

type AdvanceSecurityEventOperationsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdvanceSecurityEventOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdvanceSecurityEventOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s AdvanceSecurityEventOperationsResponse) GoString() string {
	return s.String()
}

func (s *AdvanceSecurityEventOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AdvanceSecurityEventOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AdvanceSecurityEventOperationsResponse) GetBody() *AdvanceSecurityEventOperationsResponseBody {
	return s.Body
}

func (s *AdvanceSecurityEventOperationsResponse) SetHeaders(v map[string]*string) *AdvanceSecurityEventOperationsResponse {
	s.Headers = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponse) SetStatusCode(v int32) *AdvanceSecurityEventOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *AdvanceSecurityEventOperationsResponse) SetBody(v *AdvanceSecurityEventOperationsResponseBody) *AdvanceSecurityEventOperationsResponse {
	s.Body = v
	return s
}

func (s *AdvanceSecurityEventOperationsResponse) Validate() error {
	return dara.Validate(s)
}
