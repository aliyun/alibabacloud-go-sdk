// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserDeliveryTaskResponseBody) *DeleteUserDeliveryTaskResponse
	GetBody() *DeleteUserDeliveryTaskResponseBody
}

type DeleteUserDeliveryTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserDeliveryTaskResponse) GetBody() *DeleteUserDeliveryTaskResponseBody {
	return s.Body
}

func (s *DeleteUserDeliveryTaskResponse) SetHeaders(v map[string]*string) *DeleteUserDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDeliveryTaskResponse) SetStatusCode(v int32) *DeleteUserDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDeliveryTaskResponse) SetBody(v *DeleteUserDeliveryTaskResponseBody) *DeleteUserDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteUserDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
