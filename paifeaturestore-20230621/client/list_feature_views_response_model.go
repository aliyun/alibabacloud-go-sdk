// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureViewsResponseBody) *ListFeatureViewsResponse
	GetBody() *ListFeatureViewsResponseBody
}

type ListFeatureViewsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureViewsResponse) GetBody() *ListFeatureViewsResponseBody {
	return s.Body
}

func (s *ListFeatureViewsResponse) SetHeaders(v map[string]*string) *ListFeatureViewsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureViewsResponse) SetStatusCode(v int32) *ListFeatureViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureViewsResponse) SetBody(v *ListFeatureViewsResponseBody) *ListFeatureViewsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
