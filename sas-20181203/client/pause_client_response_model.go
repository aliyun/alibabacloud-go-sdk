// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseClientResponse
	GetStatusCode() *int32
	SetBody(v *PauseClientResponseBody) *PauseClientResponse
	GetBody() *PauseClientResponseBody
}

type PauseClientResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseClientResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseClientResponse) GoString() string {
	return s.String()
}

func (s *PauseClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseClientResponse) GetBody() *PauseClientResponseBody {
	return s.Body
}

func (s *PauseClientResponse) SetHeaders(v map[string]*string) *PauseClientResponse {
	s.Headers = v
	return s
}

func (s *PauseClientResponse) SetStatusCode(v int32) *PauseClientResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseClientResponse) SetBody(v *PauseClientResponseBody) *PauseClientResponse {
	s.Body = v
	return s
}

func (s *PauseClientResponse) Validate() error {
	return dara.Validate(s)
}
