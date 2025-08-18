// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserDeliveryTaskResponseBody) *CreateUserDeliveryTaskResponse
	GetBody() *CreateUserDeliveryTaskResponseBody
}

type CreateUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserDeliveryTaskResponse) GetBody() *CreateUserDeliveryTaskResponseBody {
	return s.Body
}

func (s *CreateUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *CreateUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateUserDeliveryTaskResponse) SetStatusCode(v int32) *CreateUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserDeliveryTaskResponse) SetBody(v *CreateUserDeliveryTaskResponseBody) *CreateUserDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateUserDeliveryTaskResponse) Validate() error {
	return dara.Validate(s)
}
