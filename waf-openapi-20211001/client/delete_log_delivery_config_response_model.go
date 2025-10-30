// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogDeliveryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogDeliveryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogDeliveryConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogDeliveryConfigResponseBody) *DeleteLogDeliveryConfigResponse
	GetBody() *DeleteLogDeliveryConfigResponseBody
}

type DeleteLogDeliveryConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogDeliveryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogDeliveryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogDeliveryConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogDeliveryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogDeliveryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogDeliveryConfigResponse) GetBody() *DeleteLogDeliveryConfigResponseBody {
	return s.Body
}

func (s *DeleteLogDeliveryConfigResponse) SetHeaders(v map[string]*string) *DeleteLogDeliveryConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogDeliveryConfigResponse) SetStatusCode(v int32) *DeleteLogDeliveryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogDeliveryConfigResponse) SetBody(v *DeleteLogDeliveryConfigResponseBody) *DeleteLogDeliveryConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLogDeliveryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
