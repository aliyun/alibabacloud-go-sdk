// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPollTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PollTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PollTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *PollTaskResultResponseBody) *PollTaskResultResponse
	GetBody() *PollTaskResultResponseBody
}

type PollTaskResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PollTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PollTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s PollTaskResultResponse) GoString() string {
	return s.String()
}

func (s *PollTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PollTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PollTaskResultResponse) GetBody() *PollTaskResultResponseBody {
	return s.Body
}

func (s *PollTaskResultResponse) SetHeaders(v map[string]*string) *PollTaskResultResponse {
	s.Headers = v
	return s
}

func (s *PollTaskResultResponse) SetStatusCode(v int32) *PollTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *PollTaskResultResponse) SetBody(v *PollTaskResultResponseBody) *PollTaskResultResponse {
	s.Body = v
	return s
}

func (s *PollTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
