// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTicketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTicketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTicketsResponse
	GetStatusCode() *int32
	SetBody(v *QueryTicketsResponseBody) *QueryTicketsResponse
	GetBody() *QueryTicketsResponseBody
}

type QueryTicketsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTicketsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTicketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTicketsResponse) GetBody() *QueryTicketsResponseBody {
	return s.Body
}

func (s *QueryTicketsResponse) SetHeaders(v map[string]*string) *QueryTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketsResponse) SetStatusCode(v int32) *QueryTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketsResponse) SetBody(v *QueryTicketsResponseBody) *QueryTicketsResponse {
	s.Body = v
	return s
}

func (s *QueryTicketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
