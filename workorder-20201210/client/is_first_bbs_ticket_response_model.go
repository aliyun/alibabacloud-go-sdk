// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsFirstBbsTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsFirstBbsTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsFirstBbsTicketResponse
	GetStatusCode() *int32
	SetBody(v *IsFirstBbsTicketResponseBody) *IsFirstBbsTicketResponse
	GetBody() *IsFirstBbsTicketResponseBody
}

type IsFirstBbsTicketResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsFirstBbsTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsFirstBbsTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s IsFirstBbsTicketResponse) GoString() string {
	return s.String()
}

func (s *IsFirstBbsTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsFirstBbsTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsFirstBbsTicketResponse) GetBody() *IsFirstBbsTicketResponseBody {
	return s.Body
}

func (s *IsFirstBbsTicketResponse) SetHeaders(v map[string]*string) *IsFirstBbsTicketResponse {
	s.Headers = v
	return s
}

func (s *IsFirstBbsTicketResponse) SetStatusCode(v int32) *IsFirstBbsTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *IsFirstBbsTicketResponse) SetBody(v *IsFirstBbsTicketResponseBody) *IsFirstBbsTicketResponse {
	s.Body = v
	return s
}

func (s *IsFirstBbsTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
