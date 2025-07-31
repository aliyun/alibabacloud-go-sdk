// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *PublishDataServiceApiResponseBody) *PublishDataServiceApiResponse
	GetBody() *PublishDataServiceApiResponseBody
}

type PublishDataServiceApiResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *PublishDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishDataServiceApiResponse) GetBody() *PublishDataServiceApiResponseBody {
	return s.Body
}

func (s *PublishDataServiceApiResponse) SetHeaders(v map[string]*string) *PublishDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *PublishDataServiceApiResponse) SetStatusCode(v int32) *PublishDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishDataServiceApiResponse) SetBody(v *PublishDataServiceApiResponseBody) *PublishDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *PublishDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
