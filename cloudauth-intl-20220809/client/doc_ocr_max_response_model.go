// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DocOcrMaxResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DocOcrMaxResponse
	GetStatusCode() *int32
	SetBody(v *DocOcrMaxResponseBody) *DocOcrMaxResponse
	GetBody() *DocOcrMaxResponseBody
}

type DocOcrMaxResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DocOcrMaxResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DocOcrMaxResponse) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxResponse) GoString() string {
	return s.String()
}

func (s *DocOcrMaxResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DocOcrMaxResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DocOcrMaxResponse) GetBody() *DocOcrMaxResponseBody {
	return s.Body
}

func (s *DocOcrMaxResponse) SetHeaders(v map[string]*string) *DocOcrMaxResponse {
	s.Headers = v
	return s
}

func (s *DocOcrMaxResponse) SetStatusCode(v int32) *DocOcrMaxResponse {
	s.StatusCode = &v
	return s
}

func (s *DocOcrMaxResponse) SetBody(v *DocOcrMaxResponseBody) *DocOcrMaxResponse {
	s.Body = v
	return s
}

func (s *DocOcrMaxResponse) Validate() error {
	return dara.Validate(s)
}
