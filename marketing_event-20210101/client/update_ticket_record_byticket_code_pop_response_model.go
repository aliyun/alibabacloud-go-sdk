// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketRecordByticketCodePopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTicketRecordByticketCodePopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTicketRecordByticketCodePopResponseBody) *UpdateTicketRecordByticketCodePopResponse
	GetBody() *UpdateTicketRecordByticketCodePopResponseBody
}

type UpdateTicketRecordByticketCodePopResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketRecordByticketCodePopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTicketRecordByticketCodePopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTicketRecordByticketCodePopResponse) GetBody() *UpdateTicketRecordByticketCodePopResponseBody {
	return s.Body
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetHeaders(v map[string]*string) *UpdateTicketRecordByticketCodePopResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetStatusCode(v int32) *UpdateTicketRecordByticketCodePopResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponse) SetBody(v *UpdateTicketRecordByticketCodePopResponseBody) *UpdateTicketRecordByticketCodePopResponse {
	s.Body = v
	return s
}

func (s *UpdateTicketRecordByticketCodePopResponse) Validate() error {
	return dara.Validate(s)
}
