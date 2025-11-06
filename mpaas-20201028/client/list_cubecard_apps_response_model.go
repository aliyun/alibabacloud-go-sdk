// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCubecardAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCubecardAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCubecardAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListCubecardAppsResponseBody) *ListCubecardAppsResponse
	GetBody() *ListCubecardAppsResponseBody
}

type ListCubecardAppsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCubecardAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCubecardAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCubecardAppsResponse) GoString() string {
	return s.String()
}

func (s *ListCubecardAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCubecardAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCubecardAppsResponse) GetBody() *ListCubecardAppsResponseBody {
	return s.Body
}

func (s *ListCubecardAppsResponse) SetHeaders(v map[string]*string) *ListCubecardAppsResponse {
	s.Headers = v
	return s
}

func (s *ListCubecardAppsResponse) SetStatusCode(v int32) *ListCubecardAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCubecardAppsResponse) SetBody(v *ListCubecardAppsResponseBody) *ListCubecardAppsResponse {
	s.Body = v
	return s
}

func (s *ListCubecardAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
