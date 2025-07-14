// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelayTicketExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelayTicketExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelayTicketExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *DelayTicketExpireTimeResponseBody) *DelayTicketExpireTimeResponse
	GetBody() *DelayTicketExpireTimeResponseBody
}

type DelayTicketExpireTimeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelayTicketExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelayTicketExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DelayTicketExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *DelayTicketExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelayTicketExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelayTicketExpireTimeResponse) GetBody() *DelayTicketExpireTimeResponseBody {
	return s.Body
}

func (s *DelayTicketExpireTimeResponse) SetHeaders(v map[string]*string) *DelayTicketExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *DelayTicketExpireTimeResponse) SetStatusCode(v int32) *DelayTicketExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DelayTicketExpireTimeResponse) SetBody(v *DelayTicketExpireTimeResponseBody) *DelayTicketExpireTimeResponse {
	s.Body = v
	return s
}

func (s *DelayTicketExpireTimeResponse) Validate() error {
	return dara.Validate(s)
}
