// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOpenSourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOpenSourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOpenSourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOpenSourcePermissionResponseBody) *DeleteOpenSourcePermissionResponse
	GetBody() *DeleteOpenSourcePermissionResponseBody
}

type DeleteOpenSourcePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOpenSourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOpenSourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOpenSourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteOpenSourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOpenSourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOpenSourcePermissionResponse) GetBody() *DeleteOpenSourcePermissionResponseBody {
	return s.Body
}

func (s *DeleteOpenSourcePermissionResponse) SetHeaders(v map[string]*string) *DeleteOpenSourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteOpenSourcePermissionResponse) SetStatusCode(v int32) *DeleteOpenSourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOpenSourcePermissionResponse) SetBody(v *DeleteOpenSourcePermissionResponseBody) *DeleteOpenSourcePermissionResponse {
	s.Body = v
	return s
}

func (s *DeleteOpenSourcePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
