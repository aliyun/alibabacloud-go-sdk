// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagVodResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagVodResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagVodResourcesResponse
	GetStatusCode() *int32
	SetBody(v *TagVodResourcesResponseBody) *TagVodResourcesResponse
	GetBody() *TagVodResourcesResponseBody
}

type TagVodResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagVodResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagVodResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagVodResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagVodResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagVodResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagVodResourcesResponse) GetBody() *TagVodResourcesResponseBody {
	return s.Body
}

func (s *TagVodResourcesResponse) SetHeaders(v map[string]*string) *TagVodResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagVodResourcesResponse) SetStatusCode(v int32) *TagVodResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagVodResourcesResponse) SetBody(v *TagVodResourcesResponseBody) *TagVodResourcesResponse {
	s.Body = v
	return s
}

func (s *TagVodResourcesResponse) Validate() error {
	return dara.Validate(s)
}
