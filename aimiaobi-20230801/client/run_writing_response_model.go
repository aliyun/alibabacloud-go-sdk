// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunWritingResponse
	GetStatusCode() *int32
	SetBody(v *RunWritingResponseBody) *RunWritingResponse
	GetBody() *RunWritingResponseBody
}

type RunWritingResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponse) GoString() string {
	return s.String()
}

func (s *RunWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWritingResponse) GetBody() *RunWritingResponseBody {
	return s.Body
}

func (s *RunWritingResponse) SetHeaders(v map[string]*string) *RunWritingResponse {
	s.Headers = v
	return s
}

func (s *RunWritingResponse) SetStatusCode(v int32) *RunWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunWritingResponse) SetBody(v *RunWritingResponseBody) *RunWritingResponse {
	s.Body = v
	return s
}

func (s *RunWritingResponse) Validate() error {
	return dara.Validate(s)
}
