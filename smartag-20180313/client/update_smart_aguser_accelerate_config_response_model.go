// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGUserAccelerateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmartAGUserAccelerateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmartAGUserAccelerateConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmartAGUserAccelerateConfigResponseBody) *UpdateSmartAGUserAccelerateConfigResponse
	GetBody() *UpdateSmartAGUserAccelerateConfigResponseBody
}

type UpdateSmartAGUserAccelerateConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmartAGUserAccelerateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmartAGUserAccelerateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGUserAccelerateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) GetBody() *UpdateSmartAGUserAccelerateConfigResponseBody {
	return s.Body
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) SetHeaders(v map[string]*string) *UpdateSmartAGUserAccelerateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) SetStatusCode(v int32) *UpdateSmartAGUserAccelerateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) SetBody(v *UpdateSmartAGUserAccelerateConfigResponseBody) *UpdateSmartAGUserAccelerateConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
