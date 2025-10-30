// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLogDeliveryConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLogDeliveryConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLogDeliveryConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLogDeliveryConfigResponseBody) *ModifyLogDeliveryConfigResponse
	GetBody() *ModifyLogDeliveryConfigResponseBody
}

type ModifyLogDeliveryConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLogDeliveryConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLogDeliveryConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLogDeliveryConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogDeliveryConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLogDeliveryConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLogDeliveryConfigResponse) GetBody() *ModifyLogDeliveryConfigResponseBody {
	return s.Body
}

func (s *ModifyLogDeliveryConfigResponse) SetHeaders(v map[string]*string) *ModifyLogDeliveryConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogDeliveryConfigResponse) SetStatusCode(v int32) *ModifyLogDeliveryConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogDeliveryConfigResponse) SetBody(v *ModifyLogDeliveryConfigResponseBody) *ModifyLogDeliveryConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyLogDeliveryConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
