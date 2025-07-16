// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureEntitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureEntitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureEntitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureEntitiesResponseBody) *ListFeatureEntitiesResponse
	GetBody() *ListFeatureEntitiesResponseBody
}

type ListFeatureEntitiesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureEntitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureEntitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureEntitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureEntitiesResponse) GetBody() *ListFeatureEntitiesResponseBody {
	return s.Body
}

func (s *ListFeatureEntitiesResponse) SetHeaders(v map[string]*string) *ListFeatureEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureEntitiesResponse) SetStatusCode(v int32) *ListFeatureEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureEntitiesResponse) SetBody(v *ListFeatureEntitiesResponseBody) *ListFeatureEntitiesResponse {
	s.Body = v
	return s
}

func (s *ListFeatureEntitiesResponse) Validate() error {
	return dara.Validate(s)
}
