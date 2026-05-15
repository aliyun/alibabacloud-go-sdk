// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineServiceStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineServiceStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineServiceStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineServiceStatisticsResponseBody) *GetHotlineServiceStatisticsResponse
	GetBody() *GetHotlineServiceStatisticsResponseBody
}

type GetHotlineServiceStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineServiceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineServiceStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineServiceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineServiceStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineServiceStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineServiceStatisticsResponse) GetBody() *GetHotlineServiceStatisticsResponseBody {
	return s.Body
}

func (s *GetHotlineServiceStatisticsResponse) SetHeaders(v map[string]*string) *GetHotlineServiceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineServiceStatisticsResponse) SetStatusCode(v int32) *GetHotlineServiceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineServiceStatisticsResponse) SetBody(v *GetHotlineServiceStatisticsResponseBody) *GetHotlineServiceStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetHotlineServiceStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
