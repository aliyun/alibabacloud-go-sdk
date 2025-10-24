// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAutoScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceAutoScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceAutoScalerResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceAutoScalerResponseBody) *CreateServiceAutoScalerResponse
	GetBody() *CreateServiceAutoScalerResponseBody
}

type CreateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceAutoScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceAutoScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceAutoScalerResponse) GetBody() *CreateServiceAutoScalerResponseBody {
	return s.Body
}

func (s *CreateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *CreateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetStatusCode(v int32) *CreateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetBody(v *CreateServiceAutoScalerResponseBody) *CreateServiceAutoScalerResponse {
	s.Body = v
	return s
}

func (s *CreateServiceAutoScalerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
