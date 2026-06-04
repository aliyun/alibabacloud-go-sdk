// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewServiceInstanceResponseBody) *RenewServiceInstanceResponse
	GetBody() *RenewServiceInstanceResponseBody
}

type RenewServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewServiceInstanceResponse) GetBody() *RenewServiceInstanceResponseBody {
	return s.Body
}

func (s *RenewServiceInstanceResponse) SetHeaders(v map[string]*string) *RenewServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewServiceInstanceResponse) SetStatusCode(v int32) *RenewServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewServiceInstanceResponse) SetBody(v *RenewServiceInstanceResponseBody) *RenewServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewServiceInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
