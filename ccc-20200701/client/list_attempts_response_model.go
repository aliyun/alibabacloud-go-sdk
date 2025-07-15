// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAttemptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAttemptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAttemptsResponse
	GetStatusCode() *int32
	SetBody(v *ListAttemptsResponseBody) *ListAttemptsResponse
	GetBody() *ListAttemptsResponseBody
}

type ListAttemptsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAttemptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAttemptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAttemptsResponse) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAttemptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAttemptsResponse) GetBody() *ListAttemptsResponseBody {
	return s.Body
}

func (s *ListAttemptsResponse) SetHeaders(v map[string]*string) *ListAttemptsResponse {
	s.Headers = v
	return s
}

func (s *ListAttemptsResponse) SetStatusCode(v int32) *ListAttemptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttemptsResponse) SetBody(v *ListAttemptsResponseBody) *ListAttemptsResponse {
	s.Body = v
	return s
}

func (s *ListAttemptsResponse) Validate() error {
	return dara.Validate(s)
}
