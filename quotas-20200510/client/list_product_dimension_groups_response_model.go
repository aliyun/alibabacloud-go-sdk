// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductDimensionGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductDimensionGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductDimensionGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductDimensionGroupsResponseBody) *ListProductDimensionGroupsResponse
	GetBody() *ListProductDimensionGroupsResponseBody
}

type ListProductDimensionGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductDimensionGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductDimensionGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductDimensionGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListProductDimensionGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductDimensionGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductDimensionGroupsResponse) GetBody() *ListProductDimensionGroupsResponseBody {
	return s.Body
}

func (s *ListProductDimensionGroupsResponse) SetHeaders(v map[string]*string) *ListProductDimensionGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListProductDimensionGroupsResponse) SetStatusCode(v int32) *ListProductDimensionGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductDimensionGroupsResponse) SetBody(v *ListProductDimensionGroupsResponseBody) *ListProductDimensionGroupsResponse {
	s.Body = v
	return s
}

func (s *ListProductDimensionGroupsResponse) Validate() error {
	return dara.Validate(s)
}
