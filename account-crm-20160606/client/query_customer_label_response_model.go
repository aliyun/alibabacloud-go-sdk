// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomerLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomerLabelResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomerLabelResponseBody) *QueryCustomerLabelResponse
	GetBody() *QueryCustomerLabelResponseBody
}

type QueryCustomerLabelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerLabelResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomerLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomerLabelResponse) GetBody() *QueryCustomerLabelResponseBody {
	return s.Body
}

func (s *QueryCustomerLabelResponse) SetHeaders(v map[string]*string) *QueryCustomerLabelResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerLabelResponse) SetStatusCode(v int32) *QueryCustomerLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerLabelResponse) SetBody(v *QueryCustomerLabelResponseBody) *QueryCustomerLabelResponse {
	s.Body = v
	return s
}

func (s *QueryCustomerLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
