// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerlessTopNAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerlessTopNAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerlessTopNAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListServerlessTopNAppsResponseBody) *ListServerlessTopNAppsResponse
	GetBody() *ListServerlessTopNAppsResponseBody
}

type ListServerlessTopNAppsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerlessTopNAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerlessTopNAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerlessTopNAppsResponse) GoString() string {
	return s.String()
}

func (s *ListServerlessTopNAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerlessTopNAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerlessTopNAppsResponse) GetBody() *ListServerlessTopNAppsResponseBody {
	return s.Body
}

func (s *ListServerlessTopNAppsResponse) SetHeaders(v map[string]*string) *ListServerlessTopNAppsResponse {
	s.Headers = v
	return s
}

func (s *ListServerlessTopNAppsResponse) SetStatusCode(v int32) *ListServerlessTopNAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerlessTopNAppsResponse) SetBody(v *ListServerlessTopNAppsResponseBody) *ListServerlessTopNAppsResponse {
	s.Body = v
	return s
}

func (s *ListServerlessTopNAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
