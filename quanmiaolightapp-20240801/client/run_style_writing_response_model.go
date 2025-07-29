// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunStyleWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunStyleWritingResponse
	GetStatusCode() *int32
	SetBody(v *RunStyleWritingResponseBody) *RunStyleWritingResponse
	GetBody() *RunStyleWritingResponseBody
}

type RunStyleWritingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunStyleWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunStyleWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponse) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunStyleWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunStyleWritingResponse) GetBody() *RunStyleWritingResponseBody {
	return s.Body
}

func (s *RunStyleWritingResponse) SetHeaders(v map[string]*string) *RunStyleWritingResponse {
	s.Headers = v
	return s
}

func (s *RunStyleWritingResponse) SetStatusCode(v int32) *RunStyleWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunStyleWritingResponse) SetBody(v *RunStyleWritingResponseBody) *RunStyleWritingResponse {
	s.Body = v
	return s
}

func (s *RunStyleWritingResponse) Validate() error {
	return dara.Validate(s)
}
