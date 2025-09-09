// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse
	GetBody() *UnTagResourcesResponseBody
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnTagResourcesResponse) GetBody() *UnTagResourcesResponseBody {
	return s.Body
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

func (s *UnTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
