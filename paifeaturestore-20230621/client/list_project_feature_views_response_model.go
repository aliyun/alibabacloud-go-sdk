// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectFeatureViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectFeatureViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectFeatureViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectFeatureViewsResponseBody) *ListProjectFeatureViewsResponse
	GetBody() *ListProjectFeatureViewsResponseBody
}

type ListProjectFeatureViewsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectFeatureViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectFeatureViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectFeatureViewsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectFeatureViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectFeatureViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectFeatureViewsResponse) GetBody() *ListProjectFeatureViewsResponseBody {
	return s.Body
}

func (s *ListProjectFeatureViewsResponse) SetHeaders(v map[string]*string) *ListProjectFeatureViewsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectFeatureViewsResponse) SetStatusCode(v int32) *ListProjectFeatureViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectFeatureViewsResponse) SetBody(v *ListProjectFeatureViewsResponseBody) *ListProjectFeatureViewsResponse {
	s.Body = v
	return s
}

func (s *ListProjectFeatureViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
