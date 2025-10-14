// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogTicketResponse
	GetStatusCode() *int32
	SetBody(v *GetLogTicketResponseBody) *GetLogTicketResponse
	GetBody() *GetLogTicketResponseBody
}

type GetLogTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogTicketResponse) GoString() string {
	return s.String()
}

func (s *GetLogTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogTicketResponse) GetBody() *GetLogTicketResponseBody {
	return s.Body
}

func (s *GetLogTicketResponse) SetHeaders(v map[string]*string) *GetLogTicketResponse {
	s.Headers = v
	return s
}

func (s *GetLogTicketResponse) SetStatusCode(v int32) *GetLogTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogTicketResponse) SetBody(v *GetLogTicketResponseBody) *GetLogTicketResponse {
	s.Body = v
	return s
}

func (s *GetLogTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
