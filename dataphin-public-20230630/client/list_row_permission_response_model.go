// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRowPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRowPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRowPermissionResponse
	GetStatusCode() *int32
	SetBody(v *ListRowPermissionResponseBody) *ListRowPermissionResponse
	GetBody() *ListRowPermissionResponseBody
}

type ListRowPermissionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRowPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRowPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRowPermissionResponse) GoString() string {
	return s.String()
}

func (s *ListRowPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRowPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRowPermissionResponse) GetBody() *ListRowPermissionResponseBody {
	return s.Body
}

func (s *ListRowPermissionResponse) SetHeaders(v map[string]*string) *ListRowPermissionResponse {
	s.Headers = v
	return s
}

func (s *ListRowPermissionResponse) SetStatusCode(v int32) *ListRowPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRowPermissionResponse) SetBody(v *ListRowPermissionResponseBody) *ListRowPermissionResponse {
	s.Body = v
	return s
}

func (s *ListRowPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
