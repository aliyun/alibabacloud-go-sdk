// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostgresExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePostgresExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePostgresExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *CreatePostgresExtensionsResponseBody) *CreatePostgresExtensionsResponse
	GetBody() *CreatePostgresExtensionsResponseBody
}

type CreatePostgresExtensionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostgresExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostgresExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePostgresExtensionsResponse) GoString() string {
	return s.String()
}

func (s *CreatePostgresExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePostgresExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePostgresExtensionsResponse) GetBody() *CreatePostgresExtensionsResponseBody {
	return s.Body
}

func (s *CreatePostgresExtensionsResponse) SetHeaders(v map[string]*string) *CreatePostgresExtensionsResponse {
	s.Headers = v
	return s
}

func (s *CreatePostgresExtensionsResponse) SetStatusCode(v int32) *CreatePostgresExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostgresExtensionsResponse) SetBody(v *CreatePostgresExtensionsResponseBody) *CreatePostgresExtensionsResponse {
	s.Body = v
	return s
}

func (s *CreatePostgresExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
