// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTBAccountUnbindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TBAccountUnbindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TBAccountUnbindResponse
	GetStatusCode() *int32
	SetBody(v *TBAccountUnbindResponseBody) *TBAccountUnbindResponse
	GetBody() *TBAccountUnbindResponseBody
}

type TBAccountUnbindResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TBAccountUnbindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TBAccountUnbindResponse) String() string {
	return dara.Prettify(s)
}

func (s TBAccountUnbindResponse) GoString() string {
	return s.String()
}

func (s *TBAccountUnbindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TBAccountUnbindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TBAccountUnbindResponse) GetBody() *TBAccountUnbindResponseBody {
	return s.Body
}

func (s *TBAccountUnbindResponse) SetHeaders(v map[string]*string) *TBAccountUnbindResponse {
	s.Headers = v
	return s
}

func (s *TBAccountUnbindResponse) SetStatusCode(v int32) *TBAccountUnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *TBAccountUnbindResponse) SetBody(v *TBAccountUnbindResponseBody) *TBAccountUnbindResponse {
	s.Body = v
	return s
}

func (s *TBAccountUnbindResponse) Validate() error {
	return dara.Validate(s)
}
