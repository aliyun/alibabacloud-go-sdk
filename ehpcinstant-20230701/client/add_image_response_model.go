// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImageResponse
	GetStatusCode() *int32
	SetBody(v *AddImageResponseBody) *AddImageResponse
	GetBody() *AddImageResponseBody
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImageResponse) GetBody() *AddImageResponseBody {
	return s.Body
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

func (s *AddImageResponse) Validate() error {
	return dara.Validate(s)
}
