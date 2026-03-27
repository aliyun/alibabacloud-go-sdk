// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeliveryTaskResponseBody) *CreateDeliveryTaskResponse
	GetBody() *CreateDeliveryTaskResponseBody
}

type CreateDeliveryTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeliveryTaskResponse) GetBody() *CreateDeliveryTaskResponseBody {
	return s.Body
}

func (s *CreateDeliveryTaskResponse) SetHeaders(v map[string]*string) *CreateDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryTaskResponse) SetStatusCode(v int32) *CreateDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryTaskResponse) SetBody(v *CreateDeliveryTaskResponseBody) *CreateDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *CreateDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
