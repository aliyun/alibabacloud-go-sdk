// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerConfigResponseBody) *GetCustomerConfigResponse
	GetBody() *GetCustomerConfigResponseBody
}

type GetCustomerConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerConfigResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerConfigResponse) GetBody() *GetCustomerConfigResponseBody {
	return s.Body
}

func (s *GetCustomerConfigResponse) SetHeaders(v map[string]*string) *GetCustomerConfigResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerConfigResponse) SetStatusCode(v int32) *GetCustomerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerConfigResponse) SetBody(v *GetCustomerConfigResponseBody) *GetCustomerConfigResponse {
	s.Body = v
	return s
}

func (s *GetCustomerConfigResponse) Validate() error {
	return dara.Validate(s)
}
