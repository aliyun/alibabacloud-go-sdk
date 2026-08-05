// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskErrorLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOfflineTaskErrorLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOfflineTaskErrorLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListOfflineTaskErrorLogsResponseBody) *ListOfflineTaskErrorLogsResponse
	GetBody() *ListOfflineTaskErrorLogsResponseBody
}

type ListOfflineTaskErrorLogsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOfflineTaskErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOfflineTaskErrorLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskErrorLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskErrorLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOfflineTaskErrorLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOfflineTaskErrorLogsResponse) GetBody() *ListOfflineTaskErrorLogsResponseBody {
	return s.Body
}

func (s *ListOfflineTaskErrorLogsResponse) SetHeaders(v map[string]*string) *ListOfflineTaskErrorLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOfflineTaskErrorLogsResponse) SetStatusCode(v int32) *ListOfflineTaskErrorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfflineTaskErrorLogsResponse) SetBody(v *ListOfflineTaskErrorLogsResponseBody) *ListOfflineTaskErrorLogsResponse {
	s.Body = v
	return s
}

func (s *ListOfflineTaskErrorLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
