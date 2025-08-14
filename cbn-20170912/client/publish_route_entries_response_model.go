// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *PublishRouteEntriesResponseBody) *PublishRouteEntriesResponse
	GetBody() *PublishRouteEntriesResponseBody
}

type PublishRouteEntriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *PublishRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishRouteEntriesResponse) GetBody() *PublishRouteEntriesResponseBody {
	return s.Body
}

func (s *PublishRouteEntriesResponse) SetHeaders(v map[string]*string) *PublishRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *PublishRouteEntriesResponse) SetStatusCode(v int32) *PublishRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishRouteEntriesResponse) SetBody(v *PublishRouteEntriesResponseBody) *PublishRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *PublishRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
