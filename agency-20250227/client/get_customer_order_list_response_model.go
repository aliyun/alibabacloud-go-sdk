// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrderListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerOrderListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerOrderListResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerOrderListResponseBody) *GetCustomerOrderListResponse
	GetBody() *GetCustomerOrderListResponseBody
}

type GetCustomerOrderListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerOrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerOrderListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrderListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerOrderListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerOrderListResponse) GetBody() *GetCustomerOrderListResponseBody {
	return s.Body
}

func (s *GetCustomerOrderListResponse) SetHeaders(v map[string]*string) *GetCustomerOrderListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerOrderListResponse) SetStatusCode(v int32) *GetCustomerOrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerOrderListResponse) SetBody(v *GetCustomerOrderListResponseBody) *GetCustomerOrderListResponse {
	s.Body = v
	return s
}

func (s *GetCustomerOrderListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
