// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagResourcesForRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagResourcesForRegionResponse
	GetStatusCode() *int32
	SetBody(v *ListTagResourcesForRegionResponseBody) *ListTagResourcesForRegionResponse
	GetBody() *ListTagResourcesForRegionResponseBody
}

type ListTagResourcesForRegionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesForRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesForRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForRegionResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagResourcesForRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagResourcesForRegionResponse) GetBody() *ListTagResourcesForRegionResponseBody {
	return s.Body
}

func (s *ListTagResourcesForRegionResponse) SetHeaders(v map[string]*string) *ListTagResourcesForRegionResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesForRegionResponse) SetStatusCode(v int32) *ListTagResourcesForRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesForRegionResponse) SetBody(v *ListTagResourcesForRegionResponseBody) *ListTagResourcesForRegionResponse {
	s.Body = v
	return s
}

func (s *ListTagResourcesForRegionResponse) Validate() error {
	return dara.Validate(s)
}
