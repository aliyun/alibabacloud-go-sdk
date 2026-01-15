// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RenewAppInstanceResponseBody) *RenewAppInstanceResponse
	GetBody() *RenewAppInstanceResponseBody
}

type RenewAppInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewAppInstanceResponse) GetBody() *RenewAppInstanceResponseBody {
	return s.Body
}

func (s *RenewAppInstanceResponse) SetHeaders(v map[string]*string) *RenewAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewAppInstanceResponse) SetStatusCode(v int32) *RenewAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppInstanceResponse) SetBody(v *RenewAppInstanceResponseBody) *RenewAppInstanceResponse {
	s.Body = v
	return s
}

func (s *RenewAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
