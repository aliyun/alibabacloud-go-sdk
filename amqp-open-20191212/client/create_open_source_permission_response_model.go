// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOpenSourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOpenSourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *CreateOpenSourcePermissionResponseBody) *CreateOpenSourcePermissionResponse
	GetBody() *CreateOpenSourcePermissionResponseBody
}

type CreateOpenSourcePermissionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpenSourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpenSourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *CreateOpenSourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOpenSourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOpenSourcePermissionResponse) GetBody() *CreateOpenSourcePermissionResponseBody {
	return s.Body
}

func (s *CreateOpenSourcePermissionResponse) SetHeaders(v map[string]*string) *CreateOpenSourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *CreateOpenSourcePermissionResponse) SetStatusCode(v int32) *CreateOpenSourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpenSourcePermissionResponse) SetBody(v *CreateOpenSourcePermissionResponseBody) *CreateOpenSourcePermissionResponse {
	s.Body = v
	return s
}

func (s *CreateOpenSourcePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
