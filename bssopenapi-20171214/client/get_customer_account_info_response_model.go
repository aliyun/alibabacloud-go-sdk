// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerAccountInfoResponseBody) *GetCustomerAccountInfoResponse
	GetBody() *GetCustomerAccountInfoResponseBody
}

type GetCustomerAccountInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerAccountInfoResponse) GetBody() *GetCustomerAccountInfoResponseBody {
	return s.Body
}

func (s *GetCustomerAccountInfoResponse) SetHeaders(v map[string]*string) *GetCustomerAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerAccountInfoResponse) SetStatusCode(v int32) *GetCustomerAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerAccountInfoResponse) SetBody(v *GetCustomerAccountInfoResponseBody) *GetCustomerAccountInfoResponse {
	s.Body = v
	return s
}

func (s *GetCustomerAccountInfoResponse) Validate() error {
	return dara.Validate(s)
}
