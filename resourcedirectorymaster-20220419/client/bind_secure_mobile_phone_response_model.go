// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSecureMobilePhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindSecureMobilePhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindSecureMobilePhoneResponse
	GetStatusCode() *int32
	SetBody(v *BindSecureMobilePhoneResponseBody) *BindSecureMobilePhoneResponse
	GetBody() *BindSecureMobilePhoneResponseBody
}

type BindSecureMobilePhoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindSecureMobilePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindSecureMobilePhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s BindSecureMobilePhoneResponse) GoString() string {
	return s.String()
}

func (s *BindSecureMobilePhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindSecureMobilePhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindSecureMobilePhoneResponse) GetBody() *BindSecureMobilePhoneResponseBody {
	return s.Body
}

func (s *BindSecureMobilePhoneResponse) SetHeaders(v map[string]*string) *BindSecureMobilePhoneResponse {
	s.Headers = v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetStatusCode(v int32) *BindSecureMobilePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindSecureMobilePhoneResponse) SetBody(v *BindSecureMobilePhoneResponseBody) *BindSecureMobilePhoneResponse {
	s.Body = v
	return s
}

func (s *BindSecureMobilePhoneResponse) Validate() error {
	return dara.Validate(s)
}
