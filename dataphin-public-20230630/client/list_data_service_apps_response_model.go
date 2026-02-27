// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceAppsResponseBody) *ListDataServiceAppsResponse
	GetBody() *ListDataServiceAppsResponseBody
}

type ListDataServiceAppsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAppsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceAppsResponse) GetBody() *ListDataServiceAppsResponseBody {
	return s.Body
}

func (s *ListDataServiceAppsResponse) SetHeaders(v map[string]*string) *ListDataServiceAppsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceAppsResponse) SetStatusCode(v int32) *ListDataServiceAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceAppsResponse) SetBody(v *ListDataServiceAppsResponseBody) *ListDataServiceAppsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
