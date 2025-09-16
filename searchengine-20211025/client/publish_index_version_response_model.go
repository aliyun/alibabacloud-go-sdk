// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishIndexVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishIndexVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishIndexVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishIndexVersionResponseBody) *PublishIndexVersionResponse
	GetBody() *PublishIndexVersionResponseBody
}

type PublishIndexVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishIndexVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishIndexVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishIndexVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishIndexVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishIndexVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishIndexVersionResponse) GetBody() *PublishIndexVersionResponseBody {
	return s.Body
}

func (s *PublishIndexVersionResponse) SetHeaders(v map[string]*string) *PublishIndexVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishIndexVersionResponse) SetStatusCode(v int32) *PublishIndexVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishIndexVersionResponse) SetBody(v *PublishIndexVersionResponseBody) *PublishIndexVersionResponse {
	s.Body = v
	return s
}

func (s *PublishIndexVersionResponse) Validate() error {
	return dara.Validate(s)
}
