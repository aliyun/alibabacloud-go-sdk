// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *OpenDeliveryResponseBody) *OpenDeliveryResponse
	GetBody() *OpenDeliveryResponseBody
}

type OpenDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenDeliveryResponse) GoString() string {
	return s.String()
}

func (s *OpenDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenDeliveryResponse) GetBody() *OpenDeliveryResponseBody {
	return s.Body
}

func (s *OpenDeliveryResponse) SetHeaders(v map[string]*string) *OpenDeliveryResponse {
	s.Headers = v
	return s
}

func (s *OpenDeliveryResponse) SetStatusCode(v int32) *OpenDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenDeliveryResponse) SetBody(v *OpenDeliveryResponseBody) *OpenDeliveryResponse {
	s.Body = v
	return s
}

func (s *OpenDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
