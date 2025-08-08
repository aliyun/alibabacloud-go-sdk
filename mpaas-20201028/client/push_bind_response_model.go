// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushBindResponse
	GetStatusCode() *int32
	SetBody(v *PushBindResponseBody) *PushBindResponse
	GetBody() *PushBindResponseBody
}

type PushBindResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushBindResponse) String() string {
	return dara.Prettify(s)
}

func (s PushBindResponse) GoString() string {
	return s.String()
}

func (s *PushBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushBindResponse) GetBody() *PushBindResponseBody {
	return s.Body
}

func (s *PushBindResponse) SetHeaders(v map[string]*string) *PushBindResponse {
	s.Headers = v
	return s
}

func (s *PushBindResponse) SetStatusCode(v int32) *PushBindResponse {
	s.StatusCode = &v
	return s
}

func (s *PushBindResponse) SetBody(v *PushBindResponseBody) *PushBindResponse {
	s.Body = v
	return s
}

func (s *PushBindResponse) Validate() error {
	return dara.Validate(s)
}
