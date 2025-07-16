// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceFeatureViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatasourceFeatureViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatasourceFeatureViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatasourceFeatureViewsResponseBody) *ListDatasourceFeatureViewsResponse
	GetBody() *ListDatasourceFeatureViewsResponseBody
}

type ListDatasourceFeatureViewsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasourceFeatureViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasourceFeatureViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatasourceFeatureViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatasourceFeatureViewsResponse) GetBody() *ListDatasourceFeatureViewsResponseBody {
	return s.Body
}

func (s *ListDatasourceFeatureViewsResponse) SetHeaders(v map[string]*string) *ListDatasourceFeatureViewsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourceFeatureViewsResponse) SetStatusCode(v int32) *ListDatasourceFeatureViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasourceFeatureViewsResponse) SetBody(v *ListDatasourceFeatureViewsResponseBody) *ListDatasourceFeatureViewsResponse {
	s.Body = v
	return s
}

func (s *ListDatasourceFeatureViewsResponse) Validate() error {
	return dara.Validate(s)
}
