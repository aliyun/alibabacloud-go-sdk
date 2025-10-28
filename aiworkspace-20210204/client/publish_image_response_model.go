// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishImageResponse
	GetStatusCode() *int32
	SetBody(v *PublishImageResponseBody) *PublishImageResponse
	GetBody() *PublishImageResponseBody
}

type PublishImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishImageResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishImageResponse) GoString() string {
	return s.String()
}

func (s *PublishImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishImageResponse) GetBody() *PublishImageResponseBody {
	return s.Body
}

func (s *PublishImageResponse) SetHeaders(v map[string]*string) *PublishImageResponse {
	s.Headers = v
	return s
}

func (s *PublishImageResponse) SetStatusCode(v int32) *PublishImageResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishImageResponse) SetBody(v *PublishImageResponseBody) *PublishImageResponse {
	s.Body = v
	return s
}

func (s *PublishImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
