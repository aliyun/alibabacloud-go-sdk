// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryQwenConferenceSgTicketPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryQwenConferenceSgTicketPopResponse
	GetStatusCode() *int32
	SetBody(v *QueryQwenConferenceSgTicketPopResponseBody) *QueryQwenConferenceSgTicketPopResponse
	GetBody() *QueryQwenConferenceSgTicketPopResponseBody
}

type QueryQwenConferenceSgTicketPopResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQwenConferenceSgTicketPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQwenConferenceSgTicketPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketPopResponse) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryQwenConferenceSgTicketPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryQwenConferenceSgTicketPopResponse) GetBody() *QueryQwenConferenceSgTicketPopResponseBody {
	return s.Body
}

func (s *QueryQwenConferenceSgTicketPopResponse) SetHeaders(v map[string]*string) *QueryQwenConferenceSgTicketPopResponse {
	s.Headers = v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponse) SetStatusCode(v int32) *QueryQwenConferenceSgTicketPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponse) SetBody(v *QueryQwenConferenceSgTicketPopResponseBody) *QueryQwenConferenceSgTicketPopResponse {
	s.Body = v
	return s
}

func (s *QueryQwenConferenceSgTicketPopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
