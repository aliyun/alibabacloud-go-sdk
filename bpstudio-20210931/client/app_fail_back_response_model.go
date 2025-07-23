// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppFailBackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppFailBackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppFailBackResponse
	GetStatusCode() *int32
	SetBody(v *AppFailBackResponseBody) *AppFailBackResponse
	GetBody() *AppFailBackResponseBody
}

type AppFailBackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppFailBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppFailBackResponse) String() string {
	return dara.Prettify(s)
}

func (s AppFailBackResponse) GoString() string {
	return s.String()
}

func (s *AppFailBackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppFailBackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppFailBackResponse) GetBody() *AppFailBackResponseBody {
	return s.Body
}

func (s *AppFailBackResponse) SetHeaders(v map[string]*string) *AppFailBackResponse {
	s.Headers = v
	return s
}

func (s *AppFailBackResponse) SetStatusCode(v int32) *AppFailBackResponse {
	s.StatusCode = &v
	return s
}

func (s *AppFailBackResponse) SetBody(v *AppFailBackResponseBody) *AppFailBackResponse {
	s.Body = v
	return s
}

func (s *AppFailBackResponse) Validate() error {
	return dara.Validate(s)
}
