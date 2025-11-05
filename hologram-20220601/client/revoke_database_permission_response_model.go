// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDatabasePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeDatabasePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeDatabasePermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeDatabasePermissionResponseBody) *RevokeDatabasePermissionResponse
	GetBody() *RevokeDatabasePermissionResponseBody
}

type RevokeDatabasePermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeDatabasePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeDatabasePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeDatabasePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeDatabasePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeDatabasePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeDatabasePermissionResponse) GetBody() *RevokeDatabasePermissionResponseBody {
	return s.Body
}

func (s *RevokeDatabasePermissionResponse) SetHeaders(v map[string]*string) *RevokeDatabasePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeDatabasePermissionResponse) SetStatusCode(v int32) *RevokeDatabasePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeDatabasePermissionResponse) SetBody(v *RevokeDatabasePermissionResponseBody) *RevokeDatabasePermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeDatabasePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
