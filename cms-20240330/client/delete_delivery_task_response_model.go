// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeliveryTaskResponseBody) *DeleteDeliveryTaskResponse
	GetBody() *DeleteDeliveryTaskResponseBody
}

type DeleteDeliveryTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeliveryTaskResponse) GetBody() *DeleteDeliveryTaskResponseBody {
	return s.Body
}

func (s *DeleteDeliveryTaskResponse) SetHeaders(v map[string]*string) *DeleteDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeliveryTaskResponse) SetStatusCode(v int32) *DeleteDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeliveryTaskResponse) SetBody(v *DeleteDeliveryTaskResponseBody) *DeleteDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
