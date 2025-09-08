// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDeepWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDeepWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDeepWritingResponse
	GetStatusCode() *int32
	SetBody(v *RunDeepWritingResponseBody) *RunDeepWritingResponse
	GetBody() *RunDeepWritingResponseBody
}

type RunDeepWritingResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDeepWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDeepWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDeepWritingResponse) GoString() string {
	return s.String()
}

func (s *RunDeepWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDeepWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDeepWritingResponse) GetBody() *RunDeepWritingResponseBody {
	return s.Body
}

func (s *RunDeepWritingResponse) SetHeaders(v map[string]*string) *RunDeepWritingResponse {
	s.Headers = v
	return s
}

func (s *RunDeepWritingResponse) SetStatusCode(v int32) *RunDeepWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDeepWritingResponse) SetBody(v *RunDeepWritingResponseBody) *RunDeepWritingResponse {
	s.Body = v
	return s
}

func (s *RunDeepWritingResponse) Validate() error {
	return dara.Validate(s)
}
