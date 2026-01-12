// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishArEditProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishArEditProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishArEditProjectResponse
	GetStatusCode() *int32
	SetBody(v *PublishArEditProjectResponseBody) *PublishArEditProjectResponse
	GetBody() *PublishArEditProjectResponseBody
}

type PublishArEditProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishArEditProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishArEditProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishArEditProjectResponse) GoString() string {
	return s.String()
}

func (s *PublishArEditProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishArEditProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishArEditProjectResponse) GetBody() *PublishArEditProjectResponseBody {
	return s.Body
}

func (s *PublishArEditProjectResponse) SetHeaders(v map[string]*string) *PublishArEditProjectResponse {
	s.Headers = v
	return s
}

func (s *PublishArEditProjectResponse) SetStatusCode(v int32) *PublishArEditProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishArEditProjectResponse) SetBody(v *PublishArEditProjectResponseBody) *PublishArEditProjectResponse {
	s.Body = v
	return s
}

func (s *PublishArEditProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
