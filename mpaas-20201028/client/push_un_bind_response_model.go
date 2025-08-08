// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUnBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushUnBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushUnBindResponse
	GetStatusCode() *int32
	SetBody(v *PushUnBindResponseBody) *PushUnBindResponse
	GetBody() *PushUnBindResponseBody
}

type PushUnBindResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushUnBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushUnBindResponse) String() string {
	return dara.Prettify(s)
}

func (s PushUnBindResponse) GoString() string {
	return s.String()
}

func (s *PushUnBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushUnBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushUnBindResponse) GetBody() *PushUnBindResponseBody {
	return s.Body
}

func (s *PushUnBindResponse) SetHeaders(v map[string]*string) *PushUnBindResponse {
	s.Headers = v
	return s
}

func (s *PushUnBindResponse) SetStatusCode(v int32) *PushUnBindResponse {
	s.StatusCode = &v
	return s
}

func (s *PushUnBindResponse) SetBody(v *PushUnBindResponseBody) *PushUnBindResponse {
	s.Body = v
	return s
}

func (s *PushUnBindResponse) Validate() error {
	return dara.Validate(s)
}
