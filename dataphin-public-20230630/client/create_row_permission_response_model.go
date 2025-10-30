// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRowPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRowPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRowPermissionResponse
	GetStatusCode() *int32
	SetBody(v *CreateRowPermissionResponseBody) *CreateRowPermissionResponse
	GetBody() *CreateRowPermissionResponseBody
}

type CreateRowPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRowPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRowPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionResponse) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRowPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRowPermissionResponse) GetBody() *CreateRowPermissionResponseBody {
	return s.Body
}

func (s *CreateRowPermissionResponse) SetHeaders(v map[string]*string) *CreateRowPermissionResponse {
	s.Headers = v
	return s
}

func (s *CreateRowPermissionResponse) SetStatusCode(v int32) *CreateRowPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRowPermissionResponse) SetBody(v *CreateRowPermissionResponseBody) *CreateRowPermissionResponse {
	s.Body = v
	return s
}

func (s *CreateRowPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
