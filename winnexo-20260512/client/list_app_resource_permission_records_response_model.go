// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppResourcePermissionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppResourcePermissionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppResourcePermissionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppResourcePermissionRecordsResponseBody) *ListAppResourcePermissionRecordsResponse
	GetBody() *ListAppResourcePermissionRecordsResponseBody
}

type ListAppResourcePermissionRecordsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppResourcePermissionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppResourcePermissionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppResourcePermissionRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAppResourcePermissionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppResourcePermissionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppResourcePermissionRecordsResponse) GetBody() *ListAppResourcePermissionRecordsResponseBody {
	return s.Body
}

func (s *ListAppResourcePermissionRecordsResponse) SetHeaders(v map[string]*string) *ListAppResourcePermissionRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAppResourcePermissionRecordsResponse) SetStatusCode(v int32) *ListAppResourcePermissionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResourcePermissionRecordsResponse) SetBody(v *ListAppResourcePermissionRecordsResponseBody) *ListAppResourcePermissionRecordsResponse {
	s.Body = v
	return s
}

func (s *ListAppResourcePermissionRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
