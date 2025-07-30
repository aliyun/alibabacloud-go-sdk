// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceForDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlushInstanceForDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlushInstanceForDBResponse
	GetStatusCode() *int32
	SetBody(v *FlushInstanceForDBResponseBody) *FlushInstanceForDBResponse
	GetBody() *FlushInstanceForDBResponseBody
}

type FlushInstanceForDBResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlushInstanceForDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlushInstanceForDBResponse) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceForDBResponse) GoString() string {
	return s.String()
}

func (s *FlushInstanceForDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlushInstanceForDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlushInstanceForDBResponse) GetBody() *FlushInstanceForDBResponseBody {
	return s.Body
}

func (s *FlushInstanceForDBResponse) SetHeaders(v map[string]*string) *FlushInstanceForDBResponse {
	s.Headers = v
	return s
}

func (s *FlushInstanceForDBResponse) SetStatusCode(v int32) *FlushInstanceForDBResponse {
	s.StatusCode = &v
	return s
}

func (s *FlushInstanceForDBResponse) SetBody(v *FlushInstanceForDBResponseBody) *FlushInstanceForDBResponse {
	s.Body = v
	return s
}

func (s *FlushInstanceForDBResponse) Validate() error {
	return dara.Validate(s)
}
