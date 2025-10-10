// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageBasicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadPageBasicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadPageBasicResponse
	GetStatusCode() *int32
	SetBody(v *ReadPageBasicResponseBody) *ReadPageBasicResponse
	GetBody() *ReadPageBasicResponseBody
}

type ReadPageBasicResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadPageBasicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadPageBasicResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadPageBasicResponse) GoString() string {
	return s.String()
}

func (s *ReadPageBasicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadPageBasicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadPageBasicResponse) GetBody() *ReadPageBasicResponseBody {
	return s.Body
}

func (s *ReadPageBasicResponse) SetHeaders(v map[string]*string) *ReadPageBasicResponse {
	s.Headers = v
	return s
}

func (s *ReadPageBasicResponse) SetStatusCode(v int32) *ReadPageBasicResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadPageBasicResponse) SetBody(v *ReadPageBasicResponseBody) *ReadPageBasicResponse {
	s.Body = v
	return s
}

func (s *ReadPageBasicResponse) Validate() error {
	return dara.Validate(s)
}
