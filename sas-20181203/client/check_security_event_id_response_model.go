// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSecurityEventIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSecurityEventIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSecurityEventIdResponse
	GetStatusCode() *int32
	SetBody(v *CheckSecurityEventIdResponseBody) *CheckSecurityEventIdResponse
	GetBody() *CheckSecurityEventIdResponseBody
}

type CheckSecurityEventIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSecurityEventIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSecurityEventIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSecurityEventIdResponse) GoString() string {
	return s.String()
}

func (s *CheckSecurityEventIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSecurityEventIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSecurityEventIdResponse) GetBody() *CheckSecurityEventIdResponseBody {
	return s.Body
}

func (s *CheckSecurityEventIdResponse) SetHeaders(v map[string]*string) *CheckSecurityEventIdResponse {
	s.Headers = v
	return s
}

func (s *CheckSecurityEventIdResponse) SetStatusCode(v int32) *CheckSecurityEventIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSecurityEventIdResponse) SetBody(v *CheckSecurityEventIdResponseBody) *CheckSecurityEventIdResponse {
	s.Body = v
	return s
}

func (s *CheckSecurityEventIdResponse) Validate() error {
	return dara.Validate(s)
}
