// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishGrayDomainConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishGrayDomainConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishGrayDomainConfigResponse
	GetStatusCode() *int32
	SetBody(v *PublishGrayDomainConfigResponseBody) *PublishGrayDomainConfigResponse
	GetBody() *PublishGrayDomainConfigResponseBody
}

type PublishGrayDomainConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishGrayDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishGrayDomainConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishGrayDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishGrayDomainConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishGrayDomainConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishGrayDomainConfigResponse) GetBody() *PublishGrayDomainConfigResponseBody {
	return s.Body
}

func (s *PublishGrayDomainConfigResponse) SetHeaders(v map[string]*string) *PublishGrayDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishGrayDomainConfigResponse) SetStatusCode(v int32) *PublishGrayDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishGrayDomainConfigResponse) SetBody(v *PublishGrayDomainConfigResponseBody) *PublishGrayDomainConfigResponse {
	s.Body = v
	return s
}

func (s *PublishGrayDomainConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
