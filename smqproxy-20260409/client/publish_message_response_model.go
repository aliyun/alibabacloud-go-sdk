// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishMessageResponse
	GetStatusCode() *int32
	SetBody(v *PublishMessageResponseBody) *PublishMessageResponse
	GetBody() *PublishMessageResponseBody
}

type PublishMessageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishMessageResponse) GoString() string {
	return s.String()
}

func (s *PublishMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishMessageResponse) GetBody() *PublishMessageResponseBody {
	return s.Body
}

func (s *PublishMessageResponse) SetHeaders(v map[string]*string) *PublishMessageResponse {
	s.Headers = v
	return s
}

func (s *PublishMessageResponse) SetStatusCode(v int32) *PublishMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishMessageResponse) SetBody(v *PublishMessageResponseBody) *PublishMessageResponse {
	s.Body = v
	return s
}

func (s *PublishMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
