// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableLayer7CCResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableLayer7CCResponse
	GetStatusCode() *int32
	SetBody(v *DisableLayer7CCResponseBody) *DisableLayer7CCResponse
	GetBody() *DisableLayer7CCResponseBody
}

type DisableLayer7CCResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLayer7CCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableLayer7CCResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableLayer7CCResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableLayer7CCResponse) GetBody() *DisableLayer7CCResponseBody {
	return s.Body
}

func (s *DisableLayer7CCResponse) SetHeaders(v map[string]*string) *DisableLayer7CCResponse {
	s.Headers = v
	return s
}

func (s *DisableLayer7CCResponse) SetStatusCode(v int32) *DisableLayer7CCResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLayer7CCResponse) SetBody(v *DisableLayer7CCResponseBody) *DisableLayer7CCResponse {
	s.Body = v
	return s
}

func (s *DisableLayer7CCResponse) Validate() error {
	return dara.Validate(s)
}
