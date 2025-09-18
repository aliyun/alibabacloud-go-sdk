// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMmAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishMmAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishMmAppResponse
	GetStatusCode() *int32
	SetBody(v *PublishMmAppResponseBody) *PublishMmAppResponse
	GetBody() *PublishMmAppResponseBody
}

type PublishMmAppResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishMmAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishMmAppResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishMmAppResponse) GoString() string {
	return s.String()
}

func (s *PublishMmAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishMmAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishMmAppResponse) GetBody() *PublishMmAppResponseBody {
	return s.Body
}

func (s *PublishMmAppResponse) SetHeaders(v map[string]*string) *PublishMmAppResponse {
	s.Headers = v
	return s
}

func (s *PublishMmAppResponse) SetStatusCode(v int32) *PublishMmAppResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishMmAppResponse) SetBody(v *PublishMmAppResponseBody) *PublishMmAppResponse {
	s.Body = v
	return s
}

func (s *PublishMmAppResponse) Validate() error {
	return dara.Validate(s)
}
