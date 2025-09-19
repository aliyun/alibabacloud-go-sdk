// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionsOutput) *ListSessionsResponse
	GetBody() *ListSessionsOutput
}

type ListSessionsResponse struct {
	Headers    map[string]*string  `json:"headers" xml:"headers"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionsOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionsResponse) GetBody() *ListSessionsOutput {
	return s.Body
}

func (s *ListSessionsResponse) SetHeaders(v map[string]*string) *ListSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListSessionsResponse) SetStatusCode(v int32) *ListSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionsResponse) SetBody(v *ListSessionsOutput) *ListSessionsResponse {
	s.Body = v
	return s
}

func (s *ListSessionsResponse) Validate() error {
	return dara.Validate(s)
}
