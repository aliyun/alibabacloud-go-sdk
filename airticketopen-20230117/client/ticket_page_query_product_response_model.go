// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TicketPageQueryProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TicketPageQueryProductResponse
	GetStatusCode() *int32
	SetBody(v *TicketPageQueryProductResponseBody) *TicketPageQueryProductResponse
	GetBody() *TicketPageQueryProductResponseBody
}

type TicketPageQueryProductResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TicketPageQueryProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TicketPageQueryProductResponse) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryProductResponse) GoString() string {
	return s.String()
}

func (s *TicketPageQueryProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TicketPageQueryProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TicketPageQueryProductResponse) GetBody() *TicketPageQueryProductResponseBody {
	return s.Body
}

func (s *TicketPageQueryProductResponse) SetHeaders(v map[string]*string) *TicketPageQueryProductResponse {
	s.Headers = v
	return s
}

func (s *TicketPageQueryProductResponse) SetStatusCode(v int32) *TicketPageQueryProductResponse {
	s.StatusCode = &v
	return s
}

func (s *TicketPageQueryProductResponse) SetBody(v *TicketPageQueryProductResponseBody) *TicketPageQueryProductResponse {
	s.Body = v
	return s
}

func (s *TicketPageQueryProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
