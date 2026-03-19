// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatsEventRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStatsEventRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStatsEventRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListStatsEventRecordsResponseBody) *ListStatsEventRecordsResponse
	GetBody() *ListStatsEventRecordsResponseBody
}

type ListStatsEventRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStatsEventRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStatsEventRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStatsEventRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListStatsEventRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStatsEventRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStatsEventRecordsResponse) GetBody() *ListStatsEventRecordsResponseBody {
	return s.Body
}

func (s *ListStatsEventRecordsResponse) SetHeaders(v map[string]*string) *ListStatsEventRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListStatsEventRecordsResponse) SetStatusCode(v int32) *ListStatsEventRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStatsEventRecordsResponse) SetBody(v *ListStatsEventRecordsResponseBody) *ListStatsEventRecordsResponse {
	s.Body = v
	return s
}

func (s *ListStatsEventRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
