// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserDeliveryTaskResponseBody) *UpdateUserDeliveryTaskResponse
	GetBody() *UpdateUserDeliveryTaskResponseBody
}

type UpdateUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserDeliveryTaskResponse) GetBody() *UpdateUserDeliveryTaskResponseBody {
	return s.Body
}

func (s *UpdateUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDeliveryTaskResponse) SetStatusCode(v int32) *UpdateUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDeliveryTaskResponse) SetBody(v *UpdateUserDeliveryTaskResponseBody) *UpdateUserDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateUserDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
