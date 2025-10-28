// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishDatasetResponse
	GetStatusCode() *int32
	SetBody(v *PublishDatasetResponseBody) *PublishDatasetResponse
	GetBody() *PublishDatasetResponseBody
}

type PublishDatasetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishDatasetResponse) GoString() string {
	return s.String()
}

func (s *PublishDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishDatasetResponse) GetBody() *PublishDatasetResponseBody {
	return s.Body
}

func (s *PublishDatasetResponse) SetHeaders(v map[string]*string) *PublishDatasetResponse {
	s.Headers = v
	return s
}

func (s *PublishDatasetResponse) SetStatusCode(v int32) *PublishDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishDatasetResponse) SetBody(v *PublishDatasetResponseBody) *PublishDatasetResponse {
	s.Body = v
	return s
}

func (s *PublishDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
