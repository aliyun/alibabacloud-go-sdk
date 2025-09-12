// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseLindormInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseLindormInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseLindormInstanceResponseBody) *ReleaseLindormInstanceResponse
	GetBody() *ReleaseLindormInstanceResponseBody
}

type ReleaseLindormInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseLindormInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseLindormInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseLindormInstanceResponse) GetBody() *ReleaseLindormInstanceResponseBody {
	return s.Body
}

func (s *ReleaseLindormInstanceResponse) SetHeaders(v map[string]*string) *ReleaseLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseLindormInstanceResponse) SetStatusCode(v int32) *ReleaseLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseLindormInstanceResponse) SetBody(v *ReleaseLindormInstanceResponseBody) *ReleaseLindormInstanceResponse {
	s.Body = v
	return s
}

func (s *ReleaseLindormInstanceResponse) Validate() error {
	return dara.Validate(s)
}
