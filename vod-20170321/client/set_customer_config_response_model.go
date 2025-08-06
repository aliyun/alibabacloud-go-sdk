// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCustomerConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCustomerConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCustomerConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetCustomerConfigResponseBody) *SetCustomerConfigResponse
	GetBody() *SetCustomerConfigResponseBody
}

type SetCustomerConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCustomerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCustomerConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCustomerConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCustomerConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCustomerConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCustomerConfigResponse) GetBody() *SetCustomerConfigResponseBody {
	return s.Body
}

func (s *SetCustomerConfigResponse) SetHeaders(v map[string]*string) *SetCustomerConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCustomerConfigResponse) SetStatusCode(v int32) *SetCustomerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCustomerConfigResponse) SetBody(v *SetCustomerConfigResponseBody) *SetCustomerConfigResponse {
	s.Body = v
	return s
}

func (s *SetCustomerConfigResponse) Validate() error {
	return dara.Validate(s)
}
