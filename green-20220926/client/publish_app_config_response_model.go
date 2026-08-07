// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishAppConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishAppConfigResponse
	GetStatusCode() *int32
	SetBody(v *PublishAppConfigResponseBody) *PublishAppConfigResponse
	GetBody() *PublishAppConfigResponseBody
}

type PublishAppConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAppConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishAppConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishAppConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishAppConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishAppConfigResponse) GetBody() *PublishAppConfigResponseBody {
	return s.Body
}

func (s *PublishAppConfigResponse) SetHeaders(v map[string]*string) *PublishAppConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishAppConfigResponse) SetStatusCode(v int32) *PublishAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAppConfigResponse) SetBody(v *PublishAppConfigResponseBody) *PublishAppConfigResponse {
	s.Body = v
	return s
}

func (s *PublishAppConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
