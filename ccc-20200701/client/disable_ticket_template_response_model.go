// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTicketTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableTicketTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableTicketTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DisableTicketTemplateResponseBody) *DisableTicketTemplateResponse
	GetBody() *DisableTicketTemplateResponseBody
}

type DisableTicketTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableTicketTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableTicketTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableTicketTemplateResponse) GoString() string {
	return s.String()
}

func (s *DisableTicketTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableTicketTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableTicketTemplateResponse) GetBody() *DisableTicketTemplateResponseBody {
	return s.Body
}

func (s *DisableTicketTemplateResponse) SetHeaders(v map[string]*string) *DisableTicketTemplateResponse {
	s.Headers = v
	return s
}

func (s *DisableTicketTemplateResponse) SetStatusCode(v int32) *DisableTicketTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableTicketTemplateResponse) SetBody(v *DisableTicketTemplateResponseBody) *DisableTicketTemplateResponse {
	s.Body = v
	return s
}

func (s *DisableTicketTemplateResponse) Validate() error {
	return dara.Validate(s)
}
