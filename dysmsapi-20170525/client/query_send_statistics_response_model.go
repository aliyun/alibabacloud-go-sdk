// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySendStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySendStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySendStatisticsResponseBody) *QuerySendStatisticsResponse
	GetBody() *QuerySendStatisticsResponseBody
}

type QuerySendStatisticsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySendStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySendStatisticsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySendStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySendStatisticsResponse) GetBody() *QuerySendStatisticsResponseBody {
	return s.Body
}

func (s *QuerySendStatisticsResponse) SetHeaders(v map[string]*string) *QuerySendStatisticsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendStatisticsResponse) SetStatusCode(v int32) *QuerySendStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySendStatisticsResponse) SetBody(v *QuerySendStatisticsResponseBody) *QuerySendStatisticsResponse {
	s.Body = v
	return s
}

func (s *QuerySendStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
