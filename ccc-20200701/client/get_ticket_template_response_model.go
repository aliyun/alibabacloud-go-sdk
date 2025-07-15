// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTicketTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTicketTemplateResponse
	GetStatusCode() *int32
	SetBody(v *GetTicketTemplateResponseBody) *GetTicketTemplateResponse
	GetBody() *GetTicketTemplateResponseBody
}

type GetTicketTemplateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTicketTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTicketTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTicketTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTicketTemplateResponse) GetBody() *GetTicketTemplateResponseBody {
	return s.Body
}

func (s *GetTicketTemplateResponse) SetHeaders(v map[string]*string) *GetTicketTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTicketTemplateResponse) SetStatusCode(v int32) *GetTicketTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketTemplateResponse) SetBody(v *GetTicketTemplateResponseBody) *GetTicketTemplateResponse {
	s.Body = v
	return s
}

func (s *GetTicketTemplateResponse) Validate() error {
	return dara.Validate(s)
}
