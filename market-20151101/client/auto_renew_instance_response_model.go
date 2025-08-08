// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoRenewInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AutoRenewInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AutoRenewInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AutoRenewInstanceResponseBody) *AutoRenewInstanceResponse
	GetBody() *AutoRenewInstanceResponseBody
}

type AutoRenewInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AutoRenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AutoRenewInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AutoRenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *AutoRenewInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AutoRenewInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AutoRenewInstanceResponse) GetBody() *AutoRenewInstanceResponseBody {
	return s.Body
}

func (s *AutoRenewInstanceResponse) SetHeaders(v map[string]*string) *AutoRenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *AutoRenewInstanceResponse) SetStatusCode(v int32) *AutoRenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AutoRenewInstanceResponse) SetBody(v *AutoRenewInstanceResponseBody) *AutoRenewInstanceResponse {
	s.Body = v
	return s
}

func (s *AutoRenewInstanceResponse) Validate() error {
	return dara.Validate(s)
}
