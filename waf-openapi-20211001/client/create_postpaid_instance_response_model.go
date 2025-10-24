// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostpaidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePostpaidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePostpaidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePostpaidInstanceResponseBody) *CreatePostpaidInstanceResponse
	GetBody() *CreatePostpaidInstanceResponseBody
}

type CreatePostpaidInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostpaidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostpaidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePostpaidInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePostpaidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePostpaidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePostpaidInstanceResponse) GetBody() *CreatePostpaidInstanceResponseBody {
	return s.Body
}

func (s *CreatePostpaidInstanceResponse) SetHeaders(v map[string]*string) *CreatePostpaidInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePostpaidInstanceResponse) SetStatusCode(v int32) *CreatePostpaidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostpaidInstanceResponse) SetBody(v *CreatePostpaidInstanceResponseBody) *CreatePostpaidInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePostpaidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
