// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTicketInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTicketInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryTicketInfoResponseBody) *QueryTicketInfoResponse
	GetBody() *QueryTicketInfoResponseBody
}

type QueryTicketInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTicketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTicketInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTicketInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTicketInfoResponse) GetBody() *QueryTicketInfoResponseBody {
	return s.Body
}

func (s *QueryTicketInfoResponse) SetHeaders(v map[string]*string) *QueryTicketInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketInfoResponse) SetStatusCode(v int32) *QueryTicketInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketInfoResponse) SetBody(v *QueryTicketInfoResponseBody) *QueryTicketInfoResponse {
	s.Body = v
	return s
}

func (s *QueryTicketInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
