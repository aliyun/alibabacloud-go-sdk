// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImages2LibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddImages2LibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddImages2LibResponse
	GetStatusCode() *int32
	SetBody(v *AddImages2LibResponseBody) *AddImages2LibResponse
	GetBody() *AddImages2LibResponseBody
}

type AddImages2LibResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImages2LibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImages2LibResponse) String() string {
	return dara.Prettify(s)
}

func (s AddImages2LibResponse) GoString() string {
	return s.String()
}

func (s *AddImages2LibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddImages2LibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddImages2LibResponse) GetBody() *AddImages2LibResponseBody {
	return s.Body
}

func (s *AddImages2LibResponse) SetHeaders(v map[string]*string) *AddImages2LibResponse {
	s.Headers = v
	return s
}

func (s *AddImages2LibResponse) SetStatusCode(v int32) *AddImages2LibResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImages2LibResponse) SetBody(v *AddImages2LibResponseBody) *AddImages2LibResponse {
	s.Body = v
	return s
}

func (s *AddImages2LibResponse) Validate() error {
	return dara.Validate(s)
}
