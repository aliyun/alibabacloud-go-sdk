// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualDeliveryToMsceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VirtualDeliveryToMsceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VirtualDeliveryToMsceneResponse
	GetStatusCode() *int32
	SetBody(v *VirtualDeliveryToMsceneResponseBody) *VirtualDeliveryToMsceneResponse
	GetBody() *VirtualDeliveryToMsceneResponseBody
}

type VirtualDeliveryToMsceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VirtualDeliveryToMsceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VirtualDeliveryToMsceneResponse) String() string {
	return dara.Prettify(s)
}

func (s VirtualDeliveryToMsceneResponse) GoString() string {
	return s.String()
}

func (s *VirtualDeliveryToMsceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VirtualDeliveryToMsceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VirtualDeliveryToMsceneResponse) GetBody() *VirtualDeliveryToMsceneResponseBody {
	return s.Body
}

func (s *VirtualDeliveryToMsceneResponse) SetHeaders(v map[string]*string) *VirtualDeliveryToMsceneResponse {
	s.Headers = v
	return s
}

func (s *VirtualDeliveryToMsceneResponse) SetStatusCode(v int32) *VirtualDeliveryToMsceneResponse {
	s.StatusCode = &v
	return s
}

func (s *VirtualDeliveryToMsceneResponse) SetBody(v *VirtualDeliveryToMsceneResponseBody) *VirtualDeliveryToMsceneResponse {
	s.Body = v
	return s
}

func (s *VirtualDeliveryToMsceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
