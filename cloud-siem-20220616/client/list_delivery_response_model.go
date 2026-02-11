// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *ListDeliveryResponseBody) *ListDeliveryResponse
	GetBody() *ListDeliveryResponseBody
}

type ListDeliveryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeliveryResponse) GoString() string {
	return s.String()
}

func (s *ListDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeliveryResponse) GetBody() *ListDeliveryResponseBody {
	return s.Body
}

func (s *ListDeliveryResponse) SetHeaders(v map[string]*string) *ListDeliveryResponse {
	s.Headers = v
	return s
}

func (s *ListDeliveryResponse) SetStatusCode(v int32) *ListDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeliveryResponse) SetBody(v *ListDeliveryResponseBody) *ListDeliveryResponse {
	s.Body = v
	return s
}

func (s *ListDeliveryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
