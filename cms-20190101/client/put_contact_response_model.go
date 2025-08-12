// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutContactResponse
	GetStatusCode() *int32
	SetBody(v *PutContactResponseBody) *PutContactResponse
	GetBody() *PutContactResponseBody
}

type PutContactResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutContactResponse) String() string {
	return dara.Prettify(s)
}

func (s PutContactResponse) GoString() string {
	return s.String()
}

func (s *PutContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutContactResponse) GetBody() *PutContactResponseBody {
	return s.Body
}

func (s *PutContactResponse) SetHeaders(v map[string]*string) *PutContactResponse {
	s.Headers = v
	return s
}

func (s *PutContactResponse) SetStatusCode(v int32) *PutContactResponse {
	s.StatusCode = &v
	return s
}

func (s *PutContactResponse) SetBody(v *PutContactResponseBody) *PutContactResponse {
	s.Body = v
	return s
}

func (s *PutContactResponse) Validate() error {
	return dara.Validate(s)
}
