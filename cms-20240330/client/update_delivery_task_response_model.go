// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeliveryTaskResponseBody) *UpdateDeliveryTaskResponse
	GetBody() *UpdateDeliveryTaskResponseBody
}

type UpdateDeliveryTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeliveryTaskResponse) GetBody() *UpdateDeliveryTaskResponseBody {
	return s.Body
}

func (s *UpdateDeliveryTaskResponse) SetHeaders(v map[string]*string) *UpdateDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeliveryTaskResponse) SetStatusCode(v int32) *UpdateDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeliveryTaskResponse) SetBody(v *UpdateDeliveryTaskResponseBody) *UpdateDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
