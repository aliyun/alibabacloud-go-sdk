// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomerLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomerLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomerLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomerLabelResponseBody) *DeleteCustomerLabelResponse
	GetBody() *DeleteCustomerLabelResponseBody
}

type DeleteCustomerLabelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomerLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomerLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomerLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomerLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomerLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomerLabelResponse) GetBody() *DeleteCustomerLabelResponseBody {
	return s.Body
}

func (s *DeleteCustomerLabelResponse) SetHeaders(v map[string]*string) *DeleteCustomerLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomerLabelResponse) SetStatusCode(v int32) *DeleteCustomerLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomerLabelResponse) SetBody(v *DeleteCustomerLabelResponseBody) *DeleteCustomerLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomerLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
