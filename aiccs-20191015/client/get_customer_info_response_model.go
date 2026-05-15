// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerInfoResponseBody) *GetCustomerInfoResponse
	GetBody() *GetCustomerInfoResponseBody
}

type GetCustomerInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerInfoResponse) GetBody() *GetCustomerInfoResponseBody {
	return s.Body
}

func (s *GetCustomerInfoResponse) SetHeaders(v map[string]*string) *GetCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerInfoResponse) SetStatusCode(v int32) *GetCustomerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerInfoResponse) SetBody(v *GetCustomerInfoResponseBody) *GetCustomerInfoResponse {
	s.Body = v
	return s
}

func (s *GetCustomerInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
