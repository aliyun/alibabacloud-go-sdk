// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *PublishAppInstanceResponseBody) *PublishAppInstanceResponse
	GetBody() *PublishAppInstanceResponseBody
}

type PublishAppInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *PublishAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishAppInstanceResponse) GetBody() *PublishAppInstanceResponseBody {
	return s.Body
}

func (s *PublishAppInstanceResponse) SetHeaders(v map[string]*string) *PublishAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *PublishAppInstanceResponse) SetStatusCode(v int32) *PublishAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAppInstanceResponse) SetBody(v *PublishAppInstanceResponseBody) *PublishAppInstanceResponse {
	s.Body = v
	return s
}

func (s *PublishAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
