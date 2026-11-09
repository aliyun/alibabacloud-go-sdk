// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomAuthPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomAuthPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomAuthPermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomAuthPermissionResponseBody) *DeleteCustomAuthPermissionResponse
	GetBody() *DeleteCustomAuthPermissionResponseBody
}

type DeleteCustomAuthPermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomAuthPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomAuthPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomAuthPermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomAuthPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomAuthPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomAuthPermissionResponse) GetBody() *DeleteCustomAuthPermissionResponseBody {
	return s.Body
}

func (s *DeleteCustomAuthPermissionResponse) SetHeaders(v map[string]*string) *DeleteCustomAuthPermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomAuthPermissionResponse) SetStatusCode(v int32) *DeleteCustomAuthPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomAuthPermissionResponse) SetBody(v *DeleteCustomAuthPermissionResponseBody) *DeleteCustomAuthPermissionResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomAuthPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
