// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *CloseDeliveryResponseBody) *CloseDeliveryResponse
	GetBody() *CloseDeliveryResponseBody
}

type CloseDeliveryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CloseDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseDeliveryResponse) GetBody() *CloseDeliveryResponseBody {
	return s.Body
}

func (s *CloseDeliveryResponse) SetHeaders(v map[string]*string) *CloseDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CloseDeliveryResponse) SetStatusCode(v int32) *CloseDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDeliveryResponse) SetBody(v *CloseDeliveryResponseBody) *CloseDeliveryResponse {
	s.Body = v
	return s
}

func (s *CloseDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
