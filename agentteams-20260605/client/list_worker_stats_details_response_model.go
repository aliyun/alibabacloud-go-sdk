// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerStatsDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkerStatsDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkerStatsDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkerStatsDetailsResponseBody) *ListWorkerStatsDetailsResponse
	GetBody() *ListWorkerStatsDetailsResponseBody
}

type ListWorkerStatsDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkerStatsDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkerStatsDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerStatsDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkerStatsDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkerStatsDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkerStatsDetailsResponse) GetBody() *ListWorkerStatsDetailsResponseBody {
	return s.Body
}

func (s *ListWorkerStatsDetailsResponse) SetHeaders(v map[string]*string) *ListWorkerStatsDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkerStatsDetailsResponse) SetStatusCode(v int32) *ListWorkerStatsDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkerStatsDetailsResponse) SetBody(v *ListWorkerStatsDetailsResponseBody) *ListWorkerStatsDetailsResponse {
	s.Body = v
	return s
}

func (s *ListWorkerStatsDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
