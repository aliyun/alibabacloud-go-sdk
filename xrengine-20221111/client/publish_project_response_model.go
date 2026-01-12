// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishProjectResponse
	GetStatusCode() *int32
	SetBody(v *PublishProjectResponseBody) *PublishProjectResponse
	GetBody() *PublishProjectResponseBody
}

type PublishProjectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishProjectResponse) GoString() string {
	return s.String()
}

func (s *PublishProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishProjectResponse) GetBody() *PublishProjectResponseBody {
	return s.Body
}

func (s *PublishProjectResponse) SetHeaders(v map[string]*string) *PublishProjectResponse {
	s.Headers = v
	return s
}

func (s *PublishProjectResponse) SetStatusCode(v int32) *PublishProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishProjectResponse) SetBody(v *PublishProjectResponseBody) *PublishProjectResponse {
	s.Body = v
	return s
}

func (s *PublishProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
