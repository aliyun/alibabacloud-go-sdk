// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogDeliveryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyResourceLogDeliveryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyResourceLogDeliveryStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyResourceLogDeliveryStatusResponseBody) *ModifyResourceLogDeliveryStatusResponse
	GetBody() *ModifyResourceLogDeliveryStatusResponseBody
}

type ModifyResourceLogDeliveryStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyResourceLogDeliveryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyResourceLogDeliveryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogDeliveryStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogDeliveryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyResourceLogDeliveryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyResourceLogDeliveryStatusResponse) GetBody() *ModifyResourceLogDeliveryStatusResponseBody {
	return s.Body
}

func (s *ModifyResourceLogDeliveryStatusResponse) SetHeaders(v map[string]*string) *ModifyResourceLogDeliveryStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponse) SetStatusCode(v int32) *ModifyResourceLogDeliveryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponse) SetBody(v *ModifyResourceLogDeliveryStatusResponseBody) *ModifyResourceLogDeliveryStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyResourceLogDeliveryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
