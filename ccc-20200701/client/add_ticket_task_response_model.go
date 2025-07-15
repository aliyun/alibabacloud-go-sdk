// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTicketTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTicketTaskResponse
	GetStatusCode() *int32
	SetBody(v *AddTicketTaskResponseBody) *AddTicketTaskResponse
	GetBody() *AddTicketTaskResponseBody
}

type AddTicketTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTicketTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTicketTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTicketTaskResponse) GoString() string {
	return s.String()
}

func (s *AddTicketTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTicketTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTicketTaskResponse) GetBody() *AddTicketTaskResponseBody {
	return s.Body
}

func (s *AddTicketTaskResponse) SetHeaders(v map[string]*string) *AddTicketTaskResponse {
	s.Headers = v
	return s
}

func (s *AddTicketTaskResponse) SetStatusCode(v int32) *AddTicketTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTicketTaskResponse) SetBody(v *AddTicketTaskResponseBody) *AddTicketTaskResponse {
	s.Body = v
	return s
}

func (s *AddTicketTaskResponse) Validate() error {
	return dara.Validate(s)
}
