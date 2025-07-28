// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliveryItemDetailSynResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeliveryItemDetailSynResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeliveryItemDetailSynResponse
	GetStatusCode() *int32
	SetBody(v *DeliveryItemDetailSynResponseBody) *DeliveryItemDetailSynResponse
	GetBody() *DeliveryItemDetailSynResponseBody
}

type DeliveryItemDetailSynResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeliveryItemDetailSynResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeliveryItemDetailSynResponse) String() string {
	return dara.Prettify(s)
}

func (s DeliveryItemDetailSynResponse) GoString() string {
	return s.String()
}

func (s *DeliveryItemDetailSynResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeliveryItemDetailSynResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeliveryItemDetailSynResponse) GetBody() *DeliveryItemDetailSynResponseBody {
	return s.Body
}

func (s *DeliveryItemDetailSynResponse) SetHeaders(v map[string]*string) *DeliveryItemDetailSynResponse {
	s.Headers = v
	return s
}

func (s *DeliveryItemDetailSynResponse) SetStatusCode(v int32) *DeliveryItemDetailSynResponse {
	s.StatusCode = &v
	return s
}

func (s *DeliveryItemDetailSynResponse) SetBody(v *DeliveryItemDetailSynResponseBody) *DeliveryItemDetailSynResponse {
	s.Body = v
	return s
}

func (s *DeliveryItemDetailSynResponse) Validate() error {
	return dara.Validate(s)
}
