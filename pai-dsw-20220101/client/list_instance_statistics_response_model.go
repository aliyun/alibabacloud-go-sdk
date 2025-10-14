// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceStatisticsResponseBody) *ListInstanceStatisticsResponse
	GetBody() *ListInstanceStatisticsResponseBody
}

type ListInstanceStatisticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceStatisticsResponse) GetBody() *ListInstanceStatisticsResponseBody {
	return s.Body
}

func (s *ListInstanceStatisticsResponse) SetHeaders(v map[string]*string) *ListInstanceStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStatisticsResponse) SetStatusCode(v int32) *ListInstanceStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceStatisticsResponse) SetBody(v *ListInstanceStatisticsResponseBody) *ListInstanceStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
