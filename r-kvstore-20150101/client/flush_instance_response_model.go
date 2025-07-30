// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlushInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlushInstanceResponse
	GetStatusCode() *int32
	SetBody(v *FlushInstanceResponseBody) *FlushInstanceResponse
	GetBody() *FlushInstanceResponseBody
}

type FlushInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlushInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlushInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s FlushInstanceResponse) GoString() string {
	return s.String()
}

func (s *FlushInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlushInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlushInstanceResponse) GetBody() *FlushInstanceResponseBody {
	return s.Body
}

func (s *FlushInstanceResponse) SetHeaders(v map[string]*string) *FlushInstanceResponse {
	s.Headers = v
	return s
}

func (s *FlushInstanceResponse) SetStatusCode(v int32) *FlushInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *FlushInstanceResponse) SetBody(v *FlushInstanceResponseBody) *FlushInstanceResponse {
	s.Body = v
	return s
}

func (s *FlushInstanceResponse) Validate() error {
	return dara.Validate(s)
}
