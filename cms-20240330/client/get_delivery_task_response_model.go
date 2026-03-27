// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeliveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeliveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetDeliveryTaskResponseBody) *GetDeliveryTaskResponse
	GetBody() *GetDeliveryTaskResponseBody
}

type GetDeliveryTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeliveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeliveryTaskResponse) GetBody() *GetDeliveryTaskResponseBody {
	return s.Body
}

func (s *GetDeliveryTaskResponse) SetHeaders(v map[string]*string) *GetDeliveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryTaskResponse) SetStatusCode(v int32) *GetDeliveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryTaskResponse) SetBody(v *GetDeliveryTaskResponseBody) *GetDeliveryTaskResponse {
	s.Body = v
	return s
}

func (s *GetDeliveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
