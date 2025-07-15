// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveRealtimeLogDeliveryDomainsResponseBody) *ListLiveRealtimeLogDeliveryDomainsResponse
	GetBody() *ListLiveRealtimeLogDeliveryDomainsResponseBody
}

type ListLiveRealtimeLogDeliveryDomainsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveRealtimeLogDeliveryDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveRealtimeLogDeliveryDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) GetBody() *ListLiveRealtimeLogDeliveryDomainsResponseBody {
	return s.Body
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) SetHeaders(v map[string]*string) *ListLiveRealtimeLogDeliveryDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) SetStatusCode(v int32) *ListLiveRealtimeLogDeliveryDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) SetBody(v *ListLiveRealtimeLogDeliveryDomainsResponseBody) *ListLiveRealtimeLogDeliveryDomainsResponse {
	s.Body = v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsResponse) Validate() error {
	return dara.Validate(s)
}
