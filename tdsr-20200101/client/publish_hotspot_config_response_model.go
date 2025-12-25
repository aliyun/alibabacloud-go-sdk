// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishHotspotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishHotspotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishHotspotConfigResponse
	GetStatusCode() *int32
	SetBody(v *PublishHotspotConfigResponseBody) *PublishHotspotConfigResponse
	GetBody() *PublishHotspotConfigResponseBody
}

type PublishHotspotConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishHotspotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishHotspotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishHotspotConfigResponse) GoString() string {
	return s.String()
}

func (s *PublishHotspotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishHotspotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishHotspotConfigResponse) GetBody() *PublishHotspotConfigResponseBody {
	return s.Body
}

func (s *PublishHotspotConfigResponse) SetHeaders(v map[string]*string) *PublishHotspotConfigResponse {
	s.Headers = v
	return s
}

func (s *PublishHotspotConfigResponse) SetStatusCode(v int32) *PublishHotspotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishHotspotConfigResponse) SetBody(v *PublishHotspotConfigResponseBody) *PublishHotspotConfigResponse {
	s.Body = v
	return s
}

func (s *PublishHotspotConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
