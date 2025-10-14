// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserDeliveryTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserDeliveryTaskStatusResponseBody) *UpdateUserDeliveryTaskStatusResponse
	GetBody() *UpdateUserDeliveryTaskStatusResponseBody
}

type UpdateUserDeliveryTaskStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDeliveryTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserDeliveryTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserDeliveryTaskStatusResponse) GetBody() *UpdateUserDeliveryTaskStatusResponseBody {
	return s.Body
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDeliveryTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetStatusCode(v int32) *UpdateUserDeliveryTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponse) SetBody(v *UpdateUserDeliveryTaskStatusResponseBody) *UpdateUserDeliveryTaskStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
