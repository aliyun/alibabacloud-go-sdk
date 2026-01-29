// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewChangeInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewChangeInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewChangeInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewChangeInstanceResponseBody) *RenewChangeInstanceResponse
	GetBody() *RenewChangeInstanceResponseBody
}

type RenewChangeInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewChangeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewChangeInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewChangeInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewChangeInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewChangeInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewChangeInstanceResponse) GetBody() *RenewChangeInstanceResponseBody {
	return s.Body
}

func (s *RenewChangeInstanceResponse) SetHeaders(v map[string]*string) *RenewChangeInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewChangeInstanceResponse) SetStatusCode(v int32) *RenewChangeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewChangeInstanceResponse) SetBody(v *RenewChangeInstanceResponseBody) *RenewChangeInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewChangeInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
