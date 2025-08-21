// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSubResponse
	GetStatusCode() *int32
	SetBody(v *AddSubResponseBody) *AddSubResponse
	GetBody() *AddSubResponseBody
}

type AddSubResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSubResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSubResponse) GoString() string {
	return s.String()
}

func (s *AddSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSubResponse) GetBody() *AddSubResponseBody {
	return s.Body
}

func (s *AddSubResponse) SetHeaders(v map[string]*string) *AddSubResponse {
	s.Headers = v
	return s
}

func (s *AddSubResponse) SetStatusCode(v int32) *AddSubResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSubResponse) SetBody(v *AddSubResponseBody) *AddSubResponse {
	s.Body = v
	return s
}

func (s *AddSubResponse) Validate() error {
	return dara.Validate(s)
}
