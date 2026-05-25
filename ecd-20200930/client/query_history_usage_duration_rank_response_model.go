// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryUsageDurationRankResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoryUsageDurationRankResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoryUsageDurationRankResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoryUsageDurationRankResponseBody) *QueryHistoryUsageDurationRankResponse
	GetBody() *QueryHistoryUsageDurationRankResponseBody
}

type QueryHistoryUsageDurationRankResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoryUsageDurationRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoryUsageDurationRankResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryUsageDurationRankResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoryUsageDurationRankResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoryUsageDurationRankResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoryUsageDurationRankResponse) GetBody() *QueryHistoryUsageDurationRankResponseBody {
	return s.Body
}

func (s *QueryHistoryUsageDurationRankResponse) SetHeaders(v map[string]*string) *QueryHistoryUsageDurationRankResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoryUsageDurationRankResponse) SetStatusCode(v int32) *QueryHistoryUsageDurationRankResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoryUsageDurationRankResponse) SetBody(v *QueryHistoryUsageDurationRankResponseBody) *QueryHistoryUsageDurationRankResponse {
	s.Body = v
	return s
}

func (s *QueryHistoryUsageDurationRankResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
