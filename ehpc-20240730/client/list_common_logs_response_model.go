// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCommonLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCommonLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListCommonLogsResponseBody) *ListCommonLogsResponse
	GetBody() *ListCommonLogsResponseBody
}

type ListCommonLogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCommonLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCommonLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCommonLogsResponse) GoString() string {
	return s.String()
}

func (s *ListCommonLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCommonLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCommonLogsResponse) GetBody() *ListCommonLogsResponseBody {
	return s.Body
}

func (s *ListCommonLogsResponse) SetHeaders(v map[string]*string) *ListCommonLogsResponse {
	s.Headers = v
	return s
}

func (s *ListCommonLogsResponse) SetStatusCode(v int32) *ListCommonLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonLogsResponse) SetBody(v *ListCommonLogsResponseBody) *ListCommonLogsResponse {
	s.Body = v
	return s
}

func (s *ListCommonLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
