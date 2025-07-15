// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTicketTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTicketTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTicketTemplateResponseBody) *DeleteTicketTemplateResponse
	GetBody() *DeleteTicketTemplateResponseBody
}

type DeleteTicketTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTicketTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTicketTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTicketTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTicketTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTicketTemplateResponse) GetBody() *DeleteTicketTemplateResponseBody {
	return s.Body
}

func (s *DeleteTicketTemplateResponse) SetHeaders(v map[string]*string) *DeleteTicketTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTicketTemplateResponse) SetStatusCode(v int32) *DeleteTicketTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTicketTemplateResponse) SetBody(v *DeleteTicketTemplateResponseBody) *DeleteTicketTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteTicketTemplateResponse) Validate() error {
	return dara.Validate(s)
}
