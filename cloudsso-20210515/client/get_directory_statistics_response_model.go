// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDirectoryStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDirectoryStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetDirectoryStatisticsResponseBody) *GetDirectoryStatisticsResponse
	GetBody() *GetDirectoryStatisticsResponseBody
}

type GetDirectoryStatisticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDirectoryStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDirectoryStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDirectoryStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDirectoryStatisticsResponse) GetBody() *GetDirectoryStatisticsResponseBody {
	return s.Body
}

func (s *GetDirectoryStatisticsResponse) SetHeaders(v map[string]*string) *GetDirectoryStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDirectoryStatisticsResponse) SetStatusCode(v int32) *GetDirectoryStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDirectoryStatisticsResponse) SetBody(v *GetDirectoryStatisticsResponseBody) *GetDirectoryStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetDirectoryStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
