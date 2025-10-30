// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogDeliveryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogDeliveryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogDeliveryConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogDeliveryConfigResponseBody) *CreateLogDeliveryConfigResponse
	GetBody() *CreateLogDeliveryConfigResponseBody
}

type CreateLogDeliveryConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogDeliveryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogDeliveryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogDeliveryConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLogDeliveryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogDeliveryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogDeliveryConfigResponse) GetBody() *CreateLogDeliveryConfigResponseBody {
	return s.Body
}

func (s *CreateLogDeliveryConfigResponse) SetHeaders(v map[string]*string) *CreateLogDeliveryConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLogDeliveryConfigResponse) SetStatusCode(v int32) *CreateLogDeliveryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogDeliveryConfigResponse) SetBody(v *CreateLogDeliveryConfigResponseBody) *CreateLogDeliveryConfigResponse {
	s.Body = v
	return s
}

func (s *CreateLogDeliveryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
