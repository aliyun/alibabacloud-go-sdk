// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicAcceleratorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicAcceleratorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicAcceleratorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicAcceleratorResponseBody) *UpdateBasicAcceleratorResponse
	GetBody() *UpdateBasicAcceleratorResponseBody
}

type UpdateBasicAcceleratorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicAcceleratorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicAcceleratorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicAcceleratorResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicAcceleratorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicAcceleratorResponse) GetBody() *UpdateBasicAcceleratorResponseBody {
	return s.Body
}

func (s *UpdateBasicAcceleratorResponse) SetHeaders(v map[string]*string) *UpdateBasicAcceleratorResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicAcceleratorResponse) SetStatusCode(v int32) *UpdateBasicAcceleratorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicAcceleratorResponse) SetBody(v *UpdateBasicAcceleratorResponseBody) *UpdateBasicAcceleratorResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicAcceleratorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
