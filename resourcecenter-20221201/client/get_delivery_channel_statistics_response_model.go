// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeliveryChannelStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeliveryChannelStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetDeliveryChannelStatisticsResponseBody) *GetDeliveryChannelStatisticsResponse
	GetBody() *GetDeliveryChannelStatisticsResponseBody
}

type GetDeliveryChannelStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryChannelStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryChannelStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeliveryChannelStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeliveryChannelStatisticsResponse) GetBody() *GetDeliveryChannelStatisticsResponseBody {
	return s.Body
}

func (s *GetDeliveryChannelStatisticsResponse) SetHeaders(v map[string]*string) *GetDeliveryChannelStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryChannelStatisticsResponse) SetStatusCode(v int32) *GetDeliveryChannelStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryChannelStatisticsResponse) SetBody(v *GetDeliveryChannelStatisticsResponseBody) *GetDeliveryChannelStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetDeliveryChannelStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
