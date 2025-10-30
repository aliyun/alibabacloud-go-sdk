// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceAuthorizedAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceAuthorizedAppsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceAuthorizedAppsResponseBody) *ListDataServiceAuthorizedAppsResponse
	GetBody() *ListDataServiceAuthorizedAppsResponseBody
}

type ListDataServiceAuthorizedAppsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceAuthorizedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceAuthorizedAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedAppsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceAuthorizedAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceAuthorizedAppsResponse) GetBody() *ListDataServiceAuthorizedAppsResponseBody {
	return s.Body
}

func (s *ListDataServiceAuthorizedAppsResponse) SetHeaders(v map[string]*string) *ListDataServiceAuthorizedAppsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponse) SetStatusCode(v int32) *ListDataServiceAuthorizedAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponse) SetBody(v *ListDataServiceAuthorizedAppsResponseBody) *ListDataServiceAuthorizedAppsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceAuthorizedAppsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
