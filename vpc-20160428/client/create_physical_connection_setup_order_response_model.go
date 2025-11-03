// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionSetupOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePhysicalConnectionSetupOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePhysicalConnectionSetupOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreatePhysicalConnectionSetupOrderResponseBody) *CreatePhysicalConnectionSetupOrderResponse
	GetBody() *CreatePhysicalConnectionSetupOrderResponseBody
}

type CreatePhysicalConnectionSetupOrderResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePhysicalConnectionSetupOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePhysicalConnectionSetupOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionSetupOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionSetupOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePhysicalConnectionSetupOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePhysicalConnectionSetupOrderResponse) GetBody() *CreatePhysicalConnectionSetupOrderResponseBody {
	return s.Body
}

func (s *CreatePhysicalConnectionSetupOrderResponse) SetHeaders(v map[string]*string) *CreatePhysicalConnectionSetupOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponse) SetStatusCode(v int32) *CreatePhysicalConnectionSetupOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponse) SetBody(v *CreatePhysicalConnectionSetupOrderResponseBody) *CreatePhysicalConnectionSetupOrderResponse {
	s.Body = v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
