// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveMessageAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveMessageAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveMessageAppsResponseBody) *ListLiveMessageAppsResponse
	GetBody() *ListLiveMessageAppsResponseBody
}

type ListLiveMessageAppsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveMessageAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveMessageAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageAppsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveMessageAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveMessageAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveMessageAppsResponse) GetBody() *ListLiveMessageAppsResponseBody {
	return s.Body
}

func (s *ListLiveMessageAppsResponse) SetHeaders(v map[string]*string) *ListLiveMessageAppsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveMessageAppsResponse) SetStatusCode(v int32) *ListLiveMessageAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveMessageAppsResponse) SetBody(v *ListLiveMessageAppsResponseBody) *ListLiveMessageAppsResponse {
	s.Body = v
	return s
}

func (s *ListLiveMessageAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
