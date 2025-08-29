// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureConsistencyCheckJobConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureConsistencyCheckJobConfigsResponseBody) *ListFeatureConsistencyCheckJobConfigsResponse
	GetBody() *ListFeatureConsistencyCheckJobConfigsResponseBody
}

type ListFeatureConsistencyCheckJobConfigsResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) GetBody() *ListFeatureConsistencyCheckJobConfigsResponseBody {
	return s.Body
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) SetBody(v *ListFeatureConsistencyCheckJobConfigsResponseBody) *ListFeatureConsistencyCheckJobConfigsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureConsistencyCheckJobConfigsResponse) Validate() error {
	return dara.Validate(s)
}
