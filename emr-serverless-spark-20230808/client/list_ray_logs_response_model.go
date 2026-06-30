// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRayLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRayLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListRayLogsResponseBody) *ListRayLogsResponse
	GetBody() *ListRayLogsResponseBody
}

type ListRayLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRayLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRayLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRayLogsResponse) GoString() string {
	return s.String()
}

func (s *ListRayLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRayLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRayLogsResponse) GetBody() *ListRayLogsResponseBody {
	return s.Body
}

func (s *ListRayLogsResponse) SetHeaders(v map[string]*string) *ListRayLogsResponse {
	s.Headers = v
	return s
}

func (s *ListRayLogsResponse) SetStatusCode(v int32) *ListRayLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRayLogsResponse) SetBody(v *ListRayLogsResponseBody) *ListRayLogsResponse {
	s.Body = v
	return s
}

func (s *ListRayLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
