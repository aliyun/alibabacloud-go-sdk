// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTicketNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTicketNumResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTicketNumResponseBody) *UpdateTicketNumResponse
	GetBody() *UpdateTicketNumResponseBody
}

type UpdateTicketNumResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTicketNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTicketNumResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketNumResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTicketNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTicketNumResponse) GetBody() *UpdateTicketNumResponseBody {
	return s.Body
}

func (s *UpdateTicketNumResponse) SetHeaders(v map[string]*string) *UpdateTicketNumResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketNumResponse) SetStatusCode(v int32) *UpdateTicketNumResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketNumResponse) SetBody(v *UpdateTicketNumResponseBody) *UpdateTicketNumResponse {
	s.Body = v
	return s
}

func (s *UpdateTicketNumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
