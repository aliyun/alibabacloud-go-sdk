// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupByClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadNumGroupByClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadNumGroupByClassResponse
	GetStatusCode() *int32
	SetBody(v *ReadNumGroupByClassResponseBody) *ReadNumGroupByClassResponse
	GetBody() *ReadNumGroupByClassResponseBody
}

type ReadNumGroupByClassResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadNumGroupByClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadNumGroupByClassResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupByClassResponse) GoString() string {
	return s.String()
}

func (s *ReadNumGroupByClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadNumGroupByClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadNumGroupByClassResponse) GetBody() *ReadNumGroupByClassResponseBody {
	return s.Body
}

func (s *ReadNumGroupByClassResponse) SetHeaders(v map[string]*string) *ReadNumGroupByClassResponse {
	s.Headers = v
	return s
}

func (s *ReadNumGroupByClassResponse) SetStatusCode(v int32) *ReadNumGroupByClassResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadNumGroupByClassResponse) SetBody(v *ReadNumGroupByClassResponseBody) *ReadNumGroupByClassResponse {
	s.Body = v
	return s
}

func (s *ReadNumGroupByClassResponse) Validate() error {
	return dara.Validate(s)
}
