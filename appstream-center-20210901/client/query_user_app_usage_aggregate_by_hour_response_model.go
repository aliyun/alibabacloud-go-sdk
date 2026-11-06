// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageAggregateByHourResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserAppUsageAggregateByHourResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserAppUsageAggregateByHourResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserAppUsageAggregateByHourResponseBody) *QueryUserAppUsageAggregateByHourResponse
	GetBody() *QueryUserAppUsageAggregateByHourResponseBody
}

type QueryUserAppUsageAggregateByHourResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserAppUsageAggregateByHourResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserAppUsageAggregateByHourResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageAggregateByHourResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageAggregateByHourResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserAppUsageAggregateByHourResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserAppUsageAggregateByHourResponse) GetBody() *QueryUserAppUsageAggregateByHourResponseBody {
	return s.Body
}

func (s *QueryUserAppUsageAggregateByHourResponse) SetHeaders(v map[string]*string) *QueryUserAppUsageAggregateByHourResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponse) SetStatusCode(v int32) *QueryUserAppUsageAggregateByHourResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponse) SetBody(v *QueryUserAppUsageAggregateByHourResponseBody) *QueryUserAppUsageAggregateByHourResponse {
	s.Body = v
	return s
}

func (s *QueryUserAppUsageAggregateByHourResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
