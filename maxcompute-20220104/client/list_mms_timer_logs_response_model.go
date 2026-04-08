// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTimerLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsTimerLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsTimerLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsTimerLogsResponseBody) *ListMmsTimerLogsResponse
	GetBody() *ListMmsTimerLogsResponseBody
}

type ListMmsTimerLogsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTimerLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTimerLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTimerLogsResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTimerLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsTimerLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsTimerLogsResponse) GetBody() *ListMmsTimerLogsResponseBody {
	return s.Body
}

func (s *ListMmsTimerLogsResponse) SetHeaders(v map[string]*string) *ListMmsTimerLogsResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTimerLogsResponse) SetStatusCode(v int32) *ListMmsTimerLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTimerLogsResponse) SetBody(v *ListMmsTimerLogsResponseBody) *ListMmsTimerLogsResponse {
	s.Body = v
	return s
}

func (s *ListMmsTimerLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
