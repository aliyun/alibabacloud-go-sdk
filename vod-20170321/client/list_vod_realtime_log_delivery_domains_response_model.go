// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodRealtimeLogDeliveryDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodRealtimeLogDeliveryDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListVodRealtimeLogDeliveryDomainsResponseBody) *ListVodRealtimeLogDeliveryDomainsResponse
	GetBody() *ListVodRealtimeLogDeliveryDomainsResponseBody
}

type ListVodRealtimeLogDeliveryDomainsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodRealtimeLogDeliveryDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) GetBody() *ListVodRealtimeLogDeliveryDomainsResponseBody {
	return s.Body
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) SetStatusCode(v int32) *ListVodRealtimeLogDeliveryDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) SetBody(v *ListVodRealtimeLogDeliveryDomainsResponseBody) *ListVodRealtimeLogDeliveryDomainsResponse {
	s.Body = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) Validate() error {
	return dara.Validate(s)
}
