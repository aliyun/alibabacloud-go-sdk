// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishCodeSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishCodeSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishCodeSourceResponse
	GetStatusCode() *int32
	SetBody(v *PublishCodeSourceResponseBody) *PublishCodeSourceResponse
	GetBody() *PublishCodeSourceResponseBody
}

type PublishCodeSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishCodeSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *PublishCodeSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishCodeSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishCodeSourceResponse) GetBody() *PublishCodeSourceResponseBody {
	return s.Body
}

func (s *PublishCodeSourceResponse) SetHeaders(v map[string]*string) *PublishCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *PublishCodeSourceResponse) SetStatusCode(v int32) *PublishCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishCodeSourceResponse) SetBody(v *PublishCodeSourceResponseBody) *PublishCodeSourceResponse {
	s.Body = v
	return s
}

func (s *PublishCodeSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
