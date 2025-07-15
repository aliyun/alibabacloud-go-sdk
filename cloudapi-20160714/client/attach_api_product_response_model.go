// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApiProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachApiProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachApiProductResponse
	GetStatusCode() *int32
	SetBody(v *AttachApiProductResponseBody) *AttachApiProductResponse
	GetBody() *AttachApiProductResponseBody
}

type AttachApiProductResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachApiProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachApiProductResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachApiProductResponse) GoString() string {
	return s.String()
}

func (s *AttachApiProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachApiProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachApiProductResponse) GetBody() *AttachApiProductResponseBody {
	return s.Body
}

func (s *AttachApiProductResponse) SetHeaders(v map[string]*string) *AttachApiProductResponse {
	s.Headers = v
	return s
}

func (s *AttachApiProductResponse) SetStatusCode(v int32) *AttachApiProductResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachApiProductResponse) SetBody(v *AttachApiProductResponseBody) *AttachApiProductResponse {
	s.Body = v
	return s
}

func (s *AttachApiProductResponse) Validate() error {
	return dara.Validate(s)
}
