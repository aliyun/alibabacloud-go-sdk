// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *PutDcdnKvResponseBody) *PutDcdnKvResponse
	GetBody() *PutDcdnKvResponseBody
}

type PutDcdnKvResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *PutDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDcdnKvResponse) GetBody() *PutDcdnKvResponseBody {
	return s.Body
}

func (s *PutDcdnKvResponse) SetHeaders(v map[string]*string) *PutDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *PutDcdnKvResponse) SetStatusCode(v int32) *PutDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDcdnKvResponse) SetBody(v *PutDcdnKvResponseBody) *PutDcdnKvResponse {
	s.Body = v
	return s
}

func (s *PutDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
