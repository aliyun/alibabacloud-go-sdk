// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PullTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PullTaskResponse
	GetStatusCode() *int32
	SetBody(v *PullTaskResponseBody) *PullTaskResponse
	GetBody() *PullTaskResponseBody
}

type PullTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s PullTaskResponse) GoString() string {
	return s.String()
}

func (s *PullTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PullTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PullTaskResponse) GetBody() *PullTaskResponseBody {
	return s.Body
}

func (s *PullTaskResponse) SetHeaders(v map[string]*string) *PullTaskResponse {
	s.Headers = v
	return s
}

func (s *PullTaskResponse) SetStatusCode(v int32) *PullTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PullTaskResponse) SetBody(v *PullTaskResponseBody) *PullTaskResponse {
	s.Body = v
	return s
}

func (s *PullTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
