// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApisecLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApisecLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApisecLogDeliveryResponseBody) *ModifyApisecLogDeliveryResponse
	GetBody() *ModifyApisecLogDeliveryResponseBody
}

type ModifyApisecLogDeliveryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApisecLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApisecLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApisecLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApisecLogDeliveryResponse) GetBody() *ModifyApisecLogDeliveryResponseBody {
	return s.Body
}

func (s *ModifyApisecLogDeliveryResponse) SetHeaders(v map[string]*string) *ModifyApisecLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ModifyApisecLogDeliveryResponse) SetStatusCode(v int32) *ModifyApisecLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApisecLogDeliveryResponse) SetBody(v *ModifyApisecLogDeliveryResponseBody) *ModifyApisecLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *ModifyApisecLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
