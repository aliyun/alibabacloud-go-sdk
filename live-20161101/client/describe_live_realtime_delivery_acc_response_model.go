// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeDeliveryAccResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRealtimeDeliveryAccResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRealtimeDeliveryAccResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRealtimeDeliveryAccResponseBody) *DescribeLiveRealtimeDeliveryAccResponse
	GetBody() *DescribeLiveRealtimeDeliveryAccResponseBody
}

type DescribeLiveRealtimeDeliveryAccResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRealtimeDeliveryAccResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRealtimeDeliveryAccResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) GetBody() *DescribeLiveRealtimeDeliveryAccResponseBody {
	return s.Body
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) SetHeaders(v map[string]*string) *DescribeLiveRealtimeDeliveryAccResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) SetStatusCode(v int32) *DescribeLiveRealtimeDeliveryAccResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) SetBody(v *DescribeLiveRealtimeDeliveryAccResponseBody) *DescribeLiveRealtimeDeliveryAccResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccResponse) Validate() error {
	return dara.Validate(s)
}
