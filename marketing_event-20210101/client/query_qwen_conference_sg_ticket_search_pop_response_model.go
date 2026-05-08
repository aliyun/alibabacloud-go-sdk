// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQwenConferenceSgTicketSearchPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryQwenConferenceSgTicketSearchPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryQwenConferenceSgTicketSearchPopResponse
	GetStatusCode() *int32
	SetBody(v *QueryQwenConferenceSgTicketSearchPopResponseBody) *QueryQwenConferenceSgTicketSearchPopResponse
	GetBody() *QueryQwenConferenceSgTicketSearchPopResponseBody
}

type QueryQwenConferenceSgTicketSearchPopResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryQwenConferenceSgTicketSearchPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryQwenConferenceSgTicketSearchPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryQwenConferenceSgTicketSearchPopResponse) GoString() string {
	return s.String()
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) GetBody() *QueryQwenConferenceSgTicketSearchPopResponseBody {
	return s.Body
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) SetHeaders(v map[string]*string) *QueryQwenConferenceSgTicketSearchPopResponse {
	s.Headers = v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) SetStatusCode(v int32) *QueryQwenConferenceSgTicketSearchPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) SetBody(v *QueryQwenConferenceSgTicketSearchPopResponseBody) *QueryQwenConferenceSgTicketSearchPopResponse {
	s.Body = v
	return s
}

func (s *QueryQwenConferenceSgTicketSearchPopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
