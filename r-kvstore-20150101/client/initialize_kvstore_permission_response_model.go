// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeKvstorePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeKvstorePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeKvstorePermissionResponse
	GetStatusCode() *int32
	SetBody(v *InitializeKvstorePermissionResponseBody) *InitializeKvstorePermissionResponse
	GetBody() *InitializeKvstorePermissionResponseBody
}

type InitializeKvstorePermissionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeKvstorePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeKvstorePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeKvstorePermissionResponse) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeKvstorePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeKvstorePermissionResponse) GetBody() *InitializeKvstorePermissionResponseBody {
	return s.Body
}

func (s *InitializeKvstorePermissionResponse) SetHeaders(v map[string]*string) *InitializeKvstorePermissionResponse {
	s.Headers = v
	return s
}

func (s *InitializeKvstorePermissionResponse) SetStatusCode(v int32) *InitializeKvstorePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeKvstorePermissionResponse) SetBody(v *InitializeKvstorePermissionResponseBody) *InitializeKvstorePermissionResponse {
	s.Body = v
	return s
}

func (s *InitializeKvstorePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
