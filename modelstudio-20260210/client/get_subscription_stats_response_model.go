// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubscriptionStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubscriptionStatsResponse
	GetStatusCode() *int32
	SetBody(v *GetSubscriptionStatsResponseBody) *GetSubscriptionStatsResponse
	GetBody() *GetSubscriptionStatsResponseBody
}

type GetSubscriptionStatsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionStatsResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubscriptionStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubscriptionStatsResponse) GetBody() *GetSubscriptionStatsResponseBody {
	return s.Body
}

func (s *GetSubscriptionStatsResponse) SetHeaders(v map[string]*string) *GetSubscriptionStatsResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionStatsResponse) SetStatusCode(v int32) *GetSubscriptionStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionStatsResponse) SetBody(v *GetSubscriptionStatsResponseBody) *GetSubscriptionStatsResponse {
	s.Body = v
	return s
}

func (s *GetSubscriptionStatsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
