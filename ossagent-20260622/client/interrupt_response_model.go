// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterruptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterruptResponse
	GetStatusCode() *int32
	SetBody(v string) *InterruptResponse
	GetBody() *string
}

type InterruptResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterruptResponse) String() string {
	return dara.Prettify(s)
}

func (s InterruptResponse) GoString() string {
	return s.String()
}

func (s *InterruptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterruptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterruptResponse) GetBody() *string {
	return s.Body
}

func (s *InterruptResponse) SetHeaders(v map[string]*string) *InterruptResponse {
	s.Headers = v
	return s
}

func (s *InterruptResponse) SetStatusCode(v int32) *InterruptResponse {
	s.StatusCode = &v
	return s
}

func (s *InterruptResponse) SetBody(v string) *InterruptResponse {
	s.Body = &v
	return s
}

func (s *InterruptResponse) Validate() error {
	return dara.Validate(s)
}
