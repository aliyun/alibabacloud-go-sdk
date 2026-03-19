// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicStatisticsResponseBody) *GetBasicStatisticsResponse
	GetBody() *GetBasicStatisticsResponseBody
}

type GetBasicStatisticsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetBasicStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicStatisticsResponse) GetBody() *GetBasicStatisticsResponseBody {
	return s.Body
}

func (s *GetBasicStatisticsResponse) SetHeaders(v map[string]*string) *GetBasicStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetBasicStatisticsResponse) SetStatusCode(v int32) *GetBasicStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicStatisticsResponse) SetBody(v *GetBasicStatisticsResponseBody) *GetBasicStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetBasicStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
