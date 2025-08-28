// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseVideoFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseVideoFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseVideoFileResponse
	GetStatusCode() *int32
	SetBody(v *PauseVideoFileResponseBody) *PauseVideoFileResponse
	GetBody() *PauseVideoFileResponseBody
}

type PauseVideoFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseVideoFileResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseVideoFileResponse) GoString() string {
	return s.String()
}

func (s *PauseVideoFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseVideoFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseVideoFileResponse) GetBody() *PauseVideoFileResponseBody {
	return s.Body
}

func (s *PauseVideoFileResponse) SetHeaders(v map[string]*string) *PauseVideoFileResponse {
	s.Headers = v
	return s
}

func (s *PauseVideoFileResponse) SetStatusCode(v int32) *PauseVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseVideoFileResponse) SetBody(v *PauseVideoFileResponseBody) *PauseVideoFileResponse {
	s.Body = v
	return s
}

func (s *PauseVideoFileResponse) Validate() error {
	return dara.Validate(s)
}
