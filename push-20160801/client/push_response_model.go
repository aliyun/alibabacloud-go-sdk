// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushResponse
	GetStatusCode() *int32
	SetBody(v *PushResponseBody) *PushResponse
	GetBody() *PushResponseBody
}

type PushResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResponseBody  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResponse) String() string {
	return dara.Prettify(s)
}

func (s PushResponse) GoString() string {
	return s.String()
}

func (s *PushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushResponse) GetBody() *PushResponseBody {
	return s.Body
}

func (s *PushResponse) SetHeaders(v map[string]*string) *PushResponse {
	s.Headers = v
	return s
}

func (s *PushResponse) SetStatusCode(v int32) *PushResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResponse) SetBody(v *PushResponseBody) *PushResponse {
	s.Body = v
	return s
}

func (s *PushResponse) Validate() error {
	return dara.Validate(s)
}
