// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecLogDeliveryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecLogDeliveryStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecLogDeliveryStatusResponseBody) *ModifyApisecLogDeliveryStatusResponse
	GetBody() *ModifyApisecLogDeliveryStatusResponseBody
}

type ModifyApisecLogDeliveryStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecLogDeliveryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecLogDeliveryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecLogDeliveryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecLogDeliveryStatusResponse) GetBody() *ModifyApisecLogDeliveryStatusResponseBody {
	return s.Body
}

func (s *ModifyApisecLogDeliveryStatusResponse) SetHeaders(v map[string]*string) *ModifyApisecLogDeliveryStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecLogDeliveryStatusResponse) SetStatusCode(v int32) *ModifyApisecLogDeliveryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusResponse) SetBody(v *ModifyApisecLogDeliveryStatusResponseBody) *ModifyApisecLogDeliveryStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecLogDeliveryStatusResponse) Validate() error {
	return dara.Validate(s)
}
