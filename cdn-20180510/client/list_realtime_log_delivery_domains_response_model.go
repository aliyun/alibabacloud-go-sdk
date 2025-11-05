// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeLogDeliveryDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRealtimeLogDeliveryDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListRealtimeLogDeliveryDomainsResponseBody) *ListRealtimeLogDeliveryDomainsResponse
	GetBody() *ListRealtimeLogDeliveryDomainsResponseBody
}

type ListRealtimeLogDeliveryDomainsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRealtimeLogDeliveryDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRealtimeLogDeliveryDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListRealtimeLogDeliveryDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRealtimeLogDeliveryDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRealtimeLogDeliveryDomainsResponse) GetBody() *ListRealtimeLogDeliveryDomainsResponseBody {
	return s.Body
}

func (s *ListRealtimeLogDeliveryDomainsResponse) SetHeaders(v map[string]*string) *ListRealtimeLogDeliveryDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponse) SetStatusCode(v int32) *ListRealtimeLogDeliveryDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponse) SetBody(v *ListRealtimeLogDeliveryDomainsResponseBody) *ListRealtimeLogDeliveryDomainsResponse {
	s.Body = v
	return s
}

func (s *ListRealtimeLogDeliveryDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
