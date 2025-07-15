// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApiProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachApiProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachApiProductResponse
	GetStatusCode() *int32
	SetBody(v *DetachApiProductResponseBody) *DetachApiProductResponse
	GetBody() *DetachApiProductResponseBody
}

type DetachApiProductResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachApiProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachApiProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachApiProductResponse) GoString() string {
	return s.String()
}

func (s *DetachApiProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachApiProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachApiProductResponse) GetBody() *DetachApiProductResponseBody {
	return s.Body
}

func (s *DetachApiProductResponse) SetHeaders(v map[string]*string) *DetachApiProductResponse {
	s.Headers = v
	return s
}

func (s *DetachApiProductResponse) SetStatusCode(v int32) *DetachApiProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachApiProductResponse) SetBody(v *DetachApiProductResponseBody) *DetachApiProductResponse {
	s.Body = v
	return s
}

func (s *DetachApiProductResponse) Validate() error {
	return dara.Validate(s)
}
