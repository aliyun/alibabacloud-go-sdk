// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoScalingConfigResponseBody) *CreateAutoScalingConfigResponse
	GetBody() *CreateAutoScalingConfigResponseBody
}

type CreateAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoScalingConfigResponse) GetBody() *CreateAutoScalingConfigResponseBody {
	return s.Body
}

func (s *CreateAutoScalingConfigResponse) SetHeaders(v map[string]*string) *CreateAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoScalingConfigResponse) SetStatusCode(v int32) *CreateAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoScalingConfigResponse) SetBody(v *CreateAutoScalingConfigResponseBody) *CreateAutoScalingConfigResponse {
	s.Body = v
	return s
}

func (s *CreateAutoScalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
