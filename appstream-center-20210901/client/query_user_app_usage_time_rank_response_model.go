// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserAppUsageTimeRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserAppUsageTimeRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserAppUsageTimeRankResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserAppUsageTimeRankResponseBody) *QueryUserAppUsageTimeRankResponse
	GetBody() *QueryUserAppUsageTimeRankResponseBody
}

type QueryUserAppUsageTimeRankResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserAppUsageTimeRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserAppUsageTimeRankResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserAppUsageTimeRankResponse) GoString() string {
	return s.String()
}

func (s *QueryUserAppUsageTimeRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserAppUsageTimeRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserAppUsageTimeRankResponse) GetBody() *QueryUserAppUsageTimeRankResponseBody {
	return s.Body
}

func (s *QueryUserAppUsageTimeRankResponse) SetHeaders(v map[string]*string) *QueryUserAppUsageTimeRankResponse {
	s.Headers = v
	return s
}

func (s *QueryUserAppUsageTimeRankResponse) SetStatusCode(v int32) *QueryUserAppUsageTimeRankResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserAppUsageTimeRankResponse) SetBody(v *QueryUserAppUsageTimeRankResponseBody) *QueryUserAppUsageTimeRankResponse {
	s.Body = v
	return s
}

func (s *QueryUserAppUsageTimeRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
