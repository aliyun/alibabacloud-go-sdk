// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSearchShadeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSearchShadeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSearchShadeResponse
	GetStatusCode() *int32
	SetBody(v *SendSearchShadeResponseBody) *SendSearchShadeResponse
	GetBody() *SendSearchShadeResponseBody
}

type SendSearchShadeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSearchShadeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSearchShadeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSearchShadeResponse) GoString() string {
	return s.String()
}

func (s *SendSearchShadeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSearchShadeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSearchShadeResponse) GetBody() *SendSearchShadeResponseBody {
	return s.Body
}

func (s *SendSearchShadeResponse) SetHeaders(v map[string]*string) *SendSearchShadeResponse {
	s.Headers = v
	return s
}

func (s *SendSearchShadeResponse) SetStatusCode(v int32) *SendSearchShadeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSearchShadeResponse) SetBody(v *SendSearchShadeResponseBody) *SendSearchShadeResponse {
	s.Body = v
	return s
}

func (s *SendSearchShadeResponse) Validate() error {
	return dara.Validate(s)
}
