// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceAutoScalerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceAutoScalerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceAutoScalerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceAutoScalerResponseBody) *UpdateServiceAutoScalerResponse
	GetBody() *UpdateServiceAutoScalerResponseBody
}

type UpdateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceAutoScalerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceAutoScalerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceAutoScalerResponse) GetBody() *UpdateServiceAutoScalerResponseBody {
	return s.Body
}

func (s *UpdateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetStatusCode(v int32) *UpdateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetBody(v *UpdateServiceAutoScalerResponseBody) *UpdateServiceAutoScalerResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceAutoScalerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
