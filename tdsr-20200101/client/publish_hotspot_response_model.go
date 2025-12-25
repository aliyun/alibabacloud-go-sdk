// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishHotspotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishHotspotResponse
	GetStatusCode() *int32
	SetBody(v *PublishHotspotResponseBody) *PublishHotspotResponse
	GetBody() *PublishHotspotResponseBody
}

type PublishHotspotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishHotspotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishHotspotResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotResponse) GoString() string {
	return s.String()
}

func (s *PublishHotspotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishHotspotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishHotspotResponse) GetBody() *PublishHotspotResponseBody {
	return s.Body
}

func (s *PublishHotspotResponse) SetHeaders(v map[string]*string) *PublishHotspotResponse {
	s.Headers = v
	return s
}

func (s *PublishHotspotResponse) SetStatusCode(v int32) *PublishHotspotResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishHotspotResponse) SetBody(v *PublishHotspotResponseBody) *PublishHotspotResponse {
	s.Body = v
	return s
}

func (s *PublishHotspotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
