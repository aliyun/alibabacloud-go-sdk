// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStatisticLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStatisticLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListStatisticLogsResponseBody) *ListStatisticLogsResponse
	GetBody() *ListStatisticLogsResponseBody
}

type ListStatisticLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStatisticLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStatisticLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticLogsResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStatisticLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStatisticLogsResponse) GetBody() *ListStatisticLogsResponseBody {
	return s.Body
}

func (s *ListStatisticLogsResponse) SetHeaders(v map[string]*string) *ListStatisticLogsResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticLogsResponse) SetStatusCode(v int32) *ListStatisticLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStatisticLogsResponse) SetBody(v *ListStatisticLogsResponseBody) *ListStatisticLogsResponse {
	s.Body = v
	return s
}

func (s *ListStatisticLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
