// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutKvResponse
	GetStatusCode() *int32
	SetBody(v *PutKvResponseBody) *PutKvResponse
	GetBody() *PutKvResponseBody
}

type PutKvResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutKvResponse) String() string {
	return dara.Prettify(s)
}

func (s PutKvResponse) GoString() string {
	return s.String()
}

func (s *PutKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutKvResponse) GetBody() *PutKvResponseBody {
	return s.Body
}

func (s *PutKvResponse) SetHeaders(v map[string]*string) *PutKvResponse {
	s.Headers = v
	return s
}

func (s *PutKvResponse) SetStatusCode(v int32) *PutKvResponse {
	s.StatusCode = &v
	return s
}

func (s *PutKvResponse) SetBody(v *PutKvResponseBody) *PutKvResponse {
	s.Body = v
	return s
}

func (s *PutKvResponse) Validate() error {
	return dara.Validate(s)
}
