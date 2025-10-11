// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiAccountDeliveryChannelStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiAccountDeliveryChannelStatisticsResponseBody) *GetMultiAccountDeliveryChannelStatisticsResponse
	GetBody() *GetMultiAccountDeliveryChannelStatisticsResponseBody
}

type GetMultiAccountDeliveryChannelStatisticsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiAccountDeliveryChannelStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiAccountDeliveryChannelStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) GetBody() *GetMultiAccountDeliveryChannelStatisticsResponseBody {
	return s.Body
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetHeaders(v map[string]*string) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetStatusCode(v int32) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) SetBody(v *GetMultiAccountDeliveryChannelStatisticsResponseBody) *GetMultiAccountDeliveryChannelStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetMultiAccountDeliveryChannelStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
