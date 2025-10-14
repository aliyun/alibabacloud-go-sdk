// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossBorderOptimizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCrossBorderOptimizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCrossBorderOptimizationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCrossBorderOptimizationResponseBody) *UpdateCrossBorderOptimizationResponse
	GetBody() *UpdateCrossBorderOptimizationResponseBody
}

type UpdateCrossBorderOptimizationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrossBorderOptimizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrossBorderOptimizationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossBorderOptimizationResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrossBorderOptimizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCrossBorderOptimizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCrossBorderOptimizationResponse) GetBody() *UpdateCrossBorderOptimizationResponseBody {
	return s.Body
}

func (s *UpdateCrossBorderOptimizationResponse) SetHeaders(v map[string]*string) *UpdateCrossBorderOptimizationResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrossBorderOptimizationResponse) SetStatusCode(v int32) *UpdateCrossBorderOptimizationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrossBorderOptimizationResponse) SetBody(v *UpdateCrossBorderOptimizationResponseBody) *UpdateCrossBorderOptimizationResponse {
	s.Body = v
	return s
}

func (s *UpdateCrossBorderOptimizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
