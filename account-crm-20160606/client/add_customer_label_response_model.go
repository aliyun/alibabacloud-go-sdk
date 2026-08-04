// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomerLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomerLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomerLabelResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomerLabelResponseBody) *AddCustomerLabelResponse
	GetBody() *AddCustomerLabelResponseBody
}

type AddCustomerLabelResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomerLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomerLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomerLabelResponse) GoString() string {
	return s.String()
}

func (s *AddCustomerLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomerLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomerLabelResponse) GetBody() *AddCustomerLabelResponseBody {
	return s.Body
}

func (s *AddCustomerLabelResponse) SetHeaders(v map[string]*string) *AddCustomerLabelResponse {
	s.Headers = v
	return s
}

func (s *AddCustomerLabelResponse) SetStatusCode(v int32) *AddCustomerLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomerLabelResponse) SetBody(v *AddCustomerLabelResponseBody) *AddCustomerLabelResponse {
	s.Body = v
	return s
}

func (s *AddCustomerLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
