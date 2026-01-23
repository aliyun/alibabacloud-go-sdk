// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStandardStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStandardStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetStandardStatisticsResponseBody) *GetStandardStatisticsResponse
	GetBody() *GetStandardStatisticsResponseBody
}

type GetStandardStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStandardStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStandardStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStandardStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetStandardStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStandardStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStandardStatisticsResponse) GetBody() *GetStandardStatisticsResponseBody {
	return s.Body
}

func (s *GetStandardStatisticsResponse) SetHeaders(v map[string]*string) *GetStandardStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetStandardStatisticsResponse) SetStatusCode(v int32) *GetStandardStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStandardStatisticsResponse) SetBody(v *GetStandardStatisticsResponseBody) *GetStandardStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetStandardStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
