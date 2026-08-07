// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserDocumentPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserDocumentPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GetUserDocumentPermissionResponseBody) *GetUserDocumentPermissionResponse
	GetBody() *GetUserDocumentPermissionResponseBody
}

type GetUserDocumentPermissionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDocumentPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDocumentPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserDocumentPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserDocumentPermissionResponse) GetBody() *GetUserDocumentPermissionResponseBody {
	return s.Body
}

func (s *GetUserDocumentPermissionResponse) SetHeaders(v map[string]*string) *GetUserDocumentPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetUserDocumentPermissionResponse) SetStatusCode(v int32) *GetUserDocumentPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDocumentPermissionResponse) SetBody(v *GetUserDocumentPermissionResponseBody) *GetUserDocumentPermissionResponse {
	s.Body = v
	return s
}

func (s *GetUserDocumentPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
