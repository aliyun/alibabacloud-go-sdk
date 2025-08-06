// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagVodResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnTagVodResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnTagVodResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UnTagVodResourcesResponseBody) *UnTagVodResourcesResponse
	GetBody() *UnTagVodResourcesResponseBody
}

type UnTagVodResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagVodResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagVodResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnTagVodResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnTagVodResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnTagVodResourcesResponse) GetBody() *UnTagVodResourcesResponseBody {
	return s.Body
}

func (s *UnTagVodResourcesResponse) SetHeaders(v map[string]*string) *UnTagVodResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagVodResourcesResponse) SetStatusCode(v int32) *UnTagVodResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagVodResourcesResponse) SetBody(v *UnTagVodResourcesResponseBody) *UnTagVodResourcesResponse {
	s.Body = v
	return s
}

func (s *UnTagVodResourcesResponse) Validate() error {
	return dara.Validate(s)
}
