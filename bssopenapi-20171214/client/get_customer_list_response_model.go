// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerListResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerListResponseBody) *GetCustomerListResponse
	GetBody() *GetCustomerListResponseBody
}

type GetCustomerListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerListResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerListResponse) GetBody() *GetCustomerListResponseBody {
	return s.Body
}

func (s *GetCustomerListResponse) SetHeaders(v map[string]*string) *GetCustomerListResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerListResponse) SetStatusCode(v int32) *GetCustomerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerListResponse) SetBody(v *GetCustomerListResponseBody) *GetCustomerListResponse {
	s.Body = v
	return s
}

func (s *GetCustomerListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
