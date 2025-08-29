// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureConsistencyCheckJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureConsistencyCheckJobsResponseBody) *ListFeatureConsistencyCheckJobsResponse
	GetBody() *ListFeatureConsistencyCheckJobsResponseBody
}

type ListFeatureConsistencyCheckJobsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureConsistencyCheckJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureConsistencyCheckJobsResponse) GetBody() *ListFeatureConsistencyCheckJobsResponseBody {
	return s.Body
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponse) SetBody(v *ListFeatureConsistencyCheckJobsResponseBody) *ListFeatureConsistencyCheckJobsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureConsistencyCheckJobsResponse) Validate() error {
	return dara.Validate(s)
}
