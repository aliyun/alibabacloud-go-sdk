// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollAskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PollAskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PollAskResultResponse
	GetStatusCode() *int32
	SetBody(v *PollAskResultResponseBody) *PollAskResultResponse
	GetBody() *PollAskResultResponseBody
}

type PollAskResultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PollAskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PollAskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s PollAskResultResponse) GoString() string {
	return s.String()
}

func (s *PollAskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PollAskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PollAskResultResponse) GetBody() *PollAskResultResponseBody {
	return s.Body
}

func (s *PollAskResultResponse) SetHeaders(v map[string]*string) *PollAskResultResponse {
	s.Headers = v
	return s
}

func (s *PollAskResultResponse) SetStatusCode(v int32) *PollAskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *PollAskResultResponse) SetBody(v *PollAskResultResponseBody) *PollAskResultResponse {
	s.Body = v
	return s
}

func (s *PollAskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
