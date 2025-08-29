// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobFeatureReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobFeatureReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureConsistencyCheckJobFeatureReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) *ListFeatureConsistencyCheckJobFeatureReportsResponse
	GetBody() *ListFeatureConsistencyCheckJobFeatureReportsResponseBody
}

type ListFeatureConsistencyCheckJobFeatureReportsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobFeatureReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) GetBody() *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	return s.Body
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) SetBody(v *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) *ListFeatureConsistencyCheckJobFeatureReportsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponse) Validate() error {
	return dara.Validate(s)
}
