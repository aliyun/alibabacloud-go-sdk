// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishVpcRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishVpcRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishVpcRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *PublishVpcRouteEntriesResponseBody) *PublishVpcRouteEntriesResponse
	GetBody() *PublishVpcRouteEntriesResponseBody
}

type PublishVpcRouteEntriesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishVpcRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishVpcRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishVpcRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *PublishVpcRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishVpcRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishVpcRouteEntriesResponse) GetBody() *PublishVpcRouteEntriesResponseBody {
	return s.Body
}

func (s *PublishVpcRouteEntriesResponse) SetHeaders(v map[string]*string) *PublishVpcRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *PublishVpcRouteEntriesResponse) SetStatusCode(v int32) *PublishVpcRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishVpcRouteEntriesResponse) SetBody(v *PublishVpcRouteEntriesResponseBody) *PublishVpcRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *PublishVpcRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
