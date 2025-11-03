// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpnRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishVpnRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishVpnRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *PublishVpnRouteEntryResponseBody) *PublishVpnRouteEntryResponse
	GetBody() *PublishVpnRouteEntryResponseBody
}

type PublishVpnRouteEntryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishVpnRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishVpnRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishVpnRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *PublishVpnRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishVpnRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishVpnRouteEntryResponse) GetBody() *PublishVpnRouteEntryResponseBody {
	return s.Body
}

func (s *PublishVpnRouteEntryResponse) SetHeaders(v map[string]*string) *PublishVpnRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *PublishVpnRouteEntryResponse) SetStatusCode(v int32) *PublishVpnRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishVpnRouteEntryResponse) SetBody(v *PublishVpnRouteEntryResponseBody) *PublishVpnRouteEntryResponse {
	s.Body = v
	return s
}

func (s *PublishVpnRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
