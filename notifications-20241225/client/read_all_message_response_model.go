// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadAllMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadAllMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadAllMessageResponse
	GetStatusCode() *int32
	SetBody(v *ReadAllMessageResponseBody) *ReadAllMessageResponse
	GetBody() *ReadAllMessageResponseBody
}

type ReadAllMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadAllMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadAllMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadAllMessageResponse) GoString() string {
	return s.String()
}

func (s *ReadAllMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadAllMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadAllMessageResponse) GetBody() *ReadAllMessageResponseBody {
	return s.Body
}

func (s *ReadAllMessageResponse) SetHeaders(v map[string]*string) *ReadAllMessageResponse {
	s.Headers = v
	return s
}

func (s *ReadAllMessageResponse) SetStatusCode(v int32) *ReadAllMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadAllMessageResponse) SetBody(v *ReadAllMessageResponseBody) *ReadAllMessageResponse {
	s.Body = v
	return s
}

func (s *ReadAllMessageResponse) Validate() error {
	return dara.Validate(s)
}
