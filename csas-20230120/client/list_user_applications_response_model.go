// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserApplicationsResponseBody) *ListUserApplicationsResponse
	GetBody() *ListUserApplicationsResponseBody
}

type ListUserApplicationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserApplicationsResponse) GetBody() *ListUserApplicationsResponseBody {
	return s.Body
}

func (s *ListUserApplicationsResponse) SetHeaders(v map[string]*string) *ListUserApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListUserApplicationsResponse) SetStatusCode(v int32) *ListUserApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserApplicationsResponse) SetBody(v *ListUserApplicationsResponseBody) *ListUserApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListUserApplicationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
