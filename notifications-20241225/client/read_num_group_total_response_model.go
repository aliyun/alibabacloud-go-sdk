// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadNumGroupTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadNumGroupTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadNumGroupTotalResponse
	GetStatusCode() *int32
	SetBody(v *ReadNumGroupTotalResponseBody) *ReadNumGroupTotalResponse
	GetBody() *ReadNumGroupTotalResponseBody
}

type ReadNumGroupTotalResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadNumGroupTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadNumGroupTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadNumGroupTotalResponse) GoString() string {
	return s.String()
}

func (s *ReadNumGroupTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadNumGroupTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadNumGroupTotalResponse) GetBody() *ReadNumGroupTotalResponseBody {
	return s.Body
}

func (s *ReadNumGroupTotalResponse) SetHeaders(v map[string]*string) *ReadNumGroupTotalResponse {
	s.Headers = v
	return s
}

func (s *ReadNumGroupTotalResponse) SetStatusCode(v int32) *ReadNumGroupTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadNumGroupTotalResponse) SetBody(v *ReadNumGroupTotalResponseBody) *ReadNumGroupTotalResponse {
	s.Body = v
	return s
}

func (s *ReadNumGroupTotalResponse) Validate() error {
	return dara.Validate(s)
}
