// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceWorkOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceWorkOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceWorkOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceWorkOrderResponseBody) *CreateServiceWorkOrderResponse
	GetBody() *CreateServiceWorkOrderResponseBody
}

type CreateServiceWorkOrderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceWorkOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceWorkOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceWorkOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceWorkOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceWorkOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceWorkOrderResponse) GetBody() *CreateServiceWorkOrderResponseBody {
	return s.Body
}

func (s *CreateServiceWorkOrderResponse) SetHeaders(v map[string]*string) *CreateServiceWorkOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetStatusCode(v int32) *CreateServiceWorkOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceWorkOrderResponse) SetBody(v *CreateServiceWorkOrderResponseBody) *CreateServiceWorkOrderResponse {
	s.Body = v
	return s
}

func (s *CreateServiceWorkOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
