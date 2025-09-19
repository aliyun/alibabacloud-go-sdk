// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAegisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindAegisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindAegisResponse
	GetStatusCode() *int32
	SetBody(v *UnbindAegisResponseBody) *UnbindAegisResponse
	GetBody() *UnbindAegisResponseBody
}

type UnbindAegisResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindAegisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindAegisResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindAegisResponse) GoString() string {
	return s.String()
}

func (s *UnbindAegisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindAegisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindAegisResponse) GetBody() *UnbindAegisResponseBody {
	return s.Body
}

func (s *UnbindAegisResponse) SetHeaders(v map[string]*string) *UnbindAegisResponse {
	s.Headers = v
	return s
}

func (s *UnbindAegisResponse) SetStatusCode(v int32) *UnbindAegisResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAegisResponse) SetBody(v *UnbindAegisResponseBody) *UnbindAegisResponse {
	s.Body = v
	return s
}

func (s *UnbindAegisResponse) Validate() error {
	return dara.Validate(s)
}
