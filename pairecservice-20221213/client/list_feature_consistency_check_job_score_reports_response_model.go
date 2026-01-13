// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobScoreReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobScoreReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFeatureConsistencyCheckJobScoreReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListFeatureConsistencyCheckJobScoreReportsResponseBody) *ListFeatureConsistencyCheckJobScoreReportsResponse
	GetBody() *ListFeatureConsistencyCheckJobScoreReportsResponseBody
}

type ListFeatureConsistencyCheckJobScoreReportsResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFeatureConsistencyCheckJobScoreReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponse) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) GetBody() *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	return s.Body
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetHeaders(v map[string]*string) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.Headers = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetStatusCode(v int32) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) SetBody(v *ListFeatureConsistencyCheckJobScoreReportsResponseBody) *ListFeatureConsistencyCheckJobScoreReportsResponse {
	s.Body = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
