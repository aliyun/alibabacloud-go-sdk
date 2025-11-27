// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePostPayInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePostPayInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePostPayInstanceResponseBody) *CreatePostPayInstanceResponse
	GetBody() *CreatePostPayInstanceResponseBody
}

type CreatePostPayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostPayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostPayInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePostPayInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePostPayInstanceResponse) GetBody() *CreatePostPayInstanceResponseBody {
	return s.Body
}

func (s *CreatePostPayInstanceResponse) SetHeaders(v map[string]*string) *CreatePostPayInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePostPayInstanceResponse) SetStatusCode(v int32) *CreatePostPayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostPayInstanceResponse) SetBody(v *CreatePostPayInstanceResponseBody) *CreatePostPayInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePostPayInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
