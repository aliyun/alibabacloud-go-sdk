// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishApplicationVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishApplicationVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishApplicationVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishApplicationVersionResponseBody) *PublishApplicationVersionResponse
	GetBody() *PublishApplicationVersionResponseBody
}

type PublishApplicationVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishApplicationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishApplicationVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishApplicationVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishApplicationVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishApplicationVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishApplicationVersionResponse) GetBody() *PublishApplicationVersionResponseBody {
	return s.Body
}

func (s *PublishApplicationVersionResponse) SetHeaders(v map[string]*string) *PublishApplicationVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishApplicationVersionResponse) SetStatusCode(v int32) *PublishApplicationVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishApplicationVersionResponse) SetBody(v *PublishApplicationVersionResponseBody) *PublishApplicationVersionResponse {
	s.Body = v
	return s
}

func (s *PublishApplicationVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
