// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTraceAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTraceAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListTraceAppsResponseBody) *ListTraceAppsResponse
	GetBody() *ListTraceAppsResponseBody
}

type ListTraceAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTraceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTraceAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTraceAppsResponse) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTraceAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTraceAppsResponse) GetBody() *ListTraceAppsResponseBody {
	return s.Body
}

func (s *ListTraceAppsResponse) SetHeaders(v map[string]*string) *ListTraceAppsResponse {
	s.Headers = v
	return s
}

func (s *ListTraceAppsResponse) SetStatusCode(v int32) *ListTraceAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTraceAppsResponse) SetBody(v *ListTraceAppsResponseBody) *ListTraceAppsResponse {
	s.Body = v
	return s
}

func (s *ListTraceAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
